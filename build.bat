@echo off
go build main.go
move main.exe ./bin/castle.exe
cls
.\bin\castle.exe test/castle.yaml