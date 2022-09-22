/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Invoice } from "../invoice/invoice";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { SentInvoice } from "../invoice/sent_invoice";
import { TimedoutInvoice } from "../invoice/timedout_invoice";
export const protobufPackage = "stateset.core.invoice";
const baseQueryGetInvoiceRequest = { id: 0 };
export const QueryGetInvoiceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetInvoiceRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetInvoiceRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetInvoiceRequest };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseQueryGetInvoiceResponse = {};
export const QueryGetInvoiceResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Invoice !== undefined) {
            Invoice.encode(message.Invoice, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetInvoiceResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetInvoiceResponse,
        };
        if (object.Invoice !== undefined && object.Invoice !== null) {
            message.Invoice = Invoice.fromJSON(object.Invoice);
        }
        else {
            message.Invoice = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Invoice !== undefined &&
            (obj.Invoice = message.Invoice
                ? Invoice.toJSON(message.Invoice)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetInvoiceResponse,
        };
        if (object.Invoice !== undefined && object.Invoice !== null) {
            message.Invoice = Invoice.fromPartial(object.Invoice);
        }
        else {
            message.Invoice = undefined;
        }
        return message;
    },
};
const baseQueryAllInvoiceRequest = {};
export const QueryAllInvoiceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllInvoiceRequest };
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
    fromJSON(object) {
        const message = { ...baseQueryAllInvoiceRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllInvoiceRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllInvoiceResponse = {};
export const QueryAllInvoiceResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Invoice) {
            Invoice.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllInvoiceResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllInvoiceResponse,
        };
        message.Invoice = [];
        if (object.Invoice !== undefined && object.Invoice !== null) {
            for (const e of object.Invoice) {
                message.Invoice.push(Invoice.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.Invoice) {
            obj.Invoice = message.Invoice.map((e) => e ? Invoice.toJSON(e) : undefined);
        }
        else {
            obj.Invoice = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllInvoiceResponse,
        };
        message.Invoice = [];
        if (object.Invoice !== undefined && object.Invoice !== null) {
            for (const e of object.Invoice) {
                message.Invoice.push(Invoice.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryGetSentInvoiceRequest = { id: 0 };
export const QueryGetSentInvoiceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetSentInvoiceRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryGetSentInvoiceRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetSentInvoiceRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseQueryGetSentInvoiceResponse = {};
export const QueryGetSentInvoiceResponse = {
    encode(message, writer = Writer.create()) {
        if (message.SentInvoice !== undefined) {
            SentInvoice.encode(message.SentInvoice, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetSentInvoiceResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SentInvoice = SentInvoice.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryGetSentInvoiceResponse,
        };
        if (object.SentInvoice !== undefined && object.SentInvoice !== null) {
            message.SentInvoice = SentInvoice.fromJSON(object.SentInvoice);
        }
        else {
            message.SentInvoice = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.SentInvoice !== undefined &&
            (obj.SentInvoice = message.SentInvoice
                ? SentInvoice.toJSON(message.SentInvoice)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetSentInvoiceResponse,
        };
        if (object.SentInvoice !== undefined && object.SentInvoice !== null) {
            message.SentInvoice = SentInvoice.fromPartial(object.SentInvoice);
        }
        else {
            message.SentInvoice = undefined;
        }
        return message;
    },
};
const baseQueryAllSentInvoiceRequest = {};
export const QueryAllSentInvoiceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllSentInvoiceRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllSentInvoiceRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSentInvoiceRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllSentInvoiceResponse = {};
export const QueryAllSentInvoiceResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.SentInvoice) {
            SentInvoice.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllSentInvoiceResponse,
        };
        message.SentInvoice = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SentInvoice.push(SentInvoice.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllSentInvoiceResponse,
        };
        message.SentInvoice = [];
        if (object.SentInvoice !== undefined && object.SentInvoice !== null) {
            for (const e of object.SentInvoice) {
                message.SentInvoice.push(SentInvoice.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.SentInvoice) {
            obj.SentInvoice = message.SentInvoice.map((e) => e ? SentInvoice.toJSON(e) : undefined);
        }
        else {
            obj.SentInvoice = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSentInvoiceResponse,
        };
        message.SentInvoice = [];
        if (object.SentInvoice !== undefined && object.SentInvoice !== null) {
            for (const e of object.SentInvoice) {
                message.SentInvoice.push(SentInvoice.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryGetTimedoutInvoiceRequest = { id: 0 };
export const QueryGetTimedoutInvoiceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.id !== 0) {
            writer.uint32(8).uint64(message.id);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetTimedoutInvoiceRequest,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.id = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryGetTimedoutInvoiceRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.id !== undefined && (obj.id = message.id);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetTimedoutInvoiceRequest,
        };
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        return message;
    },
};
const baseQueryGetTimedoutInvoiceResponse = {};
export const QueryGetTimedoutInvoiceResponse = {
    encode(message, writer = Writer.create()) {
        if (message.TimedoutInvoice !== undefined) {
            TimedoutInvoice.encode(message.TimedoutInvoice, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetTimedoutInvoiceResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.TimedoutInvoice = TimedoutInvoice.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryGetTimedoutInvoiceResponse,
        };
        if (object.TimedoutInvoice !== undefined &&
            object.TimedoutInvoice !== null) {
            message.TimedoutInvoice = TimedoutInvoice.fromJSON(object.TimedoutInvoice);
        }
        else {
            message.TimedoutInvoice = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.TimedoutInvoice !== undefined &&
            (obj.TimedoutInvoice = message.TimedoutInvoice
                ? TimedoutInvoice.toJSON(message.TimedoutInvoice)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetTimedoutInvoiceResponse,
        };
        if (object.TimedoutInvoice !== undefined &&
            object.TimedoutInvoice !== null) {
            message.TimedoutInvoice = TimedoutInvoice.fromPartial(object.TimedoutInvoice);
        }
        else {
            message.TimedoutInvoice = undefined;
        }
        return message;
    },
};
const baseQueryAllTimedoutInvoiceRequest = {};
export const QueryAllTimedoutInvoiceRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllTimedoutInvoiceRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllTimedoutInvoiceRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageRequest.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllTimedoutInvoiceRequest,
        };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
const baseQueryAllTimedoutInvoiceResponse = {};
export const QueryAllTimedoutInvoiceResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.TimedoutInvoice) {
            TimedoutInvoice.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryAllTimedoutInvoiceResponse,
        };
        message.TimedoutInvoice = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.TimedoutInvoice.push(TimedoutInvoice.decode(reader, reader.uint32()));
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllTimedoutInvoiceResponse,
        };
        message.TimedoutInvoice = [];
        if (object.TimedoutInvoice !== undefined &&
            object.TimedoutInvoice !== null) {
            for (const e of object.TimedoutInvoice) {
                message.TimedoutInvoice.push(TimedoutInvoice.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.TimedoutInvoice) {
            obj.TimedoutInvoice = message.TimedoutInvoice.map((e) => e ? TimedoutInvoice.toJSON(e) : undefined);
        }
        else {
            obj.TimedoutInvoice = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllTimedoutInvoiceResponse,
        };
        message.TimedoutInvoice = [];
        if (object.TimedoutInvoice !== undefined &&
            object.TimedoutInvoice !== null) {
            for (const e of object.TimedoutInvoice) {
                message.TimedoutInvoice.push(TimedoutInvoice.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Invoice(request) {
        const data = QueryGetInvoiceRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.invoice.Query", "Invoice", data);
        return promise.then((data) => QueryGetInvoiceResponse.decode(new Reader(data)));
    }
    InvoiceAll(request) {
        const data = QueryAllInvoiceRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.invoice.Query", "InvoiceAll", data);
        return promise.then((data) => QueryAllInvoiceResponse.decode(new Reader(data)));
    }
    SentInvoice(request) {
        const data = QueryGetSentInvoiceRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.invoice.Query", "SentInvoice", data);
        return promise.then((data) => QueryGetSentInvoiceResponse.decode(new Reader(data)));
    }
    SentInvoiceAll(request) {
        const data = QueryAllSentInvoiceRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.invoice.Query", "SentInvoiceAll", data);
        return promise.then((data) => QueryAllSentInvoiceResponse.decode(new Reader(data)));
    }
    TimedoutInvoice(request) {
        const data = QueryGetTimedoutInvoiceRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.invoice.Query", "TimedoutInvoice", data);
        return promise.then((data) => QueryGetTimedoutInvoiceResponse.decode(new Reader(data)));
    }
    TimedoutInvoiceAll(request) {
        const data = QueryAllTimedoutInvoiceRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.invoice.Query", "TimedoutInvoiceAll", data);
        return promise.then((data) => QueryAllTimedoutInvoiceResponse.decode(new Reader(data)));
    }
}
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
