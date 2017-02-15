FROM debian:jessie

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

COPY install_peos.sh install_peos.sh

RUN ./install_peos.sh
RUN rm ./install_peos.sh

