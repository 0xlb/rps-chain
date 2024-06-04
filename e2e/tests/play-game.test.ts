import { OfflineDirectSigner } from "@cosmjs/proto-signing";
import { DeliverTxResponse, GasPrice } from "@cosmjs/stargate";
import { config } from "dotenv";
import _ from "../environment";
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

  expect(response.events).toHaveLength(5);
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