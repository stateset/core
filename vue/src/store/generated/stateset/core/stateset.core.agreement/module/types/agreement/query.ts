/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { SentAgreement } from "../agreement/sent_agreement";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { TimedoutAgreement } from "../agreement/timedout_agreement";
import { Agreement } from "../agreement/agreement";

export const protobufPackage = "stateset.core.agreement";

export interface QueryGetSentAgreementRequest {
  id: number;
}

export interface QueryGetSentAgreementResponse {
  SentAgreement: SentAgreement | undefined;
}

export interface QueryAllSentAgreementRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllSentAgreementResponse {
  SentAgreement: SentAgreement[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTimedoutAgreementRequest {
  id: number;
}

export interface QueryGetTimedoutAgreementResponse {
  TimedoutAgreement: TimedoutAgreement | undefined;
}

export interface QueryAllTimedoutAgreementRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTimedoutAgreementResponse {
  TimedoutAgreement: TimedoutAgreement[];
  pagination: PageResponse | undefined;
}

export interface QueryGetAgreementRequest {
  id: number;
}

export interface QueryGetAgreementResponse {
  Agreement: Agreement | undefined;
}

export interface QueryAllAgreementRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllAgreementResponse {
  Agreement: Agreement[];
  pagination: PageResponse | undefined;
}

const baseQueryGetSentAgreementRequest: object = { id: 0 };

export const QueryGetSentAgreementRequest = {
  encode(
    message: QueryGetSentAgreementRequest,
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
  ): QueryGetSentAgreementRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSentAgreementRequest,
    } as QueryGetSentAgreementRequest;
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

  fromJSON(object: any): QueryGetSentAgreementRequest {
    const message = {
      ...baseQueryGetSentAgreementRequest,
    } as QueryGetSentAgreementRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetSentAgreementRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSentAgreementRequest>
  ): QueryGetSentAgreementRequest {
    const message = {
      ...baseQueryGetSentAgreementRequest,
    } as QueryGetSentAgreementRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetSentAgreementResponse: object = {};

export const QueryGetSentAgreementResponse = {
  encode(
    message: QueryGetSentAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SentAgreement !== undefined) {
      SentAgreement.encode(
        message.SentAgreement,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSentAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSentAgreementResponse,
    } as QueryGetSentAgreementResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SentAgreement = SentAgreement.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSentAgreementResponse {
    const message = {
      ...baseQueryGetSentAgreementResponse,
    } as QueryGetSentAgreementResponse;
    if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
      message.SentAgreement = SentAgreement.fromJSON(object.SentAgreement);
    } else {
      message.SentAgreement = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSentAgreementResponse): unknown {
    const obj: any = {};
    message.SentAgreement !== undefined &&
      (obj.SentAgreement = message.SentAgreement
        ? SentAgreement.toJSON(message.SentAgreement)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSentAgreementResponse>
  ): QueryGetSentAgreementResponse {
    const message = {
      ...baseQueryGetSentAgreementResponse,
    } as QueryGetSentAgreementResponse;
    if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
      message.SentAgreement = SentAgreement.fromPartial(object.SentAgreement);
    } else {
      message.SentAgreement = undefined;
    }
    return message;
  },
};

const baseQueryAllSentAgreementRequest: object = {};

export const QueryAllSentAgreementRequest = {
  encode(
    message: QueryAllSentAgreementRequest,
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
  ): QueryAllSentAgreementRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSentAgreementRequest,
    } as QueryAllSentAgreementRequest;
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

  fromJSON(object: any): QueryAllSentAgreementRequest {
    const message = {
      ...baseQueryAllSentAgreementRequest,
    } as QueryAllSentAgreementRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSentAgreementRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSentAgreementRequest>
  ): QueryAllSentAgreementRequest {
    const message = {
      ...baseQueryAllSentAgreementRequest,
    } as QueryAllSentAgreementRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllSentAgreementResponse: object = {};

export const QueryAllSentAgreementResponse = {
  encode(
    message: QueryAllSentAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.SentAgreement) {
      SentAgreement.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllSentAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllSentAgreementResponse,
    } as QueryAllSentAgreementResponse;
    message.SentAgreement = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SentAgreement.push(
            SentAgreement.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllSentAgreementResponse {
    const message = {
      ...baseQueryAllSentAgreementResponse,
    } as QueryAllSentAgreementResponse;
    message.SentAgreement = [];
    if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
      for (const e of object.SentAgreement) {
        message.SentAgreement.push(SentAgreement.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllSentAgreementResponse): unknown {
    const obj: any = {};
    if (message.SentAgreement) {
      obj.SentAgreement = message.SentAgreement.map((e) =>
        e ? SentAgreement.toJSON(e) : undefined
      );
    } else {
      obj.SentAgreement = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllSentAgreementResponse>
  ): QueryAllSentAgreementResponse {
    const message = {
      ...baseQueryAllSentAgreementResponse,
    } as QueryAllSentAgreementResponse;
    message.SentAgreement = [];
    if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
      for (const e of object.SentAgreement) {
        message.SentAgreement.push(SentAgreement.fromPartial(e));
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

const baseQueryGetTimedoutAgreementRequest: object = { id: 0 };

export const QueryGetTimedoutAgreementRequest = {
  encode(
    message: QueryGetTimedoutAgreementRequest,
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
  ): QueryGetTimedoutAgreementRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTimedoutAgreementRequest,
    } as QueryGetTimedoutAgreementRequest;
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

  fromJSON(object: any): QueryGetTimedoutAgreementRequest {
    const message = {
      ...baseQueryGetTimedoutAgreementRequest,
    } as QueryGetTimedoutAgreementRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetTimedoutAgreementRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTimedoutAgreementRequest>
  ): QueryGetTimedoutAgreementRequest {
    const message = {
      ...baseQueryGetTimedoutAgreementRequest,
    } as QueryGetTimedoutAgreementRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetTimedoutAgreementResponse: object = {};

export const QueryGetTimedoutAgreementResponse = {
  encode(
    message: QueryGetTimedoutAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.TimedoutAgreement !== undefined) {
      TimedoutAgreement.encode(
        message.TimedoutAgreement,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTimedoutAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTimedoutAgreementResponse,
    } as QueryGetTimedoutAgreementResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TimedoutAgreement = TimedoutAgreement.decode(
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

  fromJSON(object: any): QueryGetTimedoutAgreementResponse {
    const message = {
      ...baseQueryGetTimedoutAgreementResponse,
    } as QueryGetTimedoutAgreementResponse;
    if (
      object.TimedoutAgreement !== undefined &&
      object.TimedoutAgreement !== null
    ) {
      message.TimedoutAgreement = TimedoutAgreement.fromJSON(
        object.TimedoutAgreement
      );
    } else {
      message.TimedoutAgreement = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTimedoutAgreementResponse): unknown {
    const obj: any = {};
    message.TimedoutAgreement !== undefined &&
      (obj.TimedoutAgreement = message.TimedoutAgreement
        ? TimedoutAgreement.toJSON(message.TimedoutAgreement)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTimedoutAgreementResponse>
  ): QueryGetTimedoutAgreementResponse {
    const message = {
      ...baseQueryGetTimedoutAgreementResponse,
    } as QueryGetTimedoutAgreementResponse;
    if (
      object.TimedoutAgreement !== undefined &&
      object.TimedoutAgreement !== null
    ) {
      message.TimedoutAgreement = TimedoutAgreement.fromPartial(
        object.TimedoutAgreement
      );
    } else {
      message.TimedoutAgreement = undefined;
    }
    return message;
  },
};

const baseQueryAllTimedoutAgreementRequest: object = {};

export const QueryAllTimedoutAgreementRequest = {
  encode(
    message: QueryAllTimedoutAgreementRequest,
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
  ): QueryAllTimedoutAgreementRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTimedoutAgreementRequest,
    } as QueryAllTimedoutAgreementRequest;
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

  fromJSON(object: any): QueryAllTimedoutAgreementRequest {
    const message = {
      ...baseQueryAllTimedoutAgreementRequest,
    } as QueryAllTimedoutAgreementRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTimedoutAgreementRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTimedoutAgreementRequest>
  ): QueryAllTimedoutAgreementRequest {
    const message = {
      ...baseQueryAllTimedoutAgreementRequest,
    } as QueryAllTimedoutAgreementRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTimedoutAgreementResponse: object = {};

export const QueryAllTimedoutAgreementResponse = {
  encode(
    message: QueryAllTimedoutAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.TimedoutAgreement) {
      TimedoutAgreement.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllTimedoutAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTimedoutAgreementResponse,
    } as QueryAllTimedoutAgreementResponse;
    message.TimedoutAgreement = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.TimedoutAgreement.push(
            TimedoutAgreement.decode(reader, reader.uint32())
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

  fromJSON(object: any): QueryAllTimedoutAgreementResponse {
    const message = {
      ...baseQueryAllTimedoutAgreementResponse,
    } as QueryAllTimedoutAgreementResponse;
    message.TimedoutAgreement = [];
    if (
      object.TimedoutAgreement !== undefined &&
      object.TimedoutAgreement !== null
    ) {
      for (const e of object.TimedoutAgreement) {
        message.TimedoutAgreement.push(TimedoutAgreement.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTimedoutAgreementResponse): unknown {
    const obj: any = {};
    if (message.TimedoutAgreement) {
      obj.TimedoutAgreement = message.TimedoutAgreement.map((e) =>
        e ? TimedoutAgreement.toJSON(e) : undefined
      );
    } else {
      obj.TimedoutAgreement = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTimedoutAgreementResponse>
  ): QueryAllTimedoutAgreementResponse {
    const message = {
      ...baseQueryAllTimedoutAgreementResponse,
    } as QueryAllTimedoutAgreementResponse;
    message.TimedoutAgreement = [];
    if (
      object.TimedoutAgreement !== undefined &&
      object.TimedoutAgreement !== null
    ) {
      for (const e of object.TimedoutAgreement) {
        message.TimedoutAgreement.push(TimedoutAgreement.fromPartial(e));
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

const baseQueryGetAgreementRequest: object = { id: 0 };

export const QueryGetAgreementRequest = {
  encode(
    message: QueryGetAgreementRequest,
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
  ): QueryGetAgreementRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAgreementRequest,
    } as QueryGetAgreementRequest;
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

  fromJSON(object: any): QueryGetAgreementRequest {
    const message = {
      ...baseQueryGetAgreementRequest,
    } as QueryGetAgreementRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetAgreementRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAgreementRequest>
  ): QueryGetAgreementRequest {
    const message = {
      ...baseQueryGetAgreementRequest,
    } as QueryGetAgreementRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetAgreementResponse: object = {};

export const QueryGetAgreementResponse = {
  encode(
    message: QueryGetAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Agreement !== undefined) {
      Agreement.encode(message.Agreement, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAgreementResponse,
    } as QueryGetAgreementResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Agreement = Agreement.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAgreementResponse {
    const message = {
      ...baseQueryGetAgreementResponse,
    } as QueryGetAgreementResponse;
    if (object.Agreement !== undefined && object.Agreement !== null) {
      message.Agreement = Agreement.fromJSON(object.Agreement);
    } else {
      message.Agreement = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetAgreementResponse): unknown {
    const obj: any = {};
    message.Agreement !== undefined &&
      (obj.Agreement = message.Agreement
        ? Agreement.toJSON(message.Agreement)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAgreementResponse>
  ): QueryGetAgreementResponse {
    const message = {
      ...baseQueryGetAgreementResponse,
    } as QueryGetAgreementResponse;
    if (object.Agreement !== undefined && object.Agreement !== null) {
      message.Agreement = Agreement.fromPartial(object.Agreement);
    } else {
      message.Agreement = undefined;
    }
    return message;
  },
};

const baseQueryAllAgreementRequest: object = {};

export const QueryAllAgreementRequest = {
  encode(
    message: QueryAllAgreementRequest,
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
  ): QueryAllAgreementRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAgreementRequest,
    } as QueryAllAgreementRequest;
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

  fromJSON(object: any): QueryAllAgreementRequest {
    const message = {
      ...baseQueryAllAgreementRequest,
    } as QueryAllAgreementRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAgreementRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAgreementRequest>
  ): QueryAllAgreementRequest {
    const message = {
      ...baseQueryAllAgreementRequest,
    } as QueryAllAgreementRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllAgreementResponse: object = {};

export const QueryAllAgreementResponse = {
  encode(
    message: QueryAllAgreementResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Agreement) {
      Agreement.encode(v!, writer.uint32(10).fork()).ldelim();
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
  ): QueryAllAgreementResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllAgreementResponse,
    } as QueryAllAgreementResponse;
    message.Agreement = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Agreement.push(Agreement.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllAgreementResponse {
    const message = {
      ...baseQueryAllAgreementResponse,
    } as QueryAllAgreementResponse;
    message.Agreement = [];
    if (object.Agreement !== undefined && object.Agreement !== null) {
      for (const e of object.Agreement) {
        message.Agreement.push(Agreement.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllAgreementResponse): unknown {
    const obj: any = {};
    if (message.Agreement) {
      obj.Agreement = message.Agreement.map((e) =>
        e ? Agreement.toJSON(e) : undefined
      );
    } else {
      obj.Agreement = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllAgreementResponse>
  ): QueryAllAgreementResponse {
    const message = {
      ...baseQueryAllAgreementResponse,
    } as QueryAllAgreementResponse;
    message.Agreement = [];
    if (object.Agreement !== undefined && object.Agreement !== null) {
      for (const e of object.Agreement) {
        message.Agreement.push(Agreement.fromPartial(e));
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
  /** Queries a sentAgreement by id. */
  SentAgreement(
    request: QueryGetSentAgreementRequest
  ): Promise<QueryGetSentAgreementResponse>;
  /** Queries a list of sentAgreement items. */
  SentAgreementAll(
    request: QueryAllSentAgreementRequest
  ): Promise<QueryAllSentAgreementResponse>;
  /** Queries a timedoutAgreement by id. */
  TimedoutAgreement(
    request: QueryGetTimedoutAgreementRequest
  ): Promise<QueryGetTimedoutAgreementResponse>;
  /** Queries a list of timedoutAgreement items. */
  TimedoutAgreementAll(
    request: QueryAllTimedoutAgreementRequest
  ): Promise<QueryAllTimedoutAgreementResponse>;
  /** Queries a agreement by id. */
  Agreement(
    request: QueryGetAgreementRequest
  ): Promise<QueryGetAgreementResponse>;
  /** Queries a list of agreement items. */
  AgreementAll(
    request: QueryAllAgreementRequest
  ): Promise<QueryAllAgreementResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SentAgreement(
    request: QueryGetSentAgreementRequest
  ): Promise<QueryGetSentAgreementResponse> {
    const data = QueryGetSentAgreementRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Query",
      "SentAgreement",
      data
    );
    return promise.then((data) =>
      QueryGetSentAgreementResponse.decode(new Reader(data))
    );
  }

  SentAgreementAll(
    request: QueryAllSentAgreementRequest
  ): Promise<QueryAllSentAgreementResponse> {
    const data = QueryAllSentAgreementRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Query",
      "SentAgreementAll",
      data
    );
    return promise.then((data) =>
      QueryAllSentAgreementResponse.decode(new Reader(data))
    );
  }

  TimedoutAgreement(
    request: QueryGetTimedoutAgreementRequest
  ): Promise<QueryGetTimedoutAgreementResponse> {
    const data = QueryGetTimedoutAgreementRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Query",
      "TimedoutAgreement",
      data
    );
    return promise.then((data) =>
      QueryGetTimedoutAgreementResponse.decode(new Reader(data))
    );
  }

  TimedoutAgreementAll(
    request: QueryAllTimedoutAgreementRequest
  ): Promise<QueryAllTimedoutAgreementResponse> {
    const data = QueryAllTimedoutAgreementRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Query",
      "TimedoutAgreementAll",
      data
    );
    return promise.then((data) =>
      QueryAllTimedoutAgreementResponse.decode(new Reader(data))
    );
  }

  Agreement(
    request: QueryGetAgreementRequest
  ): Promise<QueryGetAgreementResponse> {
    const data = QueryGetAgreementRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Query",
      "Agreement",
      data
    );
    return promise.then((data) =>
      QueryGetAgreementResponse.decode(new Reader(data))
    );
  }

  AgreementAll(
    request: QueryAllAgreementRequest
  ): Promise<QueryAllAgreementResponse> {
    const data = QueryAllAgreementRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.agreement.Query",
      "AgreementAll",
      data
    );
    return promise.then((data) =>
      QueryAllAgreementResponse.decode(new Reader(data))
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
