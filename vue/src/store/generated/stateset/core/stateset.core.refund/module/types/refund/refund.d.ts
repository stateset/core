import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.refund";
export interface Refund {
    id: number;
    creator: string;
    did: string;
    amount: string;
    fee: string;
    deadline: string;
    state: string;
}
export declare const Refund: {
    encode(message: Refund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Refund;
    fromJSON(object: any): Refund;
    toJSON(message: Refund): unknown;
    fromPartial(object: DeepPartial<Refund>): Refund;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
