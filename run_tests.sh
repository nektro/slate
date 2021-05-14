#!/usr/bin/env bash

set -e

err_report() {
    echo "exited with non-zero exit code: $1"
}

trap 'err_report $?' ERR

go build

for file in $(find tests | sort | grep \\.slate$)
do
    echo $file
    ./slate -run $file
    llvm-as-11 out.ll
    lli-11 out.bc
done
