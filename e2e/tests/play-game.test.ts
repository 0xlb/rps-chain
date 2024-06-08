import { OfflineDirectSigner } from "@cosmjs/proto-signing";
import { DeliverTxResponse, GasPrice } from "@cosmjs/stargate";
import { config } from "dotenv";
import _ from "../environment";
import { TxRaw } from "cosmjs-types/cosmos/tx/v1beta1/tx";
import { beforeAll, it, expect } from "@jest/globals";
import { RPSSigningStargateClient } from "rps-chain-client/src/rps_signingstargateclient";
import { RPSExtension } from "rps-chain-client/src/modules/rps/queries";
import {
  getCreatedGameId,
  getCreateGameEvent,
  getMakeMoveEvent,
} from "rps-chain-client/src/types/rps/events";
import { Game } from "rps-chain-client/src/types/generated/lb/rps/v1/rps";
import { getSignerFromMnemonic } from "rps-chain-client/src/util/signer";
import { SALT, generateHash } from "../util/hash";
import { RPSStargateClient } from "rps-chain-client/src/rps_stargateclient";
import {
  MsgMakeMove,
  MsgRevealMove,
} from "rps-chain-client/src/types/generated/lb/rps/v1/tx";
import Long from "long";
import { completeGame } from "../util/game-sequence";
import {
  typeUrlMsgMakeMove,
  typeUrlMsgRevealMove,
} from "rps-chain-client/src/types/rps/messages";

config();

const { RPC_URL, ADDRESS_ALICE: alice, ADDRESS_BOB: bob } = process.env;
let aliceSigner: OfflineDirectSigner, bobSigner: OfflineDirectSigner;

// create signers
beforeAll(async function () {
  aliceSigner = await getSignerFromMnemonic(process.env.MNEMONIC_ALICE);
  bobSigner = await getSignerFromMnemonic(process.env.MNEMONIC_BOB);
  expect((await aliceSigner.getAccounts())[0].address).toEqual(alice);
  expect((await bobSigner.getAccounts())[0].address).toEqual(bob);
});

let aliceClient: RPSSigningStargateClient,
  bobClient: RPSSigningStargateClient,
  rps: RPSExtension["rps"];

// create signing clients
beforeAll(async function () {
  aliceClient = await RPSSigningStargateClient.connectWithSigner(
    RPC_URL,
    aliceSigner,
    {
      gasPrice: GasPrice.fromString("0rps"),
    }
  );
  bobClient = await RPSSigningStargateClient.connectWithSigner(
    RPC_URL,
    bobSigner,
    {
      gasPrice: GasPrice.fromString("0rps"),
    }
  );
  rps = aliceClient.rpsQueryClient!.rps;
});

const initialBalances = {
  alice: 99999000000,
  bob: 100000000000,
};

// assert alice and bob have funds for the tests
beforeAll(async function () {
  expect(
    parseInt((await aliceClient.getBalance(alice, "rps")).amount, 10)
  ).toEqual(initialBalances.alice);
  expect(parseInt((await bobClient.getBalance(bob, "rps")).amount, 10)).toEqual(
    initialBalances.bob
  );
});

let gameId: number;

it("can create a game", async function () {
  const response: DeliverTxResponse = await aliceClient.createGame(
    alice,
    bob,
    3,
    "auto"
  );

  expect(response.events).not.toHaveLength(0);
  const event = getCreateGameEvent(response);
  expect(event).not.toBeUndefined();
  gameId = getCreatedGameId(event!)!;
  const game: Game = (await rps.getGame(gameId))!;
  expect(game.gameNumber.toNumber()).toEqual(gameId);
  expect(game.playerA).toEqual(alice);
  expect(game.playerB).toEqual(bob);
  expect(game.rounds.toNumber()).toEqual(3);
  expect(game.status).toEqual("Waiting");
});

