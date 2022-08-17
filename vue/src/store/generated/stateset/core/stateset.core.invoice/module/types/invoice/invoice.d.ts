import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.invoice";
export interface Invoice {
    id: number;
    did: string;
    uri: string;
    amount: string;
    state: string;
}
export declare const Invoice: {
    encode(message: Invoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Invoice;
    fromJSON(object: any): Invoice;
    toJSON(message: Invoice): unknown;
    fromPartial(object: DeepPartial<Invoice>): Invoice;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
