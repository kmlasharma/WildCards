#!/bin/bash 

# Running this shell script requires root access 

# Change to $HOME as PEOS depends on relative paths
cd $HOME

# Clone PEOS
git clone https://github.com/jnoll/peos

# Install dependencies
sudo DEBIAN_FRONTEND=noninteractive \
apt-get install -y build-essential \
			bison \
			check \
			flex \
			libncurses5 \
			libncurses5-dev \
			make \
			openssl \
			libssl-dev \
			readline-common \
			libreadline6 \
			libreadline6-dev \
			libxml2 \
			libxml2-dev \
			libxml2-doc \
			util-linux \
			tcl \
			tcl-dev \
			libxslt1.1 \
			libxslt1-dev \
			expect \
			expect-dev
sudo DEBIAN_FRONTEND=noninteractive apt-get update

# Make PEOS and run tests
cd $HOME/peos
make

