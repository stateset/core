/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "stateset.core.loan";

/** EventCreateLoan is an event emitted when an loan is created. */
export interface EventCreateLoan {
  /** loan_id is the unique ID of loan */
  loanId: string;
  /** creator is the creator of the loan */
  creator: string;
}

/** EventLoanRequested is an event emitted when an loan is requested. */
export interface EventLoanRequested {
  /** loan_id is the unique ID of loan */
  loanId: string;
  /** creator is the creator of the loan */
  creator: string;
}

/** EventApproved is an event emitted when an loan is approved. */
export interface EventApproved {
  /** loan_id is the unique ID of loan */
  loanId: string;
  /** creator is the creator of the loan */
  creator: string;
}

/** EventRepaid is an event emitted when an loan is repaid. */
export interface EventRepaid {
  /** loan_id is the unique ID of loan */
  loanId: string;
  /** creator is the creator of the loan */
  creator: string;
}

/** EventLiquidated is an event emitted when an loan is liquidated. */
export interface EventLiquidated {
  /** loan_id is the unique ID of loan */
  loanId: string;
  /** creator is the creator of the loan */
  creator: string;
}

/** EventCancelled is an event emitted when an loan is cancelled. */
export interface EventCancelled {
  /** loan_id is the unique ID of loan */
  loanId: string;
  /** creator is the creator of the loan */
  creator: string;
}

const baseEventCreateLoan: object = { loanId: "", creator: "" };

export const EventCreateLoan = {
  encode(message: EventCreateLoan, writer: Writer = Writer.create()): Writer {
    if (message.loanId !== "") {
      writer.uint32(10).string(message.loanId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventCreateLoan {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventCreateLoan } as EventCreateLoan;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.loanId = reader.string();
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

  fromJSON(object: any): EventCreateLoan {
    const message = { ...baseEventCreateLoan } as EventCreateLoan;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = String(object.loanId);
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventCreateLoan): unknown {
    const obj: any = {};
    message.loanId !== undefined && (obj.loanId = message.loanId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventCreateLoan>): EventCreateLoan {
    const message = { ...baseEventCreateLoan } as EventCreateLoan;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = object.loanId;
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventLoanRequested: object = { loanId: "", creator: "" };

export const EventLoanRequested = {
  encode(
    message: EventLoanRequested,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.loanId !== "") {
      writer.uint32(10).string(message.loanId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventLoanRequested {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventLoanRequested } as EventLoanRequested;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.loanId = reader.string();
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

  fromJSON(object: any): EventLoanRequested {
    const message = { ...baseEventLoanRequested } as EventLoanRequested;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = String(object.loanId);
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventLoanRequested): unknown {
    const obj: any = {};
    message.loanId !== undefined && (obj.loanId = message.loanId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventLoanRequested>): EventLoanRequested {
    const message = { ...baseEventLoanRequested } as EventLoanRequested;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = object.loanId;
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventApproved: object = { loanId: "", creator: "" };

export const EventApproved = {
  encode(message: EventApproved, writer: Writer = Writer.create()): Writer {
    if (message.loanId !== "") {
      writer.uint32(10).string(message.loanId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventApproved {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventApproved } as EventApproved;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.loanId = reader.string();
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

  fromJSON(object: any): EventApproved {
    const message = { ...baseEventApproved } as EventApproved;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = String(object.loanId);
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventApproved): unknown {
    const obj: any = {};
    message.loanId !== undefined && (obj.loanId = message.loanId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventApproved>): EventApproved {
    const message = { ...baseEventApproved } as EventApproved;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = object.loanId;
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventRepaid: object = { loanId: "", creator: "" };

export const EventRepaid = {
  encode(message: EventRepaid, writer: Writer = Writer.create()): Writer {
    if (message.loanId !== "") {
      writer.uint32(10).string(message.loanId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventRepaid {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventRepaid } as EventRepaid;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.loanId = reader.string();
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

  fromJSON(object: any): EventRepaid {
    const message = { ...baseEventRepaid } as EventRepaid;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = String(object.loanId);
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventRepaid): unknown {
    const obj: any = {};
    message.loanId !== undefined && (obj.loanId = message.loanId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventRepaid>): EventRepaid {
    const message = { ...baseEventRepaid } as EventRepaid;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = object.loanId;
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventLiquidated: object = { loanId: "", creator: "" };

export const EventLiquidated = {
  encode(message: EventLiquidated, writer: Writer = Writer.create()): Writer {
    if (message.loanId !== "") {
      writer.uint32(10).string(message.loanId);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): EventLiquidated {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseEventLiquidated } as EventLiquidated;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.loanId = reader.string();
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

  fromJSON(object: any): EventLiquidated {
    const message = { ...baseEventLiquidated } as EventLiquidated;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = String(object.loanId);
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: EventLiquidated): unknown {
    const obj: any = {};
    message.loanId !== undefined && (obj.loanId = message.loanId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventLiquidated>): EventLiquidated {
    const message = { ...baseEventLiquidated } as EventLiquidated;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = object.loanId;
    } else {
      message.loanId = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseEventCancelled: object = { loanId: "", creator: "" };

export const EventCancelled = {
  encode(message: EventCancelled, writer: Writer = Writer.create()): Writer {
    if (message.loanId !== "") {
      writer.uint32(10).string(message.loanId);
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
          message.loanId = reader.string();
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
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = String(object.loanId);
    } else {
      message.loanId = "";
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
    message.loanId !== undefined && (obj.loanId = message.loanId);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<EventCancelled>): EventCancelled {
    const message = { ...baseEventCancelled } as EventCancelled;
    if (object.loanId !== undefined && object.loanId !== null) {
      message.loanId = object.loanId;
    } else {
      message.loanId = "";
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
