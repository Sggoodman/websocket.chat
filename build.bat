@echo off
REM Build backend using goreleaser
docker build -t websocket-server:latest -f Dockerfile.goreleaser .

REM Start services
docker-compose up -d
