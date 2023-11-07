#!/bin/bash

set -euxo pipefail

#######

server_ip=$1



sudo helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/


sudo helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner --set nfs.server=$server_ip --set nfs.path=/npia-data


sudo kubectl apply -f ./default-storage-class.yaml


sudo helm repo add prometheus-community https://prometheus-community.github.io/helm-charts

sudo helm repo update  

sudo helm install -f ./default-kubeprom-custom-value.yaml kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 42.2.0
