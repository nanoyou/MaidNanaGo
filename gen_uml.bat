@echo off
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
goplantuml model >> dist/model.pu