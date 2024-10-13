#!/bin/bash

go install

go install github.com/segmentio/golines@v0.12.2
go install golang.org/x/tools/cmd/godoc@v0.26.0

go mod vendor

