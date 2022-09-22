/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "stateset.core.purchaseorder";

export interface Purchaseorder {
  id: number;
  did: string;
  uri: string;
  amount: string;
  state: string;
  purchaser: string;
  seller: string;
  financer: string;
}

const basePurchaseorder: object = {
  id: 0,
  did: "",
  uri: "",
  amount: "",
  state: "",
  purchaser: "",
  seller: "",
  financer: "",
};

export const Purchaseorder = {
  encode(message: Purchaseorder, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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
    if (message.purchaser !== "") {
      writer.uint32(50).string(message.purchaser);
    }
    if (message.seller !== "") {
      writer.uint32(58).string(message.seller);
    }
    if (message.financer !== "") {
      writer.uint32(66).string(message.financer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Purchaseorder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePurchaseorder } as Purchaseorder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
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
        case 6:
          message.purchaser = reader.string();
          break;
        case 7:
          message.seller = reader.string();
          break;
        case 8:
          message.financer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Purchaseorder {
    const message = { ...basePurchaseorder } as Purchaseorder;
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
    if (object.purchaser !== undefined && object.purchaser !== null) {
      message.purchaser = String(object.purchaser);
    } else {
      message.purchaser = "";
    }
    if (object.seller !== undefined && object.seller !== null) {
      message.seller = String(object.seller);
    } else {
      message.seller = "";
    }
    if (object.financer !== undefined && object.financer !== null) {
      message.financer = String(object.financer);
    } else {
      message.financer = "";
    }
    return message;
  },

  toJSON(message: Purchaseorder): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.did !== undefined && (obj.did = message.did);
    message.uri !== undefined && (obj.uri = message.uri);
    message.amount !== undefined && (obj.amount = message.amount);
    message.state !== undefined && (obj.state = message.state);
    message.purchaser !== undefined && (obj.purchaser = message.purchaser);
    message.seller !== undefined && (obj.seller = message.seller);
    message.financer !== undefined && (obj.financer = message.financer);
    return obj;
  },

  fromPartial(object: DeepPartial<Purchaseorder>): Purchaseorder {
    const message = { ...basePurchaseorder } as Purchaseorder;
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
    if (object.purchaser !== undefined && object.purchaser !== null) {
      message.purchaser = object.purchaser;
    } else {
      message.purchaser = "";
    }
    if (object.seller !== undefined && object.seller !== null) {
      message.seller = object.seller;
    } else {
      message.seller = "";
    }
    if (object.financer !== undefined && object.financer !== null) {
      message.financer = object.financer;
    } else {
      message.financer = "";
    }
    return message;
  },
};

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
