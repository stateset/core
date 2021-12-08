import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.invoice";
export interface MsgFactorInvoice {
    creator: string;
    id: number;
}
export interface MsgFactorInvoiceResponse {
}
export declare const MsgFactorInvoice: {
    encode(message: MsgFactorInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFactorInvoice;
    fromJSON(object: any): MsgFactorInvoice;
    toJSON(message: MsgFactorInvoice): unknown;
    fromPartial(object: DeepPartial<MsgFactorInvoice>): MsgFactorInvoice;
};
export declare const MsgFactorInvoiceResponse: {
    encode(_: MsgFactorInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFactorInvoiceResponse;
    fromJSON(_: any): MsgFactorInvoiceResponse;
    toJSON(_: MsgFactorInvoiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgFactorInvoiceResponse>): MsgFactorInvoiceResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    FactorInvoice(request: MsgFactorInvoice): Promise<MsgFactorInvoiceResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    FactorInvoice(request: MsgFactorInvoice): Promise<MsgFactorInvoiceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
