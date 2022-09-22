/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { SentAgreement } from "../agreement/sent_agreement";
import { TimedoutAgreement } from "../agreement/timedout_agreement";
import { Agreement } from "../agreement/agreement";

export const protobufPackage = "stateset.core.agreement";

/** GenesisState defines the agreement module's genesis state. */
export interface GenesisState {
  sentAgreementList: SentAgreement[];
  sentAgreementCount: number;
  timedoutAgreementList: TimedoutAgreement[];
  timedoutAgreementCount: number;
  agreementList: Agreement[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  agreementCount: number;
}

const baseGenesisState: object = {
  sentAgreementCount: 0,
  timedoutAgreementCount: 0,
  agreementCount: 0,
};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.sentAgreementList) {
      SentAgreement.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.sentAgreementCount !== 0) {
      writer.uint32(16).uint64(message.sentAgreementCount);
    }
    for (const v of message.timedoutAgreementList) {
      TimedoutAgreement.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.timedoutAgreementCount !== 0) {
      writer.uint32(32).uint64(message.timedoutAgreementCount);
    }
    for (const v of message.agreementList) {
      Agreement.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.agreementCount !== 0) {
      writer.uint32(48).uint64(message.agreementCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.sentAgreementList = [];
    message.timedoutAgreementList = [];
    message.agreementList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sentAgreementList.push(
            SentAgreement.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.sentAgreementCount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.timedoutAgreementList.push(
            TimedoutAgreement.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.timedoutAgreementCount = longToNumber(
            reader.uint64() as Long
          );
          break;
        case 5:
          message.agreementList.push(Agreement.decode(reader, reader.uint32()));
          break;
        case 6:
          message.agreementCount = longToNumber(reader.uint64() as Long);
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
    message.sentAgreementList = [];
    message.timedoutAgreementList = [];
    message.agreementList = [];
    if (
      object.sentAgreementList !== undefined &&
      object.sentAgreementList !== null
    ) {
      for (const e of object.sentAgreementList) {
        message.sentAgreementList.push(SentAgreement.fromJSON(e));
      }
    }
    if (
      object.sentAgreementCount !== undefined &&
      object.sentAgreementCount !== null
    ) {
      message.sentAgreementCount = Number(object.sentAgreementCount);
    } else {
      message.sentAgreementCount = 0;
    }
    if (
      object.timedoutAgreementList !== undefined &&
      object.timedoutAgreementList !== null
    ) {
      for (const e of object.timedoutAgreementList) {
        message.timedoutAgreementList.push(TimedoutAgreement.fromJSON(e));
      }
    }
    if (
      object.timedoutAgreementCount !== undefined &&
      object.timedoutAgreementCount !== null
    ) {
      message.timedoutAgreementCount = Number(object.timedoutAgreementCount);
    } else {
      message.timedoutAgreementCount = 0;
    }
    if (object.agreementList !== undefined && object.agreementList !== null) {
      for (const e of object.agreementList) {
        message.agreementList.push(Agreement.fromJSON(e));
      }
    }
    if (object.agreementCount !== undefined && object.agreementCount !== null) {
      message.agreementCount = Number(object.agreementCount);
    } else {
      message.agreementCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.sentAgreementList) {
      obj.sentAgreementList = message.sentAgreementList.map((e) =>
        e ? SentAgreement.toJSON(e) : undefined
      );
    } else {
      obj.sentAgreementList = [];
    }
    message.sentAgreementCount !== undefined &&
      (obj.sentAgreementCount = message.sentAgreementCount);
    if (message.timedoutAgreementList) {
      obj.timedoutAgreementList = message.timedoutAgreementList.map((e) =>
        e ? TimedoutAgreement.toJSON(e) : undefined
      );
    } else {
      obj.timedoutAgreementList = [];
    }
    message.timedoutAgreementCount !== undefined &&
      (obj.timedoutAgreementCount = message.timedoutAgreementCount);
    if (message.agreementList) {
      obj.agreementList = message.agreementList.map((e) =>
        e ? Agreement.toJSON(e) : undefined
      );
    } else {
      obj.agreementList = [];
    }
    message.agreementCount !== undefined &&
      (obj.agreementCount = message.agreementCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.sentAgreementList = [];
    message.timedoutAgreementList = [];
    message.agreementList = [];
    if (
      object.sentAgreementList !== undefined &&
      object.sentAgreementList !== null
    ) {
      for (const e of object.sentAgreementList) {
        message.sentAgreementList.push(SentAgreement.fromPartial(e));
      }
    }
    if (
      object.sentAgreementCount !== undefined &&
      object.sentAgreementCount !== null
    ) {
      message.sentAgreementCount = object.sentAgreementCount;
    } else {
      message.sentAgreementCount = 0;
    }
    if (
      object.timedoutAgreementList !== undefined &&
      object.timedoutAgreementList !== null
    ) {
      for (const e of object.timedoutAgreementList) {
        message.timedoutAgreementList.push(TimedoutAgreement.fromPartial(e));
      }
    }
    if (
      object.timedoutAgreementCount !== undefined &&
      object.timedoutAgreementCount !== null
    ) {
      message.timedoutAgreementCount = object.timedoutAgreementCount;
    } else {
      message.timedoutAgreementCount = 0;
    }
    if (object.agreementList !== undefined && object.agreementList !== null) {
      for (const e of object.agreementList) {
        message.agreementList.push(Agreement.fromPartial(e));
      }
    }
    if (object.agreementCount !== undefined && object.agreementCount !== null) {
      message.agreementCount = object.agreementCount;
    } else {
      message.agreementCount = 0;
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
