#!/bin/bash

RESULT=0
echo "Running test suite. This may take a minute.."
for dir in `ls pkgs/`;
do
    cd pkgs/$dir
    go test
    PASSED=$?
    RESULT=$(($RESULT+$PASSED))
    cd ../..
done

cd src
go test
PASSED=$?
RESULT=$(($RESULT+$PASSED))

echo "####################################################################"
echo "####################################################################"
if [ $RESULT -eq 0 ]; then
	echo "All tests passed"
else
	echo "Some tests failed"
fi
echo "####################################################################"
echo "####################################################################"

