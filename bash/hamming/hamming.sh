#!/usr/bin/env bash

if [ $# -ne 2 ];then
  echo "Usage: hamming.sh <string1> <string2>"
  exit 1
fi
if [[ "$1" == "$2" ]]; then
    echo 0
    exit 0
elif [[ "${#1}" != "${#2}" ]]; then
    echo "strands must be of equal length"
    exit 1
else
counter=0
for (( i=0; i<${#1}; i++ )); do
  if [[ "${1:$i:1}" != "${2:$i:1}" ]]; then
    counter=$((counter+1))
  fi
done
    echo $counter
    exit 0
fi
