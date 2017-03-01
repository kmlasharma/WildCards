# Image rooted at debian:jessie
FROM golang:1.6

# Build dinto components
COPY ./src/ /go/src/app
WORKDIR /go/src/app/
RUN go get -d -v
RUN go install -v

# Install Python 3
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y python3 python3-dev python3-pip
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN pip3 install ontospy
RUN DEBIAN_FRONTEND=noninteractive apt-get update

RUN echo 'ln -s -f $HOME/peos/xpml/test.pml $HOME/test.pml' >> ~/.bashrc
RUN echo 'ln -s -f /go/src/app/dinto/data/test.owl $HOME/test.owl' >> ~/.bashrc

