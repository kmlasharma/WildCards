# CS4098 Group Project
## Team Name: Wildcards

## Installation:

All that is required to run the project is Docker:

1. Running `make install` will ensure that all docker components needed (Docker, Docker Engine and Docker Compose) are installed.

## Setting up the Project:

1. Run `make setup` to ensure your computer is setup to run the project.

2.
If running on Linux
On Linux, this will start the docker service. 
If you are on Linux, at this point you need to physically restart your machine. This will apply the changes in the above steps properly and will allow you to run the following steps with ease.

If Running on Mac
Mac does not have the kernel features required to run Docker containers natively, so we use virtualbox to provision a Docker machine VM. This will all be handled using `make setup` (Including using a VM if it's already provisioned).

3. Finally, whether you are on Mac or Linux, run `docker-compose up --build` to build the project.

## Running PEOS:

1. Run `docker-compose run peos` to enter the Docker container for PEOS
2. Run `peos/pml/check/pmlcheck peos/xpml/test.pml` to test PEOS on a sample pml file
   * Note that `test.pml` does not pass pmlcheck and it should not
   * `test.pml` should be used to verify that the pmlcheck tool works
3. Create your own `.pml` file and run pmlcheck on the file to run it through PEOS
4. Run `exit` to get out of the Docker container

## Running DINTO:

1. Run `docker-compose run dinto-tests` to run our DINTO tests found in `./dinto/main_test.go`
2. Run `docker-compose run dinto` to run the project which will output analysis on the dinto file.

## Tearing down the Docker container

1. Run `docker-compose stop` to completely tear down the Docker container
