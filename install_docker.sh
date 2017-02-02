#!/bin/bash

# Root permissions are required to run this script

# Install dependencies
sudo apt-get install -y curl
sudo apt-get update

# Install Docker if not already installed
which docker
if [ $? -ne 0 ] 
	then
		curl -fsSL https://get.docker.com/ | sh
fi

