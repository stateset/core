import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.purchaseorder";
export interface MsgFinancePurchaseorder {
    creator: string;
    id: number;
}
export interface MsgFinancePurchaseorderResponse {
}
export declare const MsgFinancePurchaseorder: {
    encode(message: MsgFinancePurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFinancePurchaseorder;
    fromJSON(object: any): MsgFinancePurchaseorder;
    toJSON(message: MsgFinancePurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgFinancePurchaseorder>): MsgFinancePurchaseorder;
};
export declare const MsgFinancePurchaseorderResponse: {
    encode(_: MsgFinancePurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFinancePurchaseorderResponse;
    fromJSON(_: any): MsgFinancePurchaseorderResponse;
    toJSON(_: MsgFinancePurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgFinancePurchaseorderResponse>): MsgFinancePurchaseorderResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    FinancePurchaseorder(request: MsgFinancePurchaseorder): Promise<MsgFinancePurchaseorderResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    FinancePurchaseorder(request: MsgFinancePurchaseorder): Promise<MsgFinancePurchaseorderResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
