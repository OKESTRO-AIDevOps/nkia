#!/bin/bash

if [ ! -d "VENV" ]; then
  echo "VENV does not exist."
  exit -1
fi

source VENV/bin/activate


mdpdf -o ../../doc/nkia-api/api.pdf ../../doc/nkia-api/index.md 


mdpdf -o ../../doc/nkia-server/server.pdf ../../doc/nkia-server/index.md 


rm mdpdf.log