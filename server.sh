#!/bin/bash
go run cmd/server/main.go -db-host localhost -db-schema todo -db-user root -db-password root -grpc-port 9090 -http-port=8080 $*
