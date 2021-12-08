/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "stateset.core.purchaseorder";

export interface MsgFinancePurchaseorder {
  creator: string;
  id: number;
}

export interface MsgFinancePurchaseorderResponse {}

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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  FinancePurchaseorder(
    request: MsgFinancePurchaseorder
  ): Promise<MsgFinancePurchaseorderResponse>;
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
