# Image rooted at debian:jessie
FROM golang:1.6

# Build dinto components
COPY ./src/ /go/src/app
WORKDIR /go/src/app/
RUN go get -d -v
RUN go install -v

# Build peos components
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends apt-utils
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends apt-utils 
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y sudo git
RUN DEBIAN_FRONTEND=noninteractive apt-get update


# Has to be /root as $HOME is set to /root
# This is a restriction in place by PEOS
WORKDIR /root
COPY ./src/peos/install_peos.sh install_peos.sh
RUN ./install_peos.sh
RUN rm ./install_peos.sh


# Install Python 3
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y python2.7 python2.7-dev python-pip
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN pip install ontospy
RUN DEBIAN_FRONTEND=noninteractive apt-get update

