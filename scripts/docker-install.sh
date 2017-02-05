#!/bin/bash

if [ "$(uname)" = "Darwin" ]; then
  ./docker-install.mac.sh
else
  ./docker-install.linux.sh
fi