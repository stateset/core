/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "stateset.core.invoice";

export interface MsgFactorInvoice {
  creator: string;
  id: number;
}

export interface MsgFactorInvoiceResponse {}

export interface MsgCreateSentInvoice {
  creator: string;
  did: string;
  chain: string;
}

export interface MsgCreateSentInvoiceResponse {
  id: number;
}

export interface MsgUpdateSentInvoice {
  creator: string;
  id: number;
  did: string;
  chain: string;
}

export interface MsgUpdateSentInvoiceResponse {}

export interface MsgDeleteSentInvoice {
  creator: string;
  id: number;
}

export interface MsgDeleteSentInvoiceResponse {}

export interface MsgCreateTimedoutInvoice {
  creator: string;
  did: string;
  chain: string;
}

export interface MsgCreateTimedoutInvoiceResponse {
  id: number;
}

export interface MsgUpdateTimedoutInvoice {
  creator: string;
  id: number;
  did: string;
  chain: string;
}

export interface MsgUpdateTimedoutInvoiceResponse {}

export interface MsgDeleteTimedoutInvoice {
  creator: string;
  id: number;
}

export interface MsgDeleteTimedoutInvoiceResponse {}

export interface MsgCreateInvoice {
  creator: string;
  id: string;
  did: string;
  amount: string;
  state: string;
}

export interface MsgCreateInvoiceResponse {}

export interface MsgPayInvoice {
  creator: string;
  id: number;
}

export interface MsgPayInvoiceResponse {}

export interface MsgVoidInvoice {
  creator: string;
  id: number;
}

export interface MsgVoidInvoiceResponse {}

const baseMsgFactorInvoice: object = { creator: "", id: 0 };

