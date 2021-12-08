/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
import { PageRequest, PageResponse, } from "../cosmos/base/query/v1beta1/pagination";
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
