#!/bin/sh

go build
./osch doc
cd ../build/docs
mdbook build -d ../../docs
cd ../../osch
./osch rdf
