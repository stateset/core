/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "stateset.core.refund";

export interface MsgRequestRefund {
  creator: string;
  did: string;
  amount: string;
  fee: string;
  deadline: string;
}

export interface MsgRequestRefundResponse {}

export interface MsgApproveRefund {
  creator: string;
  id: number;
}

export interface MsgApproveRefundResponse {}

export interface MsgRejectRefund {
  creator: string;
  id: number;
}

export interface MsgRejectRefundResponse {}

const baseMsgRequestRefund: object = {
  creator: "",
  did: "",
  amount: "",
  fee: "",
  deadline: "",
};

export const MsgRequestRefund = {
  encode(message: MsgRequestRefund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.did !== "") {
      writer.uint32(18).string(message.did);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.fee !== "") {
      writer.uint32(34).string(message.fee);
    }
    if (message.deadline !== "") {
      writer.uint32(42).string(message.deadline);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestRefund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestRefund } as MsgRequestRefund;
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
          message.amount = reader.string();
          break;
        case 4:
          message.fee = reader.string();
          break;
        case 5:
          message.deadline = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestRefund {
    const message = { ...baseMsgRequestRefund } as MsgRequestRefund;
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
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = String(object.deadline);
    } else {
      message.deadline = "";
    }
    return message;
  },

  toJSON(message: MsgRequestRefund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.did !== undefined && (obj.did = message.did);
    message.amount !== undefined && (obj.amount = message.amount);
    message.fee !== undefined && (obj.fee = message.fee);
    message.deadline !== undefined && (obj.deadline = message.deadline);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRequestRefund>): MsgRequestRefund {
    const message = { ...baseMsgRequestRefund } as MsgRequestRefund;
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
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = object.deadline;
    } else {
      message.deadline = "";
    }
    return message;
  },
};

const baseMsgRequestRefundResponse: object = {};

export const MsgRequestRefundResponse = {
  encode(
    _: MsgRequestRefundResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRequestRefundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRequestRefundResponse,
    } as MsgRequestRefundResponse;
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

  fromJSON(_: any): MsgRequestRefundResponse {
    const message = {
      ...baseMsgRequestRefundResponse,
    } as MsgRequestRefundResponse;
    return message;
  },

  toJSON(_: MsgRequestRefundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRequestRefundResponse>
  ): MsgRequestRefundResponse {
    const message = {
      ...baseMsgRequestRefundResponse,
    } as MsgRequestRefundResponse;
    return message;
  },
};

const baseMsgApproveRefund: object = { creator: "", id: 0 };

export const MsgApproveRefund = {
  encode(message: MsgApproveRefund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveRefund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApproveRefund } as MsgApproveRefund;
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

  fromJSON(object: any): MsgApproveRefund {
    const message = { ...baseMsgApproveRefund } as MsgApproveRefund;
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

  toJSON(message: MsgApproveRefund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgApproveRefund>): MsgApproveRefund {
    const message = { ...baseMsgApproveRefund } as MsgApproveRefund;
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

const baseMsgApproveRefundResponse: object = {};

export const MsgApproveRefundResponse = {
  encode(
    _: MsgApproveRefundResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgApproveRefundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgApproveRefundResponse,
    } as MsgApproveRefundResponse;
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

  fromJSON(_: any): MsgApproveRefundResponse {
    const message = {
      ...baseMsgApproveRefundResponse,
    } as MsgApproveRefundResponse;
    return message;
  },

  toJSON(_: MsgApproveRefundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgApproveRefundResponse>
  ): MsgApproveRefundResponse {
    const message = {
      ...baseMsgApproveRefundResponse,
    } as MsgApproveRefundResponse;
    return message;
  },
};

const baseMsgRejectRefund: object = { creator: "", id: 0 };

export const MsgRejectRefund = {
  encode(message: MsgRejectRefund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectRefund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRejectRefund } as MsgRejectRefund;
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

  fromJSON(object: any): MsgRejectRefund {
    const message = { ...baseMsgRejectRefund } as MsgRejectRefund;
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

  toJSON(message: MsgRejectRefund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRejectRefund>): MsgRejectRefund {
    const message = { ...baseMsgRejectRefund } as MsgRejectRefund;
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

const baseMsgRejectRefundResponse: object = {};

export const MsgRejectRefundResponse = {
  encode(_: MsgRejectRefundResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRejectRefundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRejectRefundResponse,
    } as MsgRejectRefundResponse;
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

  fromJSON(_: any): MsgRejectRefundResponse {
    const message = {
      ...baseMsgRejectRefundResponse,
    } as MsgRejectRefundResponse;
    return message;
  },

  toJSON(_: MsgRejectRefundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRejectRefundResponse>
  ): MsgRejectRefundResponse {
    const message = {
      ...baseMsgRejectRefundResponse,
    } as MsgRejectRefundResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RequestRefund(request: MsgRequestRefund): Promise<MsgRequestRefundResponse>;
  ApproveRefund(request: MsgApproveRefund): Promise<MsgApproveRefundResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RejectRefund(request: MsgRejectRefund): Promise<MsgRejectRefundResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RequestRefund(request: MsgRequestRefund): Promise<MsgRequestRefundResponse> {
    const data = MsgRequestRefund.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.refund.Msg",
      "RequestRefund",
      data
    );
    return promise.then((data) =>
      MsgRequestRefundResponse.decode(new Reader(data))
    );
  }

  ApproveRefund(request: MsgApproveRefund): Promise<MsgApproveRefundResponse> {
    const data = MsgApproveRefund.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.refund.Msg",
      "ApproveRefund",
      data
    );
    return promise.then((data) =>
      MsgApproveRefundResponse.decode(new Reader(data))
    );
  }

  RejectRefund(request: MsgRejectRefund): Promise<MsgRejectRefundResponse> {
    const data = MsgRejectRefund.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.refund.Msg",
      "RejectRefund",
      data
    );
    return promise.then((data) =>
      MsgRejectRefundResponse.decode(new Reader(data))
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
