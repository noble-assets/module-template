#!/bin/bash

# NOTE: this file must be executed from the root of the project.
cd ./proto || exit 1
buf generate --template buf.gen.gogo.yaml
buf generate --template buf.gen.pulsar.yaml
cd ..

cp -r template.dev/* ./
cp -r api/noble/template/* api/

rm -rf template.dev
rm -rf api/noble
rm -rf noble
