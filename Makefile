.PHONY: run

run:
	go run cmd/sso/main.go --config=./config/local.yaml

.DEFAULT_GOAL := run