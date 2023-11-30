#!/bin/bash

curl "https://adventofcode.com/2022/day/$1/input" \
    -H "Cookie: session=$AOC_COOKIE" \
    --output $dir/input
