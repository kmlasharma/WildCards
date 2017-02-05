#!/bin/bash

if [ "$(uname)" = "Darwin" ]; then
  chmod +x ./docker-setup.mac.sh
  ./docker-setup.mac.sh
else
  chmod +x ./docker-setup.linux.sh
  ./docker-setup.linux.sh
fi