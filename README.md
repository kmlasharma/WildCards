# CS4098 Group Project
## Team Name: Wildcards

If you are cloning this repo for the first time, after cloning the repo:

1. Run the steps in `Installation`.
2. Run the steps in `Setting Up The Project`.
3. Run the steps in `Running The Project`.

If you have already cloned the repo, and would like the most up to date version:

1. Run the steps in `Setting Up The Project`.
2. Run the steps in `Running The Project`.


## Installation:

Docker is all you need for the project. To install Docker:

1. Run `make install` (this will ensure that all docker components needed (Docker, Docker Engine and Docker Compose) are installed. Accept any prompts during installation. When the command completes you should see the following:

![](http://i.imgur.com/OaLaivJ.png)

2. Restart your machine. This will apply the changes in the above steps properly and will allow you to run the following steps. If you do not restart your machine you will run into issues. You need to restart your machine to apply user group changes.


## Setting up the Project:

1. To ensure your computer is setup to run the project, run one of the following:
 * `make setup` if you are outside of the Trinity College network.

![](http://i.imgur.com/DUNCMhy.png)

 * `make setuptrinity` if you are inside of the Trinity College network.

2. Run `docker-compose build --no-cache` to build the project (the no-cache flag ensures a clean build).

![](http://i.imgur.com/GqkHsLI.png)


## Running the project:

1. Run `docker-compose run project` to enter the Docker container.

![](http://i.imgur.com/5ittEkl.png)

2. Run `app` to start the application.

![](http://i.imgur.com/HFsNKtK.png)


## Features:

It is assumed that you have followed the instructions above before testing the features listed below. Each feature will clarify it's own starting context.

### ✅ PML File Selection

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing instructions:
   * Hit enter at the prompt to select the default PML file

![](http://i.imgur.com/HFsNKtK.png)

### ✅ PML File Loading

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection feature step.
* Testing Instructions:
   * By doing the previous step, the PML file will automatically be loaded.
   
### ✅ Running PML Analysis 
* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and PML File Loading steps.
* Testing Instructions:
   * PML Analysis will automatically be run, and will output the process name along with all drugs in the process.

![](http://i.imgur.com/HWhIMMi.png)

### ✅ On Screen PML Reporting

* Status: **Complete**
* Starting context:
   * You have completed the Running PML Analysis step.
* Testing Instructions:
   * Any issues with the PML file will be reported on screen automatically at this point.

![](http://i.imgur.com/HWhIMMi.png)

### ✅ PML Log File Generation

* Status: **Complete**
* Starting context:
   * You have completed the On Screen PML Reporting step.
* Testing Instructions:
   * The PML reporting will also be logged to the `/log/output.log` file.
   
### ✅ Identify drugs in PML

* Status: **Complete**
* Starting context:
   * You have completed the PML Log File Generation step.
* Testing Instructions:
   * By doing the previous step, the drugs in the PML file will automatically be identified and printed

### ✅ PML Error and Warning Highlights

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project` and have run `app`
* Testing Instructions:
   * Enter 'error.pml' at the prompt to select the use a sample PML file with syntax errors.
   * If there is an error in the OWL file, the details will be logged to the `log/error.log` file, and an error message will show up on screen automatically.

### ✅ Select Specific OWL Ontology

* Status: **Complete**
* Starting context:
   * You have completed the PML Error and Warning Highlights step.
* Testing Instructions:
   * Hit enter at the prompt to select the default OWL file.

![](http://i.imgur.com/BBfmUdO.png)

### ✅ Load Selected Ontology

* Status: **Complete**
* Starting context:
   * You have completed the Select Specific OWL Ontology step.
* Testing Instructions:
   * By doing the previous step, the OWL Ontology will automatically be loaded.
   
### ✅ On Screen Dinto Reporting

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology steps.
* Testing Instructions:
   * Details of the OWL file will be reported on screen automatically at this point.

![](http://i.imgur.com/HWhIMMi.png)

### ✅ Dinto Log File Generation

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology steps.
* Testing Instructions:
   * The OWL reporting will also be logged to the `/log/output.log` file.

### ✅ Dinto Error Warning and Highlighting

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection step.
* Testing Instructions:
   * Enter `errortest.owl` at the prompt for the OWL Ontology file. `errortest.owl` is an incorrectly structured OWL Ontology file which will cause errors when parsed.
   * If there is an error in the OWL file, the details will be logged to the `log/error.log` file, and an error message will show up on screen automatically.


## Tearing down the Docker container

1. Make sure you have exited the container by typing `exit` ("Your prompt will change from `root@<some code>:~#` to your normal prompt).
2. Run `docker-compose stop` to completely tear down the Docker container (run this command from outside of the container).


## Automated Testing Instructions

When inside the container, the automated tests can be run to test each of the above features automatically. 

1. cd to `/root/tests` 
2. Run `nosetests`

