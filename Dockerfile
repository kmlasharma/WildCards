# Image rooted at debian:jessie
FROM golang:1.6

# Set environment variables
ENV RES_DIR /go/src/app
ENV LOG_DIR /go/src/app/log

WORKDIR /go/src/app

# Install testing dependencies
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y python3 \
			python3-dev \
			python3-pip \
			libncurses5 \
			libncurses5-dev \
			readline-common \
			libreadline6 \
			libreadline6-dev \
			vim 
RUN pip3 install nose

# Copy app assets
COPY ./src/ /go/src/app
COPY ./res/ /go/src/app
COPY ./tests/ /go/src/app/tests
COPY ./pkgs/ /go/src/app/pkgs
COPY ./utils/scripts/test.sh /go/bin/tests

# Install app dependencies
RUN go get -d -v
RUN go install -v

# Working directory of the project is where all resources are
WORKDIR /go/src/app

