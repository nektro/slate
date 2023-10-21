#!/usr/bin/env bash

assert() {
    expected="$1"
    input="$2"

    ./chibicc "$input" > tmp.s || exit
    gcc -o tmp tmp.s
    ./tmp
    actual="$?"

    # PATCH save all the .s files for later inspection
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
assert 21 '5+20-4'
assert 41 ' 12 + 34 - 5 '

echo OK
