#!/bin/bash

# Root permissions are required to run this script

# Only install Docker if not already installed.
# If Docker is already on the machine then redownloading it can cause issues
if [ "$(uname)" != "Darwin" ]; then
	# Linux.
	# Install dependencies
	sudo apt-get install curl \
				linux-image-extra-$(uname -r) \
				linux-image-extra-virtual \
				software-properties-common \
				apt-transport-https \
				ca-certificates
	sudo apt-get update

	# Install Docker
	which docker
	if [ $? -ne 0 ]; then
		curl -fsSL https://get.docker.com/ | sh
	fi

	# Add Docker's GPG key
	curl -fsSL https://yum.dockerproject.org/gpg | sudo apt-key add -

	# Add Docker as a known repository
	sudo add-apt-repository \
		"deb https://apt.dockerproject.org/repo/ \
			ubuntu-$(lsb_release -cs) \
			main"
	sudo apt-get update

	# Install Docker Engine
	sudo apt-get -y install docker-engine
	sudo apt-get update
	
	# Install Docker Compose
	which docker-compose
	if [ $? -ne 0 ]; then
		sudo curl -L "https://github.com/docker/compose/releases/download/1.10.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
		sudo chmod +x /usr/local/bin/docker-compose
	fi
else
	ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	# Install Cask
	brew install caskroom/cask/brew-cask
	# Install docker toolbox
	brew cask install docker-toolbox
	brew cask install virtualbox
fi

