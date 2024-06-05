import {
  QueryClient,
  StargateClient,
  StargateClientOptions,
} from "@cosmjs/stargate";
import { BroadcastTxSyncResponse, Comet38Client } from "@cosmjs/tendermint-rpc";
import { RPSExtension, setupRPSExtension } from "./modules/rps/queries";
import * as responses from "@cosmjs/tendermint-rpc/build/comet38/responses";

export class RPSStargateClient extends StargateClient {
  public readonly rpsQueryClient: RPSExtension | undefined;

  protected constructor(
    cmtClient: Comet38Client | undefined,
    options: StargateClientOptions = {}
  ) {
    super(cmtClient, options);
    if (cmtClient) {
      this.rpsQueryClient = QueryClient.withExtensions(
        cmtClient,
        setupRPSExtension
      );
    }
  }

  public static async connect(
    endpoint: string,
    options?: StargateClientOptions
  ): Promise<RPSStargateClient> {
    const cmtClient = await Comet38Client.connect(endpoint);
    return new RPSStargateClient(cmtClient, options);
  }

  public async cmtBroadcastTxSync(
    tx: Uint8Array
  ): Promise<responses.BroadcastTxSyncResponse | BroadcastTxSyncResponse> {
    return this.forceGetCometClient().broadcastTxSync({ tx });
  }
}
