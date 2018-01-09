#! /bin/bash

# This file will be runned from the project root
mkdir ./build/darwin/feltix.app
mkdir ./build/darwin/feltix.app/Contents
mkdir ./build/darwin/feltix.app/Contents/MacOS
mkdir ./build/darwin/feltix.app/Contents/Resources

cp ./data/Info.plist ./build/darwin/feltix.app/Contents/Info.plist
cp ./build/darwin/feltix ./build/darwin/feltix.app/Contents/MacOS/feltix
cp ./icons/feltix.icns ./build/darwin/feltix.app/Contents/Resources/feltix.icns

cd ./build/darwin
zip -r Feltix_macOS.zip feltix.app
cd ../..

rm -rf ./build/darwin/feltix.app
rm -rf ./build/darwin/feltix
