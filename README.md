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
2. Run `app` to start the application.

## Features:

### PML File Selection

* Status: Complete
* Testing instructions:
   * Enter `test.pml` at the prompt

### PML File Loading

* Status: Complete
* Testing Instructions:
   * By doing the previous step, the PML file will automatically be loaded

### Select Specific OWL Ontology

* Status: Complete
* Testing Instructions:
   * Enter `test.owl` at the prompt

### Load Selected Ontology

* Status: Complete
* Testing Instructions:
   * By doing the previous step, the OWL Ontology will automatically be loaded

## Tearing down the Docker container

1. Run `docker-compose stop` to completely tear down the Docker container
