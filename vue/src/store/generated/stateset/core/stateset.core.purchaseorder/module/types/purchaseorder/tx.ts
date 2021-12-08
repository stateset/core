/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "stateset.core.purchaseorder";

export interface MsgFinancePurchaseorder {
  creator: string;
  id: number;
}

export interface MsgFinancePurchaseorderResponse {}

export interface MsgCancelPurchaseorder {
  creator: string;
  id: number;
}

export interface MsgCancelPurchaseorderResponse {}

export interface MsgCompletePurchaseorder {
  creator: string;
  id: number;
}

export interface MsgCompletePurchaseorderResponse {}

export interface MsgCreateSentPurchaseorder {
  creator: string;
  did: string;
  chain: string;
}

export interface MsgCreateSentPurchaseorderResponse {
  id: number;
}

export interface MsgUpdateSentPurchaseorder {
  creator: string;
  id: number;
  did: string;
  chain: string;
}

export interface MsgUpdateSentPurchaseorderResponse {}

export interface MsgDeleteSentPurchaseorder {
  creator: string;
  id: number;
}

export interface MsgDeleteSentPurchaseorderResponse {}

export interface MsgCreateTimedoutPurchaseorder {
  creator: string;
  did: string;
  chain: string;
}

export interface MsgCreateTimedoutPurchaseorderResponse {
  id: number;
}

export interface MsgUpdateTimedoutPurchaseorder {
  creator: string;
  id: number;
  did: string;
  chain: string;
}

export interface MsgUpdateTimedoutPurchaseorderResponse {}

export interface MsgDeleteTimedoutPurchaseorder {
  creator: string;
  id: number;
}

export interface MsgDeleteTimedoutPurchaseorderResponse {}

export interface MsgRequestPurchaseorder {
  creator: string;
  did: string;
  uri: string;
  amount: string;
  state: string;
}

export interface MsgRequestPurchaseorderResponse {}

const baseMsgFinancePurchaseorder: object = { creator: "", id: 0 };

