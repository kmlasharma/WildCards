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
   * We require a CSV with 4 columns - Drug A, Drug B, Whether the interaction is adverse, and the duration of the interaction.
   * Here is the following format and example:
   
| Drug 1       | Drug 2         | DDI Type | Time | Unit|
|:------------:|:--------------:| --------:|-----:|----:|
| alcohol      | coke           | bad      | 1    | week|
| 7up          | pepsi          | good     | 3    | days|

   
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
   * After looking up the drugs from the PML file in the Mock data file, the program will return the valid interactions for this set of drugs after selecting 'Show All Interactions' in the menu by entering '1'.
   * For example, using the default files (`test.pml` and ddi.csv), coke and 7up will be an interaction because both are in the PML file and there is an entry in the csv file for these drugs. 
   * You can adjust the Interactions returned by changing the drugs in the DDI file or the PML file.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/ddi/ddi_test.go`.
  
![](http://i.imgur.com/DuVHsbN.png)

### ✅ Report PML Construct Name-Clash

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `sequence_clashes.pml`. 
   * The program will report that there is a PML Construct name clash for this PML file. 
   * This is because there is two sequences called 'Andy'. 
   * You can remove this error by changing one of the sequence names and rerunning this test, or recreate it by changing another pml to have clashing names within the same namespace.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`.
  
![](http://i.imgur.com/KLNmf40.png)

### ✅ Report Use Of Task Construct

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `subtasks.pml`. 
   * This file has two tasks in it, so two tasks will be summarised on the screen.
   * If you add/remove a task from this file, or indeed any other file, and re run this test, you will see the results reflecting that.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`.
  
![](http://i.imgur.com/6AgKC0V.png)


### ✅ Identify Sequential DDIs

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `sequential_ddi.pml`. 
   * When you are asked what operation you would like to complete, select 'Show Sequential DDIs' by entering '3'
   * This file has a sequential DDI which will be displayed. 
   * There are a number of other files, such as `test.pml` and `sequence_in_branch.pml` that also have sequential DDIs. 
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`.
  
  ![](http://i.imgur.com/hqVmgHg.png)

### ✅ Identify Parallel DDIs

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `parallel_ddi.pml`. 
   * When you are asked what operation you would like to complete, select 'Show Parallel DDIs' by entering '4'
   * An Parallel DDI will be displayed due to the fact that there is one in the pml file.
   * There are a number of other files, such as `sequence_in_branch.pml` that also have parallel DDIs. 
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`
  
   ![](http://i.imgur.com/VOzWqpg.png)
   
### ✅ Report Alternative Non-DDIs

* Status: **Complete**
* Starting context:
  * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `alternative_non_ddi.pml`.
   * When you are asked what operation you would like to complete, select 'Show Alternative Non DDIs' by entering '5'
   * An Alternative Non DDI will be displayed due to the fact that there is one in the pml file.
* Unit Testing:
   * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`

### ✅ Report Repeated Alternative DDIs

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `selection_in_iteration.pml`. 
   * When you are asked what operation you would like to complete, select 'Show Alternative Repeated DDIs' by entering '6'
   * An Alternative Non DDI will be displayed due to the fact that there is one in the pml file.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`
  
  ![](http://i.imgur.com/XPDfP9A.png)


### ✅ Specify Periodic Drug Use

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `periodic_use.pml`. 
   * This file has an iteration with a `loops` and `delay` constructs, along us to specify periodic drug use. 
   * It will be run every 3 days (due to a delay of 3 at the end) and run 5 times before finishing. 
   * Periodic drug use will be reported immediately.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`
   
   
###  ✅ Specify a Delay

* Status: **Complete**
* Starting context:
    * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `delays.pml`. 
   * The 'delay' construct will be correctly parsed, and the delays correcetly assigned.
   * Use of delays will be reported immediately.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`
   
### ✅ Specify a Time Interval Offset

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `time_interval_offset.pml`.
   * The 'wait' construct will be correctly parsed. It will wait to Monday due to the `wait{ "Monday" }` construct.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`


### ✅ Identify DDI Closest Approach

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter `closest_approach.pml`. 
   * When you are asked what operation you would like to complete, select 'Show all closest approaches' by entering '10'.
   * The DDI Closest Approach will be displayed for any interactions. In this case, there will be three, and one of them will have a closest approach of 'infinite' due to it being in a selection.
* Unit Testing:
  * This feature is tested in `/go/src/app/pkgs/pml/pml_test.go`


### ✅ Merging Clinical Pathways Written In PML

* Status: **Complete**
* Starting context:
    * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter any PML file (e.g `test.pml`). 
   * When you are asked what operation you would like to complete, select 'Merge PML Files' by entering '8'.
   * You will be asked to enter another PML file, which again, can be any PML file.
   * Once you hit enter, the processes will be merged using a branch construct.
   * An easy way to see this in action is to then save the pml file using the instructions below.
 
 
### ✅ PML-TX Save PML To File

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`, and have run `app`
* Testing Instructions:
   * When prompted to enter a pml file, enter any PML file (e.g `test.pml`). 
   * When you are asked what operation you would like to complete, select 'Save PML to File' by entering '7'.
   * You will then be asked what you would like the name the saved file. 
   * It will then be saved in the same directory as the other pml files.


## Tearing down the Docker container

1. Make sure you have exited the container by typing `exit` ("Your prompt will change from `root@<some code>:~#` to your normal prompt).
2. Run `docker-compose stop` to completely tear down the Docker container (run this command from outside of the container).


# RELEASE 1 Features:

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
