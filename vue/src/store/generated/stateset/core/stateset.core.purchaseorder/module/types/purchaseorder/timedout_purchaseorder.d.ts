import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.purchaseorder";
export interface TimedoutPurchaseorder {
    id: number;
    did: string;
    chain: string;
    creator: string;
}
export declare const TimedoutPurchaseorder: {
    encode(message: TimedoutPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): TimedoutPurchaseorder;
    fromJSON(object: any): TimedoutPurchaseorder;
    toJSON(message: TimedoutPurchaseorder): unknown;
    fromPartial(object: DeepPartial<TimedoutPurchaseorder>): TimedoutPurchaseorder;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
