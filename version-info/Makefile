
VERSION?="v0.0.1"
GIT_COMMIT?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell date)

.PHONY: build

build:
	go build \
	-ldflags="-X 'main.GitCommit=${GIT_COMMIT}' \
	-X 'main.Version=${VERSION}' \
	-X 'main.BuildDate=${BUILD_DATE}'" main.go