import { Writer, Reader } from "protobufjs/minimal";
import { Invoice } from "../invoice/invoice";
export declare const protobufPackage = "stateset.core.invoice";
/** GenesisState defines the invoice module's genesis state. */
export interface GenesisState {
    invoiceList: Invoice[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    invoiceCount: number;
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
