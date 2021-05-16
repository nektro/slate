#!/usr/bin/env bash

$@ 2> /dev/null

if [ "$?" == "0" ]
then
    exit 1
fi
