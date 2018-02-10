#! /bin/bash

# This file will be run from the project root
mkdir -p ./dist/macOS/wrap.app/Contents/MacOS
mkdir -p ./dist/macOS/wrap.app/Contents/Resources

cp ./scripts/Info.plist ./dist/macOS/wrap.app/Contents/Info.plist
cp ./build/darwin/wrap ./dist/macOS/wrap.app/Contents/MacOS/wrap
cp ./assets/wrap.icns ./dist/macOS/wrap.app/Contents/Resources/wrap.icns

cd ./dist/macOS
zip -q -r ../Wrap_macOS.zip wrap.app
cd ../..

rm -rf ./dist/macOS

