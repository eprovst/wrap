#! /bin/bash

echo "Removing previous build..."
rm -rf build

# Note: packaging and building for Linux
#       is done through build.snapcraft.io

# Build for older 32bit processors.
echo "Building for Windows 32 bit..."
GOARCH=386 GOOS=windows go build -o build/windows/386/feltix.exe

# Build for 64bit.
echo "Building for Windows 64 bit..."
GOARCH=amd64 GOOS=windows go build -o build/windows/amd64/feltix.exe
echo "Building for Darwin 64 bit..."
GOARCH=amd64 GOOS=darwin go build -o build/darwin/feltix

# Packaging
echo "Packaging for macOS..."
bash scripts/macos.sh

echo "Build finished."
