#!/bin/bash


source ./hack/VENV/bin/activate

echo "generating pdf from md...."

mdpdf -o ./docs/nkia-api/api.pdf ./docs/nkia-api/index.md 2>/dev/null

mdpdf -o ./docs/nkia-server/server.pdf ./docs/nkia-server/index.md 2>/dev/null

mdpdf -o ./apix.d/apix.pdf ./apix.d/README.md

rm mdpdf.log

echo "done"
