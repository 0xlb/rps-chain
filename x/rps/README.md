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
- expirationHeight (uint64)

### Param

- ttl (uint64): life of a game in blocks count

## Msg service

### MsgCreateGame

- creator (address - signer)
- oponent (address)
- rounds (uint)

### MsgMakeMove

- player (address - signer)
- gameIndex (uint)
- move (string - hash): SHA256 of the move and a salt (random string)

### MsgRevealMove

- player (address - signer)
- gameIndex (uint)
- revealedMove (string - Rock, Paper, Scissors)
- salt

hashedMove == sha256(revealedMove+salt)


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
- move (string - hash)

### EventMakeMove

- gameNumber (uint)
- player (address)
- revealedMove (string - Rock, Paper & Scissors)