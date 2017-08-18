#! /bin/bash

rm -rf build

# Build for older 32bit processors.
GOARCH=386

GOOS=linux
go build -o build/linux/386/feltix

GOOS=windows
go build -o build/windows/386/feltix.exe

# Build for 64bit.
GOARCH=amd64

GOOS=linux
go build -o build/linux/amd64/feltix

GOOS=windows
go build -o build/windows/amd64/feltix.exe


GOOS=darwin
go build -o build/darwin/feltix


# TODO: packaging...