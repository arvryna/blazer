#! /bin/bash

rm -rf build
mkdir build
go build -o build/blazer main.go
sudo cp build/blazer /usr/local/bin/
