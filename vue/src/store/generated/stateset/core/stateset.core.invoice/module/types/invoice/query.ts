/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Invoice } from "../invoice/invoice";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "stateset.core.invoice";

export interface QueryGetInvoiceRequest {
  id: number;
}

export interface QueryGetInvoiceResponse {
  Invoice: Invoice | undefined;
}

export interface QueryAllInvoiceRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllInvoiceResponse {
  Invoice: Invoice[];
  pagination: PageResponse | undefined;
}

const baseQueryGetInvoiceRequest: object = { id: 0 };

export const QueryGetInvoiceRequest = {
  encode(
    message: QueryGetInvoiceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetInvoiceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetInvoiceRequest } as QueryGetInvoiceRequest;
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

  fromJSON(object: any): QueryGetInvoiceRequest {
    const message = { ...baseQueryGetInvoiceRequest } as QueryGetInvoiceRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetInvoiceRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInvoiceRequest>
  ): QueryGetInvoiceRequest {
    const message = { ...baseQueryGetInvoiceRequest } as QueryGetInvoiceRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetInvoiceResponse: object = {};

export const QueryGetInvoiceResponse = {
  encode(
    message: QueryGetInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Invoice !== undefined) {
      Invoice.encode(message.Invoice, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetInvoiceResponse,
    } as QueryGetInvoiceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Invoice = Invoice.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetInvoiceResponse {
    const message = {
      ...baseQueryGetInvoiceResponse,
    } as QueryGetInvoiceResponse;
    if (object.Invoice !== undefined && object.Invoice !== null) {
      message.Invoice = Invoice.fromJSON(object.Invoice);
    } else {
      message.Invoice = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetInvoiceResponse): unknown {
    const obj: any = {};
    message.Invoice !== undefined &&
      (obj.Invoice = message.Invoice
        ? Invoice.toJSON(message.Invoice)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetInvoiceResponse>
  ): QueryGetInvoiceResponse {
    const message = {
      ...baseQueryGetInvoiceResponse,
    } as QueryGetInvoiceResponse;
    if (object.Invoice !== undefined && object.Invoice !== null) {
      message.Invoice = Invoice.fromPartial(object.Invoice);
    } else {
      message.Invoice = undefined;
    }
    return message;
  },
};

const baseQueryAllInvoiceRequest: object = {};

export const QueryAllInvoiceRequest = {
  encode(
    message: QueryAllInvoiceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllInvoiceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllInvoiceRequest } as QueryAllInvoiceRequest;
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

  fromJSON(object: any): QueryAllInvoiceRequest {
    const message = { ...baseQueryAllInvoiceRequest } as QueryAllInvoiceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInvoiceRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInvoiceRequest>
  ): QueryAllInvoiceRequest {
    const message = { ...baseQueryAllInvoiceRequest } as QueryAllInvoiceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllInvoiceResponse: object = {};

export const QueryAllInvoiceResponse = {
  encode(
    message: QueryAllInvoiceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Invoice) {
      Invoice.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllInvoiceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllInvoiceResponse,
    } as QueryAllInvoiceResponse;
    message.Invoice = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Invoice.push(Invoice.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllInvoiceResponse {
    const message = {
      ...baseQueryAllInvoiceResponse,
    } as QueryAllInvoiceResponse;
    message.Invoice = [];
    if (object.Invoice !== undefined && object.Invoice !== null) {
      for (const e of object.Invoice) {
        message.Invoice.push(Invoice.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllInvoiceResponse): unknown {
    const obj: any = {};
    if (message.Invoice) {
      obj.Invoice = message.Invoice.map((e) =>
        e ? Invoice.toJSON(e) : undefined
      );
    } else {
      obj.Invoice = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllInvoiceResponse>
  ): QueryAllInvoiceResponse {
    const message = {
      ...baseQueryAllInvoiceResponse,
    } as QueryAllInvoiceResponse;
    message.Invoice = [];
    if (object.Invoice !== undefined && object.Invoice !== null) {
      for (const e of object.Invoice) {
        message.Invoice.push(Invoice.fromPartial(e));
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
  /** Queries a invoice by id. */
  Invoice(request: QueryGetInvoiceRequest): Promise<QueryGetInvoiceResponse>;
  /** Queries a list of invoice items. */
  InvoiceAll(request: QueryAllInvoiceRequest): Promise<QueryAllInvoiceResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Invoice(request: QueryGetInvoiceRequest): Promise<QueryGetInvoiceResponse> {
    const data = QueryGetInvoiceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Query",
      "Invoice",
      data
    );
    return promise.then((data) =>
      QueryGetInvoiceResponse.decode(new Reader(data))
    );
  }

  InvoiceAll(
    request: QueryAllInvoiceRequest
  ): Promise<QueryAllInvoiceResponse> {
    const data = QueryAllInvoiceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "stateset.core.invoice.Query",
      "InvoiceAll",
      data
    );
    return promise.then((data) =>
      QueryAllInvoiceResponse.decode(new Reader(data))
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
