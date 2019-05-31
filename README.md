# adhocio

A CLI tool for running adhoc pipelines on Jenkins.

Currently not very extensible beyond the current use case but just needs some TLC :)

## Installation

To download the package:

    $ go get -u github.com/eimlav/adhocio

The following will install a compiled binary:

    $ go install -i github.com/eimlav/adhocio

## Usage

### Authentication

You will need to export credentials to access your Jenkins instance:

    $ export JENKINS_USERNAME=<your-username>

    $ export JENKINS_API_TOKEN=<your-api-token>

To get an API token, go to the Jenkins dashboard. At the top right corner click the arrow drop-down beside your username then select the `Configure` option. This will direct you to a page where you can generate an token in the section titled `API Token`. Ensure you save this token somewhere as you will not be able to reveal it after you have created it.

### Commands

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
