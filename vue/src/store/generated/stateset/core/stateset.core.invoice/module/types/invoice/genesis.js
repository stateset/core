/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Invoice } from "../invoice/invoice";
import { SentInvoice } from "../invoice/sent_invoice";
import { TimedoutInvoice } from "../invoice/timedout_invoice";
export const protobufPackage = "stateset.core.invoice";
const baseGenesisState = {
    invoiceCount: 0,
    sentInvoiceCount: 0,
    timedoutInvoiceCount: 0,
};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.invoiceList) {
            Invoice.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.invoiceCount !== 0) {
            writer.uint32(16).uint64(message.invoiceCount);
        }
        for (const v of message.sentInvoiceList) {
            SentInvoice.encode(v, writer.uint32(26).fork()).ldelim();
        }
        if (message.sentInvoiceCount !== 0) {
            writer.uint32(32).uint64(message.sentInvoiceCount);
        }
        for (const v of message.timedoutInvoiceList) {
            TimedoutInvoice.encode(v, writer.uint32(42).fork()).ldelim();
        }
        if (message.timedoutInvoiceCount !== 0) {
            writer.uint32(48).uint64(message.timedoutInvoiceCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.invoiceList = [];
        message.sentInvoiceList = [];
        message.timedoutInvoiceList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.invoiceList.push(Invoice.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.invoiceCount = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.sentInvoiceList.push(SentInvoice.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.sentInvoiceCount = longToNumber(reader.uint64());
                    break;
                case 5:
                    message.timedoutInvoiceList.push(TimedoutInvoice.decode(reader, reader.uint32()));
                    break;
                case 6:
                    message.timedoutInvoiceCount = longToNumber(reader.uint64());
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
        message.sentInvoiceList = [];
        message.timedoutInvoiceList = [];
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
        if (object.sentInvoiceList !== undefined &&
            object.sentInvoiceList !== null) {
            for (const e of object.sentInvoiceList) {
                message.sentInvoiceList.push(SentInvoice.fromJSON(e));
            }
        }
        if (object.sentInvoiceCount !== undefined &&
            object.sentInvoiceCount !== null) {
            message.sentInvoiceCount = Number(object.sentInvoiceCount);
        }
        else {
            message.sentInvoiceCount = 0;
        }
        if (object.timedoutInvoiceList !== undefined &&
            object.timedoutInvoiceList !== null) {
            for (const e of object.timedoutInvoiceList) {
                message.timedoutInvoiceList.push(TimedoutInvoice.fromJSON(e));
            }
        }
        if (object.timedoutInvoiceCount !== undefined &&
            object.timedoutInvoiceCount !== null) {
            message.timedoutInvoiceCount = Number(object.timedoutInvoiceCount);
        }
        else {
            message.timedoutInvoiceCount = 0;
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
        if (message.sentInvoiceList) {
            obj.sentInvoiceList = message.sentInvoiceList.map((e) => e ? SentInvoice.toJSON(e) : undefined);
        }
        else {
            obj.sentInvoiceList = [];
        }
        message.sentInvoiceCount !== undefined &&
            (obj.sentInvoiceCount = message.sentInvoiceCount);
        if (message.timedoutInvoiceList) {
            obj.timedoutInvoiceList = message.timedoutInvoiceList.map((e) => e ? TimedoutInvoice.toJSON(e) : undefined);
        }
        else {
            obj.timedoutInvoiceList = [];
        }
        message.timedoutInvoiceCount !== undefined &&
            (obj.timedoutInvoiceCount = message.timedoutInvoiceCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.invoiceList = [];
        message.sentInvoiceList = [];
        message.timedoutInvoiceList = [];
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
        if (object.sentInvoiceList !== undefined &&
            object.sentInvoiceList !== null) {
            for (const e of object.sentInvoiceList) {
                message.sentInvoiceList.push(SentInvoice.fromPartial(e));
            }
        }
        if (object.sentInvoiceCount !== undefined &&
            object.sentInvoiceCount !== null) {
            message.sentInvoiceCount = object.sentInvoiceCount;
        }
        else {
            message.sentInvoiceCount = 0;
        }
        if (object.timedoutInvoiceList !== undefined &&
            object.timedoutInvoiceList !== null) {
            for (const e of object.timedoutInvoiceList) {
                message.timedoutInvoiceList.push(TimedoutInvoice.fromPartial(e));
            }
        }
        if (object.timedoutInvoiceCount !== undefined &&
            object.timedoutInvoiceCount !== null) {
            message.timedoutInvoiceCount = object.timedoutInvoiceCount;
        }
        else {
            message.timedoutInvoiceCount = 0;
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
