#!/bin/bash

curl "https://adventofcode.com/2023/day/$1/input" \
    -H "Cookie: session=$AOC_COOKIE" \
    -o "input.txt"
