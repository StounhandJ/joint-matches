#!/bin/bash

set GOOS=linux
set GOARCH=amd64
go build -o joint-linix64.exe main.go