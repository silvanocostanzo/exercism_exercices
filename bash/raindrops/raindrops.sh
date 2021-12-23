#!/bin/bash

FACTOR_THREE=$(($1 % 3 == 0))
FACTOR_FIVE=$(($1 % 5 == 0))
FACTOR_SEVEN=$(($1 % 7 == 0))

if [[ $FACTOR_THREE == 0 && $FACTOR_FIVE == 0 && $FACTOR_SEVEN == 0 ]]; then
    echo $1
    exit 0
fi


string=""

if [[ $FACTOR_THREE == 1 ]]; then
    string+="Pling"
fi

if [[ $FACTOR_FIVE == 1 ]]; then
    string+="Plang"
fi

if [[ $FACTOR_SEVEN == 1 ]]; then
    string+="Plong"
fi

echo $string