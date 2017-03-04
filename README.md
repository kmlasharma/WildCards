# CS4098 Group Project
## Team Name: Wildcards

If you are cloning this repo for the first time:

1. Run the steps in `Installation`.
2. Run the steps in `Setting Up The Project`.
3. Run the steps in `Running The Project`.

If you have already cloned the repo, and would like the most up to date version:

1. Run the steps in `Setting Up The Project`.
2. Run the steps in `Running The Project`.

## Installation:

Docker is all you need for the project. To install Docker:

1. Run `make install` (this will ensure that all docker components needed (Docker, Docker Engine and Docker Compose) are installed. Accept any prompts during installation.
2. Restart your machine. This will apply the changes in the above steps properly and will allow you to run the following steps. If you do not restart your machine you will run into issues. You need to restart your machine to apply user group changes.

## Setting up the Project:

1. To ensure your computer is setup to run the project, run one of the following:
 * `make setup` if you are outside of the Trinity College network.
 * `make setuptrinity` if you are inside of the Trinity College network.
2. Run `docker-compose build --no-cache` to build the project (the no-cache flag ensures a clean build).

## Running the project:

1. Run `docker-compose run project` to enter the Docker container.
2. Run `app` to start the application.

## Features:

It is assumed that you have followed the instructions above before testing the features listed below. Each feature will clarify it's own starting context.

### PML File Selection

* Status: **Complete**
* Starting context:
   * You are in the container and have run `app`
* Testing instructions:
   * Enter `test.pml` at the prompt

### PML File Loading

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection feature step
* Testing Instructions:
   * By doing the previous step, the PML file will automatically be loaded

### Select Specific OWL Ontology

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection feature step
* Testing Instructions:
   * Enter `test.owl` at the prompt

### Load Selected Ontology

* Status: **Complete**
* Starting context:
   * You have completed the Select Specific OWL Ontology feature step
* Testing Instructions:
   * By doing the previous step, the OWL Ontology will automatically be loaded

### On Screen PML Reporting

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO

### PML Log File Generation

* Status: **OIn Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO

### PML Error Warning and Highlighting

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO

### On Screen Dinto Reporting

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO

### Dinto Log File Generation

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO

### Dinto Error Warning and Highlighting

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO

## Tearing down the Docker container

1. Run `docker-compose stop` to completely tear down the Docker container
