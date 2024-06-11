# RPS - A Rock, Paper & Scissors Cosmos SDK app-chain

This repository contains a Cosmos SDK app chain implementation of the Rock, Paper & Scissors game.
It uses the least modules possible and is intended to be used as learning material.

`rpsd` uses the **v0.50.4** version of the [Cosmos-SDK](https://github.com/cosmos/cosmos-sdk).

## System requirements

- [Golang >= v1.21.0](https://go.dev/doc/install)

Make sure to add `$HOME/go/bin` to your path to easily execute
the installed go packages binaries

```
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

## How to use

In addition to learn how to build a chain thanks to `rpsd`, you can as well directly run `rpsd`.

### Installation

Install and run `rpsd`:

```sh
git clone https://github.com/0xlb/rpschain
cd rps-chain
make install # install the rpsd binary
make init # initialize the chain
rpsd start # start the chain
```

## Useful links

- [Cosmos-SDK Documentation](https://docs.cosmos.network/)
- [Mini - A minimal Cosmos-SDK chain](https://github.com/cosmosregistry/chain-minimal)
- [Cosmos-SDK module template](https://github.com/cosmosregistry/example)

-----------

## Unlocking the Interchain potential

Connect your chain to the Interchain ecosystem using the 
[go relayer](https://github.com/cosmos/relayer) (make sure to use a version >v2.5.0)

### Initialize the relayer's configuration directory

```
rly config init
```

To customize the memo when relaying packets:

```
rly config init --memo "My custom memo"
```

### Update the config.yaml file

Add the configuration for the two chains we want to connect.
In this example I'll connect our RPS chain to the Evmos testnet.

```yaml
chains:
    evmos-testnet:
        type: cosmos
        value:
            key-directory: /home/lb/.relayer/keys/evmos_9000-4
            key: rly1
            chain-id: evmos_9000-4
            rpc-addr: https://evmos-testnet-rpc.polkachu.com:443
            account-prefix: evmos
            keyring-backend: test
            gas-adjustment: 1.5
            gas-prices: 40000000000atevmos
            min-gas-amount: 1
            max-gas-amount: 0
            debug: false
            timeout: 20s
            block-timeout: ""
            output-format: json
            sign-mode: direct
            extra-codecs:
                - ethermint
            coin-type: 60
            signing-algorithm: ""
            broadcast-mode: batch
            min-loop-duration: 0s
            extension-options: []
            feegrants: null
      rps:
        type: cosmos
        value:
            key-directory: /home/lb/.relayer/keys/rps-1
            key: rly1
            chain-id: rps-1
            rpc-addr: http://localhost:26657
            account-prefix: rps
            keyring-backend: test
            gas-adjustment: 1.5
            gas-prices: 0rps
            min-gas-amount: 1
            max-gas-amount: 0
            debug: false
            timeout: 20s
            block-timeout: ""
            output-format: json
            sign-mode: direct
            extra-codecs: []
            coin-type: null
            signing-algorithm: ""
            broadcast-mode: batch
            min-loop-duration: 0s
            extension-options: []
            feegrants: null            
```

### Setup accounts for each of the chains

```
rly keys add evmos-testnet rly1
rly keys add rps rly1
```

### Fund accounts

Send some funds to the relayer's accounts.
For the evmos-testnet account, you can use [the faucet](https://faucet.evmos.dev/).
Then, check the balance of these using the following commands:

```
rly q balance evmos-testnet
rly q balance rps
```

### Create a new path for the desired chains

Use the chain ids to create the new path

```
rly paths new evmos_9000-4 rps-1 demo-ibc-path
```

### Connect the chains

Create the clients and open the connections using the command:

```
rly tx link demo-ibc-path
```

### Start relaying

```
rly start demo-ibc-path
```
