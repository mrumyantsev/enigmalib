.SILENT:
.DEFAULT_GOAL := run-demo

include ./.env

export ALPINE_VER
export ENABLE_DEBUG_LOGS
export GO_VER
export HTTP_SERVER_LISTEN_PORT
export NGINX_VER
export SERVER_APP_NAME
export WEB_LOCAL_DIR
export WEB_PORT

.PHONY: build
build:
	go build -o ./build/${SERVER_APP_NAME} ./cmd/${SERVER_APP_NAME}/main.go

.PHONY: run
run:
	./build/${SERVER_APP_NAME}

.PHONY: run-demo
run-demo:
	go run ./cmd/demo
