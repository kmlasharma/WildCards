#!/bin/bash

which brew
if [ $? -ne 0 ]; then
  # Install homebrew
  ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
fi

# Install Cask
brew install caskroom/cask/brew-cask
brew cask install docker-toolbox
brew cask install virtualbox