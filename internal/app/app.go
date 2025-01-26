package app

import (
	"log/slog"
	"sso/internal/app/grpc"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int,storagePath string,tokenTTL time.Duration) *App{
	//TODO storage

	//TODO auth service

	grpcApp:=grpcapp.New(log,grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