export const MsgFinancePurchaseorder = {
  encode(
    message: MsgFinancePurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFinancePurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgFinancePurchaseorder,
    } as MsgFinancePurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFinancePurchaseorder {
    const message = {
      ...baseMsgFinancePurchaseorder,
    } as MsgFinancePurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgFinancePurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgFinancePurchaseorder>
  ): MsgFinancePurchaseorder {
    const message = {
      ...baseMsgFinancePurchaseorder,
    } as MsgFinancePurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgFinancePurchaseorderResponse: object = {};

export const MsgFinancePurchaseorderResponse = {
  encode(
    _: MsgFinancePurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgFinancePurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgFinancePurchaseorderResponse,
    } as MsgFinancePurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgFinancePurchaseorderResponse {
    const message = {
      ...baseMsgFinancePurchaseorderResponse,
    } as MsgFinancePurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgFinancePurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgFinancePurchaseorderResponse>
  ): MsgFinancePurchaseorderResponse {
    const message = {
      ...baseMsgFinancePurchaseorderResponse,
    } as MsgFinancePurchaseorderResponse;
    return message;
  },
};

const baseMsgCancelPurchaseorder: object = { creator: "", id: 0 };

export const MsgCancelPurchaseorder = {
  encode(
    message: MsgCancelPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelPurchaseorder } as MsgCancelPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelPurchaseorder {
    const message = { ...baseMsgCancelPurchaseorder } as MsgCancelPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCancelPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCancelPurchaseorder>
  ): MsgCancelPurchaseorder {
    const message = { ...baseMsgCancelPurchaseorder } as MsgCancelPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCancelPurchaseorderResponse: object = {};

export const MsgCancelPurchaseorderResponse = {
  encode(
    _: MsgCancelPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCancelPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCancelPurchaseorderResponse,
    } as MsgCancelPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCancelPurchaseorderResponse {
    const message = {
      ...baseMsgCancelPurchaseorderResponse,
    } as MsgCancelPurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgCancelPurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCancelPurchaseorderResponse>
  ): MsgCancelPurchaseorderResponse {
    const message = {
      ...baseMsgCancelPurchaseorderResponse,
    } as MsgCancelPurchaseorderResponse;
    return message;
  },
};

const baseMsgCompletePurchaseorder: object = { creator: "", id: 0 };

export const MsgCompletePurchaseorder = {
  encode(
    message: MsgCompletePurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCompletePurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCompletePurchaseorder,
    } as MsgCompletePurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCompletePurchaseorder {
    const message = {
      ...baseMsgCompletePurchaseorder,
    } as MsgCompletePurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCompletePurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCompletePurchaseorder>
  ): MsgCompletePurchaseorder {
    const message = {
      ...baseMsgCompletePurchaseorder,
    } as MsgCompletePurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCompletePurchaseorderResponse: object = {};

export const MsgCompletePurchaseorderResponse = {
  encode(
    _: MsgCompletePurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCompletePurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCompletePurchaseorderResponse,
    } as MsgCompletePurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCompletePurchaseorderResponse {
    const message = {
      ...baseMsgCompletePurchaseorderResponse,
    } as MsgCompletePurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgCompletePurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCompletePurchaseorderResponse>
  ): MsgCompletePurchaseorderResponse {
    const message = {
      ...baseMsgCompletePurchaseorderResponse,
    } as MsgCompletePurchaseorderResponse;
    return message;
  },
};

const baseMsgCreateSentPurchaseorder: object = {
  creator: "",
  did: "",
  chain: "",
};

export const MsgCreateSentPurchaseorder = {
  encode(
    message: MsgCreateSentPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSentPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSentPurchaseorder,
    } as MsgCreateSentPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.chain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSentPurchaseorder {
    const message = {
      ...baseMsgCreateSentPurchaseorder,
    } as MsgCreateSentPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    return message;
  },

  toJSON(message: MsgCreateSentPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSentPurchaseorder>
  ): MsgCreateSentPurchaseorder {
    const message = {
      ...baseMsgCreateSentPurchaseorder,
    } as MsgCreateSentPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    return message;
  },
};

const baseMsgCreateSentPurchaseorderResponse: object = { id: 0 };

export const MsgCreateSentPurchaseorderResponse = {
  encode(
    message: MsgCreateSentPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateSentPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSentPurchaseorderResponse,
    } as MsgCreateSentPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateSentPurchaseorderResponse {
    const message = {
      ...baseMsgCreateSentPurchaseorderResponse,
    } as MsgCreateSentPurchaseorderResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateSentPurchaseorderResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSentPurchaseorderResponse>
  ): MsgCreateSentPurchaseorderResponse {
    const message = {
      ...baseMsgCreateSentPurchaseorderResponse,
    } as MsgCreateSentPurchaseorderResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateSentPurchaseorder: object = {
  creator: "",
  id: 0,
  did: "",
  chain: "",
};

export const MsgUpdateSentPurchaseorder = {
  encode(
    message: MsgUpdateSentPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.chain !== "") {
      writer.uint32(34).string(message.chain);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSentPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSentPurchaseorder,
    } as MsgUpdateSentPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.chain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateSentPurchaseorder {
    const message = {
      ...baseMsgUpdateSentPurchaseorder,
    } as MsgUpdateSentPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateSentPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateSentPurchaseorder>
  ): MsgUpdateSentPurchaseorder {
    const message = {
      ...baseMsgUpdateSentPurchaseorder,
    } as MsgUpdateSentPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    return message;
  },
};

const baseMsgUpdateSentPurchaseorderResponse: object = {};

export const MsgUpdateSentPurchaseorderResponse = {
  encode(
    _: MsgUpdateSentPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSentPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSentPurchaseorderResponse,
    } as MsgUpdateSentPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateSentPurchaseorderResponse {
    const message = {
      ...baseMsgUpdateSentPurchaseorderResponse,
    } as MsgUpdateSentPurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgUpdateSentPurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateSentPurchaseorderResponse>
  ): MsgUpdateSentPurchaseorderResponse {
    const message = {
      ...baseMsgUpdateSentPurchaseorderResponse,
    } as MsgUpdateSentPurchaseorderResponse;
    return message;
  },
};

const baseMsgDeleteSentPurchaseorder: object = { creator: "", id: 0 };

export const MsgDeleteSentPurchaseorder = {
  encode(
    message: MsgDeleteSentPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSentPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSentPurchaseorder,
    } as MsgDeleteSentPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteSentPurchaseorder {
    const message = {
      ...baseMsgDeleteSentPurchaseorder,
    } as MsgDeleteSentPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteSentPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteSentPurchaseorder>
  ): MsgDeleteSentPurchaseorder {
    const message = {
      ...baseMsgDeleteSentPurchaseorder,
    } as MsgDeleteSentPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteSentPurchaseorderResponse: object = {};

export const MsgDeleteSentPurchaseorderResponse = {
  encode(
    _: MsgDeleteSentPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSentPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSentPurchaseorderResponse,
    } as MsgDeleteSentPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteSentPurchaseorderResponse {
    const message = {
      ...baseMsgDeleteSentPurchaseorderResponse,
    } as MsgDeleteSentPurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgDeleteSentPurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteSentPurchaseorderResponse>
  ): MsgDeleteSentPurchaseorderResponse {
    const message = {
      ...baseMsgDeleteSentPurchaseorderResponse,
    } as MsgDeleteSentPurchaseorderResponse;
    return message;
  },
};

const baseMsgCreateTimedoutPurchaseorder: object = {
  creator: "",
  did: "",
  chain: "",
};

export const MsgCreateTimedoutPurchaseorder = {
  encode(
    message: MsgCreateTimedoutPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateTimedoutPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTimedoutPurchaseorder,
    } as MsgCreateTimedoutPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.chain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTimedoutPurchaseorder {
    const message = {
      ...baseMsgCreateTimedoutPurchaseorder,
    } as MsgCreateTimedoutPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    return message;
  },

  toJSON(message: MsgCreateTimedoutPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTimedoutPurchaseorder>
  ): MsgCreateTimedoutPurchaseorder {
    const message = {
      ...baseMsgCreateTimedoutPurchaseorder,
    } as MsgCreateTimedoutPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    return message;
  },
};

const baseMsgCreateTimedoutPurchaseorderResponse: object = { id: 0 };

export const MsgCreateTimedoutPurchaseorderResponse = {
  encode(
    message: MsgCreateTimedoutPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateTimedoutPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTimedoutPurchaseorderResponse,
    } as MsgCreateTimedoutPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateTimedoutPurchaseorderResponse {
    const message = {
      ...baseMsgCreateTimedoutPurchaseorderResponse,
    } as MsgCreateTimedoutPurchaseorderResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateTimedoutPurchaseorderResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTimedoutPurchaseorderResponse>
  ): MsgCreateTimedoutPurchaseorderResponse {
    const message = {
      ...baseMsgCreateTimedoutPurchaseorderResponse,
    } as MsgCreateTimedoutPurchaseorderResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateTimedoutPurchaseorder: object = {
  creator: "",
  id: 0,
  did: "",
  chain: "",
};

export const MsgUpdateTimedoutPurchaseorder = {
  encode(
    message: MsgUpdateTimedoutPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.chain !== "") {
      writer.uint32(34).string(message.chain);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateTimedoutPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateTimedoutPurchaseorder,
    } as MsgUpdateTimedoutPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.did = reader.string();
          break;
        case 4:
          message.chain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateTimedoutPurchaseorder {
    const message = {
      ...baseMsgUpdateTimedoutPurchaseorder,
    } as MsgUpdateTimedoutPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    return message;
  },

  toJSON(message: MsgUpdateTimedoutPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateTimedoutPurchaseorder>
  ): MsgUpdateTimedoutPurchaseorder {
    const message = {
      ...baseMsgUpdateTimedoutPurchaseorder,
    } as MsgUpdateTimedoutPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    return message;
  },
};

const baseMsgUpdateTimedoutPurchaseorderResponse: object = {};

export const MsgUpdateTimedoutPurchaseorderResponse = {
  encode(
    _: MsgUpdateTimedoutPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateTimedoutPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateTimedoutPurchaseorderResponse,
    } as MsgUpdateTimedoutPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateTimedoutPurchaseorderResponse {
    const message = {
      ...baseMsgUpdateTimedoutPurchaseorderResponse,
    } as MsgUpdateTimedoutPurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgUpdateTimedoutPurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateTimedoutPurchaseorderResponse>
  ): MsgUpdateTimedoutPurchaseorderResponse {
    const message = {
      ...baseMsgUpdateTimedoutPurchaseorderResponse,
    } as MsgUpdateTimedoutPurchaseorderResponse;
    return message;
  },
};

const baseMsgDeleteTimedoutPurchaseorder: object = { creator: "", id: 0 };

export const MsgDeleteTimedoutPurchaseorder = {
  encode(
    message: MsgDeleteTimedoutPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteTimedoutPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteTimedoutPurchaseorder,
    } as MsgDeleteTimedoutPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteTimedoutPurchaseorder {
    const message = {
      ...baseMsgDeleteTimedoutPurchaseorder,
    } as MsgDeleteTimedoutPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteTimedoutPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteTimedoutPurchaseorder>
  ): MsgDeleteTimedoutPurchaseorder {
    const message = {
      ...baseMsgDeleteTimedoutPurchaseorder,
    } as MsgDeleteTimedoutPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteTimedoutPurchaseorderResponse: object = {};

export const MsgDeleteTimedoutPurchaseorderResponse = {
  encode(
    _: MsgDeleteTimedoutPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteTimedoutPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteTimedoutPurchaseorderResponse,
    } as MsgDeleteTimedoutPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteTimedoutPurchaseorderResponse {
    const message = {
      ...baseMsgDeleteTimedoutPurchaseorderResponse,
    } as MsgDeleteTimedoutPurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgDeleteTimedoutPurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteTimedoutPurchaseorderResponse>
  ): MsgDeleteTimedoutPurchaseorderResponse {
    const message = {
      ...baseMsgDeleteTimedoutPurchaseorderResponse,
    } as MsgDeleteTimedoutPurchaseorderResponse;
    return message;
  },
};

const baseMsgRequestPurchaseorder: object = {
  creator: "",
  did: "",
  uri: "",
  amount: "",
  state: "",
};

export const MsgRequestPurchaseorder = {
  encode(
    message: MsgRequestPurchaseorder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.uri !== "") {
      writer.uint32(26).string(message.uri);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.state !== "") {
      writer.uint32(42).string(message.state);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestPurchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRequestPurchaseorder,
    } as MsgRequestPurchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.did = reader.string();
          break;
        case 3:
          message.uri = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.state = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestPurchaseorder {
    const message = {
      ...baseMsgRequestPurchaseorder,
    } as MsgRequestPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri);
    } else {
      message.uri = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    return message;
  },

  toJSON(message: MsgRequestPurchaseorder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.uri !== undefined && (obj.uri = message.uri);
    message.amount !== undefined && (obj.amount = message.amount);
    message.state !== undefined && (obj.state = message.state);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRequestPurchaseorder>
  ): MsgRequestPurchaseorder {
    const message = {
      ...baseMsgRequestPurchaseorder,
    } as MsgRequestPurchaseorder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri;
    } else {
      message.uri = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    return message;
  },
};

const baseMsgRequestPurchaseorderResponse: object = {};

export const MsgRequestPurchaseorderResponse = {
  encode(
    _: MsgRequestPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRequestPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRequestPurchaseorderResponse,
    } as MsgRequestPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRequestPurchaseorderResponse {
    const message = {
      ...baseMsgRequestPurchaseorderResponse,
    } as MsgRequestPurchaseorderResponse;
    return message;
  },

  toJSON(_: MsgRequestPurchaseorderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRequestPurchaseorderResponse>
  ): MsgRequestPurchaseorderResponse {
    const message = {
      ...baseMsgRequestPurchaseorderResponse,
    } as MsgRequestPurchaseorderResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  FinancePurchaseorder(
    request: MsgFinancePurchaseorder
  ): Promise<MsgFinancePurchaseorderResponse>;
  CancelPurchaseorder(
    request: MsgCancelPurchaseorder
  ): Promise<MsgCancelPurchaseorderResponse>;
  CompletePurchaseorder(
    request: MsgCompletePurchaseorder
  ): Promise<MsgCompletePurchaseorderResponse>;
  CreateSentPurchaseorder(
    request: MsgCreateSentPurchaseorder
  ): Promise<MsgCreateSentPurchaseorderResponse>;
  UpdateSentPurchaseorder(
    request: MsgUpdateSentPurchaseorder
  ): Promise<MsgUpdateSentPurchaseorderResponse>;
  DeleteSentPurchaseorder(
    request: MsgDeleteSentPurchaseorder
  ): Promise<MsgDeleteSentPurchaseorderResponse>;
  CreateTimedoutPurchaseorder(
    request: MsgCreateTimedoutPurchaseorder
  ): Promise<MsgCreateTimedoutPurchaseorderResponse>;
  UpdateTimedoutPurchaseorder(
    request: MsgUpdateTimedoutPurchaseorder
  ): Promise<MsgUpdateTimedoutPurchaseorderResponse>;
  DeleteTimedoutPurchaseorder(
    request: MsgDeleteTimedoutPurchaseorder
  ): Promise<MsgDeleteTimedoutPurchaseorderResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RequestPurchaseorder(
    request: MsgRequestPurchaseorder
  ): Promise<MsgRequestPurchaseorderResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  FinancePurchaseorder(
    request: MsgFinancePurchaseorder
  ): Promise<MsgFinancePurchaseorderResponse> {
    const data = MsgFinancePurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "FinancePurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgFinancePurchaseorderResponse.decode(new Reader(data))
    );
  }

  CancelPurchaseorder(
    request: MsgCancelPurchaseorder
  ): Promise<MsgCancelPurchaseorderResponse> {
    const data = MsgCancelPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "CancelPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgCancelPurchaseorderResponse.decode(new Reader(data))
    );
  }

  CompletePurchaseorder(
    request: MsgCompletePurchaseorder
  ): Promise<MsgCompletePurchaseorderResponse> {
    const data = MsgCompletePurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "CompletePurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgCompletePurchaseorderResponse.decode(new Reader(data))
    );
  }

  CreateSentPurchaseorder(
    request: MsgCreateSentPurchaseorder
  ): Promise<MsgCreateSentPurchaseorderResponse> {
    const data = MsgCreateSentPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "CreateSentPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgCreateSentPurchaseorderResponse.decode(new Reader(data))
    );
  }

  UpdateSentPurchaseorder(
    request: MsgUpdateSentPurchaseorder
  ): Promise<MsgUpdateSentPurchaseorderResponse> {
    const data = MsgUpdateSentPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "UpdateSentPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgUpdateSentPurchaseorderResponse.decode(new Reader(data))
    );
  }

  DeleteSentPurchaseorder(
    request: MsgDeleteSentPurchaseorder
  ): Promise<MsgDeleteSentPurchaseorderResponse> {
    const data = MsgDeleteSentPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "DeleteSentPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgDeleteSentPurchaseorderResponse.decode(new Reader(data))
    );
  }

  CreateTimedoutPurchaseorder(
    request: MsgCreateTimedoutPurchaseorder
  ): Promise<MsgCreateTimedoutPurchaseorderResponse> {
    const data = MsgCreateTimedoutPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "CreateTimedoutPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgCreateTimedoutPurchaseorderResponse.decode(new Reader(data))
    );
  }

  UpdateTimedoutPurchaseorder(
    request: MsgUpdateTimedoutPurchaseorder
  ): Promise<MsgUpdateTimedoutPurchaseorderResponse> {
    const data = MsgUpdateTimedoutPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "UpdateTimedoutPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgUpdateTimedoutPurchaseorderResponse.decode(new Reader(data))
    );
  }

  DeleteTimedoutPurchaseorder(
    request: MsgDeleteTimedoutPurchaseorder
  ): Promise<MsgDeleteTimedoutPurchaseorderResponse> {
    const data = MsgDeleteTimedoutPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "DeleteTimedoutPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgDeleteTimedoutPurchaseorderResponse.decode(new Reader(data))
    );
  }

  RequestPurchaseorder(
    request: MsgRequestPurchaseorder
  ): Promise<MsgRequestPurchaseorderResponse> {
    const data = MsgRequestPurchaseorder.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Msg",
      "RequestPurchaseorder",
      data
    );
    return promise.then((data) =>
      MsgRequestPurchaseorderResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
