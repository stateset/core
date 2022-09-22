/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
import { SentPurchaseorder } from "../purchaseorder/sent_purchaseorder";
import { TimedoutPurchaseorder } from "../purchaseorder/timedout_purchaseorder";
export const protobufPackage = "stateset.core.purchaseorder";
const baseGenesisState = {
    purchaseorderCount: 0,
    sentPurchaseorderCount: 0,
    timedoutPurchaseorderCount: 0,
};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.purchaseorderList) {
            Purchaseorder.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.purchaseorderCount !== 0) {
            writer.uint32(16).uint64(message.purchaseorderCount);
        }
        for (const v of message.sentPurchaseorderList) {
            SentPurchaseorder.encode(v, writer.uint32(26).fork()).ldelim();
        }
        if (message.sentPurchaseorderCount !== 0) {
            writer.uint32(32).uint64(message.sentPurchaseorderCount);
        }
        for (const v of message.timedoutPurchaseorderList) {
            TimedoutPurchaseorder.encode(v, writer.uint32(42).fork()).ldelim();
        }
        if (message.timedoutPurchaseorderCount !== 0) {
            writer.uint32(48).uint64(message.timedoutPurchaseorderCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.purchaseorderList = [];
        message.sentPurchaseorderList = [];
        message.timedoutPurchaseorderList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.purchaseorderList.push(Purchaseorder.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.purchaseorderCount = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.sentPurchaseorderList.push(SentPurchaseorder.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.sentPurchaseorderCount = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.timedoutPurchaseorderList.push(TimedoutPurchaseorder.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.timedoutPurchaseorderCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.purchaseorderList = [];
        message.sentPurchaseorderList = [];
        message.timedoutPurchaseorderList = [];
        if (object.purchaseorderList !== undefined &&
            object.purchaseorderList !== null) {
            for (const e of object.purchaseorderList) {
                message.purchaseorderList.push(Purchaseorder.fromJSON(e));
            }
        }
        if (object.purchaseorderCount !== undefined &&
            object.purchaseorderCount !== null) {
            message.purchaseorderCount = Number(object.purchaseorderCount);
        }
        else {
            message.purchaseorderCount = 0;
        }
        if (object.sentPurchaseorderList !== undefined &&
            object.sentPurchaseorderList !== null) {
            for (const e of object.sentPurchaseorderList) {
                message.sentPurchaseorderList.push(SentPurchaseorder.fromJSON(e));
            }
        }
        if (object.sentPurchaseorderCount !== undefined &&
            object.sentPurchaseorderCount !== null) {
            message.sentPurchaseorderCount = Number(object.sentPurchaseorderCount);
        }
        else {
            message.sentPurchaseorderCount = 0;
        }
        if (object.timedoutPurchaseorderList !== undefined &&
            object.timedoutPurchaseorderList !== null) {
            for (const e of object.timedoutPurchaseorderList) {
                message.timedoutPurchaseorderList.push(TimedoutPurchaseorder.fromJSON(e));
            }
        }
        if (object.timedoutPurchaseorderCount !== undefined &&
            object.timedoutPurchaseorderCount !== null) {
            message.timedoutPurchaseorderCount = Number(object.timedoutPurchaseorderCount);
        }
        else {
            message.timedoutPurchaseorderCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.purchaseorderList) {
            obj.purchaseorderList = message.purchaseorderList.map((e) => e ? Purchaseorder.toJSON(e) : undefined);
        }
        else {
            obj.purchaseorderList = [];
        }
        message.purchaseorderCount !== undefined &&
            (obj.purchaseorderCount = message.purchaseorderCount);
        if (message.sentPurchaseorderList) {
            obj.sentPurchaseorderList = message.sentPurchaseorderList.map((e) => e ? SentPurchaseorder.toJSON(e) : undefined);
        }
        else {
            obj.sentPurchaseorderList = [];
        }
        message.sentPurchaseorderCount !== undefined &&
            (obj.sentPurchaseorderCount = message.sentPurchaseorderCount);
        if (message.timedoutPurchaseorderList) {
            obj.timedoutPurchaseorderList = message.timedoutPurchaseorderList.map((e) => (e ? TimedoutPurchaseorder.toJSON(e) : undefined));
        }
        else {
            obj.timedoutPurchaseorderList = [];
        }
        message.timedoutPurchaseorderCount !== undefined &&
            (obj.timedoutPurchaseorderCount = message.timedoutPurchaseorderCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.purchaseorderList = [];
        message.sentPurchaseorderList = [];
        message.timedoutPurchaseorderList = [];
        if (object.purchaseorderList !== undefined &&
            object.purchaseorderList !== null) {
            for (const e of object.purchaseorderList) {
                message.purchaseorderList.push(Purchaseorder.fromPartial(e));
            }
        }
        if (object.purchaseorderCount !== undefined &&
            object.purchaseorderCount !== null) {
            message.purchaseorderCount = object.purchaseorderCount;
        }
        else {
            message.purchaseorderCount = 0;
        }
        if (object.sentPurchaseorderList !== undefined &&
            object.sentPurchaseorderList !== null) {
            for (const e of object.sentPurchaseorderList) {
                message.sentPurchaseorderList.push(SentPurchaseorder.fromPartial(e));
            }
        }
        if (object.sentPurchaseorderCount !== undefined &&
            object.sentPurchaseorderCount !== null) {
            message.sentPurchaseorderCount = object.sentPurchaseorderCount;
        }
        else {
            message.sentPurchaseorderCount = 0;
        }
        if (object.timedoutPurchaseorderList !== undefined &&
            object.timedoutPurchaseorderList !== null) {
            for (const e of object.timedoutPurchaseorderList) {
                message.timedoutPurchaseorderList.push(TimedoutPurchaseorder.fromPartial(e));
            }
        }
        if (object.timedoutPurchaseorderCount !== undefined &&
            object.timedoutPurchaseorderCount !== null) {
            message.timedoutPurchaseorderCount = object.timedoutPurchaseorderCount;
        }
        else {
            message.timedoutPurchaseorderCount = 0;
        }
        return message;
    },
};
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
