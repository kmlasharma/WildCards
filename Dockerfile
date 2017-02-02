FROM debian:jessie

WORKDIR /workdir

COPY Scripts/install_peos.sh install_peos.sh

RUN ./install_peos.sh

