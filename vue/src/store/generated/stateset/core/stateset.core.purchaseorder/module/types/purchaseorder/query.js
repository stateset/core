/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { SentPurchaseorder } from "../purchaseorder/sent_purchaseorder";
import { TimedoutPurchaseorder } from "../purchaseorder/timedout_purchaseorder";
export const protobufPackage = "stateset.core.purchaseorder";
const baseQueryGetPurchaseorderRequest = { id: 0 };
export const QueryGetPurchaseorderRequest = {
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
            ...baseQueryGetPurchaseorderRequest,
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
            ...baseQueryGetPurchaseorderRequest,
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
            ...baseQueryGetPurchaseorderRequest,
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
const baseQueryGetPurchaseorderResponse = {};
export const QueryGetPurchaseorderResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Purchaseorder !== undefined) {
            Purchaseorder.encode(message.Purchaseorder, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetPurchaseorderResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetPurchaseorderResponse,
        };
        if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
            message.Purchaseorder = Purchaseorder.fromJSON(object.Purchaseorder);
        }
        else {
            message.Purchaseorder = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Purchaseorder !== undefined &&
            (obj.Purchaseorder = message.Purchaseorder
                ? Purchaseorder.toJSON(message.Purchaseorder)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetPurchaseorderResponse,
        };
        if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
            message.Purchaseorder = Purchaseorder.fromPartial(object.Purchaseorder);
        }
        else {
            message.Purchaseorder = undefined;
        }
        return message;
    },
};
const baseQueryAllPurchaseorderRequest = {};
export const QueryAllPurchaseorderRequest = {
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
            ...baseQueryAllPurchaseorderRequest,
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
            ...baseQueryAllPurchaseorderRequest,
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
            ...baseQueryAllPurchaseorderRequest,
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
const baseQueryAllPurchaseorderResponse = {};
export const QueryAllPurchaseorderResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Purchaseorder) {
            Purchaseorder.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryAllPurchaseorderResponse,
        };
        message.Purchaseorder = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.Purchaseorder.push(Purchaseorder.decode(reader, reader.uint32()));
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
            ...baseQueryAllPurchaseorderResponse,
        };
        message.Purchaseorder = [];
        if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
            for (const e of object.Purchaseorder) {
                message.Purchaseorder.push(Purchaseorder.fromJSON(e));
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
        if (message.Purchaseorder) {
            obj.Purchaseorder = message.Purchaseorder.map((e) => e ? Purchaseorder.toJSON(e) : undefined);
        }
        else {
            obj.Purchaseorder = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllPurchaseorderResponse,
        };
        message.Purchaseorder = [];
        if (object.Purchaseorder !== undefined && object.Purchaseorder !== null) {
            for (const e of object.Purchaseorder) {
                message.Purchaseorder.push(Purchaseorder.fromPartial(e));
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
const baseQueryGetSentPurchaseorderRequest = { id: 0 };
export const QueryGetSentPurchaseorderRequest = {
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
            ...baseQueryGetSentPurchaseorderRequest,
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
            ...baseQueryGetSentPurchaseorderRequest,
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
            ...baseQueryGetSentPurchaseorderRequest,
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
const baseQueryGetSentPurchaseorderResponse = {};
export const QueryGetSentPurchaseorderResponse = {
    encode(message, writer = Writer.create()) {
        if (message.SentPurchaseorder !== undefined) {
            SentPurchaseorder.encode(message.SentPurchaseorder, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetSentPurchaseorderResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SentPurchaseorder = SentPurchaseorder.decode(reader, reader.uint32());
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
            ...baseQueryGetSentPurchaseorderResponse,
        };
        if (object.SentPurchaseorder !== undefined &&
            object.SentPurchaseorder !== null) {
            message.SentPurchaseorder = SentPurchaseorder.fromJSON(object.SentPurchaseorder);
        }
        else {
            message.SentPurchaseorder = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.SentPurchaseorder !== undefined &&
            (obj.SentPurchaseorder = message.SentPurchaseorder
                ? SentPurchaseorder.toJSON(message.SentPurchaseorder)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetSentPurchaseorderResponse,
        };
        if (object.SentPurchaseorder !== undefined &&
            object.SentPurchaseorder !== null) {
            message.SentPurchaseorder = SentPurchaseorder.fromPartial(object.SentPurchaseorder);
        }
        else {
            message.SentPurchaseorder = undefined;
        }
        return message;
    },
};
const baseQueryAllSentPurchaseorderRequest = {};
export const QueryAllSentPurchaseorderRequest = {
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
            ...baseQueryAllSentPurchaseorderRequest,
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
            ...baseQueryAllSentPurchaseorderRequest,
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
            ...baseQueryAllSentPurchaseorderRequest,
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
const baseQueryAllSentPurchaseorderResponse = {};
export const QueryAllSentPurchaseorderResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.SentPurchaseorder) {
            SentPurchaseorder.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryAllSentPurchaseorderResponse,
        };
        message.SentPurchaseorder = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SentPurchaseorder.push(SentPurchaseorder.decode(reader, reader.uint32()));
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
            ...baseQueryAllSentPurchaseorderResponse,
        };
        message.SentPurchaseorder = [];
        if (object.SentPurchaseorder !== undefined &&
            object.SentPurchaseorder !== null) {
            for (const e of object.SentPurchaseorder) {
                message.SentPurchaseorder.push(SentPurchaseorder.fromJSON(e));
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
        if (message.SentPurchaseorder) {
            obj.SentPurchaseorder = message.SentPurchaseorder.map((e) => e ? SentPurchaseorder.toJSON(e) : undefined);
        }
        else {
            obj.SentPurchaseorder = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSentPurchaseorderResponse,
        };
        message.SentPurchaseorder = [];
        if (object.SentPurchaseorder !== undefined &&
            object.SentPurchaseorder !== null) {
            for (const e of object.SentPurchaseorder) {
                message.SentPurchaseorder.push(SentPurchaseorder.fromPartial(e));
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
const baseQueryGetTimedoutPurchaseorderRequest = { id: 0 };
export const QueryGetTimedoutPurchaseorderRequest = {
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
            ...baseQueryGetTimedoutPurchaseorderRequest,
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
            ...baseQueryGetTimedoutPurchaseorderRequest,
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
            ...baseQueryGetTimedoutPurchaseorderRequest,
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
const baseQueryGetTimedoutPurchaseorderResponse = {};
export const QueryGetTimedoutPurchaseorderResponse = {
    encode(message, writer = Writer.create()) {
        if (message.TimedoutPurchaseorder !== undefined) {
            TimedoutPurchaseorder.encode(message.TimedoutPurchaseorder, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetTimedoutPurchaseorderResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.TimedoutPurchaseorder = TimedoutPurchaseorder.decode(reader, reader.uint32());
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
            ...baseQueryGetTimedoutPurchaseorderResponse,
        };
        if (object.TimedoutPurchaseorder !== undefined &&
            object.TimedoutPurchaseorder !== null) {
            message.TimedoutPurchaseorder = TimedoutPurchaseorder.fromJSON(object.TimedoutPurchaseorder);
        }
        else {
            message.TimedoutPurchaseorder = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.TimedoutPurchaseorder !== undefined &&
            (obj.TimedoutPurchaseorder = message.TimedoutPurchaseorder
                ? TimedoutPurchaseorder.toJSON(message.TimedoutPurchaseorder)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetTimedoutPurchaseorderResponse,
        };
        if (object.TimedoutPurchaseorder !== undefined &&
            object.TimedoutPurchaseorder !== null) {
            message.TimedoutPurchaseorder = TimedoutPurchaseorder.fromPartial(object.TimedoutPurchaseorder);
        }
        else {
            message.TimedoutPurchaseorder = undefined;
        }
        return message;
    },
};
const baseQueryAllTimedoutPurchaseorderRequest = {};
export const QueryAllTimedoutPurchaseorderRequest = {
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
            ...baseQueryAllTimedoutPurchaseorderRequest,
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
            ...baseQueryAllTimedoutPurchaseorderRequest,
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
            ...baseQueryAllTimedoutPurchaseorderRequest,
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
const baseQueryAllTimedoutPurchaseorderResponse = {};
export const QueryAllTimedoutPurchaseorderResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.TimedoutPurchaseorder) {
            TimedoutPurchaseorder.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryAllTimedoutPurchaseorderResponse,
        };
        message.TimedoutPurchaseorder = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.TimedoutPurchaseorder.push(TimedoutPurchaseorder.decode(reader, reader.uint32()));
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
            ...baseQueryAllTimedoutPurchaseorderResponse,
        };
        message.TimedoutPurchaseorder = [];
        if (object.TimedoutPurchaseorder !== undefined &&
            object.TimedoutPurchaseorder !== null) {
            for (const e of object.TimedoutPurchaseorder) {
                message.TimedoutPurchaseorder.push(TimedoutPurchaseorder.fromJSON(e));
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
        if (message.TimedoutPurchaseorder) {
            obj.TimedoutPurchaseorder = message.TimedoutPurchaseorder.map((e) => e ? TimedoutPurchaseorder.toJSON(e) : undefined);
        }
        else {
            obj.TimedoutPurchaseorder = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllTimedoutPurchaseorderResponse,
        };
        message.TimedoutPurchaseorder = [];
        if (object.TimedoutPurchaseorder !== undefined &&
            object.TimedoutPurchaseorder !== null) {
            for (const e of object.TimedoutPurchaseorder) {
                message.TimedoutPurchaseorder.push(TimedoutPurchaseorder.fromPartial(e));
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
    Purchaseorder(request) {
        const data = QueryGetPurchaseorderRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.purchaseorder.Query", "Purchaseorder", data);
        return promise.then((data) => QueryGetPurchaseorderResponse.decode(new Reader(data)));
    }
    PurchaseorderAll(request) {
        const data = QueryAllPurchaseorderRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.purchaseorder.Query", "PurchaseorderAll", data);
        return promise.then((data) => QueryAllPurchaseorderResponse.decode(new Reader(data)));
    }
    SentPurchaseorder(request) {
        const data = QueryGetSentPurchaseorderRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.purchaseorder.Query", "SentPurchaseorder", data);
        return promise.then((data) => QueryGetSentPurchaseorderResponse.decode(new Reader(data)));
    }
    SentPurchaseorderAll(request) {
        const data = QueryAllSentPurchaseorderRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.purchaseorder.Query", "SentPurchaseorderAll", data);
        return promise.then((data) => QueryAllSentPurchaseorderResponse.decode(new Reader(data)));
    }
    TimedoutPurchaseorder(request) {
        const data = QueryGetTimedoutPurchaseorderRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.purchaseorder.Query", "TimedoutPurchaseorder", data);
        return promise.then((data) => QueryGetTimedoutPurchaseorderResponse.decode(new Reader(data)));
    }
    TimedoutPurchaseorderAll(request) {
        const data = QueryAllTimedoutPurchaseorderRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.purchaseorder.Query", "TimedoutPurchaseorderAll", data);
        return promise.then((data) => QueryAllTimedoutPurchaseorderResponse.decode(new Reader(data)));
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
