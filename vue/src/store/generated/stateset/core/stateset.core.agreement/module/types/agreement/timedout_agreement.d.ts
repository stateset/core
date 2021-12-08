import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.agreement";
export interface TimedoutAgreement {
    id: number;
    did: string;
    chain: string;
    creator: string;
}
export declare const TimedoutAgreement: {
    encode(message: TimedoutAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): TimedoutAgreement;
    fromJSON(object: any): TimedoutAgreement;
    toJSON(message: TimedoutAgreement): unknown;
    fromPartial(object: DeepPartial<TimedoutAgreement>): TimedoutAgreement;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
