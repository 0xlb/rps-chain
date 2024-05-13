#!/usr/bin/env bash

rm -r ~/.rpsd || true
BIN=$(which rpsd)
CHAIN_ID="rps-1"
GEN_FILE="$HOME"/.rpsd/config/genesis.json

# configure rpsd
$BIN config set client chain-id $CHAIN_ID
$BIN config set client keyring-backend test
$BIN keys add alice
$BIN keys add bob
$BIN init test --chain-id $CHAIN_ID --default-denom rps
# update genesis
$BIN genesis add-genesis-account alice 10000000rps --keyring-backend test
$BIN genesis add-genesis-account bob 1000rps --keyring-backend test
# update governance params (voting period) for local testing
sed -i.bak 's/"voting_period": "172800s"/"voting_period": "10s"/g' $GEN_FILE 
sed -i.bak 's/"expedited_voting_period": "86400s"/"expedited_voting_period": "5s"/g' $GEN_FILE 
# create default validator
$BIN genesis gentx alice 1000000rps --chain-id $CHAIN_ID
$BIN genesis collect-gentxs
