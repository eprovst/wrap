#! /bin/bash

echo "Removing previous build..."
rm -rf build

# Note: packaging and building for Linux
#       is done through build.snapcraft.io

# Build for 64bit Windows
echo "Building for Windows 64 bit..."
go generate # Prepare resource.syso
GOARCH=amd64 GOOS=windows go build -o build/windows/feltix.exe
rm resource.syso # Remove resource.syso

# Build for 64bit Darwin
echo "Building for Darwin 64 bit..."
GOARCH=amd64 GOOS=darwin go build -o build/darwin/feltix

# Packaging
echo "Packaging for macOS..."
bash scripts/macos.sh

echo "Build finished."
