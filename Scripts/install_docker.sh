#!/bin/bash

# Root permissions are required to run this script

# Only install Docker if not already installed.
# If Docker is already on the machine then redownloading it can cause issues
which docker
if [ $? -ne 0 ]; then
		if [ "$(uname)" != "Darwin" ]; then
			# This is a linux machine. Install dependencies
			sudo apt-get install -y curl
			sudo apt-get update
		fi
		# Install Docker
		curl -fsSL https://get.docker.com/ | sh
fi

