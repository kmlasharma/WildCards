# CS4098 Group Project
## Team Name: Wildcards

## Installation:
### Ubuntu

All that is required to run the project is Docker:

1. Run `./Scripts/install_docker_stack.sh`
   * This installs Docker, Docker Engine and Docker Compose

## Running:
### Ubuntu

Once the Docker stack has been installed (see above), running the project is simple.

1. Run `docker-compose up --build`
   * This builds the Docker container, installs the project dependencies in the container and set up the project
2. Run `docker-compose run groupproject` to enter the Docker container
3. Rum `peos/pml/check/pmlcheck peos/xpml/test.pml` to test PEOS on a sample pml file
   * Note that `test.pml` does not pass pmlcheck and it should not
   * `test.pml` should be used to verify that the pmlcheck tool works
4. Create your own `.pml` file and run pmlcheck on the file to run it through PEOS
5. Run `exit` to get out of the Docker container
6. Run `docker-compose stop` to completely tear down the Docker container

