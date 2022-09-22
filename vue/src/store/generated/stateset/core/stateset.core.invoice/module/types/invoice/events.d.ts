import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.invoice";
/** EventCreateInvoice is an event emitted when an invoice is created. */
export interface EventCreateInvoice {
    /** invoice_id is the unique ID of invoice */
    invoiceId: string;
    /** creator is the creator of the transaction */
    creator: string;
}
/** EventPaid is an event emitted when an invoice is paid. */
export interface EventPaid {
    /** agreement_id is the unique ID of invoice */
    invoiceId: string;
    /** creator is the creator of the invoice */
    creator: string;
}
/** EventVoided is an event emitted when an invoice is voided. */
export interface EventVoided {
    /** agreement_id is the unique ID of invoice */
    invoiceId: string;
    /** creator is the creator of the transaction */
    creator: string;
}
/** EventFactored is an event emitted when an invoice is factored. */
export interface EventFactored {
    /** invoice_id is the unique ID of invoice */
    invoiceId: string;
    /** creator is the creator of the transaction */
    creator: string;
}
export declare const EventCreateInvoice: {
    encode(message: EventCreateInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventCreateInvoice;
    fromJSON(object: any): EventCreateInvoice;
    toJSON(message: EventCreateInvoice): unknown;
    fromPartial(object: DeepPartial<EventCreateInvoice>): EventCreateInvoice;
};
export declare const EventPaid: {
    encode(message: EventPaid, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventPaid;
    fromJSON(object: any): EventPaid;
    toJSON(message: EventPaid): unknown;
    fromPartial(object: DeepPartial<EventPaid>): EventPaid;
};
export declare const EventVoided: {
    encode(message: EventVoided, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventVoided;
    fromJSON(object: any): EventVoided;
    toJSON(message: EventVoided): unknown;
    fromPartial(object: DeepPartial<EventVoided>): EventVoided;
};
export declare const EventFactored: {
    encode(message: EventFactored, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventFactored;
    fromJSON(object: any): EventFactored;
    toJSON(message: EventFactored): unknown;
    fromPartial(object: DeepPartial<EventFactored>): EventFactored;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
