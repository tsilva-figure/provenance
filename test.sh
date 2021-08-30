#!/bin/bash
CHAIN_HOME=./build/run/provenanced
PIO_PROJECT_HOME=~/code/provenance
#PIO_VERSION=v1.5.0
METADATA_DATA_FILE=./genesis_with_metadata.json

cd ${PIO_PROJECT_HOME}
#git checkout ${PIO_VERSION}
go mod vendor
make clean
make build
make run-config

# update genesis file with better config for testing
cat ${CHAIN_HOME}/config/genesis.json | jq ' .app_state.gov.voting_params.voting_period="20s" ' | tee ${CHAIN_HOME}/config/genesis.json
metadata_data=$(cat ${METADATA_DATA_FILE} | jq .app_state.metadata)
echo ${metadata_data}
cat ${CHAIN_HOME}/config/genesis.json | jq ' .app_state.metadata='"${metadata_data}"' ' > ${CHAIN_HOME}/config/genesis-tmp.json
mv ${CHAIN_HOME}/config/genesis-tmp.json ${CHAIN_HOME}/config/genesis.json

make run