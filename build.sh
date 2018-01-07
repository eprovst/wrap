#! /bin/bash

rm -rf build

# Note: packaging and building for Linux
#       is done through build.snapcraft.io

# Build for older 32bit processors.
GOARCH=386 GOOS=windows go build -o build/windows/386/feltix.exe

# Build for 64bit.
GOARCH=amd64 GOOS=windows go build -o build/windows/amd64/feltix.exe
GOARCH=amd64 GOOS=darwin go build -o build/darwin/feltix

# TODO: packaging...