export const MsgFactorInvoice = {
  encode(message: MsgFactorInvoice, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFactorInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFactorInvoice } as MsgFactorInvoice;
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

  fromJSON(object: any): MsgFactorInvoice {
    const message = { ...baseMsgFactorInvoice } as MsgFactorInvoice;
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

  toJSON(message: MsgFactorInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFactorInvoice>): MsgFactorInvoice {
    const message = { ...baseMsgFactorInvoice } as MsgFactorInvoice;
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

const baseMsgFactorInvoiceResponse: object = {};

export const MsgFactorInvoiceResponse = {
  encode(
    _: MsgFactorInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgFactorInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgFactorInvoiceResponse,
    } as MsgFactorInvoiceResponse;
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

  fromJSON(_: any): MsgFactorInvoiceResponse {
    const message = {
      ...baseMsgFactorInvoiceResponse,
    } as MsgFactorInvoiceResponse;
    return message;
  },

  toJSON(_: MsgFactorInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgFactorInvoiceResponse>
  ): MsgFactorInvoiceResponse {
    const message = {
      ...baseMsgFactorInvoiceResponse,
    } as MsgFactorInvoiceResponse;
    return message;
  },
};

const baseMsgCreateSentInvoice: object = { creator: "", did: "", chain: "" };

export const MsgCreateSentInvoice = {
  encode(
    message: MsgCreateSentInvoice,
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

  decode(input: Reader | Uint8Array, length?: number): MsgCreateSentInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateSentInvoice } as MsgCreateSentInvoice;
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

  fromJSON(object: any): MsgCreateSentInvoice {
    const message = { ...baseMsgCreateSentInvoice } as MsgCreateSentInvoice;
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

  toJSON(message: MsgCreateSentInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateSentInvoice>): MsgCreateSentInvoice {
    const message = { ...baseMsgCreateSentInvoice } as MsgCreateSentInvoice;
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

const baseMsgCreateSentInvoiceResponse: object = { id: 0 };

export const MsgCreateSentInvoiceResponse = {
  encode(
    message: MsgCreateSentInvoiceResponse,
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
  ): MsgCreateSentInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSentInvoiceResponse,
    } as MsgCreateSentInvoiceResponse;
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

  fromJSON(object: any): MsgCreateSentInvoiceResponse {
    const message = {
      ...baseMsgCreateSentInvoiceResponse,
    } as MsgCreateSentInvoiceResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateSentInvoiceResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSentInvoiceResponse>
  ): MsgCreateSentInvoiceResponse {
    const message = {
      ...baseMsgCreateSentInvoiceResponse,
    } as MsgCreateSentInvoiceResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateSentInvoice: object = {
  creator: "",
  id: 0,
  did: "",
  chain: "",
};

export const MsgUpdateSentInvoice = {
  encode(
    message: MsgUpdateSentInvoice,
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

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateSentInvoice } as MsgUpdateSentInvoice;
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

  fromJSON(object: any): MsgUpdateSentInvoice {
    const message = { ...baseMsgUpdateSentInvoice } as MsgUpdateSentInvoice;
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

  toJSON(message: MsgUpdateSentInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdateSentInvoice>): MsgUpdateSentInvoice {
    const message = { ...baseMsgUpdateSentInvoice } as MsgUpdateSentInvoice;
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

const baseMsgUpdateSentInvoiceResponse: object = {};

export const MsgUpdateSentInvoiceResponse = {
  encode(
    _: MsgUpdateSentInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSentInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSentInvoiceResponse,
    } as MsgUpdateSentInvoiceResponse;
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

  fromJSON(_: any): MsgUpdateSentInvoiceResponse {
    const message = {
      ...baseMsgUpdateSentInvoiceResponse,
    } as MsgUpdateSentInvoiceResponse;
    return message;
  },

  toJSON(_: MsgUpdateSentInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateSentInvoiceResponse>
  ): MsgUpdateSentInvoiceResponse {
    const message = {
      ...baseMsgUpdateSentInvoiceResponse,
    } as MsgUpdateSentInvoiceResponse;
    return message;
  },
};

const baseMsgDeleteSentInvoice: object = { creator: "", id: 0 };

export const MsgDeleteSentInvoice = {
  encode(
    message: MsgDeleteSentInvoice,
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

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteSentInvoice } as MsgDeleteSentInvoice;
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

  fromJSON(object: any): MsgDeleteSentInvoice {
    const message = { ...baseMsgDeleteSentInvoice } as MsgDeleteSentInvoice;
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

  toJSON(message: MsgDeleteSentInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteSentInvoice>): MsgDeleteSentInvoice {
    const message = { ...baseMsgDeleteSentInvoice } as MsgDeleteSentInvoice;
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

const baseMsgDeleteSentInvoiceResponse: object = {};

export const MsgDeleteSentInvoiceResponse = {
  encode(
    _: MsgDeleteSentInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSentInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSentInvoiceResponse,
    } as MsgDeleteSentInvoiceResponse;
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

  fromJSON(_: any): MsgDeleteSentInvoiceResponse {
    const message = {
      ...baseMsgDeleteSentInvoiceResponse,
    } as MsgDeleteSentInvoiceResponse;
    return message;
  },

  toJSON(_: MsgDeleteSentInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteSentInvoiceResponse>
  ): MsgDeleteSentInvoiceResponse {
    const message = {
      ...baseMsgDeleteSentInvoiceResponse,
    } as MsgDeleteSentInvoiceResponse;
    return message;
  },
};

const baseMsgCreateTimedoutInvoice: object = {
  creator: "",
  did: "",
  chain: "",
};

export const MsgCreateTimedoutInvoice = {
  encode(
    message: MsgCreateTimedoutInvoice,
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
  ): MsgCreateTimedoutInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTimedoutInvoice,
    } as MsgCreateTimedoutInvoice;
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

  fromJSON(object: any): MsgCreateTimedoutInvoice {
    const message = {
      ...baseMsgCreateTimedoutInvoice,
    } as MsgCreateTimedoutInvoice;
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

  toJSON(message: MsgCreateTimedoutInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTimedoutInvoice>
  ): MsgCreateTimedoutInvoice {
    const message = {
      ...baseMsgCreateTimedoutInvoice,
    } as MsgCreateTimedoutInvoice;
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

const baseMsgCreateTimedoutInvoiceResponse: object = { id: 0 };

export const MsgCreateTimedoutInvoiceResponse = {
  encode(
    message: MsgCreateTimedoutInvoiceResponse,
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
  ): MsgCreateTimedoutInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTimedoutInvoiceResponse,
    } as MsgCreateTimedoutInvoiceResponse;
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

  fromJSON(object: any): MsgCreateTimedoutInvoiceResponse {
    const message = {
      ...baseMsgCreateTimedoutInvoiceResponse,
    } as MsgCreateTimedoutInvoiceResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateTimedoutInvoiceResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTimedoutInvoiceResponse>
  ): MsgCreateTimedoutInvoiceResponse {
    const message = {
      ...baseMsgCreateTimedoutInvoiceResponse,
    } as MsgCreateTimedoutInvoiceResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateTimedoutInvoice: object = {
  creator: "",
  id: 0,
  did: "",
  chain: "",
};

export const MsgUpdateTimedoutInvoice = {
  encode(
    message: MsgUpdateTimedoutInvoice,
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
  ): MsgUpdateTimedoutInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateTimedoutInvoice,
    } as MsgUpdateTimedoutInvoice;
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

  fromJSON(object: any): MsgUpdateTimedoutInvoice {
    const message = {
      ...baseMsgUpdateTimedoutInvoice,
    } as MsgUpdateTimedoutInvoice;
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

  toJSON(message: MsgUpdateTimedoutInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateTimedoutInvoice>
  ): MsgUpdateTimedoutInvoice {
    const message = {
      ...baseMsgUpdateTimedoutInvoice,
    } as MsgUpdateTimedoutInvoice;
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

const baseMsgUpdateTimedoutInvoiceResponse: object = {};

export const MsgUpdateTimedoutInvoiceResponse = {
  encode(
    _: MsgUpdateTimedoutInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateTimedoutInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateTimedoutInvoiceResponse,
    } as MsgUpdateTimedoutInvoiceResponse;
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

  fromJSON(_: any): MsgUpdateTimedoutInvoiceResponse {
    const message = {
      ...baseMsgUpdateTimedoutInvoiceResponse,
    } as MsgUpdateTimedoutInvoiceResponse;
    return message;
  },

  toJSON(_: MsgUpdateTimedoutInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateTimedoutInvoiceResponse>
  ): MsgUpdateTimedoutInvoiceResponse {
    const message = {
      ...baseMsgUpdateTimedoutInvoiceResponse,
    } as MsgUpdateTimedoutInvoiceResponse;
    return message;
  },
};

const baseMsgDeleteTimedoutInvoice: object = { creator: "", id: 0 };

export const MsgDeleteTimedoutInvoice = {
  encode(
    message: MsgDeleteTimedoutInvoice,
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
  ): MsgDeleteTimedoutInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteTimedoutInvoice,
    } as MsgDeleteTimedoutInvoice;
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

  fromJSON(object: any): MsgDeleteTimedoutInvoice {
    const message = {
      ...baseMsgDeleteTimedoutInvoice,
    } as MsgDeleteTimedoutInvoice;
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

  toJSON(message: MsgDeleteTimedoutInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteTimedoutInvoice>
  ): MsgDeleteTimedoutInvoice {
    const message = {
      ...baseMsgDeleteTimedoutInvoice,
    } as MsgDeleteTimedoutInvoice;
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

const baseMsgDeleteTimedoutInvoiceResponse: object = {};

export const MsgDeleteTimedoutInvoiceResponse = {
  encode(
    _: MsgDeleteTimedoutInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteTimedoutInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteTimedoutInvoiceResponse,
    } as MsgDeleteTimedoutInvoiceResponse;
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

  fromJSON(_: any): MsgDeleteTimedoutInvoiceResponse {
    const message = {
      ...baseMsgDeleteTimedoutInvoiceResponse,
    } as MsgDeleteTimedoutInvoiceResponse;
    return message;
  },

  toJSON(_: MsgDeleteTimedoutInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteTimedoutInvoiceResponse>
  ): MsgDeleteTimedoutInvoiceResponse {
    const message = {
      ...baseMsgDeleteTimedoutInvoiceResponse,
    } as MsgDeleteTimedoutInvoiceResponse;
    return message;
  },
};

const baseMsgCreateInvoice: object = {
  creator: "",
  id: "",
  did: "",
  amount: "",
  state: "",
};

export const MsgCreateInvoice = {
  encode(message: MsgCreateInvoice, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    if (message.did !== "") {
      writer.uint32(26).string(message.did);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.state !== "") {
      writer.uint32(42).string(message.state);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateInvoice } as MsgCreateInvoice;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        case 3:
          message.did = reader.string();
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

  fromJSON(object: any): MsgCreateInvoice {
    const message = { ...baseMsgCreateInvoice } as MsgCreateInvoice;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = String(object.did);
    } else {
      message.did = "";
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

  toJSON(message: MsgCreateInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.amount !== undefined && (obj.amount = message.amount);
    message.state !== undefined && (obj.state = message.state);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateInvoice>): MsgCreateInvoice {
    const message = { ...baseMsgCreateInvoice } as MsgCreateInvoice;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.did !== undefined && object.did !== null) {
      message.did = object.did;
    } else {
      message.did = "";
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

const baseMsgCreateInvoiceResponse: object = {};

export const MsgCreateInvoiceResponse = {
  encode(
    _: MsgCreateInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateInvoiceResponse,
    } as MsgCreateInvoiceResponse;
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

  fromJSON(_: any): MsgCreateInvoiceResponse {
    const message = {
      ...baseMsgCreateInvoiceResponse,
    } as MsgCreateInvoiceResponse;
    return message;
  },

  toJSON(_: MsgCreateInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateInvoiceResponse>
  ): MsgCreateInvoiceResponse {
    const message = {
      ...baseMsgCreateInvoiceResponse,
    } as MsgCreateInvoiceResponse;
    return message;
  },
};

const baseMsgPayInvoice: object = { creator: "", id: 0 };

export const MsgPayInvoice = {
  encode(message: MsgPayInvoice, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPayInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPayInvoice } as MsgPayInvoice;
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

  fromJSON(object: any): MsgPayInvoice {
    const message = { ...baseMsgPayInvoice } as MsgPayInvoice;
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

  toJSON(message: MsgPayInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgPayInvoice>): MsgPayInvoice {
    const message = { ...baseMsgPayInvoice } as MsgPayInvoice;
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

const baseMsgPayInvoiceResponse: object = {};

export const MsgPayInvoiceResponse = {
  encode(_: MsgPayInvoiceResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgPayInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgPayInvoiceResponse } as MsgPayInvoiceResponse;
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

  fromJSON(_: any): MsgPayInvoiceResponse {
    const message = { ...baseMsgPayInvoiceResponse } as MsgPayInvoiceResponse;
    return message;
  },

  toJSON(_: MsgPayInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgPayInvoiceResponse>): MsgPayInvoiceResponse {
    const message = { ...baseMsgPayInvoiceResponse } as MsgPayInvoiceResponse;
    return message;
  },
};

const baseMsgVoidInvoice: object = { creator: "", id: 0 };

export const MsgVoidInvoice = {
  encode(message: MsgVoidInvoice, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoidInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoidInvoice } as MsgVoidInvoice;
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

  fromJSON(object: any): MsgVoidInvoice {
    const message = { ...baseMsgVoidInvoice } as MsgVoidInvoice;
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

  toJSON(message: MsgVoidInvoice): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgVoidInvoice>): MsgVoidInvoice {
    const message = { ...baseMsgVoidInvoice } as MsgVoidInvoice;
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

const baseMsgVoidInvoiceResponse: object = {};

export const MsgVoidInvoiceResponse = {
  encode(_: MsgVoidInvoiceResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgVoidInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgVoidInvoiceResponse } as MsgVoidInvoiceResponse;
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

  fromJSON(_: any): MsgVoidInvoiceResponse {
    const message = { ...baseMsgVoidInvoiceResponse } as MsgVoidInvoiceResponse;
    return message;
  },

  toJSON(_: MsgVoidInvoiceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgVoidInvoiceResponse>): MsgVoidInvoiceResponse {
    const message = { ...baseMsgVoidInvoiceResponse } as MsgVoidInvoiceResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  FactorInvoice(request: MsgFactorInvoice): Promise<MsgFactorInvoiceResponse>;
  CreateSentInvoice(
    request: MsgCreateSentInvoice
  ): Promise<MsgCreateSentInvoiceResponse>;
  UpdateSentInvoice(
    request: MsgUpdateSentInvoice
  ): Promise<MsgUpdateSentInvoiceResponse>;
  DeleteSentInvoice(
    request: MsgDeleteSentInvoice
  ): Promise<MsgDeleteSentInvoiceResponse>;
  CreateTimedoutInvoice(
    request: MsgCreateTimedoutInvoice
  ): Promise<MsgCreateTimedoutInvoiceResponse>;
  UpdateTimedoutInvoice(
    request: MsgUpdateTimedoutInvoice
  ): Promise<MsgUpdateTimedoutInvoiceResponse>;
  DeleteTimedoutInvoice(
    request: MsgDeleteTimedoutInvoice
  ): Promise<MsgDeleteTimedoutInvoiceResponse>;
  CreateInvoice(request: MsgCreateInvoice): Promise<MsgCreateInvoiceResponse>;
  PayInvoice(request: MsgPayInvoice): Promise<MsgPayInvoiceResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  VoidInvoice(request: MsgVoidInvoice): Promise<MsgVoidInvoiceResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  FactorInvoice(request: MsgFactorInvoice): Promise<MsgFactorInvoiceResponse> {
    const data = MsgFactorInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "FactorInvoice",
      data
    );
    return promise.then((data) =>
      MsgFactorInvoiceResponse.decode(new Reader(data))
    );
  }

  CreateSentInvoice(
    request: MsgCreateSentInvoice
  ): Promise<MsgCreateSentInvoiceResponse> {
    const data = MsgCreateSentInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "CreateSentInvoice",
      data
    );
    return promise.then((data) =>
      MsgCreateSentInvoiceResponse.decode(new Reader(data))
    );
  }

  UpdateSentInvoice(
    request: MsgUpdateSentInvoice
  ): Promise<MsgUpdateSentInvoiceResponse> {
    const data = MsgUpdateSentInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "UpdateSentInvoice",
      data
    );
    return promise.then((data) =>
      MsgUpdateSentInvoiceResponse.decode(new Reader(data))
    );
  }

  DeleteSentInvoice(
    request: MsgDeleteSentInvoice
  ): Promise<MsgDeleteSentInvoiceResponse> {
    const data = MsgDeleteSentInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "DeleteSentInvoice",
      data
    );
    return promise.then((data) =>
      MsgDeleteSentInvoiceResponse.decode(new Reader(data))
    );
  }

  CreateTimedoutInvoice(
    request: MsgCreateTimedoutInvoice
  ): Promise<MsgCreateTimedoutInvoiceResponse> {
    const data = MsgCreateTimedoutInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "CreateTimedoutInvoice",
      data
    );
    return promise.then((data) =>
      MsgCreateTimedoutInvoiceResponse.decode(new Reader(data))
    );
  }

  UpdateTimedoutInvoice(
    request: MsgUpdateTimedoutInvoice
  ): Promise<MsgUpdateTimedoutInvoiceResponse> {
    const data = MsgUpdateTimedoutInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "UpdateTimedoutInvoice",
      data
    );
    return promise.then((data) =>
      MsgUpdateTimedoutInvoiceResponse.decode(new Reader(data))
    );
  }

  DeleteTimedoutInvoice(
    request: MsgDeleteTimedoutInvoice
  ): Promise<MsgDeleteTimedoutInvoiceResponse> {
    const data = MsgDeleteTimedoutInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "DeleteTimedoutInvoice",
      data
    );
    return promise.then((data) =>
      MsgDeleteTimedoutInvoiceResponse.decode(new Reader(data))
    );
  }

  CreateInvoice(request: MsgCreateInvoice): Promise<MsgCreateInvoiceResponse> {
    const data = MsgCreateInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "CreateInvoice",
      data
    );
    return promise.then((data) =>
      MsgCreateInvoiceResponse.decode(new Reader(data))
    );
  }

  PayInvoice(request: MsgPayInvoice): Promise<MsgPayInvoiceResponse> {
    const data = MsgPayInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "PayInvoice",
      data
    );
    return promise.then((data) =>
      MsgPayInvoiceResponse.decode(new Reader(data))
    );
  }

  VoidInvoice(request: MsgVoidInvoice): Promise<MsgVoidInvoiceResponse> {
    const data = MsgVoidInvoice.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Msg",
      "VoidInvoice",
      data
    );
    return promise.then((data) =>
      MsgVoidInvoiceResponse.decode(new Reader(data))
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
