#!/usr/bin/env bash

set -e
set -x

go build
./slate \
-run './tests/basics/01.slate' \

llvm-as-10 out.ll

lli-10 out.bc
