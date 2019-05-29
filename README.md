# adhocio

A CLI tool for running adhoc pipelines on Jenkins.

Currently not very extensible beyond the current use case but just needs some TLC :)

## Installation

The following will install a compiled binary:

    $ go install -i github.com/eimlav/adhocio

## Usage

```
config      Display config 
enable      Enable all jobs in a pipeline
help        Help about any command
run         Run a pipeline
version     Print current version of adhocio
```

To enable an experimental adhoc pipeline:

    $ adhocio enable motd -x

To run an adhoc pipeline with a specific SHA and GitHub user specified:

    $ adhocio run exec --ref cca6249c6a47494159dec8278032b4932ca1e836 --user eimlav

## Configuration

By default, configuration will be retrieved from `$HOME/.adhicio.yaml`. However you can specify a custom config file with the `-c` flag.
The file should be configured as follows:

```
---
jenkins_domain: 'my-jenkins-master.awesomecompany.com'
adhoc_prefix: 'adhoc_awesomecompany'
experimental_prefix: 'experimental_awesomecompany'
jobs:
  - 'init-manual_adhoc'
  - 'prep-smoke_adhoc'
  - 'smoke_adhoc'
```

These various configuration parameters are used to build the target URL used for making requests to the Jenkins server using the following scheme:

`https://<jenkins_domain>/job/<prefix>-<repo_name>_<job>/`

The first job you specify in your job list is the job used to trigger an adhoc build. As such, ensure that you have the correct job here.

## TODO

- [X] CLI interface
- [X] Run against normal adhoc and experimental
- [X] Supply GitHub repo, username and SHA
    - [X] Use default values when username or SHA not specified
- [X] Enable all jobs in selected pipeline
- [ ] Estimate total run time from historical data
- [~] Configuration file for default values such as GitHub username, Jenkins domain, adhoc paths etc
- [ ] Notifications on job completion (?)

```
X = Done
~ = In progress
```