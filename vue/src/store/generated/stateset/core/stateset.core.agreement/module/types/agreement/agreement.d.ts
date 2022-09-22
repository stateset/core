import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.agreement";
export interface Agreement {
    id: number;
    did: string;
    uri: string;
    amount: string;
    state: string;
}
export declare const Agreement: {
    encode(message: Agreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Agreement;
    fromJSON(object: any): Agreement;
    toJSON(message: Agreement): unknown;
    fromPartial(object: DeepPartial<Agreement>): Agreement;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
