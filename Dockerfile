# Image rooted at debian:jessie
FROM golang:1.6

# Build dinto components
COPY ./src/ /go/src/app
WORKDIR /go/src/app/
RUN go get -d -v
RUN go install -v

# Install Python 3
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y python3 \
			python3-dev \
			python3-pip \
			build-essential \
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
			expect-dev \
			vim \
			tmux
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN pip3 install ontospy
RUN DEBIAN_FRONTEND=noninteractive apt-get update

RUN echo 'ln -s -f /go/src/app/pml/test.pml $HOME/test.pml' >> $HOME/.bashrc
RUN echo 'ln -s -f /go/src/app/dinto/data/test.owl $HOME/test.owl' >> $HOME/.bashrc

WORKDIR /root/

