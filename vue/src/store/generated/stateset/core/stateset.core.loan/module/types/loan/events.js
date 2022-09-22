/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "stateset.core.loan";
const baseEventCreateLoan = { loanId: "", creator: "" };
export const EventCreateLoan = {
    encode(message, writer = Writer.create()) {
        if (message.loanId !== "") {
            writer.uint32(10).string(message.loanId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventCreateLoan };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.loanId = reader.string();
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
        const message = { ...baseEventCreateLoan };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = String(object.loanId);
        }
        else {
            message.loanId = "";
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
        message.loanId !== undefined && (obj.loanId = message.loanId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventCreateLoan };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = object.loanId;
        }
        else {
            message.loanId = "";
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
const baseEventLoanRequested = { loanId: "", creator: "" };
export const EventLoanRequested = {
    encode(message, writer = Writer.create()) {
        if (message.loanId !== "") {
            writer.uint32(10).string(message.loanId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventLoanRequested };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.loanId = reader.string();
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
        const message = { ...baseEventLoanRequested };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = String(object.loanId);
        }
        else {
            message.loanId = "";
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
        message.loanId !== undefined && (obj.loanId = message.loanId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventLoanRequested };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = object.loanId;
        }
        else {
            message.loanId = "";
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
const baseEventApproved = { loanId: "", creator: "" };
export const EventApproved = {
    encode(message, writer = Writer.create()) {
        if (message.loanId !== "") {
            writer.uint32(10).string(message.loanId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventApproved };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.loanId = reader.string();
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
        const message = { ...baseEventApproved };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = String(object.loanId);
        }
        else {
            message.loanId = "";
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
        message.loanId !== undefined && (obj.loanId = message.loanId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventApproved };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = object.loanId;
        }
        else {
            message.loanId = "";
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
const baseEventRepaid = { loanId: "", creator: "" };
export const EventRepaid = {
    encode(message, writer = Writer.create()) {
        if (message.loanId !== "") {
            writer.uint32(10).string(message.loanId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventRepaid };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.loanId = reader.string();
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
        const message = { ...baseEventRepaid };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = String(object.loanId);
        }
        else {
            message.loanId = "";
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
        message.loanId !== undefined && (obj.loanId = message.loanId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventRepaid };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = object.loanId;
        }
        else {
            message.loanId = "";
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
const baseEventLiquidated = { loanId: "", creator: "" };
export const EventLiquidated = {
    encode(message, writer = Writer.create()) {
        if (message.loanId !== "") {
            writer.uint32(10).string(message.loanId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventLiquidated };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.loanId = reader.string();
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
        const message = { ...baseEventLiquidated };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = String(object.loanId);
        }
        else {
            message.loanId = "";
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
        message.loanId !== undefined && (obj.loanId = message.loanId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventLiquidated };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = object.loanId;
        }
        else {
            message.loanId = "";
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
const baseEventCancelled = { loanId: "", creator: "" };
export const EventCancelled = {
    encode(message, writer = Writer.create()) {
        if (message.loanId !== "") {
            writer.uint32(10).string(message.loanId);
        }
        if (message.creator !== "") {
            writer.uint32(18).string(message.creator);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEventCancelled };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.loanId = reader.string();
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
        const message = { ...baseEventCancelled };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = String(object.loanId);
        }
        else {
            message.loanId = "";
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
        message.loanId !== undefined && (obj.loanId = message.loanId);
        message.creator !== undefined && (obj.creator = message.creator);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEventCancelled };
        if (object.loanId !== undefined && object.loanId !== null) {
            message.loanId = object.loanId;
        }
        else {
            message.loanId = "";
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
