#! /bin/bash

# Note: packaging and building for Linux
#       is done through build.snapcraft.io

# Run normal build
bash ./build.sh

# Rename results
mv ./dist/Wrap_macOS.zip ./dist/Wrap_macOS_nightly.zip
mv ./dist/Wrap_Win64.exe ./dist/Wrap_Win64_nightly.exe