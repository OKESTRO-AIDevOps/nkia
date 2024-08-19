# nkia-doc

[Search](#search)\
[Description](#description)\
[Where to start](#where-to-start)\
[Core components](#core-components)\
[Peripheral components](#peripheral-components)

## Search

[https://okestro-aidevops.github.io/nkia-doc/](https://okestro-aidevops.github.io/nkia-doc/)



## Description

This document holds the layout of the NKIA project

NKIA project aims to serve as an engine that can provide\
a set of API, either through command line or web browser,\
that can help its users 

1. more easily manage container lifecycle across multiple Kubernetes clusters
2. more easily build and deploy their cloud-native projects across multiple Kubernetes cluster (primarily based on docker-compose.yaml)
3. more easily balance out between better observability and enhanced security 

It's up to the users to find out what good those ever do to them


## Where to start

If you really want to try out,

- start by reading **nokubectl** document below in the section **Core components**
- then **orch.io**
- then **nokubeadm**...
- in that following order to the end

But here is the overall picture of how everything works in conjunction


***(will be added soon...)***


## Core components

[nokubectl](nokubectl)

[orch.io](orch.io)

[nokubeadm](nokubeadm)

[nokubelet](nokubelet)

[pkg](pkg)


## Peripheral components

[infra](infra)

[hack](hack)

[doc](doc)


## Getting Started

Requirements:
- go 1.21+
- make
- docker

### dev

```shell

# assumption:
# you have a host computer to run orch.io and nokubectl
# you have a target kubernetes cluster


# clone this repository on host computer
# running below will set all the development requirements

cd hack/dev

./dep.sh $(whoami)


# on host computer terminal 1
# this will build orchestrator server

make orch.io

# on host computer terminal 1
# this will start orchestrator server

cd orch.io/osock

./osock

# on host computer terminal 2
# this will build nokubectl, nokubeadm, nokubelet

make build

# on host computer terminal 2
# look EXAMPLES for $COMMANDS_AND_FLAGS

cd nokubectl

./nokubectl $COMMANDS_AND_FLAGS 


# on target computer
# also clone this repository

cd hack/dev

./dep.sh $(whoami)

# on target computer
# this will build nokubectl, nokubeadm, nokubelet

make build-noctl

# now, run below on host computer to check if server is responsive

./nokubectl orch conncheck

# if okay, now we have to connect target computer's nokubelet to host computer server

# on host computer
# now, get one-time token by adding a cluster on host computer

./nokubectl orch add cl --clusterid test-cs

# assuming we have retrieved token 384996074c73704f703759cc2c7fb4f1
# on target computer

cd ./nokubelet

sudo ./nkletd io connect update --clusterid test-cs --updatetoken 384996074c73704f703759cc2c7fb4f1


# now on host computer

# see below examples



```

## Examples

```shell

# set target cluster (or manually put --to in every command)

./nokubectl set --to test-cs

# get nodes

./nokubectl resource nodes --ns -

```