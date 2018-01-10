#! /bin/bash

# This file will be run from the project root
mkdir -p ./build/darwin/feltix.app/Contents/MacOS
mkdir -p ./build/darwin/feltix.app/Contents/Resources

cp ./data/Info.plist ./build/darwin/feltix.app/Contents/Info.plist
cp ./build/darwin/feltix ./build/darwin/feltix.app/Contents/MacOS/feltix
cp ./icons/feltix.icns ./build/darwin/feltix.app/Contents/Resources/feltix.icns

cd ./build/darwin
zip -q -r Feltix_macOS.zip feltix.app
cd ../..

rm -rf ./build/darwin/feltix.app
rm ./build/darwin/feltix
