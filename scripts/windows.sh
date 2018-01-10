#! /bin/bash

# This file will be run from the project root
makensis -V2 -DARCH=x64 data/installer.nsi
rm ./build/windows/feltix.exe
