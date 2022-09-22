/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "stateset.core.purchaseorder";

/** EventCreatePurchaseOrder is an event emitted when an purchaseorder is created. */
export interface EventCreatePurchaseOrder {
  /** purchaseorder_id is the unique ID of purchaseorder */
  purchaseorderId: string;
  /** creator is the creator of the purchaseorder */
  creator: string;
}

/** EventCompleted is an event emitted when an purchaseorder is completed. */
export interface EventCompleted {
  /** agreement_id is the unique ID of purchaseorder */
  purchaseorderId: string;
  /** creator is the creator of the purchaseorder */
  creator: string;
}

/** EventCancelled is an event emitted when an purchaseorder is cancelled. */
export interface EventCancelled {
  /** agreement_id is the unique ID of purchaseorder */
  purchaseorderId: string;
  /** creator is the creator of the purchaseorder */
  creator: string;
}

/** EventFinanced is an event emitted when an purchaseorder is financed. */
export interface EventFinanced {
  /** purchaseorder_id is the unique ID of purchaseorder */
  purchaseorderId: string;
  /** creator is the creator of the purchaseorder */
  creator: string;
}

const baseEventCreatePurchaseOrder: object = {
  purchaseorderId: "",
  creator: "",
};

export const EventCreatePurchaseOrder = {
  encode(
    message: EventCreatePurchaseOrder,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.purchaseorderId !== "") {
      writer.uint32(10).string(message.purchaseorderId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventCreatePurchaseOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventCreatePurchaseOrder,
    } as EventCreatePurchaseOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.purchaseorderId = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventCreatePurchaseOrder {
    const message = {
      ...baseEventCreatePurchaseOrder,
    } as EventCreatePurchaseOrder;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = String(object.purchaseorderId);
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventCreatePurchaseOrder): unknown {
    const obj: any = {};
    message.purchaseorderId !== undefined &&
      (obj.purchaseorderId = message.purchaseorderId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventCreatePurchaseOrder>
  ): EventCreatePurchaseOrder {
    const message = {
      ...baseEventCreatePurchaseOrder,
    } as EventCreatePurchaseOrder;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = object.purchaseorderId;
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventCompleted: object = { purchaseorderId: "", creator: "" };

export const EventCompleted = {
  encode(message: EventCompleted, writer: Writer = Writer.create()): Writer {
    if (message.purchaseorderId !== "") {
      writer.uint32(10).string(message.purchaseorderId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventCompleted {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventCompleted } as EventCompleted;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.purchaseorderId = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventCompleted {
    const message = { ...baseEventCompleted } as EventCompleted;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = String(object.purchaseorderId);
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventCompleted): unknown {
    const obj: any = {};
    message.purchaseorderId !== undefined &&
      (obj.purchaseorderId = message.purchaseorderId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventCompleted>): EventCompleted {
    const message = { ...baseEventCompleted } as EventCompleted;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = object.purchaseorderId;
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventCancelled: object = { purchaseorderId: "", creator: "" };

export const EventCancelled = {
  encode(message: EventCancelled, writer: Writer = Writer.create()): Writer {
    if (message.purchaseorderId !== "") {
      writer.uint32(10).string(message.purchaseorderId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventCancelled {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventCancelled } as EventCancelled;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.purchaseorderId = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventCancelled {
    const message = { ...baseEventCancelled } as EventCancelled;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = String(object.purchaseorderId);
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventCancelled): unknown {
    const obj: any = {};
    message.purchaseorderId !== undefined &&
      (obj.purchaseorderId = message.purchaseorderId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventCancelled>): EventCancelled {
    const message = { ...baseEventCancelled } as EventCancelled;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = object.purchaseorderId;
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventFinanced: object = { purchaseorderId: "", creator: "" };

export const EventFinanced = {
  encode(message: EventFinanced, writer: Writer = Writer.create()): Writer {
    if (message.purchaseorderId !== "") {
      writer.uint32(10).string(message.purchaseorderId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventFinanced {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventFinanced } as EventFinanced;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.purchaseorderId = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventFinanced {
    const message = { ...baseEventFinanced } as EventFinanced;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = String(object.purchaseorderId);
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventFinanced): unknown {
    const obj: any = {};
    message.purchaseorderId !== undefined &&
      (obj.purchaseorderId = message.purchaseorderId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventFinanced>): EventFinanced {
    const message = { ...baseEventFinanced } as EventFinanced;
    if (
      object.purchaseorderId !== undefined &&
      object.purchaseorderId !== null
    ) {
      message.purchaseorderId = object.purchaseorderId;
    } else {
      message.purchaseorderId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

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
