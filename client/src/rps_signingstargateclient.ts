import { GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import {
  defaultRegistryTypes,
  DeliverTxResponse,
  QueryClient,
  SigningStargateClient,
  SigningStargateClientOptions,
  StdFee,
} from "@cosmjs/stargate";
import { Comet38Client } from "@cosmjs/tendermint-rpc";
import Long from "long";
import { RPSExtension, setupRPSExtension } from "./modules/rps/queries";
import {
  rpsTypes,
  MsgCreateGameEncodeObject,
  MsgMakeMoveEncodeObject,
  MsgRevealMoveEncodeObject,
  typeUrlMsgCreateGame,
  typeUrlMsgMakeMove,
  typeUrlMsgRevealMove,
} from "./types/rps/messages";

export const rpsDefaultRegistryTypes: ReadonlyArray<[string, GeneratedType]> = [
  ...defaultRegistryTypes,
  ...rpsTypes,
];

function createDefaultRegistry(): Registry {
  return new Registry(rpsDefaultRegistryTypes);
}

export class RPSSigningStargateClient extends SigningStargateClient {
  public readonly rpsQueryClient: RPSExtension | undefined;

  protected constructor(
    cmtClient: Comet38Client | undefined,
    signer: OfflineSigner,
    options: SigningStargateClientOptions
  ) {
    super(cmtClient, signer, options);
    if (cmtClient) {
      this.rpsQueryClient = QueryClient.withExtensions(
        cmtClient,
        setupRPSExtension
      );
    }
  }

  public static async connectWithSigner(
    endpoint: string,
    signer: OfflineSigner,
    options: SigningStargateClientOptions = {}
  ): Promise<RPSSigningStargateClient> {
    const cmtClient = await Comet38Client.connect(endpoint);
    return new RPSSigningStargateClient(cmtClient, signer, {
      registry: createDefaultRegistry(),
      ...options,
    });
  }

  public async createGame(
    creator: string,
    oponent: string,
    rounds: number,
    fee: StdFee | "auto" | number,
    memo = ""
  ): Promise<DeliverTxResponse> {
    const msg: MsgCreateGameEncodeObject = {
      typeUrl: typeUrlMsgCreateGame,
      value: {
        creator: creator,
        oponent: oponent,
        rounds: Long.fromNumber(rounds),
      },
    };
    return this.signAndBroadcast(creator, [msg], fee, memo);
  }

  public async makeMove(
    player: string,
    gameIndex: number,
    move: string,
    fee: StdFee | "auto" | number,
    memo = ""
  ): Promise<DeliverTxResponse> {
    const msg: MsgMakeMoveEncodeObject = {
      typeUrl: typeUrlMsgMakeMove,
      value: {
        player: player,
        gameIndex: Long.fromNumber(gameIndex),
        move: move,
      },
    };
    return this.signAndBroadcast(player, [msg], fee, memo);
  }

  public async revealMove(
    player: string,
    gameIndex: number,
    revealedMove: string,
    salt: string,
    fee: StdFee | "auto" | number,
    memo = ""
  ): Promise<DeliverTxResponse> {
    const msg: MsgRevealMoveEncodeObject = {
      typeUrl: typeUrlMsgRevealMove,
      value: {
        player: player,
        gameIndex: Long.fromNumber(gameIndex),
        revealedMove: revealedMove,
        salt: salt,
      },
    };
    return this.signAndBroadcast(player, [msg], fee, memo);
  }
}