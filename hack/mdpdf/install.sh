#!/bin/bash


rm -r VENV

sudo apt-get update

sudo apt-get install python3-pip python3-venv

python3 -m venv VENV

source VENV/bin/activate

pip3 install mdpdf

