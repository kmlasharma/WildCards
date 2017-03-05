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

1. Run `make install` (this will ensure that all docker components needed (Docker, Docker Engine and Docker Compose) are installed. Accept any prompts during installation. When the command completes you should see the following:

![]('./res/images/make_install_success.png')

2. Restart your machine. This will apply the changes in the above steps properly and will allow you to run the following steps. If you do not restart your machine you will run into issues. You need to restart your machine to apply user group changes.


## Setting up the Project:

1. To ensure your computer is setup to run the project, run one of the following:
 * `make setup` if you are outside of the Trinity College network.

![]('./res/images/make_setup_success.png')

 * `make setuptrinity` if you are inside of the Trinity College network.

2. Run `docker-compose build --no-cache` to build the project (the no-cache flag ensures a clean build).

![]('./res/images/docker_compose_build_success.png')


## Running the project:

1. Run `docker-compose run project` to enter the Docker container.

![]('./res/images/docer_compose_run_success.png')

2. Run `app` to start the application.

![]('./res/images/app_success.png')

## Features:

It is assumed that you have followed the instructions above before testing the features listed below. Each feature will clarify it's own starting context.

### ✅ PML File Selection

* Status: **Complete**
* Starting context:
   * You are in the container having run `docker-compose run project`
* Testing instructions:
   * Hit enter at the prompt

![]('./res/images/app_prompt_pml_file_selection.png')

### ✅ PML File Loading

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection feature step
* Testing Instructions:
   * By doing the previous step, the PML file will automatically be loaded 

### ✅ Select Specific OWL Ontology

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection feature step
* Testing Instructions:
   * Hit enter at the prompt

![]('./res/images/app_prompt_owl_file_selection.png')

### ✅ Load Selected Ontology

* Status: **Complete**
* Starting context:
   * You have completed the Select Specific OWL Ontology feature step
* Testing Instructions:
   * By doing the previous step, the OWL Ontology will automatically be loaded

### ✅ Running PML Analysis 
* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology
* Testing Instructions:
   * PML Analysis will automatically be run, and will output the process name along with all drugs in the process.

![]('./res/images/app_analysis_complete.png)

### ✅ On Screen PML Reporting

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology
* Testing Instructions:
   * Any issues with the PML file will be reported on screen automatically at this point.

![]('./res/images/app_analysis_complete.png)

### ✅ PML Log File Generation

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology
* Testing Instructions:
   * The PML reporting will also be logged to the log/output.log file.

### ⏳ PML Error Warning and Highlighting

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 4

### ✅ On Screen Dinto Reporting

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology
* Testing Instructions:
   * Any issues with the OWL file will be reported on screen automatically at this point.

![]('./res/images/app_analysis_complete.png)

### ✅ Dinto Log File Generation

* Status: **Complete**
* Starting context:
   * You have completed the PML File Selection and Select Specific OWL Ontology
* Testing Instructions:
   * The OWL reporting will also be logged to the log/output.log file.

### ⏳ Dinto Error Warning and Highlighting

* Status: **In Progress**
* Starting context:
   * TODO
* Testing Instructions:
   * TODO: Planned for Iteration 4


## Tearing down the Docker container

1. Run `docker-compose stop` to completely tear down the Docker container

![]('./res/images/docker_compose_stop_success.png')