it("both players can play first moves", async function () {
  const aliceMoveRes = await aliceClient.makeMove(
    alice,
    gameId,
    generateHash("Rock", SALT),
    "auto"
  );
  // assert response code is OK (0)
  expect(aliceMoveRes.code).toEqual(0);
  const aliceEvent = getMakeMoveEvent(aliceMoveRes);
  expect(aliceEvent).not.toBeUndefined();

  const bobMoveRes = await bobClient.makeMove(
    bob,
    gameId,
    generateHash("Paper", SALT),
    "auto"
  );
  // assert response code is OK (0)
  expect(bobMoveRes.code).toEqual(0);
  const bobEvent = getMakeMoveEvent(bobMoveRes);
  expect(bobEvent).not.toBeUndefined();

  // check players move where updated
  const game: Game = (await rps.getGame(gameId))!;
  expect(game.gameNumber.toNumber()).toEqual(gameId);
  expect(game.playerAMoves).toHaveLength(1);
  expect(game.playerBMoves).toHaveLength(1);
}, 10_000); // increase time to avoid timeout error

interface AccountInfo {
  accountNumber: number;
  sequence: number;
}

it("submit many txs - can finish the game in one block", async function () {
  // create a new game and get the game id from the events
  const response: DeliverTxResponse = await aliceClient.createGame(
    alice,
    bob,
    3,
    "auto"
  );

  const event = getCreateGameEvent(response);
  expect(event).not.toBeUndefined();
  const gameId = getCreatedGameId(event!);

  const client: RPSStargateClient = await RPSStargateClient.connect(RPC_URL);
  const chainId: string = await client.getChainId();

  // store alice and bob accounts info to set the nonce (sequence)
  // accordingly when adding many msgs in the same block
  const accountInfo = {
    [alice]: (await client.getAccount(alice)) as AccountInfo,
    [bob]: (await client.getAccount(bob)) as AccountInfo,
  };

  const getSignedTx = async (
    msg: MsgMakeMove | MsgRevealMove,
    typeUrl: string
  ): Promise<TxRaw> => {
    const client = msg.player == alice ? aliceClient : bobClient;
    // update the gameIndex to the created game
    //@ts-ignore
    msg.gameIndex = Long.fromNumber(gameId);
    return client.sign(
      msg.player,
      [{ typeUrl: typeUrl, value: msg }],
      {
        amount: [{ denom: "rps", amount: "0" }],
        gas: "500000",
      },
      "",
      {
        accountNumber: accountInfo[msg.player].accountNumber,
        sequence: accountInfo[msg.player].sequence++,
        chainId: chainId,
      }
    );
  };

  const txList: TxRaw[] = [];

  for (const round of completeGame) {
    // for each round, both players need to first submit the commitment
    // and then reveal the move
    for (const msg of round.commitment) {
      txList.push(await getSignedTx(msg, typeUrlMsgMakeMove));
    }
    for (const msg of round.reveal) {
      txList.push(await getSignedTx(msg, typeUrlMsgRevealMove));
    }
  }

  // broadcast all txs
  for (const i in txList) {
    await client.cmtBroadcastTxSync(TxRaw.encode(txList[i]).finish());
  }

  const initBlockHeight = (await client.getBlock()).header.height;
  let currentBlockHeight = initBlockHeight;
  while (currentBlockHeight <= initBlockHeight + 1) {
    currentBlockHeight = (await client.getBlock()).header.height;
  }

  // check that game ended and alice won
  const game: Game | undefined = await rps.getGame(gameId!);
  expect(game).not.toBeUndefined();
  expect(game!.gameNumber.toNumber()).toEqual(gameId);
  expect(game!.playerAMoves).toHaveLength(3);
  expect(game!.playerBMoves).toHaveLength(3);
  expect(game!.score[0].toNumber()).toEqual(2);
  expect(game!.score[1].toNumber()).toEqual(1);
  expect(game!.status).toEqual("Player A Wins");
}, 25_000); // increase test timeout
