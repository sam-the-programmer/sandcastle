@echo off
go build -o=./bin/castle.exe main.go
cls
.\bin\castle.exe -c test/castle.yaml task admin