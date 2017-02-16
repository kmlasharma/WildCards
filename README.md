# CS4098 Group Project
## Team Name: Wildcards

## Installation:

All that is required to run the project is Docker:

1. Running `make install` will ensure that all docker components needed (Docker, Docker Engine and Docker Compose) are installed. Accept any prompts during installation.

## Setting up the Project:

1. To ensure your computer is setup to run the project, run one of the following (to start the docker service):
 * `make clean` then `make setup` if you are outside of the Trinity College network
 * `make clean` then `make setuptrinity` if you are inside of the Trinity College network
2. Restart your machine. This will apply the changes in the above steps properly and will allow you to run the following steps. If you do not restart your machine you will run into issues.
3. Run `docker-compose build` to build the project.

The project is now set up and ready for use.

## Running the project:

1. Run `docker-compose run project` to enter the Docker container.
2. Verify that the project is set up properly:
 * Run `app`
 * Run `cd /go/src/app` followed by `go test`
 * Run `cd` followed by `pmlcheck peos/xpml/test.pml`

## Features:

### PML File Selection

* Status: In Progress
* Testing instructions:

### PML File Loading

* Status: In Progress
* Testing Instructions:

### Select Specific OWL Ontology

* Status: In Progress
* Testing Instructions:

### Load Selected Ontology

* Status: In Progress
* Testing Instructions:

## Tearing down the Docker container

1. Run `docker-compose stop` to completely tear down the Docker container
