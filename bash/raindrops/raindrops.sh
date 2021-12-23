#!/bin/bash

string=""

if [[ $(expr $1 % 3) == 0 ]]; then
    string+="Pling"
fi

if [[ $(expr $1 % 5) == 0 ]]; then
    string+="Plang"
fi

if [[ $(expr $1 % 7) == 0 ]]; then
    string+="Plong"
fi

if [[ -z $string ]]; then
    string=$1
fi

echo $string