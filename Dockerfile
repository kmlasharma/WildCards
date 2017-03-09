# Image rooted at debian:jessie
FROM golang:1.6

# Build dinto components
COPY ./src/ /go/src/app
COPY ./res/ /go/src/app/res
WORKDIR /go/src/app/
RUN go get -d -v
RUN go install -v

# Install Python 3
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
RUN pip3 install ontospy
RUN DEBIAN_FRONTEND=noninteractive apt-get update

# testing
WORKDIR /root/
COPY ./tests/ /root/tests/
RUN pip3 install nose

RUN echo 'ln -s -f /go/src/app/res/test.pml $HOME/test.pml' >> $HOME/.bashrc
RUN echo 'ln -s -f /go/src/app/res/test.owl $HOME/test.owl' >> $HOME/.bashrc

WORKDIR /root/

