package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env          string        `yaml:"env" env-default:"local"`
	Storage_path string        `yaml:"storage_path" env-required:"true"`
	TokenTTL     time.Duration `yaml:"token_ttl" env-default:"1h"`
	GRPC         GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"44044"`
	Timeout time.Duration `yaml:"timeout" env-default:"1h"`
}

func MustLoad() *Config { // Must приставка, когда функция не будет возвращать ошибку а будет выходить
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("confug path is empy")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exists " + configPath)
	}
	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config " + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
