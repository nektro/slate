#!/usr/bin/env bash

set -e
set -x

go build
./slate \
-run 'tests/pass/behavior/01_asm_hello_world.slate' \

llvm-as-12 out.ll

lli-12 out.bc
