#!/bin/bash

for dir in `ls pkgs/`;
do
    cd pkgs/$dir
    go build -i
    rm -rf vendor
    govendor init
    govendor add +external
    cd ../..
done

cd src
rm -rf vendor
govendor init
govendor add +external
