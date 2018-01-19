#! /bin/bash

# This file will be run from the project root

# Generate bash autocompletion for snap
go run scripts/generate_bashcompletion.go
mv complete.sh prime/complete.sh
