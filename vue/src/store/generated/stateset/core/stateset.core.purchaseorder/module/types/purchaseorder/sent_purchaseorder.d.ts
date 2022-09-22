import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.purchaseorder";
export interface SentPurchaseorder {
    id: number;
    did: string;
    chain: string;
    creator: string;
}
export declare const SentPurchaseorder: {
    encode(message: SentPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SentPurchaseorder;
    fromJSON(object: any): SentPurchaseorder;
    toJSON(message: SentPurchaseorder): unknown;
    fromPartial(object: DeepPartial<SentPurchaseorder>): SentPurchaseorder;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
