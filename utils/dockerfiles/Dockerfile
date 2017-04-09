# Image rooted at debian:jessie
FROM tomasbarry/groupproject:latest

# Copy app assets
COPY ./src/ /go/src/app
COPY ./res/ /go/src/app
COPY ./pkgs/ /go/src/app/pkgs
COPY ./utils/scripts/test.sh /go/bin/tests

# Install app dependencies
RUN go get -d -v
RUN go install -v

