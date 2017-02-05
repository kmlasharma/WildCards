#!/bin/bash

if [ "$(uname)" = "Darwin" ]; then
  ./docker-setup.mac.sh
else
  ./docker-setup.linux.sh
fi