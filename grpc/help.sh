#!/bin/bash
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
protoc --go_out=plugins=grpc:. grpc/warectu/*.proto