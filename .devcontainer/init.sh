#!/bin/bash

go install

go install github.com/segmentio/golines@v0.12.2
go install golang.org/x/tools/cmd/godoc@v0.26.0
go install github.com/766b/go-outliner@latest

go mod vendor

