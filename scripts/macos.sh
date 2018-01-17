#! /bin/bash

# This file will be run from the project root
mkdir -p ./dist/macOS/feltix.app/Contents/MacOS
mkdir -p ./dist/macOS/feltix.app/Contents/Resources

cp ./data/Info.plist ./dist/macOS/feltix.app/Contents/Info.plist
cp ./build/darwin/feltix ./dist/macOS/feltix.app/Contents/MacOS/feltix
cp ./data/feltix.icns ./dist/macOS/feltix.app/Contents/Resources/feltix.icns

cd ./dist/macOS
zip -q -r ../Feltix_macOS.zip feltix.app
cd ../..

rm -rf ./dist/macOS

