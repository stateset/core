/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { SentPurchaseorder } from "../purchaseorder/sent_purchaseorder";
import { TimedoutPurchaseorder } from "../purchaseorder/timedout_purchaseorder";

export const protobufPackage = "stateset.core.purchaseorder";

export interface QueryGetPurchaseorderRequest {
  id: number;
}

export interface QueryGetPurchaseorderResponse {
  Purchaseorder: Purchaseorder | undefined;
}

export interface QueryAllPurchaseorderRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPurchaseorderResponse {
  Purchaseorder: Purchaseorder[];
  pagination: PageResponse | undefined;
}

export interface QueryGetSentPurchaseorderRequest {
  id: number;
}

export interface QueryGetSentPurchaseorderResponse {
  SentPurchaseorder: SentPurchaseorder | undefined;
}

export interface QueryAllSentPurchaseorderRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSentPurchaseorderResponse {
  SentPurchaseorder: SentPurchaseorder[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTimedoutPurchaseorderRequest {
  id: number;
}

export interface QueryGetTimedoutPurchaseorderResponse {
  TimedoutPurchaseorder: TimedoutPurchaseorder | undefined;
}

export interface QueryAllTimedoutPurchaseorderRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTimedoutPurchaseorderResponse {
  TimedoutPurchaseorder: TimedoutPurchaseorder[];
  pagination: PageResponse | undefined;
}

const baseQueryGetPurchaseorderRequest: object = { id: 0 };

export const QueryGetPurchaseorderRequest = {
  encode(
    message: QueryGetPurchaseorderRequest,
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
  ): QueryGetPurchaseorderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPurchaseorderRequest,
    } as QueryGetPurchaseorderRequest;
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

  fromJSON(object: any): QueryGetPurchaseorderRequest {
    const message = {
      ...baseQueryGetPurchaseorderRequest,
    } as QueryGetPurchaseorderRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetPurchaseorderRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPurchaseorderRequest>
  ): QueryGetPurchaseorderRequest {
    const message = {
      ...baseQueryGetPurchaseorderRequest,
    } as QueryGetPurchaseorderRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetPurchaseorderResponse: object = {};

export const QueryGetPurchaseorderResponse = {
  encode(
    message: QueryGetPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Purchaseorder !== undefined) {
      Purchaseorder.encode(
        message.Purchaseorder,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetPurchaseorderResponse,
    } as QueryGetPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Purchaseorder = Purchaseorder.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPurchaseorderResponse {
    const message = {
      ...baseQueryGetPurchaseorderResponse,
    } as QueryGetPurchaseorderResponse;
    if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
      message.Purchaseorder = Purchaseorder.fromJSON(object.Purchaseorder);
    } else {
      message.Purchaseorder = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPurchaseorderResponse): unknown {
    const obj: any = {};
    message.Purchaseorder !== undefined &&
      (obj.Purchaseorder = message.Purchaseorder
        ? Purchaseorder.toJSON(message.Purchaseorder)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetPurchaseorderResponse>
  ): QueryGetPurchaseorderResponse {
    const message = {
      ...baseQueryGetPurchaseorderResponse,
    } as QueryGetPurchaseorderResponse;
    if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
      message.Purchaseorder = Purchaseorder.fromPartial(object.Purchaseorder);
    } else {
      message.Purchaseorder = undefined;
    }
    return message;
  },
};

const baseQueryAllPurchaseorderRequest: object = {};

export const QueryAllPurchaseorderRequest = {
  encode(
    message: QueryAllPurchaseorderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPurchaseorderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPurchaseorderRequest,
    } as QueryAllPurchaseorderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPurchaseorderRequest {
    const message = {
      ...baseQueryAllPurchaseorderRequest,
    } as QueryAllPurchaseorderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPurchaseorderRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPurchaseorderRequest>
  ): QueryAllPurchaseorderRequest {
    const message = {
      ...baseQueryAllPurchaseorderRequest,
    } as QueryAllPurchaseorderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPurchaseorderResponse: object = {};

export const QueryAllPurchaseorderResponse = {
  encode(
    message: QueryAllPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Purchaseorder) {
      Purchaseorder.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllPurchaseorderResponse,
    } as QueryAllPurchaseorderResponse;
    message.Purchaseorder = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Purchaseorder.push(
            Purchaseorder.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPurchaseorderResponse {
    const message = {
      ...baseQueryAllPurchaseorderResponse,
    } as QueryAllPurchaseorderResponse;
    message.Purchaseorder = [];
    if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
      for (const e of object.Purchaseorder) {
        message.Purchaseorder.push(Purchaseorder.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPurchaseorderResponse): unknown {
    const obj: any = {};
    if (message.Purchaseorder) {
      obj.Purchaseorder = message.Purchaseorder.map((e) =>
        e ? Purchaseorder.toJSON(e) : undefined
      );
    } else {
      obj.Purchaseorder = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllPurchaseorderResponse>
  ): QueryAllPurchaseorderResponse {
    const message = {
      ...baseQueryAllPurchaseorderResponse,
    } as QueryAllPurchaseorderResponse;
    message.Purchaseorder = [];
    if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
      for (const e of object.Purchaseorder) {
        message.Purchaseorder.push(Purchaseorder.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetSentPurchaseorderRequest: object = { id: 0 };

export const QueryGetSentPurchaseorderRequest = {
  encode(
    message: QueryGetSentPurchaseorderRequest,
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
  ): QueryGetSentPurchaseorderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSentPurchaseorderRequest,
    } as QueryGetSentPurchaseorderRequest;
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

  fromJSON(object: any): QueryGetSentPurchaseorderRequest {
    const message = {
      ...baseQueryGetSentPurchaseorderRequest,
    } as QueryGetSentPurchaseorderRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetSentPurchaseorderRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSentPurchaseorderRequest>
  ): QueryGetSentPurchaseorderRequest {
    const message = {
      ...baseQueryGetSentPurchaseorderRequest,
    } as QueryGetSentPurchaseorderRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetSentPurchaseorderResponse: object = {};

export const QueryGetSentPurchaseorderResponse = {
  encode(
    message: QueryGetSentPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SentPurchaseorder !== undefined) {
      SentPurchaseorder.encode(
        message.SentPurchaseorder,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSentPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSentPurchaseorderResponse,
    } as QueryGetSentPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SentPurchaseorder = SentPurchaseorder.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSentPurchaseorderResponse {
    const message = {
      ...baseQueryGetSentPurchaseorderResponse,
    } as QueryGetSentPurchaseorderResponse;
    if (
      object.SentPurchaseorder !== undefined &&
      object.SentPurchaseorder !== null
    ) {
      message.SentPurchaseorder = SentPurchaseorder.fromJSON(
        object.SentPurchaseorder
      );
    } else {
      message.SentPurchaseorder = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSentPurchaseorderResponse): unknown {
    const obj: any = {};
    message.SentPurchaseorder !== undefined &&
      (obj.SentPurchaseorder = message.SentPurchaseorder
        ? SentPurchaseorder.toJSON(message.SentPurchaseorder)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSentPurchaseorderResponse>
  ): QueryGetSentPurchaseorderResponse {
    const message = {
      ...baseQueryGetSentPurchaseorderResponse,
    } as QueryGetSentPurchaseorderResponse;
    if (
      object.SentPurchaseorder !== undefined &&
      object.SentPurchaseorder !== null
    ) {
      message.SentPurchaseorder = SentPurchaseorder.fromPartial(
        object.SentPurchaseorder
      );
    } else {
      message.SentPurchaseorder = undefined;
    }
    return message;
  },
};

const baseQueryAllSentPurchaseorderRequest: object = {};

export const QueryAllSentPurchaseorderRequest = {
  encode(
    message: QueryAllSentPurchaseorderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSentPurchaseorderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSentPurchaseorderRequest,
    } as QueryAllSentPurchaseorderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSentPurchaseorderRequest {
    const message = {
      ...baseQueryAllSentPurchaseorderRequest,
    } as QueryAllSentPurchaseorderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSentPurchaseorderRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSentPurchaseorderRequest>
  ): QueryAllSentPurchaseorderRequest {
    const message = {
      ...baseQueryAllSentPurchaseorderRequest,
    } as QueryAllSentPurchaseorderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSentPurchaseorderResponse: object = {};

export const QueryAllSentPurchaseorderResponse = {
  encode(
    message: QueryAllSentPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.SentPurchaseorder) {
      SentPurchaseorder.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllSentPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSentPurchaseorderResponse,
    } as QueryAllSentPurchaseorderResponse;
    message.SentPurchaseorder = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SentPurchaseorder.push(
            SentPurchaseorder.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllSentPurchaseorderResponse {
    const message = {
      ...baseQueryAllSentPurchaseorderResponse,
    } as QueryAllSentPurchaseorderResponse;
    message.SentPurchaseorder = [];
    if (
      object.SentPurchaseorder !== undefined &&
      object.SentPurchaseorder !== null
    ) {
      for (const e of object.SentPurchaseorder) {
        message.SentPurchaseorder.push(SentPurchaseorder.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSentPurchaseorderResponse): unknown {
    const obj: any = {};
    if (message.SentPurchaseorder) {
      obj.SentPurchaseorder = message.SentPurchaseorder.map((e) =>
        e ? SentPurchaseorder.toJSON(e) : undefined
      );
    } else {
      obj.SentPurchaseorder = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSentPurchaseorderResponse>
  ): QueryAllSentPurchaseorderResponse {
    const message = {
      ...baseQueryAllSentPurchaseorderResponse,
    } as QueryAllSentPurchaseorderResponse;
    message.SentPurchaseorder = [];
    if (
      object.SentPurchaseorder !== undefined &&
      object.SentPurchaseorder !== null
    ) {
      for (const e of object.SentPurchaseorder) {
        message.SentPurchaseorder.push(SentPurchaseorder.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetTimedoutPurchaseorderRequest: object = { id: 0 };

export const QueryGetTimedoutPurchaseorderRequest = {
  encode(
    message: QueryGetTimedoutPurchaseorderRequest,
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
  ): QueryGetTimedoutPurchaseorderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTimedoutPurchaseorderRequest,
    } as QueryGetTimedoutPurchaseorderRequest;
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

  fromJSON(object: any): QueryGetTimedoutPurchaseorderRequest {
    const message = {
      ...baseQueryGetTimedoutPurchaseorderRequest,
    } as QueryGetTimedoutPurchaseorderRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetTimedoutPurchaseorderRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTimedoutPurchaseorderRequest>
  ): QueryGetTimedoutPurchaseorderRequest {
    const message = {
      ...baseQueryGetTimedoutPurchaseorderRequest,
    } as QueryGetTimedoutPurchaseorderRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetTimedoutPurchaseorderResponse: object = {};

export const QueryGetTimedoutPurchaseorderResponse = {
  encode(
    message: QueryGetTimedoutPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TimedoutPurchaseorder !== undefined) {
      TimedoutPurchaseorder.encode(
        message.TimedoutPurchaseorder,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTimedoutPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTimedoutPurchaseorderResponse,
    } as QueryGetTimedoutPurchaseorderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TimedoutPurchaseorder = TimedoutPurchaseorder.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTimedoutPurchaseorderResponse {
    const message = {
      ...baseQueryGetTimedoutPurchaseorderResponse,
    } as QueryGetTimedoutPurchaseorderResponse;
    if (
      object.TimedoutPurchaseorder !== undefined &&
      object.TimedoutPurchaseorder !== null
    ) {
      message.TimedoutPurchaseorder = TimedoutPurchaseorder.fromJSON(
        object.TimedoutPurchaseorder
      );
    } else {
      message.TimedoutPurchaseorder = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTimedoutPurchaseorderResponse): unknown {
    const obj: any = {};
    message.TimedoutPurchaseorder !== undefined &&
      (obj.TimedoutPurchaseorder = message.TimedoutPurchaseorder
        ? TimedoutPurchaseorder.toJSON(message.TimedoutPurchaseorder)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTimedoutPurchaseorderResponse>
  ): QueryGetTimedoutPurchaseorderResponse {
    const message = {
      ...baseQueryGetTimedoutPurchaseorderResponse,
    } as QueryGetTimedoutPurchaseorderResponse;
    if (
      object.TimedoutPurchaseorder !== undefined &&
      object.TimedoutPurchaseorder !== null
    ) {
      message.TimedoutPurchaseorder = TimedoutPurchaseorder.fromPartial(
        object.TimedoutPurchaseorder
      );
    } else {
      message.TimedoutPurchaseorder = undefined;
    }
    return message;
  },
};

const baseQueryAllTimedoutPurchaseorderRequest: object = {};

export const QueryAllTimedoutPurchaseorderRequest = {
  encode(
    message: QueryAllTimedoutPurchaseorderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllTimedoutPurchaseorderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTimedoutPurchaseorderRequest,
    } as QueryAllTimedoutPurchaseorderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTimedoutPurchaseorderRequest {
    const message = {
      ...baseQueryAllTimedoutPurchaseorderRequest,
    } as QueryAllTimedoutPurchaseorderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTimedoutPurchaseorderRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTimedoutPurchaseorderRequest>
  ): QueryAllTimedoutPurchaseorderRequest {
    const message = {
      ...baseQueryAllTimedoutPurchaseorderRequest,
    } as QueryAllTimedoutPurchaseorderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTimedoutPurchaseorderResponse: object = {};

export const QueryAllTimedoutPurchaseorderResponse = {
  encode(
    message: QueryAllTimedoutPurchaseorderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.TimedoutPurchaseorder) {
      TimedoutPurchaseorder.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllTimedoutPurchaseorderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTimedoutPurchaseorderResponse,
    } as QueryAllTimedoutPurchaseorderResponse;
    message.TimedoutPurchaseorder = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TimedoutPurchaseorder.push(
            TimedoutPurchaseorder.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllTimedoutPurchaseorderResponse {
    const message = {
      ...baseQueryAllTimedoutPurchaseorderResponse,
    } as QueryAllTimedoutPurchaseorderResponse;
    message.TimedoutPurchaseorder = [];
    if (
      object.TimedoutPurchaseorder !== undefined &&
      object.TimedoutPurchaseorder !== null
    ) {
      for (const e of object.TimedoutPurchaseorder) {
        message.TimedoutPurchaseorder.push(TimedoutPurchaseorder.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTimedoutPurchaseorderResponse): unknown {
    const obj: any = {};
    if (message.TimedoutPurchaseorder) {
      obj.TimedoutPurchaseorder = message.TimedoutPurchaseorder.map((e) =>
        e ? TimedoutPurchaseorder.toJSON(e) : undefined
      );
    } else {
      obj.TimedoutPurchaseorder = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTimedoutPurchaseorderResponse>
  ): QueryAllTimedoutPurchaseorderResponse {
    const message = {
      ...baseQueryAllTimedoutPurchaseorderResponse,
    } as QueryAllTimedoutPurchaseorderResponse;
    message.TimedoutPurchaseorder = [];
    if (
      object.TimedoutPurchaseorder !== undefined &&
      object.TimedoutPurchaseorder !== null
    ) {
      for (const e of object.TimedoutPurchaseorder) {
        message.TimedoutPurchaseorder.push(
          TimedoutPurchaseorder.fromPartial(e)
        );
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a purchaseorder by id. */
  Purchaseorder(
    request: QueryGetPurchaseorderRequest
  ): Promise<QueryGetPurchaseorderResponse>;
  /** Queries a list of purchaseorder items. */
  PurchaseorderAll(
    request: QueryAllPurchaseorderRequest
  ): Promise<QueryAllPurchaseorderResponse>;
  /** Queries a sentPurchaseorder by id. */
  SentPurchaseorder(
    request: QueryGetSentPurchaseorderRequest
  ): Promise<QueryGetSentPurchaseorderResponse>;
  /** Queries a list of sentPurchaseorder items. */
  SentPurchaseorderAll(
    request: QueryAllSentPurchaseorderRequest
  ): Promise<QueryAllSentPurchaseorderResponse>;
  /** Queries a timedoutPurchaseorder by id. */
  TimedoutPurchaseorder(
    request: QueryGetTimedoutPurchaseorderRequest
  ): Promise<QueryGetTimedoutPurchaseorderResponse>;
  /** Queries a list of timedoutPurchaseorder items. */
  TimedoutPurchaseorderAll(
    request: QueryAllTimedoutPurchaseorderRequest
  ): Promise<QueryAllTimedoutPurchaseorderResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Purchaseorder(
    request: QueryGetPurchaseorderRequest
  ): Promise<QueryGetPurchaseorderResponse> {
    const data = QueryGetPurchaseorderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Query",
      "Purchaseorder",
      data
    );
    return promise.then((data) =>
      QueryGetPurchaseorderResponse.decode(new Reader(data))
    );
  }

  PurchaseorderAll(
    request: QueryAllPurchaseorderRequest
  ): Promise<QueryAllPurchaseorderResponse> {
    const data = QueryAllPurchaseorderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Query",
      "PurchaseorderAll",
      data
    );
    return promise.then((data) =>
      QueryAllPurchaseorderResponse.decode(new Reader(data))
    );
  }

  SentPurchaseorder(
    request: QueryGetSentPurchaseorderRequest
  ): Promise<QueryGetSentPurchaseorderResponse> {
    const data = QueryGetSentPurchaseorderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Query",
      "SentPurchaseorder",
      data
    );
    return promise.then((data) =>
      QueryGetSentPurchaseorderResponse.decode(new Reader(data))
    );
  }

  SentPurchaseorderAll(
    request: QueryAllSentPurchaseorderRequest
  ): Promise<QueryAllSentPurchaseorderResponse> {
    const data = QueryAllSentPurchaseorderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Query",
      "SentPurchaseorderAll",
      data
    );
    return promise.then((data) =>
      QueryAllSentPurchaseorderResponse.decode(new Reader(data))
    );
  }

  TimedoutPurchaseorder(
    request: QueryGetTimedoutPurchaseorderRequest
  ): Promise<QueryGetTimedoutPurchaseorderResponse> {
    const data = QueryGetTimedoutPurchaseorderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Query",
      "TimedoutPurchaseorder",
      data
    );
    return promise.then((data) =>
      QueryGetTimedoutPurchaseorderResponse.decode(new Reader(data))
    );
  }

  TimedoutPurchaseorderAll(
    request: QueryAllTimedoutPurchaseorderRequest
  ): Promise<QueryAllTimedoutPurchaseorderResponse> {
    const data = QueryAllTimedoutPurchaseorderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.purchaseorder.Query",
      "TimedoutPurchaseorderAll",
      data
    );
    return promise.then((data) =>
      QueryAllTimedoutPurchaseorderResponse.decode(new Reader(data))
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
