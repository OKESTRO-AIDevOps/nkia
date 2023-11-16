#!/bin/bash

rm -r bin.tgz

curl -C - -L https://github.com/OKESTRO-AIDevOps/nkia/releases/download/latest/bin.tgz -o bin.tgz

tar -xzf bin.tgz

/bin/cp -Rf ./bin/nokubeadm/nokubeadm ./nokubeadm/nokubeadm

/bin/cp -Rf ./bin/nokubectl/nokubectl ./nokubectl/nokubectl

/bin/cp -Rf ./bin/nokubelet/nokubelet ./nokubelet/nokubelet

rm -r bin

rm -r bin.tgz