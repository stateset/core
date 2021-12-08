import { Writer, Reader } from "protobufjs/minimal";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
import { SentPurchaseorder } from "../purchaseorder/sent_purchaseorder";
import { TimedoutPurchaseorder } from "../purchaseorder/timedout_purchaseorder";
export declare const protobufPackage = "stateset.core.purchaseorder";
/** GenesisState defines the purchaseorder module's genesis state. */
export interface GenesisState {
    purchaseorderList: Purchaseorder[];
    purchaseorderCount: number;
    sentPurchaseorderList: SentPurchaseorder[];
    sentPurchaseorderCount: number;
    timedoutPurchaseorderList: TimedoutPurchaseorder[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    timedoutPurchaseorderCount: number;
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
