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
git clone https://github.com/0xlb/rps-chain
cd rps-chain
make install # install the rpsd binary
make init # initialize the chain
rpsd start # start the chain
```

## Useful links

- [Cosmos-SDK Documentation](https://docs.cosmos.network/)
- [Mini - A minimal Cosmos-SDK chain](https://github.com/cosmosregistry/chain-minimal)
- [Cosmos-SDK module template](https://github.com/cosmosregistry/example)
