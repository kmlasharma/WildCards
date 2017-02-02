# CS4098 Group Project
## Team Name: Wildcards

## Installation:

### PEOS
1. Run `install_peos.sh` to install PEOS and its dependencies


### Installing the Go programming language on macOS Sierra
1. Install Go using Homebrew
    * Run `ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"` to install Homebrew
    * Run `brew doctor` to test that Homebrew installed correctly
    * Run `brew install go` to install Go
2. Create and set your $GOPATH
    * Create a directory to use as Go's workspace, eg. '/go'
    * Add these two lines to *~/.bashrc*, with your chosen directory in place of '/go': <br /> `export GOPATH=$HOME/go` <br />
	 `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`
