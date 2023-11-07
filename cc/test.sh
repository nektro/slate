#!/usr/bin/env bash

set -eu

assert() {
    expected="$1"
    input="$2"

    ./slatecc "$input" > tmp.s || exit
    gcc -o tmp tmp.s
    ./tmp
    actual="$?"

    # all the .s files for later inspection
    cp ./tmp.s "./s/$(sha256sum tmp.s | head -c 64).s"

    if [ "$actual" = "$expected" ]; then
        echo "$input => $actual"
    else
        echo "$input => $expected expected, but got $actual"
        exit 1
    fi
}

assert 0 0
assert 42 42

echo OK
