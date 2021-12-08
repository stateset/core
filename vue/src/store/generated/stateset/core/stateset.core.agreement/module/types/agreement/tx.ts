/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "stateset.core.agreement";

export interface MsgActivateAgreement {
  creator: string;
  id: number;
}

export interface MsgActivateAgreementResponse {}

export interface MsgExpireAgreement {
  creator: string;
  id: number;
}

export interface MsgExpireAgreementResponse {}

export interface MsgRenewAgreement {
  creator: string;
  id: number;
}

export interface MsgRenewAgreementResponse {}

export interface MsgTerminateAgreement {
  creator: string;
  id: number;
}

export interface MsgTerminateAgreementResponse {}

export interface MsgCreateSentAgreement {
  creator: string;
  did: string;
  chain: string;
}

export interface MsgCreateSentAgreementResponse {
  id: number;
}

export interface MsgUpdateSentAgreement {
  creator: string;
  id: number;
  did: string;
  chain: string;
}

export interface MsgUpdateSentAgreementResponse {}

export interface MsgDeleteSentAgreement {
  creator: string;
  id: number;
}

export interface MsgDeleteSentAgreementResponse {}

export interface MsgCreateTimedoutAgreement {
  creator: string;
  did: string;
  chain: string;
}

export interface MsgCreateTimedoutAgreementResponse {
  id: number;
}

export interface MsgUpdateTimedoutAgreement {
  creator: string;
  id: number;
  did: string;
  chain: string;
}

export interface MsgUpdateTimedoutAgreementResponse {}

export interface MsgDeleteTimedoutAgreement {
  creator: string;
  id: number;
}

export interface MsgDeleteTimedoutAgreementResponse {}

const baseMsgActivateAgreement: object = { creator: "", id: 0 };

