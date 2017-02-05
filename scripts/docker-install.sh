#!/bin/bash

if [ "$(uname)" = "Darwin" ]; then
  chmod +x ./docker-install.mac.sh
  ./docker-install.mac.sh
else
  chmod +x ./docker-install.linux.sh
  ./docker-install.linux.sh
fi