# `/x/rps` - Rock, Paper and Scissors Module

This directory contains the code for your chain custom modules.

## State objects

### Game

- gameNumber (uint)
- playerA (address)
- playerB (address)
- status (string)
- rounds (uint)
- playerAMoves ([]string)
- playerBMoves ([]string)
- score ([2]uint)

### Param



## Msg service

### MsgCreateGame

- creator (address - signer)
- oponent (address)
- rounds (uint)

### MsgMakeMove

- player (address - signer)
- gameIndex (uint)
- move (string - Rock, Paper, Scissors)

## Query service

### GetGame

- gameNumber (uint)

### GetParams

## Events

### EventCreateGame

- gameNumber (uint)
- playerA (address)
- playerB (address)

### EventEndGame

- gameNumber (uint)
- status (string)

### EventMakeMove

- gameNumber (uint)
- player (address)
- move (string)