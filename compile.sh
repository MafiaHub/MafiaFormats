#!/bin/bash

# utils

function genspec {
    specgen --file=specs/$1 --lang=md | sed -e 's/## /### /'
}

function gendoc {
    cp templates/$1.md docs/$1.md
    genspec $1.gspec.go >> docs/$1.md
}

function gentool {
    mkdir -p parsers/$1/$2 2> /dev/null
    specgen --file=specs/$1.gspec.go --lang=$2 | sed -e 's/'$1'/main/' > parsers/$1/$2/$1.$2
}

function genall {
    gendoc $1
    gentool $1 go
}

# formats

genall checkbin
genall mnu
genall textdb

echo Done!

