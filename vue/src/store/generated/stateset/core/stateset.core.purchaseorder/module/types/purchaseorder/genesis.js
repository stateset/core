/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
export const protobufPackage = "stateset.core.purchaseorder";
const baseGenesisState = { purchaseorderCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.purchaseorderList) {
            Purchaseorder.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.purchaseorderCount !== 0) {
            writer.uint32(16).uint64(message.purchaseorderCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.purchaseorderList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.purchaseorderList.push(Purchaseorder.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.purchaseorderCount = longToNumber(reader.uint64());
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
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.purchaseorderList = [];
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
