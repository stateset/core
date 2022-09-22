import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.purchaseorder";
/** EventCreatePurchaseOrder is an event emitted when an purchaseorder is created. */
export interface EventCreatePurchaseOrder {
    /** purchaseorder_id is the unique ID of purchaseorder */
    purchaseorderId: string;
    /** creator is the creator of the purchaseorder */
    creator: string;
}
/** EventCompleted is an event emitted when an purchaseorder is completed. */
export interface EventCompleted {
    /** agreement_id is the unique ID of purchaseorder */
    purchaseorderId: string;
    /** creator is the creator of the purchaseorder */
    creator: string;
}
/** EventCancelled is an event emitted when an purchaseorder is cancelled. */
export interface EventCancelled {
    /** agreement_id is the unique ID of purchaseorder */
    purchaseorderId: string;
    /** creator is the creator of the purchaseorder */
    creator: string;
}
/** EventFinanced is an event emitted when an purchaseorder is financed. */
export interface EventFinanced {
    /** purchaseorder_id is the unique ID of purchaseorder */
    purchaseorderId: string;
    /** creator is the creator of the purchaseorder */
    creator: string;
}
export declare const EventCreatePurchaseOrder: {
    encode(message: EventCreatePurchaseOrder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventCreatePurchaseOrder;
    fromJSON(object: any): EventCreatePurchaseOrder;
    toJSON(message: EventCreatePurchaseOrder): unknown;
    fromPartial(object: DeepPartial<EventCreatePurchaseOrder>): EventCreatePurchaseOrder;
};
export declare const EventCompleted: {
    encode(message: EventCompleted, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventCompleted;
    fromJSON(object: any): EventCompleted;
    toJSON(message: EventCompleted): unknown;
    fromPartial(object: DeepPartial<EventCompleted>): EventCompleted;
};
export declare const EventCancelled: {
    encode(message: EventCancelled, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventCancelled;
    fromJSON(object: any): EventCancelled;
    toJSON(message: EventCancelled): unknown;
    fromPartial(object: DeepPartial<EventCancelled>): EventCancelled;
};
export declare const EventFinanced: {
    encode(message: EventFinanced, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventFinanced;
    fromJSON(object: any): EventFinanced;
    toJSON(message: EventFinanced): unknown;
    fromPartial(object: DeepPartial<EventFinanced>): EventFinanced;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
