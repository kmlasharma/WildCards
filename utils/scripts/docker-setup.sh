#!/bin/bash

printf "Setting up docker VM if one does not already exist..\n\n"
if [ "$(uname)" = "Darwin" ]; then
  ./docker-setup.mac.sh
else
  ./docker-setup.linux.sh
fi