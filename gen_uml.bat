@echo off
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
goplantuml model >> dist/model.pu
goplantuml controller >> dist/controller.pu
goplantuml service >> dist/service.pu
goplantuml model controller service >> dist/all.pu