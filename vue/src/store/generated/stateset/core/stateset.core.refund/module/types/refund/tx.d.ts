import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.refund";
export interface MsgRequestRefund {
    creator: string;
    did: string;
    amount: string;
    fee: string;
    deadline: string;
}
export interface MsgRequestRefundResponse {
}
export interface MsgApproveRefund {
    creator: string;
    id: number;
}
export interface MsgApproveRefundResponse {
}
export interface MsgRejectRefund {
    creator: string;
    id: number;
}
export interface MsgRejectRefundResponse {
}
export declare const MsgRequestRefund: {
    encode(message: MsgRequestRefund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRequestRefund;
    fromJSON(object: any): MsgRequestRefund;
    toJSON(message: MsgRequestRefund): unknown;
    fromPartial(object: DeepPartial<MsgRequestRefund>): MsgRequestRefund;
};
export declare const MsgRequestRefundResponse: {
    encode(_: MsgRequestRefundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRequestRefundResponse;
    fromJSON(_: any): MsgRequestRefundResponse;
    toJSON(_: MsgRequestRefundResponse): unknown;
    fromPartial(_: DeepPartial<MsgRequestRefundResponse>): MsgRequestRefundResponse;
};
export declare const MsgApproveRefund: {
    encode(message: MsgApproveRefund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveRefund;
    fromJSON(object: any): MsgApproveRefund;
    toJSON(message: MsgApproveRefund): unknown;
    fromPartial(object: DeepPartial<MsgApproveRefund>): MsgApproveRefund;
};
export declare const MsgApproveRefundResponse: {
    encode(_: MsgApproveRefundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveRefundResponse;
    fromJSON(_: any): MsgApproveRefundResponse;
    toJSON(_: MsgApproveRefundResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveRefundResponse>): MsgApproveRefundResponse;
};
export declare const MsgRejectRefund: {
    encode(message: MsgRejectRefund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectRefund;
    fromJSON(object: any): MsgRejectRefund;
    toJSON(message: MsgRejectRefund): unknown;
    fromPartial(object: DeepPartial<MsgRejectRefund>): MsgRejectRefund;
};
export declare const MsgRejectRefundResponse: {
    encode(_: MsgRejectRefundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectRefundResponse;
    fromJSON(_: any): MsgRejectRefundResponse;
    toJSON(_: MsgRejectRefundResponse): unknown;
    fromPartial(_: DeepPartial<MsgRejectRefundResponse>): MsgRejectRefundResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    RequestRefund(request: MsgRequestRefund): Promise<MsgRequestRefundResponse>;
    ApproveRefund(request: MsgApproveRefund): Promise<MsgApproveRefundResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RejectRefund(request: MsgRejectRefund): Promise<MsgRejectRefundResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    RequestRefund(request: MsgRequestRefund): Promise<MsgRequestRefundResponse>;
    ApproveRefund(request: MsgApproveRefund): Promise<MsgApproveRefundResponse>;
    RejectRefund(request: MsgRejectRefund): Promise<MsgRejectRefundResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
