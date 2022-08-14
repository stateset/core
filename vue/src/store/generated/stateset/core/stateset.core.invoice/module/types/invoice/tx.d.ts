import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.invoice";
export interface MsgFactorInvoice {
    creator: string;
    id: number;
}
export interface MsgFactorInvoiceResponse {
}
export interface MsgCreateSentInvoice {
    creator: string;
    did: string;
    chain: string;
}
export interface MsgCreateSentInvoiceResponse {
    id: number;
}
export interface MsgUpdateSentInvoice {
    creator: string;
    id: number;
    did: string;
    chain: string;
}
export interface MsgUpdateSentInvoiceResponse {
}
export interface MsgDeleteSentInvoice {
    creator: string;
    id: number;
}
export interface MsgDeleteSentInvoiceResponse {
}
export interface MsgCreateTimedoutInvoice {
    creator: string;
    did: string;
    chain: string;
}
export interface MsgCreateTimedoutInvoiceResponse {
    id: number;
}
export interface MsgUpdateTimedoutInvoice {
    creator: string;
    id: number;
    did: string;
    chain: string;
}
export interface MsgUpdateTimedoutInvoiceResponse {
}
export interface MsgDeleteTimedoutInvoice {
    creator: string;
    id: number;
}
export interface MsgDeleteTimedoutInvoiceResponse {
}
export interface MsgCreateInvoice {
    creator: string;
    id: string;
    did: string;
    amount: string;
    state: string;
}
export interface MsgCreateInvoiceResponse {
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
export declare const MsgCreateSentInvoice: {
    encode(message: MsgCreateSentInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSentInvoice;
    fromJSON(object: any): MsgCreateSentInvoice;
    toJSON(message: MsgCreateSentInvoice): unknown;
    fromPartial(object: DeepPartial<MsgCreateSentInvoice>): MsgCreateSentInvoice;
};
export declare const MsgCreateSentInvoiceResponse: {
    encode(message: MsgCreateSentInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSentInvoiceResponse;
    fromJSON(object: any): MsgCreateSentInvoiceResponse;
    toJSON(message: MsgCreateSentInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateSentInvoiceResponse>): MsgCreateSentInvoiceResponse;
};
export declare const MsgUpdateSentInvoice: {
    encode(message: MsgUpdateSentInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentInvoice;
    fromJSON(object: any): MsgUpdateSentInvoice;
    toJSON(message: MsgUpdateSentInvoice): unknown;
    fromPartial(object: DeepPartial<MsgUpdateSentInvoice>): MsgUpdateSentInvoice;
};
export declare const MsgUpdateSentInvoiceResponse: {
    encode(_: MsgUpdateSentInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentInvoiceResponse;
    fromJSON(_: any): MsgUpdateSentInvoiceResponse;
    toJSON(_: MsgUpdateSentInvoiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateSentInvoiceResponse>): MsgUpdateSentInvoiceResponse;
};
export declare const MsgDeleteSentInvoice: {
    encode(message: MsgDeleteSentInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentInvoice;
    fromJSON(object: any): MsgDeleteSentInvoice;
    toJSON(message: MsgDeleteSentInvoice): unknown;
    fromPartial(object: DeepPartial<MsgDeleteSentInvoice>): MsgDeleteSentInvoice;
};
export declare const MsgDeleteSentInvoiceResponse: {
    encode(_: MsgDeleteSentInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentInvoiceResponse;
    fromJSON(_: any): MsgDeleteSentInvoiceResponse;
    toJSON(_: MsgDeleteSentInvoiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteSentInvoiceResponse>): MsgDeleteSentInvoiceResponse;
};
export declare const MsgCreateTimedoutInvoice: {
    encode(message: MsgCreateTimedoutInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTimedoutInvoice;
    fromJSON(object: any): MsgCreateTimedoutInvoice;
    toJSON(message: MsgCreateTimedoutInvoice): unknown;
    fromPartial(object: DeepPartial<MsgCreateTimedoutInvoice>): MsgCreateTimedoutInvoice;
};
export declare const MsgCreateTimedoutInvoiceResponse: {
    encode(message: MsgCreateTimedoutInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTimedoutInvoiceResponse;
    fromJSON(object: any): MsgCreateTimedoutInvoiceResponse;
    toJSON(message: MsgCreateTimedoutInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateTimedoutInvoiceResponse>): MsgCreateTimedoutInvoiceResponse;
};
export declare const MsgUpdateTimedoutInvoice: {
    encode(message: MsgUpdateTimedoutInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTimedoutInvoice;
    fromJSON(object: any): MsgUpdateTimedoutInvoice;
    toJSON(message: MsgUpdateTimedoutInvoice): unknown;
    fromPartial(object: DeepPartial<MsgUpdateTimedoutInvoice>): MsgUpdateTimedoutInvoice;
};
export declare const MsgUpdateTimedoutInvoiceResponse: {
    encode(_: MsgUpdateTimedoutInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTimedoutInvoiceResponse;
    fromJSON(_: any): MsgUpdateTimedoutInvoiceResponse;
    toJSON(_: MsgUpdateTimedoutInvoiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateTimedoutInvoiceResponse>): MsgUpdateTimedoutInvoiceResponse;
};
export declare const MsgDeleteTimedoutInvoice: {
    encode(message: MsgDeleteTimedoutInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTimedoutInvoice;
    fromJSON(object: any): MsgDeleteTimedoutInvoice;
    toJSON(message: MsgDeleteTimedoutInvoice): unknown;
    fromPartial(object: DeepPartial<MsgDeleteTimedoutInvoice>): MsgDeleteTimedoutInvoice;
};
export declare const MsgDeleteTimedoutInvoiceResponse: {
    encode(_: MsgDeleteTimedoutInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTimedoutInvoiceResponse;
    fromJSON(_: any): MsgDeleteTimedoutInvoiceResponse;
    toJSON(_: MsgDeleteTimedoutInvoiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteTimedoutInvoiceResponse>): MsgDeleteTimedoutInvoiceResponse;
};
export declare const MsgCreateInvoice: {
    encode(message: MsgCreateInvoice, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateInvoice;
    fromJSON(object: any): MsgCreateInvoice;
    toJSON(message: MsgCreateInvoice): unknown;
    fromPartial(object: DeepPartial<MsgCreateInvoice>): MsgCreateInvoice;
};
export declare const MsgCreateInvoiceResponse: {
    encode(_: MsgCreateInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateInvoiceResponse;
    fromJSON(_: any): MsgCreateInvoiceResponse;
    toJSON(_: MsgCreateInvoiceResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateInvoiceResponse>): MsgCreateInvoiceResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    FactorInvoice(request: MsgFactorInvoice): Promise<MsgFactorInvoiceResponse>;
    CreateSentInvoice(request: MsgCreateSentInvoice): Promise<MsgCreateSentInvoiceResponse>;
    UpdateSentInvoice(request: MsgUpdateSentInvoice): Promise<MsgUpdateSentInvoiceResponse>;
    DeleteSentInvoice(request: MsgDeleteSentInvoice): Promise<MsgDeleteSentInvoiceResponse>;
    CreateTimedoutInvoice(request: MsgCreateTimedoutInvoice): Promise<MsgCreateTimedoutInvoiceResponse>;
    UpdateTimedoutInvoice(request: MsgUpdateTimedoutInvoice): Promise<MsgUpdateTimedoutInvoiceResponse>;
    DeleteTimedoutInvoice(request: MsgDeleteTimedoutInvoice): Promise<MsgDeleteTimedoutInvoiceResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateInvoice(request: MsgCreateInvoice): Promise<MsgCreateInvoiceResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    FactorInvoice(request: MsgFactorInvoice): Promise<MsgFactorInvoiceResponse>;
    CreateSentInvoice(request: MsgCreateSentInvoice): Promise<MsgCreateSentInvoiceResponse>;
    UpdateSentInvoice(request: MsgUpdateSentInvoice): Promise<MsgUpdateSentInvoiceResponse>;
    DeleteSentInvoice(request: MsgDeleteSentInvoice): Promise<MsgDeleteSentInvoiceResponse>;
    CreateTimedoutInvoice(request: MsgCreateTimedoutInvoice): Promise<MsgCreateTimedoutInvoiceResponse>;
    UpdateTimedoutInvoice(request: MsgUpdateTimedoutInvoice): Promise<MsgUpdateTimedoutInvoiceResponse>;
    DeleteTimedoutInvoice(request: MsgDeleteTimedoutInvoice): Promise<MsgDeleteTimedoutInvoiceResponse>;
    CreateInvoice(request: MsgCreateInvoice): Promise<MsgCreateInvoiceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
