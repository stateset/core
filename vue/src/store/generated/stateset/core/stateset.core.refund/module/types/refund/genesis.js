/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../refund/params";
import { Refund } from "../refund/refund";
export const protobufPackage = "stateset.core.refund";
const baseGenesisState = { refundCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.refundList) {
            Refund.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.refundCount !== 0) {
            writer.uint32(24).uint64(message.refundCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.refundList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.refundList.push(Refund.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.refundCount = longToNumber(reader.uint64());
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
        message.refundList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.refundList !== undefined && object.refundList !== null) {
            for (const e of object.refundList) {
                message.refundList.push(Refund.fromJSON(e));
            }
        }
        if (object.refundCount !== undefined && object.refundCount !== null) {
            message.refundCount = Number(object.refundCount);
        }
        else {
            message.refundCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.refundList) {
            obj.refundList = message.refundList.map((e) => e ? Refund.toJSON(e) : undefined);
        }
        else {
            obj.refundList = [];
        }
        message.refundCount !== undefined &&
            (obj.refundCount = message.refundCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.refundList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.refundList !== undefined && object.refundList !== null) {
            for (const e of object.refundList) {
                message.refundList.push(Refund.fromPartial(e));
            }
        }
        if (object.refundCount !== undefined && object.refundCount !== null) {
            message.refundCount = object.refundCount;
        }
        else {
            message.refundCount = 0;
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
