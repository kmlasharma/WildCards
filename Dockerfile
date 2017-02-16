# Image rooted at debian:jessie
FROM golang:1.6-onbuild

# Build dinto components
COPY ./dinto /go/src/app
WORKDIR /go/src/app
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
COPY ./peos/install_peos.sh install_peos.sh
RUN ./install_peos.sh
RUN rm ./install_peos.sh

# Tell the user where they are
CMD ["echo", "pwd"]
