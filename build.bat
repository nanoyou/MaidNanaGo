@echo off
go install github.com/swaggo/swag/cmd/swag@latest
swag init
go build -o dist/MaidNanaGo.exe