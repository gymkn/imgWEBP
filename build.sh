#!/usr/bin/env bash
set -euo pipefail

mkdir -p dist && \
GOOS=linux   GOARCH=amd64 go build -ldflags="-s -w" -o dist/imgwebp-linux-amd64 && \
GOOS=linux   GOARCH=arm64 go build -ldflags="-s -w" -o dist/imgwebp-linux-arm64 && \
GOOS=darwin  GOARCH=amd64 go build -ldflags="-s -w" -o dist/imgwebp-macos-amd64 && \
GOOS=darwin  GOARCH=arm64 go build -ldflags="-s -w" -o dist/imgwebp-macos-arm64 && \
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/imgwebp-win-amd64.exe && \
GOOS=windows GOARCH=arm64 go build -ldflags="-s -w" -o dist/imgwebp-win-arm64.exe