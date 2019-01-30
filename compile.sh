#!/bin/bash

# utils

function genspec {
    specgen --file=$1 --lang=md | sed -e 's/## /### /'
}

function gendoc {
    cp templates/$1.md docs/$1.md
    genspec specs/$1.gspec.go >> docs/$1.md
}

gendoc checkbin

echo Done!

