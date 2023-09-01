#!/bin/bash


if [ -d ./hack/VENV]
then 
    python3 -m venv ./hack/VENV
    source ./hack/VENV/bin/activate
    pip3 install -r ./hack/requirments.txt
fi

echo "hack environment initiated"


