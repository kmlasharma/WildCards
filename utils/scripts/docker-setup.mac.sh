#!/bin/bash

# IF there is no docker machines provisioned
docker-machine ls | tee >(wc -l)
if [ $? -eq 1 ]; then
  docker-machine create --driver virtualbox default
else
  printf "Docker machine already provisioned. You're good to go!\n"
fi
eval $(docker-machine env default)
docker-machine start default
