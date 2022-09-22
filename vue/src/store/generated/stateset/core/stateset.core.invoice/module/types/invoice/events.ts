/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "stateset.core.invoice";

/** EventCreateInvoice is an event emitted when an invoice is created. */
export interface EventCreateInvoice {
  /** invoice_id is the unique ID of invoice */
  invoiceId: string;
  /** creator is the creator of the transaction */
  creator: string;
}

/** EventPaid is an event emitted when an invoice is paid. */
export interface EventPaid {
  /** agreement_id is the unique ID of invoice */
  invoiceId: string;
  /** creator is the creator of the invoice */
  creator: string;
}

/** EventVoided is an event emitted when an invoice is voided. */
export interface EventVoided {
  /** agreement_id is the unique ID of invoice */
  invoiceId: string;
  /** creator is the creator of the transaction */
  creator: string;
}

/** EventFactored is an event emitted when an invoice is factored. */
export interface EventFactored {
  /** invoice_id is the unique ID of invoice */
  invoiceId: string;
  /** creator is the creator of the transaction */
  creator: string;
}

const baseEventCreateInvoice: object = { invoiceId: "", creator: "" };

export const EventCreateInvoice = {
  encode(
    message: EventCreateInvoice,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.invoiceId !== "") {
      writer.uint32(10).string(message.invoiceId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventCreateInvoice {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventCreateInvoice } as EventCreateInvoice;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.invoiceId = reader.string();
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

  fromJSON(object: any): EventCreateInvoice {
    const message = { ...baseEventCreateInvoice } as EventCreateInvoice;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = String(object.invoiceId);
    } else {
      message.invoiceId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventCreateInvoice): unknown {
    const obj: any = {};
    message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventCreateInvoice>): EventCreateInvoice {
    const message = { ...baseEventCreateInvoice } as EventCreateInvoice;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = object.invoiceId;
    } else {
      message.invoiceId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventPaid: object = { invoiceId: "", creator: "" };

export const EventPaid = {
  encode(message: EventPaid, writer: Writer = Writer.create()): Writer {
    if (message.invoiceId !== "") {
      writer.uint32(10).string(message.invoiceId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventPaid {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventPaid } as EventPaid;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.invoiceId = reader.string();
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

  fromJSON(object: any): EventPaid {
    const message = { ...baseEventPaid } as EventPaid;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = String(object.invoiceId);
    } else {
      message.invoiceId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventPaid): unknown {
    const obj: any = {};
    message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventPaid>): EventPaid {
    const message = { ...baseEventPaid } as EventPaid;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = object.invoiceId;
    } else {
      message.invoiceId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventVoided: object = { invoiceId: "", creator: "" };

export const EventVoided = {
  encode(message: EventVoided, writer: Writer = Writer.create()): Writer {
    if (message.invoiceId !== "") {
      writer.uint32(10).string(message.invoiceId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventVoided {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventVoided } as EventVoided;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.invoiceId = reader.string();
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

  fromJSON(object: any): EventVoided {
    const message = { ...baseEventVoided } as EventVoided;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = String(object.invoiceId);
    } else {
      message.invoiceId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventVoided): unknown {
    const obj: any = {};
    message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventVoided>): EventVoided {
    const message = { ...baseEventVoided } as EventVoided;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = object.invoiceId;
    } else {
      message.invoiceId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventFactored: object = { invoiceId: "", creator: "" };

export const EventFactored = {
  encode(message: EventFactored, writer: Writer = Writer.create()): Writer {
    if (message.invoiceId !== "") {
      writer.uint32(10).string(message.invoiceId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventFactored {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventFactored } as EventFactored;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.invoiceId = reader.string();
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

  fromJSON(object: any): EventFactored {
    const message = { ...baseEventFactored } as EventFactored;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = String(object.invoiceId);
    } else {
      message.invoiceId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventFactored): unknown {
    const obj: any = {};
    message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventFactored>): EventFactored {
    const message = { ...baseEventFactored } as EventFactored;
    if (object.invoiceId !== undefined && object.invoiceId !== null) {
      message.invoiceId = object.invoiceId;
    } else {
      message.invoiceId = "";
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
