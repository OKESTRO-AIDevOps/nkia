# npia-server

## Table of Contents

- [Description](#description)
- [Project Overview](#project-overview)
- [Details and Examples](#details-and-examples)
- [Security](security.md)
- [Scenario](scenario.md)

## Description

This project aims not only to demonstrate a proper use case for [npia-api](https://github.com/OKESTRO-AIDevOps/npia-api) but also to\
offer a production-ready implementation of npia-api compliant system (and as usual for me, it doesn't reach the level at this point ʘ‿ʘ).

This repository holds two examplary (and hopefully someday production-grade) systems that implement somewhat distinct \
interfaces to interact with npia-api. Each system is as follows.

1. Http server that implements KCXD-STTC Protocol (you think it's a bogus word? Well... it was until now! See [Security](security.md) section to\
   know what on earth that means) to handle secure http query from a compliant client, ex) [npia-go-client](https://github.com/seantywork/npia-go-client) 


2. Web socket hub that implements KCXD-MTSC Protocol (again, [Security](security.md)) to handle secure web socket query from a compliant client\
   ex) the code in src/sock directory implements this, ex2) also, orchestrator/ofront will implement this protocol in future release, but for now,\
   it relies on oauth2 based https connection for security. This means, oauth2 user must (**please!**) configure properly the external \
   reverse proxy (nginx, per se) to handle https:// and wss://.



## Project Overview

The blueprint for the repository is as follows.

![npia-server]()

The tree structure for the repository is as follows

```bash
├── debug_bin
├── debug_build_run
├── debug_cleanup
├── docs
├── go.mod
├── go.sum
├── LICENSE
├── orchestrator
│   ├── debug_amalgamate_bin
│   ├── debug_amalgamate_config
│   ├── docker-compose.yaml
│   ├── odb
│   ├── odebug_build_run
│   ├── odebug_cleanup
│   ├── ofront
│   │   ├── config.json
│   │   ├── ocontroller
│   │   ├── omodels
│   │   ├── omodules
│   │   ├── orouter
│   │   └── oview
│   └── osock
├── src
│   ├── controller
│   ├── modules
│   ├── router
│   └── sock
├── test
│   └── kindcluster
└── var

```

Let me guide you through the each entrypoint briefly.

- [debug_bin](#debug_bin)\
  Stores a kubeconfig file that holds information about two clusters for testing npia-server.

- debug_build_run\
  Is a Bash script for compiling and setting up the executable and base environment\ 
  for testing npia-server.  

- debug_cleanup\
  Is a Bash script for cleaning up the artifacts and environment used for testing\
  npia-server.

- docs\
  Is what it says it is.

- go.mod, go.sum\
  Define the directly and indirectly required modules for functioning apis.

- LICENSE\
  Is what it says it is.

- [orchestrator](#orchestrator)\
  Stores a docker-compose manifest file for building and running components that make up\
  orchestrator system while also has Bash scripts for testing it\
  **Caveats: if you already have a running cluster and corresponding kubeconfig file,\
  DO NOT run the script "odebug_build_run" without backing up your original\
  kubeconfig, because the script will obliterate that one**


- [orchestrator/debug_amalgamate_bin](#orchestrator-debug_amalgamate_bin)\
  Stores "amalgamation tools" used when testing the orchestrator system, which means that\
  the tools handle automatically baseline components needed for having a functioning\
  orchestrator system, including merging two test kubeconfig files into one and then\
  generating encryption key and putting encrypted merged file into the db  

- [orchestrator/debug_amalgamate_config](#orchestrator-debug_amalgamate_config)\
  Stores the two test kubeconfig files used for testing the orchestrator system


- [orchestrator/odb](#orchestrator-odb)\
  Stores image and container information for running orchestrator database.


- [orchestrator/ofront](#orchestrator-ofront)\
  Holds the front part (meaning user-facing on the browser) of the orchestrator system, including\
  configuration file needed for Google oauth.


- [orchestrator/ofront/ocontroller](#orchestrator-ofront-ocontroller)\
  Stores controller logic for handling user request regarding access and oauth authentication.


- [orchestrator/ofront/omodels](#orchestrator-ofront-omodels)\
  Stores database querying logic for handling user request regarding access and \
  oauth authentication and also storing orchestrator socket request key. 


- [orchestrator/ofront/omodules](#orchestrator-ofront-omodules)\
  Stores orchestrator gadgets that include primarily oauth modules and configuration.


- [orchestrator/ofront/orouter](#orchestrator-ofront-orouter)\
  Stores access paths information for users to access the orchestrator.
      

- [orchestrator/ofront/oview](#orchestrator-ofront-oview)\
  Stores html and template assets to render on the user's browser.


- [orchestrator/osock](#orchestrator-osock)\
  Stores socket connection endpoints and handling logic for both oauth-ed front user and \
  challenge-passed server socket client and the hub channels for both sides to talk to each\
  other. 

- [src](#src)\
  Stores entry and initiation logic for both npia-server STTC host and npia-server MTSC client.

- [src/controller](#src-controller)\
  Stores various data structures and logic for interacting with npia-server that inlcude\
  npia-api wrapper, challenge protocol handler, network communication data structures, and\
  npia-multi-mode handler so and so.
  The data structure and logic defined and implemented in here are used across,\
  but not limited to, this project. Those are exported to other projects such as\
  npia-go-client.

- [src/modules](#src-modules)\
  Stores various data structures and logic primarily related to asymmetric/symmetric\
  encryption/decryption algorithms and certificate authentication methods that are all\
  building blocks of KCXD Challenge Protocol. 

- [src/router](#src-router)\
  Stores endpoints where clients programs can access the server, get authenticated\
  through challenges and make query to npia-api 

- [src/sock](#src-sock)\
  Stores functions that make up a socket client for connecting to the orchestrator hub, handling\
  challenge protocol, and sending to/receiving from it using crypto algorithms.

- test\
  Stores testing materials

- [test/kindcluster](#test-kindcluster)\
  Stores scripts and config files for setting up working kind Kubernetes cluster for testing purpose 

- var\
  Has various this and that artifacts when developing the server




## Details and Examples

This section dives into the details of each entry point.

However, it doesn’t go so deeper that you don’t have to look at the source code to understand
how everything works in conjunction.

For even more details, refer to specific comments associated with a code block, or better,
you could just run it yourself.

### debug_bin

- config\
  Is a kubeconfig file that holds the information of two kind clusters that doesn't exist\
  on earth anymore.


### orchestrator

- docker-compose.yaml\
  Is a docker-compose manifest file that defines the build directory for each orchestrator\
  component and the way they run.

- odebug_build_run\
  Is a Bash script that consists of command lines to set up a test environment of orchestrator.\
  It merges two configs in the orchestrator/debug_amalgamate_bin into one, generates an \
  encryption key, encrypts the merged config file and puts the file into the orchestrator db.

- odebug_cleanup\
  Removes all the environments created using orchestrator/odebug_build_run
  

### orchestrator debug_amalgamate_bin

- amalgamate\
  Is a Bash script that is used by orchestrator/odebug_build_run and handles the job of\
  actually merging two kubeconfig files.

- amalgamate.go\
  Is a program that generates AES GCM key and consumes merged kubeconfig file to output \
  both the symmetric key and finally the encrypted kubeconfig file. 


### orchestrator debug_amalgamate_config


- config1\
  the first kind Kubernetes cluster config for testing out npia-server.

- config2\
  the second kind Kubernetes cluster config for testing out npia-server.


### orchestrator odb

- Dockerfile\
  Defines how the MySQL database should be built as a container image, picking up the\
  encrypted kubeconfig file along the way 

- init.sql\
  Defines how the MySQL database should be initiated inside a created container, inserting \
  pre-defined information along with the encrypted kubeconfig file into corresponding columns.


### orchestrator ofront

- Dockerfile\
  Defines how a golang server environment should be built and initiated inside a container \
  image.

- config.json\
  Holds information for later use when user is trying to authenticate through Google OAuth \
  channel.\
  It consists of two fields, GOOGLE_OAUTH_CLIENT_ID which is client id retrieved when \
  registering the oauth process on cloud.google.com and GOOGLE_OAUTH_CLIENT_SECRET which is \
  also retrievable at the time you register the process on the platform along the client id.
  
- ofront.go\
  Is a main entry for creating and running the Go Gin server that serves the front user.

### orchestrator ofront ocontroller

- ocontroller.go\
  Has functions for serving index html file and /orchestrator template page, redirecting\
  to /orchstrate if session is present and authenticated, and Google oauth initiation along\
  with callback  


### orchestrator ofront omodels

- omodels.go\
  Has functions to handle database query and check access authenticity, register front-user\
  session id with socket request key if the user has successfully passed oauth process. 



### orchestrator ofront omodules

- config.go\
  For now, has a function to read the orchestrator/ofront/config.json and a data structure to\
  store the data 

- ooauth2.go\
  Has functions to set session cookie and get user data from Google after the front user got \
  redirected to authenticate with the Google server.\
  Also it has data structures to hold the oauth2 configuration and Google oauth result to\
  check if the oauth process has been successful.


### orchestrator ofront orouter

- orouter.go\
  Has paths for front user to access which are main landing page, and another page \
  where orchestrator call is made, and two others for conducting oauth2 authentication. 


### orchestrator ofront oview

- index.html\
  Is what it says it is.

- orchestrator.tmpl\
  Is where, after the authentication, the front user makes call to the orchestrator hub to \
  retrieve api processing result from the npia-server sock agent.


### orchestrator osock

- Dockerfile

- osock_front.go

- osock_modules.go

- osock_server.go

- osock.go


### src

- main.go


### src controller

- api.go

- auth.go

- check.go

- definition.go

- multimode.go


### src modules

- auth.go

- definition.go

- utils.go

### src router

- router.go

### src sock

- module.go

- sock.go


### test kindcluster

- kindcluster.sh

- kindcluster1.yaml

- kindcluster2.yaml