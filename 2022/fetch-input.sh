#!/bin/bash

if [ -z "$1" ]
  then
    echo "No argument supplied"
fi

echo "Setting up day $1"

dir="day$1"
mkdir -p $dir

curl "https://adventofcode.com/2022/day/$1/input" \
    -H "Cookie: session=$AOC_COOKIE" \
    --output $dir/input
