// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.176.1
//   protoc               v3.6.1
// source: lb/rps/v1/query.proto

/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Game, Params } from "./rps";

export const protobufPackage = "lb.rps.v1";

/** QueryGetGameRequest is the request type for the Query/GetGame RPC method */
export interface QueryGetGameRequest {
  index: Long;
}

/** QueryGetGameResponse is the response type for the Query/GetGame RPC method */
export interface QueryGetGameResponse {
  game?: Game | undefined;
}

/** QueryGetParamsRequest is the request type for the Query/GetParams RPC method */
export interface QueryGetParamsRequest {
}

/**
 * QueryGetParamsResponse is the response type for the Query/GetParams RPC
 * method
 */
export interface QueryGetParamsResponse {
  param?: Params | undefined;
}

function createBaseQueryGetGameRequest(): QueryGetGameRequest {
  return { index: Long.UZERO };
}

export const QueryGetGameRequest = {
  encode(message: QueryGetGameRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.index.equals(Long.UZERO)) {
      writer.uint32(8).uint64(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetGameRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetGameRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.index = reader.uint64() as Long;
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetGameRequest {
    return { index: isSet(object.index) ? Long.fromValue(object.index) : Long.UZERO };
  },

  toJSON(message: QueryGetGameRequest): unknown {
    const obj: any = {};
    if (!message.index.equals(Long.UZERO)) {
      obj.index = (message.index || Long.UZERO).toString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetGameRequest>, I>>(base?: I): QueryGetGameRequest {
    return QueryGetGameRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetGameRequest>, I>>(object: I): QueryGetGameRequest {
    const message = createBaseQueryGetGameRequest();
    message.index = (object.index !== undefined && object.index !== null) ? Long.fromValue(object.index) : Long.UZERO;
    return message;
  },
};

function createBaseQueryGetGameResponse(): QueryGetGameResponse {
  return { game: undefined };
}

export const QueryGetGameResponse = {
  encode(message: QueryGetGameResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.game !== undefined) {
      Game.encode(message.game, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetGameResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetGameResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.game = Game.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetGameResponse {
    return { game: isSet(object.game) ? Game.fromJSON(object.game) : undefined };
  },

  toJSON(message: QueryGetGameResponse): unknown {
    const obj: any = {};
    if (message.game !== undefined) {
      obj.game = Game.toJSON(message.game);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetGameResponse>, I>>(base?: I): QueryGetGameResponse {
    return QueryGetGameResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetGameResponse>, I>>(object: I): QueryGetGameResponse {
    const message = createBaseQueryGetGameResponse();
    message.game = (object.game !== undefined && object.game !== null) ? Game.fromPartial(object.game) : undefined;
    return message;
  },
};

function createBaseQueryGetParamsRequest(): QueryGetParamsRequest {
  return {};
}

export const QueryGetParamsRequest = {
  encode(_: QueryGetParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetParamsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): QueryGetParamsRequest {
    return {};
  },

  toJSON(_: QueryGetParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetParamsRequest>, I>>(base?: I): QueryGetParamsRequest {
    return QueryGetParamsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetParamsRequest>, I>>(_: I): QueryGetParamsRequest {
    const message = createBaseQueryGetParamsRequest();
    return message;
  },
};

function createBaseQueryGetParamsResponse(): QueryGetParamsResponse {
  return { param: undefined };
}

export const QueryGetParamsResponse = {
  encode(message: QueryGetParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.param !== undefined) {
      Params.encode(message.param, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.param = Params.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetParamsResponse {
    return { param: isSet(object.param) ? Params.fromJSON(object.param) : undefined };
  },

  toJSON(message: QueryGetParamsResponse): unknown {
    const obj: any = {};
    if (message.param !== undefined) {
      obj.param = Params.toJSON(message.param);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetParamsResponse>, I>>(base?: I): QueryGetParamsResponse {
    return QueryGetParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetParamsResponse>, I>>(object: I): QueryGetParamsResponse {
    const message = createBaseQueryGetParamsResponse();
    message.param = (object.param !== undefined && object.param !== null)
      ? Params.fromPartial(object.param)
      : undefined;
    return message;
  },
};

/** Query is the query service for the rps module */
export interface Query {
  /** GetGame returns the game at the requested index */
  GetGame(request: QueryGetGameRequest): Promise<QueryGetGameResponse>;
  /** GetParams returns the RPS module params */
  GetParams(request: QueryGetParamsRequest): Promise<QueryGetParamsResponse>;
}

export const QueryServiceName = "lb.rps.v1.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.GetGame = this.GetGame.bind(this);
    this.GetParams = this.GetParams.bind(this);
  }
  GetGame(request: QueryGetGameRequest): Promise<QueryGetGameResponse> {
    const data = QueryGetGameRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetGame", data);
    return promise.then((data) => QueryGetGameResponse.decode(_m0.Reader.create(data)));
  }

  GetParams(request: QueryGetParamsRequest): Promise<QueryGetParamsResponse> {
    const data = QueryGetParamsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetParams", data);
    return promise.then((data) => QueryGetParamsResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}