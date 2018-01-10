#! /bin/bash

# This file will be run from the project root
wixl data/installer.wxs -o "build/windows/Feltix.msi"
rm ./build/windows/feltix.exe