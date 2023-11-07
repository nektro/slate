#!/usr/bin/env bash

set -eu

nix-shell --run 'cd chibicc && make test'
