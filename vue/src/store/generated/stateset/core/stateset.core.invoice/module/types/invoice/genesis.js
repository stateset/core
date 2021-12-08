/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Invoice } from "../invoice/invoice";
export const protobufPackage = "stateset.core.invoice";
const baseGenesisState = { invoiceCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.invoiceList) {
            Invoice.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.invoiceCount !== 0) {
            writer.uint32(16).uint64(message.invoiceCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.invoiceList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.invoiceList.push(Invoice.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.invoiceCount = longToNumber(reader.uint64());
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
        message.invoiceList = [];
        if (object.invoiceList !== undefined && object.invoiceList !== null) {
            for (const e of object.invoiceList) {
                message.invoiceList.push(Invoice.fromJSON(e));
            }
        }
        if (object.invoiceCount !== undefined && object.invoiceCount !== null) {
            message.invoiceCount = Number(object.invoiceCount);
        }
        else {
            message.invoiceCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.invoiceList) {
            obj.invoiceList = message.invoiceList.map((e) => e ? Invoice.toJSON(e) : undefined);
        }
        else {
            obj.invoiceList = [];
        }
        message.invoiceCount !== undefined &&
            (obj.invoiceCount = message.invoiceCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.invoiceList = [];
        if (object.invoiceList !== undefined && object.invoiceList !== null) {
            for (const e of object.invoiceList) {
                message.invoiceList.push(Invoice.fromPartial(e));
            }
        }
        if (object.invoiceCount !== undefined && object.invoiceCount !== null) {
            message.invoiceCount = object.invoiceCount;
        }
        else {
            message.invoiceCount = 0;
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