export const MsgActivateAgreement = {
  encode(
    message: MsgActivateAgreement,
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

  decode(input: Reader | Uint8Array, length?: number): MsgActivateAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgActivateAgreement } as MsgActivateAgreement;
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

  fromJSON(object: any): MsgActivateAgreement {
    const message = { ...baseMsgActivateAgreement } as MsgActivateAgreement;
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

  toJSON(message: MsgActivateAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgActivateAgreement>): MsgActivateAgreement {
    const message = { ...baseMsgActivateAgreement } as MsgActivateAgreement;
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

const baseMsgActivateAgreementResponse: object = {};

export const MsgActivateAgreementResponse = {
  encode(
    _: MsgActivateAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgActivateAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgActivateAgreementResponse,
    } as MsgActivateAgreementResponse;
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

  fromJSON(_: any): MsgActivateAgreementResponse {
    const message = {
      ...baseMsgActivateAgreementResponse,
    } as MsgActivateAgreementResponse;
    return message;
  },

  toJSON(_: MsgActivateAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgActivateAgreementResponse>
  ): MsgActivateAgreementResponse {
    const message = {
      ...baseMsgActivateAgreementResponse,
    } as MsgActivateAgreementResponse;
    return message;
  },
};

const baseMsgExpireAgreement: object = { creator: "", id: 0 };

export const MsgExpireAgreement = {
  encode(
    message: MsgExpireAgreement,
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

  decode(input: Reader | Uint8Array, length?: number): MsgExpireAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgExpireAgreement } as MsgExpireAgreement;
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

  fromJSON(object: any): MsgExpireAgreement {
    const message = { ...baseMsgExpireAgreement } as MsgExpireAgreement;
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

  toJSON(message: MsgExpireAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgExpireAgreement>): MsgExpireAgreement {
    const message = { ...baseMsgExpireAgreement } as MsgExpireAgreement;
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

const baseMsgExpireAgreementResponse: object = {};

export const MsgExpireAgreementResponse = {
  encode(
    _: MsgExpireAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgExpireAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgExpireAgreementResponse,
    } as MsgExpireAgreementResponse;
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

  fromJSON(_: any): MsgExpireAgreementResponse {
    const message = {
      ...baseMsgExpireAgreementResponse,
    } as MsgExpireAgreementResponse;
    return message;
  },

  toJSON(_: MsgExpireAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgExpireAgreementResponse>
  ): MsgExpireAgreementResponse {
    const message = {
      ...baseMsgExpireAgreementResponse,
    } as MsgExpireAgreementResponse;
    return message;
  },
};

const baseMsgRenewAgreement: object = { creator: "", id: 0 };

export const MsgRenewAgreement = {
  encode(message: MsgRenewAgreement, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRenewAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRenewAgreement } as MsgRenewAgreement;
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

  fromJSON(object: any): MsgRenewAgreement {
    const message = { ...baseMsgRenewAgreement } as MsgRenewAgreement;
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

  toJSON(message: MsgRenewAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRenewAgreement>): MsgRenewAgreement {
    const message = { ...baseMsgRenewAgreement } as MsgRenewAgreement;
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

const baseMsgRenewAgreementResponse: object = {};

export const MsgRenewAgreementResponse = {
  encode(
    _: MsgRenewAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRenewAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRenewAgreementResponse,
    } as MsgRenewAgreementResponse;
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

  fromJSON(_: any): MsgRenewAgreementResponse {
    const message = {
      ...baseMsgRenewAgreementResponse,
    } as MsgRenewAgreementResponse;
    return message;
  },

  toJSON(_: MsgRenewAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRenewAgreementResponse>
  ): MsgRenewAgreementResponse {
    const message = {
      ...baseMsgRenewAgreementResponse,
    } as MsgRenewAgreementResponse;
    return message;
  },
};

const baseMsgTerminateAgreement: object = { creator: "", id: 0 };

export const MsgTerminateAgreement = {
  encode(
    message: MsgTerminateAgreement,
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

  decode(input: Reader | Uint8Array, length?: number): MsgTerminateAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgTerminateAgreement } as MsgTerminateAgreement;
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

  fromJSON(object: any): MsgTerminateAgreement {
    const message = { ...baseMsgTerminateAgreement } as MsgTerminateAgreement;
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

  toJSON(message: MsgTerminateAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgTerminateAgreement>
  ): MsgTerminateAgreement {
    const message = { ...baseMsgTerminateAgreement } as MsgTerminateAgreement;
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

const baseMsgTerminateAgreementResponse: object = {};

export const MsgTerminateAgreementResponse = {
  encode(
    _: MsgTerminateAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgTerminateAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgTerminateAgreementResponse,
    } as MsgTerminateAgreementResponse;
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

  fromJSON(_: any): MsgTerminateAgreementResponse {
    const message = {
      ...baseMsgTerminateAgreementResponse,
    } as MsgTerminateAgreementResponse;
    return message;
  },

  toJSON(_: MsgTerminateAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgTerminateAgreementResponse>
  ): MsgTerminateAgreementResponse {
    const message = {
      ...baseMsgTerminateAgreementResponse,
    } as MsgTerminateAgreementResponse;
    return message;
  },
};

const baseMsgCreateSentAgreement: object = { creator: "", did: "", chain: "" };

export const MsgCreateSentAgreement = {
  encode(
    message: MsgCreateSentAgreement,
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

  decode(input: Reader | Uint8Array, length?: number): MsgCreateSentAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateSentAgreement } as MsgCreateSentAgreement;
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

  fromJSON(object: any): MsgCreateSentAgreement {
    const message = { ...baseMsgCreateSentAgreement } as MsgCreateSentAgreement;
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

  toJSON(message: MsgCreateSentAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSentAgreement>
  ): MsgCreateSentAgreement {
    const message = { ...baseMsgCreateSentAgreement } as MsgCreateSentAgreement;
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

const baseMsgCreateSentAgreementResponse: object = { id: 0 };

export const MsgCreateSentAgreementResponse = {
  encode(
    message: MsgCreateSentAgreementResponse,
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
  ): MsgCreateSentAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateSentAgreementResponse,
    } as MsgCreateSentAgreementResponse;
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

  fromJSON(object: any): MsgCreateSentAgreementResponse {
    const message = {
      ...baseMsgCreateSentAgreementResponse,
    } as MsgCreateSentAgreementResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateSentAgreementResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateSentAgreementResponse>
  ): MsgCreateSentAgreementResponse {
    const message = {
      ...baseMsgCreateSentAgreementResponse,
    } as MsgCreateSentAgreementResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateSentAgreement: object = {
  creator: "",
  id: 0,
  did: "",
  chain: "",
};

export const MsgUpdateSentAgreement = {
  encode(
    message: MsgUpdateSentAgreement,
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

  decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdateSentAgreement } as MsgUpdateSentAgreement;
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

  fromJSON(object: any): MsgUpdateSentAgreement {
    const message = { ...baseMsgUpdateSentAgreement } as MsgUpdateSentAgreement;
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

  toJSON(message: MsgUpdateSentAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateSentAgreement>
  ): MsgUpdateSentAgreement {
    const message = { ...baseMsgUpdateSentAgreement } as MsgUpdateSentAgreement;
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

const baseMsgUpdateSentAgreementResponse: object = {};

export const MsgUpdateSentAgreementResponse = {
  encode(
    _: MsgUpdateSentAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateSentAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateSentAgreementResponse,
    } as MsgUpdateSentAgreementResponse;
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

  fromJSON(_: any): MsgUpdateSentAgreementResponse {
    const message = {
      ...baseMsgUpdateSentAgreementResponse,
    } as MsgUpdateSentAgreementResponse;
    return message;
  },

  toJSON(_: MsgUpdateSentAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateSentAgreementResponse>
  ): MsgUpdateSentAgreementResponse {
    const message = {
      ...baseMsgUpdateSentAgreementResponse,
    } as MsgUpdateSentAgreementResponse;
    return message;
  },
};

const baseMsgDeleteSentAgreement: object = { creator: "", id: 0 };

export const MsgDeleteSentAgreement = {
  encode(
    message: MsgDeleteSentAgreement,
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

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteSentAgreement } as MsgDeleteSentAgreement;
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

  fromJSON(object: any): MsgDeleteSentAgreement {
    const message = { ...baseMsgDeleteSentAgreement } as MsgDeleteSentAgreement;
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

  toJSON(message: MsgDeleteSentAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteSentAgreement>
  ): MsgDeleteSentAgreement {
    const message = { ...baseMsgDeleteSentAgreement } as MsgDeleteSentAgreement;
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

const baseMsgDeleteSentAgreementResponse: object = {};

export const MsgDeleteSentAgreementResponse = {
  encode(
    _: MsgDeleteSentAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteSentAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteSentAgreementResponse,
    } as MsgDeleteSentAgreementResponse;
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

  fromJSON(_: any): MsgDeleteSentAgreementResponse {
    const message = {
      ...baseMsgDeleteSentAgreementResponse,
    } as MsgDeleteSentAgreementResponse;
    return message;
  },

  toJSON(_: MsgDeleteSentAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteSentAgreementResponse>
  ): MsgDeleteSentAgreementResponse {
    const message = {
      ...baseMsgDeleteSentAgreementResponse,
    } as MsgDeleteSentAgreementResponse;
    return message;
  },
};

const baseMsgCreateTimedoutAgreement: object = {
  creator: "",
  did: "",
  chain: "",
};

export const MsgCreateTimedoutAgreement = {
  encode(
    message: MsgCreateTimedoutAgreement,
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
  ): MsgCreateTimedoutAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTimedoutAgreement,
    } as MsgCreateTimedoutAgreement;
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

  fromJSON(object: any): MsgCreateTimedoutAgreement {
    const message = {
      ...baseMsgCreateTimedoutAgreement,
    } as MsgCreateTimedoutAgreement;
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

  toJSON(message: MsgCreateTimedoutAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTimedoutAgreement>
  ): MsgCreateTimedoutAgreement {
    const message = {
      ...baseMsgCreateTimedoutAgreement,
    } as MsgCreateTimedoutAgreement;
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

const baseMsgCreateTimedoutAgreementResponse: object = { id: 0 };

export const MsgCreateTimedoutAgreementResponse = {
  encode(
    message: MsgCreateTimedoutAgreementResponse,
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
  ): MsgCreateTimedoutAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateTimedoutAgreementResponse,
    } as MsgCreateTimedoutAgreementResponse;
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

  fromJSON(object: any): MsgCreateTimedoutAgreementResponse {
    const message = {
      ...baseMsgCreateTimedoutAgreementResponse,
    } as MsgCreateTimedoutAgreementResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCreateTimedoutAgreementResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateTimedoutAgreementResponse>
  ): MsgCreateTimedoutAgreementResponse {
    const message = {
      ...baseMsgCreateTimedoutAgreementResponse,
    } as MsgCreateTimedoutAgreementResponse;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgUpdateTimedoutAgreement: object = {
  creator: "",
  id: 0,
  did: "",
  chain: "",
};

export const MsgUpdateTimedoutAgreement = {
  encode(
    message: MsgUpdateTimedoutAgreement,
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
  ): MsgUpdateTimedoutAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateTimedoutAgreement,
    } as MsgUpdateTimedoutAgreement;
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

  fromJSON(object: any): MsgUpdateTimedoutAgreement {
    const message = {
      ...baseMsgUpdateTimedoutAgreement,
    } as MsgUpdateTimedoutAgreement;
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

  toJSON(message: MsgUpdateTimedoutAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgUpdateTimedoutAgreement>
  ): MsgUpdateTimedoutAgreement {
    const message = {
      ...baseMsgUpdateTimedoutAgreement,
    } as MsgUpdateTimedoutAgreement;
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

const baseMsgUpdateTimedoutAgreementResponse: object = {};

export const MsgUpdateTimedoutAgreementResponse = {
  encode(
    _: MsgUpdateTimedoutAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgUpdateTimedoutAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdateTimedoutAgreementResponse,
    } as MsgUpdateTimedoutAgreementResponse;
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

  fromJSON(_: any): MsgUpdateTimedoutAgreementResponse {
    const message = {
      ...baseMsgUpdateTimedoutAgreementResponse,
    } as MsgUpdateTimedoutAgreementResponse;
    return message;
  },

  toJSON(_: MsgUpdateTimedoutAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdateTimedoutAgreementResponse>
  ): MsgUpdateTimedoutAgreementResponse {
    const message = {
      ...baseMsgUpdateTimedoutAgreementResponse,
    } as MsgUpdateTimedoutAgreementResponse;
    return message;
  },
};

const baseMsgDeleteTimedoutAgreement: object = { creator: "", id: 0 };

export const MsgDeleteTimedoutAgreement = {
  encode(
    message: MsgDeleteTimedoutAgreement,
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
  ): MsgDeleteTimedoutAgreement {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteTimedoutAgreement,
    } as MsgDeleteTimedoutAgreement;
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

  fromJSON(object: any): MsgDeleteTimedoutAgreement {
    const message = {
      ...baseMsgDeleteTimedoutAgreement,
    } as MsgDeleteTimedoutAgreement;
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

  toJSON(message: MsgDeleteTimedoutAgreement): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteTimedoutAgreement>
  ): MsgDeleteTimedoutAgreement {
    const message = {
      ...baseMsgDeleteTimedoutAgreement,
    } as MsgDeleteTimedoutAgreement;
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

const baseMsgDeleteTimedoutAgreementResponse: object = {};

export const MsgDeleteTimedoutAgreementResponse = {
  encode(
    _: MsgDeleteTimedoutAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteTimedoutAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteTimedoutAgreementResponse,
    } as MsgDeleteTimedoutAgreementResponse;
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

  fromJSON(_: any): MsgDeleteTimedoutAgreementResponse {
    const message = {
      ...baseMsgDeleteTimedoutAgreementResponse,
    } as MsgDeleteTimedoutAgreementResponse;
    return message;
  },

  toJSON(_: MsgDeleteTimedoutAgreementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteTimedoutAgreementResponse>
  ): MsgDeleteTimedoutAgreementResponse {
    const message = {
      ...baseMsgDeleteTimedoutAgreementResponse,
    } as MsgDeleteTimedoutAgreementResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ActivateAgreement(
    request: MsgActivateAgreement
  ): Promise<MsgActivateAgreementResponse>;
  ExpireAgreement(
    request: MsgExpireAgreement
  ): Promise<MsgExpireAgreementResponse>;
  RenewAgreement(
    request: MsgRenewAgreement
  ): Promise<MsgRenewAgreementResponse>;
  TerminateAgreement(
    request: MsgTerminateAgreement
  ): Promise<MsgTerminateAgreementResponse>;
  CreateSentAgreement(
    request: MsgCreateSentAgreement
  ): Promise<MsgCreateSentAgreementResponse>;
  UpdateSentAgreement(
    request: MsgUpdateSentAgreement
  ): Promise<MsgUpdateSentAgreementResponse>;
  DeleteSentAgreement(
    request: MsgDeleteSentAgreement
  ): Promise<MsgDeleteSentAgreementResponse>;
  CreateTimedoutAgreement(
    request: MsgCreateTimedoutAgreement
  ): Promise<MsgCreateTimedoutAgreementResponse>;
  UpdateTimedoutAgreement(
    request: MsgUpdateTimedoutAgreement
  ): Promise<MsgUpdateTimedoutAgreementResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteTimedoutAgreement(
    request: MsgDeleteTimedoutAgreement
  ): Promise<MsgDeleteTimedoutAgreementResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ActivateAgreement(
    request: MsgActivateAgreement
  ): Promise<MsgActivateAgreementResponse> {
    const data = MsgActivateAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "ActivateAgreement",
      data
    );
    return promise.then((data) =>
      MsgActivateAgreementResponse.decode(new Reader(data))
    );
  }

  ExpireAgreement(
    request: MsgExpireAgreement
  ): Promise<MsgExpireAgreementResponse> {
    const data = MsgExpireAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "ExpireAgreement",
      data
    );
    return promise.then((data) =>
      MsgExpireAgreementResponse.decode(new Reader(data))
    );
  }

  RenewAgreement(
    request: MsgRenewAgreement
  ): Promise<MsgRenewAgreementResponse> {
    const data = MsgRenewAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "RenewAgreement",
      data
    );
    return promise.then((data) =>
      MsgRenewAgreementResponse.decode(new Reader(data))
    );
  }

  TerminateAgreement(
    request: MsgTerminateAgreement
  ): Promise<MsgTerminateAgreementResponse> {
    const data = MsgTerminateAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "TerminateAgreement",
      data
    );
    return promise.then((data) =>
      MsgTerminateAgreementResponse.decode(new Reader(data))
    );
  }

  CreateSentAgreement(
    request: MsgCreateSentAgreement
  ): Promise<MsgCreateSentAgreementResponse> {
    const data = MsgCreateSentAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "CreateSentAgreement",
      data
    );
    return promise.then((data) =>
      MsgCreateSentAgreementResponse.decode(new Reader(data))
    );
  }

  UpdateSentAgreement(
    request: MsgUpdateSentAgreement
  ): Promise<MsgUpdateSentAgreementResponse> {
    const data = MsgUpdateSentAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "UpdateSentAgreement",
      data
    );
    return promise.then((data) =>
      MsgUpdateSentAgreementResponse.decode(new Reader(data))
    );
  }

  DeleteSentAgreement(
    request: MsgDeleteSentAgreement
  ): Promise<MsgDeleteSentAgreementResponse> {
    const data = MsgDeleteSentAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "DeleteSentAgreement",
      data
    );
    return promise.then((data) =>
      MsgDeleteSentAgreementResponse.decode(new Reader(data))
    );
  }

  CreateTimedoutAgreement(
    request: MsgCreateTimedoutAgreement
  ): Promise<MsgCreateTimedoutAgreementResponse> {
    const data = MsgCreateTimedoutAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "CreateTimedoutAgreement",
      data
    );
    return promise.then((data) =>
      MsgCreateTimedoutAgreementResponse.decode(new Reader(data))
    );
  }

  UpdateTimedoutAgreement(
    request: MsgUpdateTimedoutAgreement
  ): Promise<MsgUpdateTimedoutAgreementResponse> {
    const data = MsgUpdateTimedoutAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "UpdateTimedoutAgreement",
      data
    );
    return promise.then((data) =>
      MsgUpdateTimedoutAgreementResponse.decode(new Reader(data))
    );
  }

  DeleteTimedoutAgreement(
    request: MsgDeleteTimedoutAgreement
  ): Promise<MsgDeleteTimedoutAgreementResponse> {
    const data = MsgDeleteTimedoutAgreement.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Msg",
      "DeleteTimedoutAgreement",
      data
    );
    return promise.then((data) =>
      MsgDeleteTimedoutAgreementResponse.decode(new Reader(data))
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
