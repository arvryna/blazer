#! /bin/bash

rm -rf out build
mkdir out build
go build -o build/blazer main.go

