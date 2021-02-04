#!/bin/bash

export GO111MODULE=on
go mod init
go mod tidy

go run cmd/playing-with-golang-on-k8s.go