#!/bin/bash

echo "Running test suite. This may take a minute.."
for dir in `ls pkgs/`;
do
    cd pkgs/$dir
    go test
    cd ../..
done
