#!/bin/bash

for dir in `ls pkgs/`;
do
    cd pkgs/$dir
    go test
    cd ../..
done