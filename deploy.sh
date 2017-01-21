#!/usr/bin/env bash
rm -rf ./dist/*
GOOS=linux go build -o ./dist/http-fs main.go
docker build -t 0x4139/http-fs .
docker push  0x4139/http-fs