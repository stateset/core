import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.agreement";
export interface SentAgreement {
    id: number;
    did: string;
    chain: string;
    creator: string;
}
export declare const SentAgreement: {
    encode(message: SentAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): SentAgreement;
    fromJSON(object: any): SentAgreement;
    toJSON(message: SentAgreement): unknown;
    fromPartial(object: DeepPartial<SentAgreement>): SentAgreement;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
