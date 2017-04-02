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

## Testing:

A suite of automated unit tests can be run to test all features of the app. This can be done by navigating to `/go/src/app` (the default starting point of the app), and running `tests`. The test files are primarly for PML and DDI features, so are found in `/go/src/app/pkgs/pml/pml_test.go` and `/go/src/app/pkgs/ddi/ddi_test.go`. The `tests` command will automatically include these tests. The output will look like below (this is a subset of the tests and there will be more output):

![](http://i.imgur.com/kmi1Pt5.png)

## Features:

It is assumed that you have followed the instructions above before testing the features listed below. Each feature will clarify it's own starting context. Unit and behavioural tests are also outlined for each feature.

# Release 2:
#### Release 2 features are outlined below

### ✅ Identify drugs in PML

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing instructions:
   * Hit enter at the prompt to select the default PML file, which is `test.pml`. 
   * This will display the drugs in that process. 
   * Files that have drugs in them include `test.pml` and `multi_drugs.pml`. They will be displayed as shown.
   * Files that have no drugs include `no_drugs.pml` and `no_subtasks.pml`. These files will be parsed and will show no drugs in them. 
   * Drugs are encoded in stringified JSON format within the script tag of an action. The results of this step can be changed by changing the script tags of actions in the pml files and rerunning this test.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`.
   
![](http://i.imgur.com/r5czJYf.png)

### ✅ Report Un-named PML Construct

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `missing_pml_construct.pml`. 
   * An error detailing an un-named PML construct will be returned, because the file starts with `process {`. 
   * To fix this and remove this error, change line 1 to `process process_name {`.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`.
  
![](http://i.imgur.com/iVQGo8m.png)
   
### ✅ Mock DDI Characterisation Data

* Status: **Complete**
* Testing Instructions:
   * This feature is not directly testable, but reviewing the ddi.csv will show our defined format for DDI characterisation Data. 
   * We require a CSV with 4 columns - Drug A, Drug B, Whether the interaction is adverse, and the duration of the interactoin.

   
### ✅ Lookup Drugs in Mock Data File

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
    * Hit enter at the prompt to select the default PML file. You can also use an alternative PML file.
    * Hit enter at the prompt to select the default DDI file. You can also use an alternative DDI file.
    * The program will automatically lookup the drugs from the PML file in the Mock data file.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/ddi/ddi_test.go`.
  
![](http://i.imgur.com/DuVHsbN.png)
   
### ✅ Identify DDIs

* Status: **Complete**
* Starting context:
   * You have completed the Lookup Drugs in Mock Data File step.
* Testing Instructions:
   * After looking up the drugs from the PML file in the Mock data file, the program will return the valid interactions for this set of drugs. 
   * For example, using the default files (test.pml and ddi.csv), coke and 7up will be an interaction because both are in the PML file and there is an entry in the csv file for these drugs. 
   * You can adjust the Interactions returned by changing the drugs in the DDI file or the PML file.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/ddi/ddi_test.go`.
  
![](http://i.imgur.com/DuVHsbN.png)

### ✅ Report PML Construct Name-Clash

* Status: **In Progress**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `sequence_clashes.pml`. 
   * The program will report that there is a PML Construct name clash for this PML file. 
   * This is because there is two sequences called 'Andy'. 
   * You can remove this error by changing one of the sequence names and rerunning this test, or recreate it by changing another pml to have clashing names within the same namespace.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`.
  
![](http://i.imgur.com/Advfand.png)

### ✅ Report Use Of Task Construct

* Status: **In Progress**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `subtasks.pml`. 
   * This feature has two tasks in it, so two tasks will be summarised on the screen.
   * If you add/remove a task from this file, or indeed any other file, and re run this test, you will see the results reflecting that.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`.
  
![](http://i.imgur.com/2CbrRvY.png)


### ⏳ Identify Sequential DDIs

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7


### ⏳ Identify Parallel DDIs

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7
   
### ⏳ Report Alternative Non-DDIs

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7

### ⏳ Report Repeated Alternative DDIs

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7


### ⏳ Specify Periodic Drug Use

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7
   
   
### ⏳ Specify a Delay

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7
   
   
### ⏳ Specify a Time Interval Offset

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7


### ⏳ Identify DDI Closest Approach

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7


### ⏳ Merging Clinical Pathways Written In PML

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7
   
   
### ⏳ PML-TX Save PML To File

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 7


## Tearing down the Docker container

1. Make sure you have exited the container by typing `exit` ("Your prompt will change from `root@<some code>:~#` to your normal prompt).
2. Run `docker-compose stop` to completely tear down the Docker container (run this command from outside of the container).


## Automated Testing Instructions

When inside the container, the automated tests can be run to test each of the above features automatically. 

1. cd to `/root/tests` 
2. Run `nosetests`
3. 


RELEASE 1 Features:

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

![](http://imgur.com/GUBevCH.png)

### ✅ On Screen PML Reporting

* Status: **Complete**
* Starting context:
   * You have completed the Running PML Analysis step.
* Testing Instructions:
   * Any issues with the PML file will be reported on screen automatically at this point.

![](http://imgur.com/GUBevCH.png)

### ✅ PML Log File Generation

* Status: **Complete**
* Starting context:
   * You have completed the On Screen PML Reporting step.
* Testing Instructions:
   * The PML reporting will also be logged to the `/root/log/output.log` file.
   
### ✅ Identify drugs in PML

* Status: **Complete**
* Starting context:
   * You have completed the PML Log File Generation step.
* Testing Instructions:
   * By doing the previous step, the drugs in the PML file will automatically be identified and printed
   
![](http://imgur.com/GUBevCH.png)

### ✅ PML Error and Warning Highlights

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project` and have run `app`
* Testing Instructions:
   * Enter `errortest.pml` at the prompt to select the use a sample PML file with syntax errors.
   * If there is an error in the PML file, the details will be logged to the `/root/log/output.log` file, and an error message will show up on screen automatically.
   
![](http://imgur.com/sbWeBlf.png)

### ✅ Select Specific OWL Ontology

* Status: **Complete**
* Starting context:
   * You have completed the PML Error and Warning Highlights step.
* Testing Instructions:
   * Hit enter at the prompt to select the default OWL file.

![](http://imgur.com/UPaZMbB.png)

### ✅ Load Selected Ontology

* Status: **Complete**
* Starting context:
   * You have completed the Select Specific OWL Ontology step.
* Testing Instructions:
   * By doing the previous step, the OWL Ontology will automatically be loaded.
   
![](http://imgur.com/UPaZMbB.png)

> **Note:** The loading bar displayed is not 100% accurate in forecasting the actual time to load the ontology. The ontology may load before the loading bar reaches 100%, or even a few seconds after the loading bar reaches 100%.
   
### ✅ On Screen Dinto Reporting

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology steps.
* Testing Instructions:
   * Details of the OWL file will be reported on screen automatically at this point.

![](http://imgur.com/UPaZMbB.png)

### ✅ Dinto Log File Generation

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology steps.
* Testing Instructions:
   * The OWL reporting will also be logged to the `/root/log/output.log` file.

### ✅ Dinto Error Warning and Highlighting

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection step.
* Testing Instructions:
   * Enter `errortest.owl` at the prompt for the OWL Ontology file. `errortest.owl` is an incorrectly structured OWL Ontology file which will cause errors when parsed.
   * If there is an error in the OWL file, the details will be logged to the `/root/log/output.log` file, and an error message will show up on screen automatically.

![](http://imgur.com/8w7FaNJ.png)
