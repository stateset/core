/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Invoice } from "../invoice/invoice";
import { SentInvoice } from "../invoice/sent_invoice";
import { TimedoutInvoice } from "../invoice/timedout_invoice";

export const protobufPackage = "stateset.core.invoice";

/** GenesisState defines the invoice module's genesis state. */
export interface GenesisState {
  invoiceList: Invoice[];
  invoiceCount: number;
  sentInvoiceList: SentInvoice[];
  sentInvoiceCount: number;
  timedoutInvoiceList: TimedoutInvoice[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  timedoutInvoiceCount: number;
}

const baseGenesisState: object = {
  invoiceCount: 0,
  sentInvoiceCount: 0,
  timedoutInvoiceCount: 0,
};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.invoiceList) {
      Invoice.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.invoiceCount !== 0) {
      writer.uint32(16).uint64(message.invoiceCount);
    }
    for (const v of message.sentInvoiceList) {
      SentInvoice.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.sentInvoiceCount !== 0) {
      writer.uint32(32).uint64(message.sentInvoiceCount);
    }
    for (const v of message.timedoutInvoiceList) {
      TimedoutInvoice.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.timedoutInvoiceCount !== 0) {
      writer.uint32(48).uint64(message.timedoutInvoiceCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.invoiceList = [];
    message.sentInvoiceList = [];
    message.timedoutInvoiceList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.invoiceList.push(Invoice.decode(reader, reader.uint32()));
          break;
        case 2:
          message.invoiceCount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.sentInvoiceList.push(
            SentInvoice.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.sentInvoiceCount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.timedoutInvoiceList.push(
            TimedoutInvoice.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.timedoutInvoiceCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.invoiceList = [];
    message.sentInvoiceList = [];
    message.timedoutInvoiceList = [];
    if (object.invoiceList !== undefined && object.invoiceList !== null) {
      for (const e of object.invoiceList) {
        message.invoiceList.push(Invoice.fromJSON(e));
      }
    }
    if (object.invoiceCount !== undefined && object.invoiceCount !== null) {
      message.invoiceCount = Number(object.invoiceCount);
    } else {
      message.invoiceCount = 0;
    }
    if (
      object.sentInvoiceList !== undefined &&
      object.sentInvoiceList !== null
    ) {
      for (const e of object.sentInvoiceList) {
        message.sentInvoiceList.push(SentInvoice.fromJSON(e));
      }
    }
    if (
      object.sentInvoiceCount !== undefined &&
      object.sentInvoiceCount !== null
    ) {
      message.sentInvoiceCount = Number(object.sentInvoiceCount);
    } else {
      message.sentInvoiceCount = 0;
    }
    if (
      object.timedoutInvoiceList !== undefined &&
      object.timedoutInvoiceList !== null
    ) {
      for (const e of object.timedoutInvoiceList) {
        message.timedoutInvoiceList.push(TimedoutInvoice.fromJSON(e));
      }
    }
    if (
      object.timedoutInvoiceCount !== undefined &&
      object.timedoutInvoiceCount !== null
    ) {
      message.timedoutInvoiceCount = Number(object.timedoutInvoiceCount);
    } else {
      message.timedoutInvoiceCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.invoiceList) {
      obj.invoiceList = message.invoiceList.map((e) =>
        e ? Invoice.toJSON(e) : undefined
      );
    } else {
      obj.invoiceList = [];
    }
    message.invoiceCount !== undefined &&
      (obj.invoiceCount = message.invoiceCount);
    if (message.sentInvoiceList) {
      obj.sentInvoiceList = message.sentInvoiceList.map((e) =>
        e ? SentInvoice.toJSON(e) : undefined
      );
    } else {
      obj.sentInvoiceList = [];
    }
    message.sentInvoiceCount !== undefined &&
      (obj.sentInvoiceCount = message.sentInvoiceCount);
    if (message.timedoutInvoiceList) {
      obj.timedoutInvoiceList = message.timedoutInvoiceList.map((e) =>
        e ? TimedoutInvoice.toJSON(e) : undefined
      );
    } else {
      obj.timedoutInvoiceList = [];
    }
    message.timedoutInvoiceCount !== undefined &&
      (obj.timedoutInvoiceCount = message.timedoutInvoiceCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.invoiceList = [];
    message.sentInvoiceList = [];
    message.timedoutInvoiceList = [];
    if (object.invoiceList !== undefined && object.invoiceList !== null) {
      for (const e of object.invoiceList) {
        message.invoiceList.push(Invoice.fromPartial(e));
      }
    }
    if (object.invoiceCount !== undefined && object.invoiceCount !== null) {
      message.invoiceCount = object.invoiceCount;
    } else {
      message.invoiceCount = 0;
    }
    if (
      object.sentInvoiceList !== undefined &&
      object.sentInvoiceList !== null
    ) {
      for (const e of object.sentInvoiceList) {
        message.sentInvoiceList.push(SentInvoice.fromPartial(e));
      }
    }
    if (
      object.sentInvoiceCount !== undefined &&
      object.sentInvoiceCount !== null
    ) {
      message.sentInvoiceCount = object.sentInvoiceCount;
    } else {
      message.sentInvoiceCount = 0;
    }
    if (
      object.timedoutInvoiceList !== undefined &&
      object.timedoutInvoiceList !== null
    ) {
      for (const e of object.timedoutInvoiceList) {
        message.timedoutInvoiceList.push(TimedoutInvoice.fromPartial(e));
      }
    }
    if (
      object.timedoutInvoiceCount !== undefined &&
      object.timedoutInvoiceCount !== null
    ) {
      message.timedoutInvoiceCount = object.timedoutInvoiceCount;
    } else {
      message.timedoutInvoiceCount = 0;
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
