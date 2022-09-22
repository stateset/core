/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "stateset.core.invoice";
const baseEventCreateInvoice = { invoiceId: "", creator: "" };
export const EventCreateInvoice = {
    encode(message, writer = Writer.create()) {
        if (message.invoiceId !== "") {
            writer.uint32(10).string(message.invoiceId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventCreateInvoice };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.invoiceId = reader.string();
                    break;
                case 2:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseEventCreateInvoice };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = String(object.invoiceId);
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventCreateInvoice };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = object.invoiceId;
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseEventPaid = { invoiceId: "", creator: "" };
export const EventPaid = {
    encode(message, writer = Writer.create()) {
        if (message.invoiceId !== "") {
            writer.uint32(10).string(message.invoiceId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventPaid };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.invoiceId = reader.string();
                    break;
                case 2:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseEventPaid };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = String(object.invoiceId);
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventPaid };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = object.invoiceId;
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseEventVoided = { invoiceId: "", creator: "" };
export const EventVoided = {
    encode(message, writer = Writer.create()) {
        if (message.invoiceId !== "") {
            writer.uint32(10).string(message.invoiceId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventVoided };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.invoiceId = reader.string();
                    break;
                case 2:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseEventVoided };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = String(object.invoiceId);
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventVoided };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = object.invoiceId;
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
const baseEventFactored = { invoiceId: "", creator: "" };
export const EventFactored = {
    encode(message, writer = Writer.create()) {
        if (message.invoiceId !== "") {
            writer.uint32(10).string(message.invoiceId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventFactored };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.invoiceId = reader.string();
                    break;
                case 2:
                    message.creator = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseEventFactored };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = String(object.invoiceId);
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.invoiceId !== undefined && (obj.invoiceId = message.invoiceId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventFactored };
        if (object.invoiceId !== undefined && object.invoiceId !== null) {
            message.invoiceId = object.invoiceId;
        }
        else {
            message.invoiceId = "";
        }
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        return message;
    },
};
