#!/bin/bash


rm -r lib.tgz

curl -C - -L https://github.com/OKESTRO-AIDevOps/nkia/releases/download/latest/lib.tgz -o lib.tgz

tar -xzf lib.tgz

rm -r ./nokubeadm/lib

rm -r ./nokubectl/lib

rm -r ./nokubelet/lib

/bin/cp -R lib ./nokubeadm/

/bin/cp -R lib ./nokubectl/

/bin/cp -R lib ./nokubelet/

rm -r lib

rm -r lib.tgz