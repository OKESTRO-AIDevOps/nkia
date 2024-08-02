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

# clone this repository
# running below will set all the development requirements

cd hack/dev

./dep.sh $(whoami)


# on terminal 1

make orch.io

# use below if you want orch.io in a container 
# make orch.io-up

make build

cd orch.io/osock

./osock

# on target computer
# also clone this repository

cd hack/dev

./dep.sh $(whoami)

# on terminal 2

cd nokubectl

./nokubectl $COMMANDS_AND_FLAGS 

```

## Examples

```shell

# check connection info

./nokubectl --as admin orch conncheck

# create cluster

./nokubectl --as admin orch add cl --clusterid test-cs

# install cluster

./nokubectl --as admin orch install cl --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu --localip 192.168.50.94 --osnm ubuntu --cv 1.27 --updatetoken 13e24636f0e94334fbbaa25d24113aa9

# get log for installing cluter

./nokubectl --as admin orch install cl log --clusterid test-cs --targetip 192.168.50.94:22 --targetid ubuntu --targetpw ubuntu

```