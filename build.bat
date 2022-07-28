@echo off

set GOOS=windows
go build -o=bin/castle.exe main.go
echo "Built for windows."

set GOOS=linux
go build -o=bin/castle main.go
echo "Built for linux."

set GOOS=darwin
go build -o=bin/castle-macos main.go
echo "Built for mac."