#!/bin/bash

# IF there is no docker machines provisioned
docker-machine ls | tee >(wc -l)
if [ $? -eq 1 ]; then
  docker-machine create --driver virtualbox default
  eval $(docker-machine env default)
else
  printf "Docker machine already provisioned. You're good to go!\n"
fi
docker-machine start default
