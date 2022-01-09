#!/usr/bin/env bash

if [[ "$1" == "$2" ]]; then
    echo 0
    exit 0
elif [[ "${#1}" != "${#2}" ]]; then
    echo "strands must be of equal length"
else
    echo 1
    exit 0
fi