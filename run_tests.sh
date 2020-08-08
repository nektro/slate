#!/usr/bin/env bash

set -e

go build

for file in $(find ./tests/ -type f)
do
    ./slate -run $file
done
