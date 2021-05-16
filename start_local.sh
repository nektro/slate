#!/usr/bin/env bash

set -e
set -x

go build
./slate \
-run 'tests/basics/01_asm_hello_world.slate' \

llvm-as-11 out.ll

lli-11 out.bc
