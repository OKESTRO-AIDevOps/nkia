#!/bin/bash

sudo kind create cluster --name kindcluster1 --config ./kindcluster1.yaml --image=kindest/node:v1.27.2

sudo kind create cluster --name kindcluster2 --config ./kindcluster2.yaml --image=kindest/node:v1.27.2