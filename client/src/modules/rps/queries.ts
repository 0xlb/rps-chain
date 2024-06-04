import { QueryClient, createProtobufRpcClient } from "@cosmjs/stargate";
import { QueryClientImpl } from "./../../types/generated/lb/rps/v1/query";
import { Game, Params } from "./../../types/generated/lb/rps/v1/rps";
import Long from "long";

export interface RPSExtension {
  readonly rps: {
    readonly getGame: (index: number) => Promise<Game | undefined>;
    readonly getParams: () => Promise<Params | undefined>;
  };
}

export function setupRPSExtension(base: QueryClient): RPSExtension {
  const rpc = createProtobufRpcClient(base);
  const queryService = new QueryClientImpl(rpc);
  return {
    rps: {
      getGame: async (index: number): Promise<Game | undefined> => {
        const res = await queryService.GetGame({
          index: Long.fromNumber(index),
        });
        return res.game;
      },
      getParams: async (): Promise<Params | undefined> => {
        const res = await queryService.GetParams({});
        return res.param;
      },
    },
  };
}
