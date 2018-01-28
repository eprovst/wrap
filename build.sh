#! /bin/bash

# Note: packaging and building for Linux
#       is done through build.snapcraft.io

echo "Updating data..."
bash scripts/update_data.sh

echo "Removing old installers and packages..."
rm -rf dist

# Build for 64bit Windows
echo "Building for Windows 64 bit..."
go generate # Prepare resource.syso
GOARCH=amd64 GOOS=windows go build -o build/windows/wrap.exe
rm resource.syso # Remove resource.syso

# Build for 64bit Darwin
echo "Building for Darwin 64 bit..."
GOARCH=amd64 GOOS=darwin go build -o build/darwin/wrap

echo "Build finished. Start packaging..."

# Packaging
echo "Packaging for macOS..."
bash scripts/macos.sh

echo "Packaging for Windows..."
bash scripts/windows.sh

echo "Done packaging. Removing build files..."
rm -rf build

echo "Done."
