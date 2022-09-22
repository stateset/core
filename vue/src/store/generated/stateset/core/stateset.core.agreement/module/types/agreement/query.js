/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { SentAgreement } from "../agreement/sent_agreement";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
import { TimedoutAgreement } from "../agreement/timedout_agreement";
import { Agreement } from "../agreement/agreement";
export const protobufPackage = "stateset.core.agreement";
const baseQueryGetSentAgreementRequest = { id: 0 };
export const QueryGetSentAgreementRequest = {
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
            ...baseQueryGetSentAgreementRequest,
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
            ...baseQueryGetSentAgreementRequest,
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
            ...baseQueryGetSentAgreementRequest,
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
const baseQueryGetSentAgreementResponse = {};
export const QueryGetSentAgreementResponse = {
    encode(message, writer = Writer.create()) {
        if (message.SentAgreement !== undefined) {
            SentAgreement.encode(message.SentAgreement, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetSentAgreementResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetSentAgreementResponse,
        };
        if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
            message.SentAgreement = SentAgreement.fromJSON(object.SentAgreement);
        }
        else {
            message.SentAgreement = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.SentAgreement !== undefined &&
            (obj.SentAgreement = message.SentAgreement
                ? SentAgreement.toJSON(message.SentAgreement)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetSentAgreementResponse,
        };
        if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
            message.SentAgreement = SentAgreement.fromPartial(object.SentAgreement);
        }
        else {
            message.SentAgreement = undefined;
        }
        return message;
    },
};
const baseQueryAllSentAgreementRequest = {};
export const QueryAllSentAgreementRequest = {
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
            ...baseQueryAllSentAgreementRequest,
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
            ...baseQueryAllSentAgreementRequest,
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
            ...baseQueryAllSentAgreementRequest,
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
const baseQueryAllSentAgreementResponse = {};
export const QueryAllSentAgreementResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.SentAgreement) {
            SentAgreement.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryAllSentAgreementResponse,
        };
        message.SentAgreement = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.SentAgreement.push(SentAgreement.decode(reader, reader.uint32()));
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
            ...baseQueryAllSentAgreementResponse,
        };
        message.SentAgreement = [];
        if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
            for (const e of object.SentAgreement) {
                message.SentAgreement.push(SentAgreement.fromJSON(e));
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
        if (message.SentAgreement) {
            obj.SentAgreement = message.SentAgreement.map((e) => e ? SentAgreement.toJSON(e) : undefined);
        }
        else {
            obj.SentAgreement = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllSentAgreementResponse,
        };
        message.SentAgreement = [];
        if (object.SentAgreement !== undefined && object.SentAgreement !== null) {
            for (const e of object.SentAgreement) {
                message.SentAgreement.push(SentAgreement.fromPartial(e));
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
const baseQueryGetTimedoutAgreementRequest = { id: 0 };
export const QueryGetTimedoutAgreementRequest = {
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
            ...baseQueryGetTimedoutAgreementRequest,
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
            ...baseQueryGetTimedoutAgreementRequest,
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
            ...baseQueryGetTimedoutAgreementRequest,
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
const baseQueryGetTimedoutAgreementResponse = {};
export const QueryGetTimedoutAgreementResponse = {
    encode(message, writer = Writer.create()) {
        if (message.TimedoutAgreement !== undefined) {
            TimedoutAgreement.encode(message.TimedoutAgreement, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetTimedoutAgreementResponse,
        };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.TimedoutAgreement = TimedoutAgreement.decode(reader, reader.uint32());
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
            ...baseQueryGetTimedoutAgreementResponse,
        };
        if (object.TimedoutAgreement !== undefined &&
            object.TimedoutAgreement !== null) {
            message.TimedoutAgreement = TimedoutAgreement.fromJSON(object.TimedoutAgreement);
        }
        else {
            message.TimedoutAgreement = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.TimedoutAgreement !== undefined &&
            (obj.TimedoutAgreement = message.TimedoutAgreement
                ? TimedoutAgreement.toJSON(message.TimedoutAgreement)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetTimedoutAgreementResponse,
        };
        if (object.TimedoutAgreement !== undefined &&
            object.TimedoutAgreement !== null) {
            message.TimedoutAgreement = TimedoutAgreement.fromPartial(object.TimedoutAgreement);
        }
        else {
            message.TimedoutAgreement = undefined;
        }
        return message;
    },
};
const baseQueryAllTimedoutAgreementRequest = {};
export const QueryAllTimedoutAgreementRequest = {
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
            ...baseQueryAllTimedoutAgreementRequest,
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
            ...baseQueryAllTimedoutAgreementRequest,
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
            ...baseQueryAllTimedoutAgreementRequest,
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
const baseQueryAllTimedoutAgreementResponse = {};
export const QueryAllTimedoutAgreementResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.TimedoutAgreement) {
            TimedoutAgreement.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryAllTimedoutAgreementResponse,
        };
        message.TimedoutAgreement = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.TimedoutAgreement.push(TimedoutAgreement.decode(reader, reader.uint32()));
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
            ...baseQueryAllTimedoutAgreementResponse,
        };
        message.TimedoutAgreement = [];
        if (object.TimedoutAgreement !== undefined &&
            object.TimedoutAgreement !== null) {
            for (const e of object.TimedoutAgreement) {
                message.TimedoutAgreement.push(TimedoutAgreement.fromJSON(e));
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
        if (message.TimedoutAgreement) {
            obj.TimedoutAgreement = message.TimedoutAgreement.map((e) => e ? TimedoutAgreement.toJSON(e) : undefined);
        }
        else {
            obj.TimedoutAgreement = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllTimedoutAgreementResponse,
        };
        message.TimedoutAgreement = [];
        if (object.TimedoutAgreement !== undefined &&
            object.TimedoutAgreement !== null) {
            for (const e of object.TimedoutAgreement) {
                message.TimedoutAgreement.push(TimedoutAgreement.fromPartial(e));
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
const baseQueryGetAgreementRequest = { id: 0 };
export const QueryGetAgreementRequest = {
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
            ...baseQueryGetAgreementRequest,
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
            ...baseQueryGetAgreementRequest,
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
            ...baseQueryGetAgreementRequest,
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
const baseQueryGetAgreementResponse = {};
export const QueryGetAgreementResponse = {
    encode(message, writer = Writer.create()) {
        if (message.Agreement !== undefined) {
            Agreement.encode(message.Agreement, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetAgreementResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetAgreementResponse,
        };
        if (object.Agreement !== undefined && object.Agreement !== null) {
            message.Agreement = Agreement.fromJSON(object.Agreement);
        }
        else {
            message.Agreement = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.Agreement !== undefined &&
            (obj.Agreement = message.Agreement
                ? Agreement.toJSON(message.Agreement)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetAgreementResponse,
        };
        if (object.Agreement !== undefined && object.Agreement !== null) {
            message.Agreement = Agreement.fromPartial(object.Agreement);
        }
        else {
            message.Agreement = undefined;
        }
        return message;
    },
};
const baseQueryAllAgreementRequest = {};
export const QueryAllAgreementRequest = {
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
            ...baseQueryAllAgreementRequest,
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
            ...baseQueryAllAgreementRequest,
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
            ...baseQueryAllAgreementRequest,
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
const baseQueryAllAgreementResponse = {};
export const QueryAllAgreementResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.Agreement) {
            Agreement.encode(v, writer.uint32(10).fork()).ldelim();
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
            ...baseQueryAllAgreementResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryAllAgreementResponse,
        };
        message.Agreement = [];
        if (object.Agreement !== undefined && object.Agreement !== null) {
            for (const e of object.Agreement) {
                message.Agreement.push(Agreement.fromJSON(e));
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
        if (message.Agreement) {
            obj.Agreement = message.Agreement.map((e) => e ? Agreement.toJSON(e) : undefined);
        }
        else {
            obj.Agreement = [];
        }
        message.pagination !== undefined &&
            (obj.pagination = message.pagination
                ? PageResponse.toJSON(message.pagination)
                : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryAllAgreementResponse,
        };
        message.Agreement = [];
        if (object.Agreement !== undefined && object.Agreement !== null) {
            for (const e of object.Agreement) {
                message.Agreement.push(Agreement.fromPartial(e));
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
    SentAgreement(request) {
        const data = QueryGetSentAgreementRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.agreement.Query", "SentAgreement", data);
        return promise.then((data) => QueryGetSentAgreementResponse.decode(new Reader(data)));
    }
    SentAgreementAll(request) {
        const data = QueryAllSentAgreementRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.agreement.Query", "SentAgreementAll", data);
        return promise.then((data) => QueryAllSentAgreementResponse.decode(new Reader(data)));
    }
    TimedoutAgreement(request) {
        const data = QueryGetTimedoutAgreementRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.agreement.Query", "TimedoutAgreement", data);
        return promise.then((data) => QueryGetTimedoutAgreementResponse.decode(new Reader(data)));
    }
    TimedoutAgreementAll(request) {
        const data = QueryAllTimedoutAgreementRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.agreement.Query", "TimedoutAgreementAll", data);
        return promise.then((data) => QueryAllTimedoutAgreementResponse.decode(new Reader(data)));
    }
    Agreement(request) {
        const data = QueryGetAgreementRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.agreement.Query", "Agreement", data);
        return promise.then((data) => QueryGetAgreementResponse.decode(new Reader(data)));
    }
    AgreementAll(request) {
        const data = QueryAllAgreementRequest.encode(request).finish();
        const promise = this.rpc.request("stateset.core.agreement.Query", "AgreementAll", data);
        return promise.then((data) => QueryAllAgreementResponse.decode(new Reader(data)));
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
