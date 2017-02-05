#!/bin/bash

sudo service docker stop
sudo groupadd docker
sudo gpasswd -a ${USER} docker
sudo service docker start
