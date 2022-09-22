import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.invoice";
export interface TimedoutInvoice {
    id: number;
    did: string;
    chain: string;
    creator: string;
}
export declare const TimedoutInvoice: {
    encode(message: TimedoutInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): TimedoutInvoice;
    fromJSON(object: any): TimedoutInvoice;
    toJSON(message: TimedoutInvoice): unknown;
    fromPartial(object: DeepPartial<TimedoutInvoice>): TimedoutInvoice;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
