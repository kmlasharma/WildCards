# CS4098 Group Project
## Team Name: Wildcards

## Installation:

All that is required to run the project is Docker:

1. Run `./Scripts/install_docker_stack.sh`
   * This installs Docker, Docker Engine and Docker Compose if it is not already installed.

## Building Project:

Once the Docker stack has been installed (see above), building the project is simple.

### Ubuntu

1. Run the following commands as root - `sudo su` to change to root user
2. Run `sudo usermod -aG root && sudo service docker start`
3. Run `docker-compose up --build`

### Mac

1. TODO

## Running PEOS:

1. Run `docker-compose run peos` to enter the Docker container for PEOS
2. Run `peos/pml/check/pmlcheck peos/xpml/test.pml` to test PEOS on a sample pml file
   * Note that `test.pml` does not pass pmlcheck and it should not
   * `test.pml` should be used to verify that the pmlcheck tool works
3. Create your own `.pml` file and run pmlcheck on the file to run it through PEOS
4. Run `exit` to get out of the Docker container

## Running DINTO:

1. Run `docker-compose run dinto` to run our DINTO tests found in `./dinto/main_test.go`
2. Run `dinto` to run the project which will output analysis on the dinto file.

## Tearing down the Docker container

1. Run `docker-compose stop` to completely tear down the Docker container
