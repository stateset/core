import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.purchaseorder";
export interface MsgFinancePurchaseorder {
    creator: string;
    id: number;
}
export interface MsgFinancePurchaseorderResponse {
}
export interface MsgCancelPurchaseorder {
    creator: string;
    id: number;
}
export interface MsgCancelPurchaseorderResponse {
}
export interface MsgCompletePurchaseorder {
    creator: string;
    id: number;
}
export interface MsgCompletePurchaseorderResponse {
}
export interface MsgCreateSentPurchaseorder {
    creator: string;
    did: string;
    chain: string;
}
export interface MsgCreateSentPurchaseorderResponse {
    id: number;
}
export interface MsgUpdateSentPurchaseorder {
    creator: string;
    id: number;
    did: string;
    chain: string;
}
export interface MsgUpdateSentPurchaseorderResponse {
}
export interface MsgDeleteSentPurchaseorder {
    creator: string;
    id: number;
}
export interface MsgDeleteSentPurchaseorderResponse {
}
export interface MsgCreateTimedoutPurchaseorder {
    creator: string;
    did: string;
    chain: string;
}
export interface MsgCreateTimedoutPurchaseorderResponse {
    id: number;
}
export interface MsgUpdateTimedoutPurchaseorder {
    creator: string;
    id: number;
    did: string;
    chain: string;
}
export interface MsgUpdateTimedoutPurchaseorderResponse {
}
export interface MsgDeleteTimedoutPurchaseorder {
    creator: string;
    id: number;
}
export interface MsgDeleteTimedoutPurchaseorderResponse {
}
export interface MsgRequestPurchaseorder {
    creator: string;
    did: string;
    uri: string;
    amount: string;
    state: string;
}
export interface MsgRequestPurchaseorderResponse {
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
export declare const MsgCancelPurchaseorder: {
    encode(message: MsgCancelPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCancelPurchaseorder;
    fromJSON(object: any): MsgCancelPurchaseorder;
    toJSON(message: MsgCancelPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgCancelPurchaseorder>): MsgCancelPurchaseorder;
};
export declare const MsgCancelPurchaseorderResponse: {
    encode(_: MsgCancelPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCancelPurchaseorderResponse;
    fromJSON(_: any): MsgCancelPurchaseorderResponse;
    toJSON(_: MsgCancelPurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgCancelPurchaseorderResponse>): MsgCancelPurchaseorderResponse;
};
export declare const MsgCompletePurchaseorder: {
    encode(message: MsgCompletePurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCompletePurchaseorder;
    fromJSON(object: any): MsgCompletePurchaseorder;
    toJSON(message: MsgCompletePurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgCompletePurchaseorder>): MsgCompletePurchaseorder;
};
export declare const MsgCompletePurchaseorderResponse: {
    encode(_: MsgCompletePurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCompletePurchaseorderResponse;
    fromJSON(_: any): MsgCompletePurchaseorderResponse;
    toJSON(_: MsgCompletePurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgCompletePurchaseorderResponse>): MsgCompletePurchaseorderResponse;
};
export declare const MsgCreateSentPurchaseorder: {
    encode(message: MsgCreateSentPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSentPurchaseorder;
    fromJSON(object: any): MsgCreateSentPurchaseorder;
    toJSON(message: MsgCreateSentPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgCreateSentPurchaseorder>): MsgCreateSentPurchaseorder;
};
export declare const MsgCreateSentPurchaseorderResponse: {
    encode(message: MsgCreateSentPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSentPurchaseorderResponse;
    fromJSON(object: any): MsgCreateSentPurchaseorderResponse;
    toJSON(message: MsgCreateSentPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateSentPurchaseorderResponse>): MsgCreateSentPurchaseorderResponse;
};
export declare const MsgUpdateSentPurchaseorder: {
    encode(message: MsgUpdateSentPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentPurchaseorder;
    fromJSON(object: any): MsgUpdateSentPurchaseorder;
    toJSON(message: MsgUpdateSentPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgUpdateSentPurchaseorder>): MsgUpdateSentPurchaseorder;
};
export declare const MsgUpdateSentPurchaseorderResponse: {
    encode(_: MsgUpdateSentPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentPurchaseorderResponse;
    fromJSON(_: any): MsgUpdateSentPurchaseorderResponse;
    toJSON(_: MsgUpdateSentPurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateSentPurchaseorderResponse>): MsgUpdateSentPurchaseorderResponse;
};
export declare const MsgDeleteSentPurchaseorder: {
    encode(message: MsgDeleteSentPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentPurchaseorder;
    fromJSON(object: any): MsgDeleteSentPurchaseorder;
    toJSON(message: MsgDeleteSentPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgDeleteSentPurchaseorder>): MsgDeleteSentPurchaseorder;
};
export declare const MsgDeleteSentPurchaseorderResponse: {
    encode(_: MsgDeleteSentPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentPurchaseorderResponse;
    fromJSON(_: any): MsgDeleteSentPurchaseorderResponse;
    toJSON(_: MsgDeleteSentPurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteSentPurchaseorderResponse>): MsgDeleteSentPurchaseorderResponse;
};
export declare const MsgCreateTimedoutPurchaseorder: {
    encode(message: MsgCreateTimedoutPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTimedoutPurchaseorder;
    fromJSON(object: any): MsgCreateTimedoutPurchaseorder;
    toJSON(message: MsgCreateTimedoutPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgCreateTimedoutPurchaseorder>): MsgCreateTimedoutPurchaseorder;
};
export declare const MsgCreateTimedoutPurchaseorderResponse: {
    encode(message: MsgCreateTimedoutPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTimedoutPurchaseorderResponse;
    fromJSON(object: any): MsgCreateTimedoutPurchaseorderResponse;
    toJSON(message: MsgCreateTimedoutPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateTimedoutPurchaseorderResponse>): MsgCreateTimedoutPurchaseorderResponse;
};
export declare const MsgUpdateTimedoutPurchaseorder: {
    encode(message: MsgUpdateTimedoutPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTimedoutPurchaseorder;
    fromJSON(object: any): MsgUpdateTimedoutPurchaseorder;
    toJSON(message: MsgUpdateTimedoutPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgUpdateTimedoutPurchaseorder>): MsgUpdateTimedoutPurchaseorder;
};
export declare const MsgUpdateTimedoutPurchaseorderResponse: {
    encode(_: MsgUpdateTimedoutPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTimedoutPurchaseorderResponse;
    fromJSON(_: any): MsgUpdateTimedoutPurchaseorderResponse;
    toJSON(_: MsgUpdateTimedoutPurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateTimedoutPurchaseorderResponse>): MsgUpdateTimedoutPurchaseorderResponse;
};
export declare const MsgDeleteTimedoutPurchaseorder: {
    encode(message: MsgDeleteTimedoutPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTimedoutPurchaseorder;
    fromJSON(object: any): MsgDeleteTimedoutPurchaseorder;
    toJSON(message: MsgDeleteTimedoutPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgDeleteTimedoutPurchaseorder>): MsgDeleteTimedoutPurchaseorder;
};
export declare const MsgDeleteTimedoutPurchaseorderResponse: {
    encode(_: MsgDeleteTimedoutPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTimedoutPurchaseorderResponse;
    fromJSON(_: any): MsgDeleteTimedoutPurchaseorderResponse;
    toJSON(_: MsgDeleteTimedoutPurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteTimedoutPurchaseorderResponse>): MsgDeleteTimedoutPurchaseorderResponse;
};
export declare const MsgRequestPurchaseorder: {
    encode(message: MsgRequestPurchaseorder, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRequestPurchaseorder;
    fromJSON(object: any): MsgRequestPurchaseorder;
    toJSON(message: MsgRequestPurchaseorder): unknown;
    fromPartial(object: DeepPartial<MsgRequestPurchaseorder>): MsgRequestPurchaseorder;
};
export declare const MsgRequestPurchaseorderResponse: {
    encode(_: MsgRequestPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRequestPurchaseorderResponse;
    fromJSON(_: any): MsgRequestPurchaseorderResponse;
    toJSON(_: MsgRequestPurchaseorderResponse): unknown;
    fromPartial(_: DeepPartial<MsgRequestPurchaseorderResponse>): MsgRequestPurchaseorderResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    FinancePurchaseorder(request: MsgFinancePurchaseorder): Promise<MsgFinancePurchaseorderResponse>;
    CancelPurchaseorder(request: MsgCancelPurchaseorder): Promise<MsgCancelPurchaseorderResponse>;
    CompletePurchaseorder(request: MsgCompletePurchaseorder): Promise<MsgCompletePurchaseorderResponse>;
    CreateSentPurchaseorder(request: MsgCreateSentPurchaseorder): Promise<MsgCreateSentPurchaseorderResponse>;
    UpdateSentPurchaseorder(request: MsgUpdateSentPurchaseorder): Promise<MsgUpdateSentPurchaseorderResponse>;
    DeleteSentPurchaseorder(request: MsgDeleteSentPurchaseorder): Promise<MsgDeleteSentPurchaseorderResponse>;
    CreateTimedoutPurchaseorder(request: MsgCreateTimedoutPurchaseorder): Promise<MsgCreateTimedoutPurchaseorderResponse>;
    UpdateTimedoutPurchaseorder(request: MsgUpdateTimedoutPurchaseorder): Promise<MsgUpdateTimedoutPurchaseorderResponse>;
    DeleteTimedoutPurchaseorder(request: MsgDeleteTimedoutPurchaseorder): Promise<MsgDeleteTimedoutPurchaseorderResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RequestPurchaseorder(request: MsgRequestPurchaseorder): Promise<MsgRequestPurchaseorderResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    FinancePurchaseorder(request: MsgFinancePurchaseorder): Promise<MsgFinancePurchaseorderResponse>;
    CancelPurchaseorder(request: MsgCancelPurchaseorder): Promise<MsgCancelPurchaseorderResponse>;
    CompletePurchaseorder(request: MsgCompletePurchaseorder): Promise<MsgCompletePurchaseorderResponse>;
    CreateSentPurchaseorder(request: MsgCreateSentPurchaseorder): Promise<MsgCreateSentPurchaseorderResponse>;
    UpdateSentPurchaseorder(request: MsgUpdateSentPurchaseorder): Promise<MsgUpdateSentPurchaseorderResponse>;
    DeleteSentPurchaseorder(request: MsgDeleteSentPurchaseorder): Promise<MsgDeleteSentPurchaseorderResponse>;
    CreateTimedoutPurchaseorder(request: MsgCreateTimedoutPurchaseorder): Promise<MsgCreateTimedoutPurchaseorderResponse>;
    UpdateTimedoutPurchaseorder(request: MsgUpdateTimedoutPurchaseorder): Promise<MsgUpdateTimedoutPurchaseorderResponse>;
    DeleteTimedoutPurchaseorder(request: MsgDeleteTimedoutPurchaseorder): Promise<MsgDeleteTimedoutPurchaseorderResponse>;
    RequestPurchaseorder(request: MsgRequestPurchaseorder): Promise<MsgRequestPurchaseorderResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
