import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.agreement";
export interface MsgActivateAgreement {
    creator: string;
    id: number;
}
export interface MsgActivateAgreementResponse {
}
export interface MsgExpireAgreement {
    creator: string;
    id: number;
}
export interface MsgExpireAgreementResponse {
}
export interface MsgRenewAgreement {
    creator: string;
    id: number;
}
export interface MsgRenewAgreementResponse {
}
export interface MsgTerminateAgreement {
    creator: string;
    id: number;
}
export interface MsgTerminateAgreementResponse {
}
export interface MsgCreateSentAgreement {
    creator: string;
    did: string;
    chain: string;
}
export interface MsgCreateSentAgreementResponse {
    id: number;
}
export interface MsgUpdateSentAgreement {
    creator: string;
    id: number;
    did: string;
    chain: string;
}
export interface MsgUpdateSentAgreementResponse {
}
export interface MsgDeleteSentAgreement {
    creator: string;
    id: number;
}
export interface MsgDeleteSentAgreementResponse {
}
export interface MsgCreateTimedoutAgreement {
    creator: string;
    did: string;
    chain: string;
}
export interface MsgCreateTimedoutAgreementResponse {
    id: number;
}
export interface MsgUpdateTimedoutAgreement {
    creator: string;
    id: number;
    did: string;
    chain: string;
}
export interface MsgUpdateTimedoutAgreementResponse {
}
export interface MsgDeleteTimedoutAgreement {
    creator: string;
    id: number;
}
export interface MsgDeleteTimedoutAgreementResponse {
}
export declare const MsgActivateAgreement: {
    encode(message: MsgActivateAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgActivateAgreement;
    fromJSON(object: any): MsgActivateAgreement;
    toJSON(message: MsgActivateAgreement): unknown;
    fromPartial(object: DeepPartial<MsgActivateAgreement>): MsgActivateAgreement;
};
export declare const MsgActivateAgreementResponse: {
    encode(_: MsgActivateAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgActivateAgreementResponse;
    fromJSON(_: any): MsgActivateAgreementResponse;
    toJSON(_: MsgActivateAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgActivateAgreementResponse>): MsgActivateAgreementResponse;
};
export declare const MsgExpireAgreement: {
    encode(message: MsgExpireAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgExpireAgreement;
    fromJSON(object: any): MsgExpireAgreement;
    toJSON(message: MsgExpireAgreement): unknown;
    fromPartial(object: DeepPartial<MsgExpireAgreement>): MsgExpireAgreement;
};
export declare const MsgExpireAgreementResponse: {
    encode(_: MsgExpireAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgExpireAgreementResponse;
    fromJSON(_: any): MsgExpireAgreementResponse;
    toJSON(_: MsgExpireAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgExpireAgreementResponse>): MsgExpireAgreementResponse;
};
export declare const MsgRenewAgreement: {
    encode(message: MsgRenewAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRenewAgreement;
    fromJSON(object: any): MsgRenewAgreement;
    toJSON(message: MsgRenewAgreement): unknown;
    fromPartial(object: DeepPartial<MsgRenewAgreement>): MsgRenewAgreement;
};
export declare const MsgRenewAgreementResponse: {
    encode(_: MsgRenewAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRenewAgreementResponse;
    fromJSON(_: any): MsgRenewAgreementResponse;
    toJSON(_: MsgRenewAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgRenewAgreementResponse>): MsgRenewAgreementResponse;
};
export declare const MsgTerminateAgreement: {
    encode(message: MsgTerminateAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTerminateAgreement;
    fromJSON(object: any): MsgTerminateAgreement;
    toJSON(message: MsgTerminateAgreement): unknown;
    fromPartial(object: DeepPartial<MsgTerminateAgreement>): MsgTerminateAgreement;
};
export declare const MsgTerminateAgreementResponse: {
    encode(_: MsgTerminateAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgTerminateAgreementResponse;
    fromJSON(_: any): MsgTerminateAgreementResponse;
    toJSON(_: MsgTerminateAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgTerminateAgreementResponse>): MsgTerminateAgreementResponse;
};
export declare const MsgCreateSentAgreement: {
    encode(message: MsgCreateSentAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSentAgreement;
    fromJSON(object: any): MsgCreateSentAgreement;
    toJSON(message: MsgCreateSentAgreement): unknown;
    fromPartial(object: DeepPartial<MsgCreateSentAgreement>): MsgCreateSentAgreement;
};
export declare const MsgCreateSentAgreementResponse: {
    encode(message: MsgCreateSentAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateSentAgreementResponse;
    fromJSON(object: any): MsgCreateSentAgreementResponse;
    toJSON(message: MsgCreateSentAgreementResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateSentAgreementResponse>): MsgCreateSentAgreementResponse;
};
export declare const MsgUpdateSentAgreement: {
    encode(message: MsgUpdateSentAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentAgreement;
    fromJSON(object: any): MsgUpdateSentAgreement;
    toJSON(message: MsgUpdateSentAgreement): unknown;
    fromPartial(object: DeepPartial<MsgUpdateSentAgreement>): MsgUpdateSentAgreement;
};
export declare const MsgUpdateSentAgreementResponse: {
    encode(_: MsgUpdateSentAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateSentAgreementResponse;
    fromJSON(_: any): MsgUpdateSentAgreementResponse;
    toJSON(_: MsgUpdateSentAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateSentAgreementResponse>): MsgUpdateSentAgreementResponse;
};
export declare const MsgDeleteSentAgreement: {
    encode(message: MsgDeleteSentAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentAgreement;
    fromJSON(object: any): MsgDeleteSentAgreement;
    toJSON(message: MsgDeleteSentAgreement): unknown;
    fromPartial(object: DeepPartial<MsgDeleteSentAgreement>): MsgDeleteSentAgreement;
};
export declare const MsgDeleteSentAgreementResponse: {
    encode(_: MsgDeleteSentAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteSentAgreementResponse;
    fromJSON(_: any): MsgDeleteSentAgreementResponse;
    toJSON(_: MsgDeleteSentAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteSentAgreementResponse>): MsgDeleteSentAgreementResponse;
};
export declare const MsgCreateTimedoutAgreement: {
    encode(message: MsgCreateTimedoutAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTimedoutAgreement;
    fromJSON(object: any): MsgCreateTimedoutAgreement;
    toJSON(message: MsgCreateTimedoutAgreement): unknown;
    fromPartial(object: DeepPartial<MsgCreateTimedoutAgreement>): MsgCreateTimedoutAgreement;
};
export declare const MsgCreateTimedoutAgreementResponse: {
    encode(message: MsgCreateTimedoutAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateTimedoutAgreementResponse;
    fromJSON(object: any): MsgCreateTimedoutAgreementResponse;
    toJSON(message: MsgCreateTimedoutAgreementResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateTimedoutAgreementResponse>): MsgCreateTimedoutAgreementResponse;
};
export declare const MsgUpdateTimedoutAgreement: {
    encode(message: MsgUpdateTimedoutAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTimedoutAgreement;
    fromJSON(object: any): MsgUpdateTimedoutAgreement;
    toJSON(message: MsgUpdateTimedoutAgreement): unknown;
    fromPartial(object: DeepPartial<MsgUpdateTimedoutAgreement>): MsgUpdateTimedoutAgreement;
};
export declare const MsgUpdateTimedoutAgreementResponse: {
    encode(_: MsgUpdateTimedoutAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateTimedoutAgreementResponse;
    fromJSON(_: any): MsgUpdateTimedoutAgreementResponse;
    toJSON(_: MsgUpdateTimedoutAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateTimedoutAgreementResponse>): MsgUpdateTimedoutAgreementResponse;
};
export declare const MsgDeleteTimedoutAgreement: {
    encode(message: MsgDeleteTimedoutAgreement, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTimedoutAgreement;
    fromJSON(object: any): MsgDeleteTimedoutAgreement;
    toJSON(message: MsgDeleteTimedoutAgreement): unknown;
    fromPartial(object: DeepPartial<MsgDeleteTimedoutAgreement>): MsgDeleteTimedoutAgreement;
};
export declare const MsgDeleteTimedoutAgreementResponse: {
    encode(_: MsgDeleteTimedoutAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteTimedoutAgreementResponse;
    fromJSON(_: any): MsgDeleteTimedoutAgreementResponse;
    toJSON(_: MsgDeleteTimedoutAgreementResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteTimedoutAgreementResponse>): MsgDeleteTimedoutAgreementResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    ActivateAgreement(request: MsgActivateAgreement): Promise<MsgActivateAgreementResponse>;
    ExpireAgreement(request: MsgExpireAgreement): Promise<MsgExpireAgreementResponse>;
    RenewAgreement(request: MsgRenewAgreement): Promise<MsgRenewAgreementResponse>;
    TerminateAgreement(request: MsgTerminateAgreement): Promise<MsgTerminateAgreementResponse>;
    CreateSentAgreement(request: MsgCreateSentAgreement): Promise<MsgCreateSentAgreementResponse>;
    UpdateSentAgreement(request: MsgUpdateSentAgreement): Promise<MsgUpdateSentAgreementResponse>;
    DeleteSentAgreement(request: MsgDeleteSentAgreement): Promise<MsgDeleteSentAgreementResponse>;
    CreateTimedoutAgreement(request: MsgCreateTimedoutAgreement): Promise<MsgCreateTimedoutAgreementResponse>;
    UpdateTimedoutAgreement(request: MsgUpdateTimedoutAgreement): Promise<MsgUpdateTimedoutAgreementResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteTimedoutAgreement(request: MsgDeleteTimedoutAgreement): Promise<MsgDeleteTimedoutAgreementResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    ActivateAgreement(request: MsgActivateAgreement): Promise<MsgActivateAgreementResponse>;
    ExpireAgreement(request: MsgExpireAgreement): Promise<MsgExpireAgreementResponse>;
    RenewAgreement(request: MsgRenewAgreement): Promise<MsgRenewAgreementResponse>;
    TerminateAgreement(request: MsgTerminateAgreement): Promise<MsgTerminateAgreementResponse>;
    CreateSentAgreement(request: MsgCreateSentAgreement): Promise<MsgCreateSentAgreementResponse>;
    UpdateSentAgreement(request: MsgUpdateSentAgreement): Promise<MsgUpdateSentAgreementResponse>;
    DeleteSentAgreement(request: MsgDeleteSentAgreement): Promise<MsgDeleteSentAgreementResponse>;
    CreateTimedoutAgreement(request: MsgCreateTimedoutAgreement): Promise<MsgCreateTimedoutAgreementResponse>;
    UpdateTimedoutAgreement(request: MsgUpdateTimedoutAgreement): Promise<MsgUpdateTimedoutAgreementResponse>;
    DeleteTimedoutAgreement(request: MsgDeleteTimedoutAgreement): Promise<MsgDeleteTimedoutAgreementResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
