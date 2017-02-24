#!/bin/bash

printf "Installing docker components if not already installed..\n\n"
if [ "$(uname)" = "Darwin" ]; then
  ./docker-install.mac.sh
else
  ./docker-install.linux.sh
fi