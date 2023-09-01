#!/bin/bash

source ./hack/VENV/bin/activate

PWD=$(pwd)

python3 $PWD/staging_downstream/repo-arg.py $PWD

rm -rf $PWD/staging_downstream/nkia-*