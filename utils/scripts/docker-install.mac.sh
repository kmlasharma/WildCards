#!/bin/bash

printf "Checking if homebrew is installed..\n"
which brew
if [ $? -ne 0 ]; then
  printf "Installing homebrew..\n"
  # Install homebrew
  ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
else
  printf "Already installed!\n"
fi

printf "Installing docker-toolbox and virtualbox..\n\n"
# Install Cask
brew install caskroom/cask/brew-cask
brew cask install docker-toolbox
brew cask install virtualbox