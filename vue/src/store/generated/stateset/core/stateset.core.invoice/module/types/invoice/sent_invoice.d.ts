import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.invoice";
export interface SentInvoice {
    id: number;
    did: string;
    chain: string;
    creator: string;
}
export declare const SentInvoice: {
    encode(message: SentInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SentInvoice;
    fromJSON(object: any): SentInvoice;
    toJSON(message: SentInvoice): unknown;
    fromPartial(object: DeepPartial<SentInvoice>): SentInvoice;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
