#!/usr/bin/env bash

# load the mnemonics on the .env file used for tests
source ./e2e/.env

rm -r ~/.rpsd || true
BIN=$(which rpsd)
CHAIN_ID="lb-rps-1"
GEN_FILE="$HOME"/.rpsd/config/genesis.json

# configure rpsd
$BIN config set client chain-id $CHAIN_ID
$BIN config set client keyring-backend test
# setup keys using the mnemoics on the .env file
echo "$MNEMONIC_ALICE" | $BIN keys add alice --recover --keyring-backend test
echo "$MNEMONIC_BOB" | $BIN keys add bob --recover --keyring-backend test
# initialize the chain directory
$BIN init test --chain-id $CHAIN_ID --default-denom rps
# update genesis
$BIN genesis add-genesis-account alice 100000000000rps --keyring-backend test
$BIN genesis add-genesis-account bob 100000000000rps --keyring-backend test
# update governance params (voting period) for local testing
sed -i.bak 's/"voting_period": "172800s"/"voting_period": "10s"/g' $GEN_FILE 
sed -i.bak 's/"expedited_voting_period": "86400s"/"expedited_voting_period": "5s"/g' $GEN_FILE 
# create default validator
$BIN genesis gentx alice 1000000rps --chain-id $CHAIN_ID
$BIN genesis collect-gentxs
