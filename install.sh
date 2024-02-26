#!/bin/bash
apt-get update && apt-get install -y protobuf-compiler make
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
#for macos
#brew install golang-migrate
export PATH="$PATH:$(go env GOPATH)/bin"
