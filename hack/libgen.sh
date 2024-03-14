#!/bin/bash

rm -r ../lib

rm -r ../lib.tgz

mkdir -p ../lib

/bin/cp -Rf libfactory/base ../lib/

/bin/cp -Rf libfactory/bin  ../lib/
