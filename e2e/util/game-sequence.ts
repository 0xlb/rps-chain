import { config } from "dotenv";
import _ from "../environment";
import {
  MsgMakeMove,
  MsgRevealMove,
} from "rps-chain-client/src/types/generated/lb/rps/v1/tx";
import { SALT, generateHash } from "./hash";

config();

const { ADDRESS_ALICE: alice, ADDRESS_BOB: bob } = process.env;

export interface GameRound {
  commitment: MsgMakeMove[];
  reveal: MsgRevealMove[];
}

// round 1. Bob starts winning
const Round1: GameRound = {
  commitment: [
    {
      player: alice,
      move: generateHash("Rock", SALT),
    } as MsgMakeMove,
    {
      player: bob,
      move: generateHash("Paper", SALT),
    } as MsgMakeMove,
  ],
  reveal: [
    {
      player: alice,
      revealedMove: "Rock",
      salt: SALT,
    } as MsgRevealMove,
    {
      player: bob,
      revealedMove: "Paper",
      salt: SALT,
    } as MsgRevealMove,
  ],
};

// round 2. Alice draws the game
const Round2: GameRound = {
  commitment: [
    {
      player: alice,
      move: generateHash("Rock", SALT),
    } as MsgMakeMove,
    {
      player: bob,
      move: generateHash("Scissors", SALT),
    } as MsgMakeMove,
  ],
  reveal: [
    {
      player: alice,
      revealedMove: "Rock",
      salt: SALT,
    } as MsgRevealMove,
    {
      player: bob,
      revealedMove: "Scissors",
      salt: SALT,
    } as MsgRevealMove,
  ],
};

// round 3. Alice wins the game
const Round3: GameRound = {
  commitment: [
    {
      player: alice,
      move: generateHash("Scissors", SALT),
    } as MsgMakeMove,
    {
      player: bob,
      move: generateHash("Paper", SALT),
    } as MsgMakeMove,
  ],
  reveal: [
    {
      player: alice,
      revealedMove: "Scissors",
      salt: SALT,
    } as MsgRevealMove,
    {
      player: bob,
      revealedMove: "Paper",
      salt: SALT,
    } as MsgRevealMove,
  ],
};

export const completeGame: GameRound[] = [Round1, Round2, Round3];