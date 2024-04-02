#!/bin/bash

if [ ! -d "VENV" ]; then
  echo "VENV does not exist."
  exit -1
fi

source VENV/bin/activate


mdpdf -o ../../doc/nokubectl/index.pdf ../../doc/nokubectl/index.md 

mdpdf -o ../../doc/orch.io/index.pdf ../../doc/orch.io/index.md 

mdpdf -o ../../doc/nokubeadm/index.pdf ../../doc/nokubeadm/index.md 

mdpdf -o ../../doc/nokubelet/index.pdf ../../doc/nokubelet/index.md 

mdpdf -o ../../doc/pkg/index.pdf ../../doc/pkg/index.md 

mdpdf -o ../../doc/infra/index.pdf ../../doc/infra/index.md 

mdpdf -o ../../doc/hack/index.pdf ../../doc/hack/index.md 

mdpdf -o ../../doc/doc/index.pdf ../../doc/doc/index.md 



rm mdpdf.log