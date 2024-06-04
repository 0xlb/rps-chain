import {
    MsgCreateGame,
    MsgCreateGameResponse,
    MsgMakeMove,
    MsgMakeMoveResponse,
    MsgRevealMove,
    MsgRevealMoveResponse,
    MsgUpdateParams,
    MsgUpdateParamsResponse,
  } from "../generated/lb/rps/v1/tx";
  import { EncodeObject, GeneratedType } from "@cosmjs/proto-signing";
  
  export const typeUrlMsgCreateGame = "/lb.rps.v1.MsgCreateGame";
  export const typeUrlMsgCreateGameResponse = "/lb.rps.v1.MsgCreateGameResponse";
  export const typeUrlMsgMakeMove = "/lb.rps.v1.MsgMakeMove";
  export const typeUrlMsgMakeMoveResponse = "/lb.rps.v1.MsgMakeMoveResponse";
  export const typeUrlMsgRevealMove = "/lb.rps.v1.MsgRevealMove";
  export const typeUrlMsgRevealMoveResponse = "/lb.rps.v1.MsgRevealMoveResponse";
  export const typeUrlMsgUpdateParams = "/lb.rps.v1.MsgUpdateParams";
  export const typeUrlMsgUpdateParamsResponse =
    "/lb.rps.v1.MsgUpdateParamsResponse";
  
  export const rpsTypes: ReadonlyArray<[string, GeneratedType]> = [
    [typeUrlMsgCreateGame, MsgCreateGame],
    [typeUrlMsgCreateGameResponse, MsgCreateGameResponse],
    [typeUrlMsgMakeMove, MsgMakeMove],
    [typeUrlMsgMakeMoveResponse, MsgMakeMoveResponse],
    [typeUrlMsgRevealMove, MsgRevealMove],
    [typeUrlMsgRevealMoveResponse, MsgRevealMoveResponse],
    [typeUrlMsgRevealMoveResponse, MsgRevealMoveResponse],
    [typeUrlMsgUpdateParams, MsgUpdateParams],
    [typeUrlMsgUpdateParamsResponse, MsgUpdateParamsResponse],
  ];
  
  export interface MsgCreateGameEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgCreateGame";
    readonly value: Partial<MsgCreateGame>;
  }
  
  export function isMsgCreateGameEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgCreateGameEncodeObject {
    return (
      (encodeObject as MsgCreateGameEncodeObject).typeUrl === typeUrlMsgCreateGame
    );
  }
  
  export interface MsgCreateGameResponseEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgCreateGameResponse";
    readonly value: Partial<MsgCreateGame>;
  }
  
  export function isMsgCreateGameResponseEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgCreateGameResponseEncodeObject {
    return (
      (encodeObject as MsgCreateGameResponseEncodeObject).typeUrl ===
      typeUrlMsgCreateGameResponse
    );
  }
  
  export interface MsgMakeMoveEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgMakeMove";
    readonly value: Partial<MsgMakeMove>;
  }
  
  export function isMsgMakeMoveEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgMakeMoveEncodeObject {
    return (
      (encodeObject as MsgMakeMoveEncodeObject).typeUrl === typeUrlMsgMakeMove
    );
  }
  
  export interface MsgMakeMoveResponseEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgMakeMoveResponse";
    readonly value: Partial<MsgMakeMove>;
  }
  
  export function isMsgMakeMoveResponseEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgMakeMoveResponseEncodeObject {
    return (
      (encodeObject as MsgMakeMoveResponseEncodeObject).typeUrl ===
      typeUrlMsgMakeMoveResponse
    );
  }
  
  export interface MsgRevealMoveEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgRevealMove";
    readonly value: Partial<MsgRevealMove>;
  }
  
  export function isMsgRevealMoveEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgRevealMoveEncodeObject {
    return (
      (encodeObject as MsgRevealMoveEncodeObject).typeUrl === typeUrlMsgRevealMove
    );
  }
  
  export interface MsgRevealMoveResponseEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgRevealMoveResponse";
    readonly value: Partial<MsgRevealMove>;
  }
  
  export function isMsgRevealMoveResponseEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgRevealMoveResponseEncodeObject {
    return (
      (encodeObject as MsgRevealMoveResponseEncodeObject).typeUrl ===
      typeUrlMsgRevealMoveResponse
    );
  }
  
  export interface MsgUpdateParamsEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgUpdateParams";
    readonly value: Partial<MsgUpdateParams>;
  }
  
  export function isMsgUpdateParamsEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgUpdateParamsEncodeObject {
    return (
      (encodeObject as MsgUpdateParamsEncodeObject).typeUrl ===
      typeUrlMsgUpdateParams
    );
  }
  
  export interface MsgUpdateParamsResponseEncodeObject extends EncodeObject {
    readonly typeUrl: "/lb.rps.v1.MsgUpdateParamsResponse";
    readonly value: Partial<MsgUpdateParams>;
  }
  
  export function isMsgUpdateParamsResponseEncodeObject(
    encodeObject: EncodeObject
  ): encodeObject is MsgUpdateParamsResponseEncodeObject {
    return (
      (encodeObject as MsgUpdateParamsResponseEncodeObject).typeUrl ===
      typeUrlMsgUpdateParamsResponse
    );
  }