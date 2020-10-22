#!/bin/bash

export GO111MODULE=on
go mod init
go mod tidy

go run cmd/main.go