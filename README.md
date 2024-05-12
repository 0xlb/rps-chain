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

# Lesson 2

In this lesson we will include all our module's boilerplate code needed
to be used in a Cosmos-SDK chain:

- `/types` directory contains all the module's types and stateless functions:
    - Protobuf generated files that includes the store objects, genesis state, 
    messages, events and `Query` and `Msg` services.
    - `keys.go`: includes the module's store keys and module name.
    - `codec.go`: includes the module's store keys and module name.
    - `params.go`: includes functions to get the module's default params and params validation.
- `/keeper` directory contains the keeper type definition and methods to read/write the module's store/s:
    - `keeper.go`: includes the keeper type definition and function to instanciate it.
    - `genesis.go`: includes keeper's functions related to genesis (`InitGenesis` & `ExportGenesis`).
- `module.go`: contains the module type that complies with the interfaces required
to allow the correct wiring of the custom module (`AppModuleBasic` & `AppModule`).
- `depinject.go`: includes the necessary functions of the custom module to allow it to be wired up
using dependency injection.
- `autocli.go`: includes the `AutoCLIOptions` function of the custom module to register its corresponding
CLI commands in the app.

Additionally, we'll wire up our custom `x/rps` module with our app.
