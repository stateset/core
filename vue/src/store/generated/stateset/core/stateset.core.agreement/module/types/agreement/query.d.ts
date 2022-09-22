import { Reader, Writer } from "protobufjs/minimal";
import { SentAgreement } from "../agreement/sent_agreement";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { TimedoutAgreement } from "../agreement/timedout_agreement";
import { Agreement } from "../agreement/agreement";
export declare const protobufPackage = "stateset.core.agreement";
export interface QueryGetSentAgreementRequest {
    id: number;
}
export interface QueryGetSentAgreementResponse {
    SentAgreement: SentAgreement | undefined;
}
export interface QueryAllSentAgreementRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSentAgreementResponse {
    SentAgreement: SentAgreement[];
    pagination: PageResponse | undefined;
}
export interface QueryGetTimedoutAgreementRequest {
    id: number;
}
export interface QueryGetTimedoutAgreementResponse {
    TimedoutAgreement: TimedoutAgreement | undefined;
}
export interface QueryAllTimedoutAgreementRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllTimedoutAgreementResponse {
    TimedoutAgreement: TimedoutAgreement[];
    pagination: PageResponse | undefined;
}
export interface QueryGetAgreementRequest {
    id: number;
}
export interface QueryGetAgreementResponse {
    Agreement: Agreement | undefined;
}
export interface QueryAllAgreementRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllAgreementResponse {
    Agreement: Agreement[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetSentAgreementRequest: {
    encode(message: QueryGetSentAgreementRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSentAgreementRequest;
    fromJSON(object: any): QueryGetSentAgreementRequest;
    toJSON(message: QueryGetSentAgreementRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSentAgreementRequest>): QueryGetSentAgreementRequest;
};
export declare const QueryGetSentAgreementResponse: {
    encode(message: QueryGetSentAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSentAgreementResponse;
    fromJSON(object: any): QueryGetSentAgreementResponse;
    toJSON(message: QueryGetSentAgreementResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSentAgreementResponse>): QueryGetSentAgreementResponse;
};
export declare const QueryAllSentAgreementRequest: {
    encode(message: QueryAllSentAgreementRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSentAgreementRequest;
    fromJSON(object: any): QueryAllSentAgreementRequest;
    toJSON(message: QueryAllSentAgreementRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSentAgreementRequest>): QueryAllSentAgreementRequest;
};
export declare const QueryAllSentAgreementResponse: {
    encode(message: QueryAllSentAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSentAgreementResponse;
    fromJSON(object: any): QueryAllSentAgreementResponse;
    toJSON(message: QueryAllSentAgreementResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSentAgreementResponse>): QueryAllSentAgreementResponse;
};
export declare const QueryGetTimedoutAgreementRequest: {
    encode(message: QueryGetTimedoutAgreementRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutAgreementRequest;
    fromJSON(object: any): QueryGetTimedoutAgreementRequest;
    toJSON(message: QueryGetTimedoutAgreementRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTimedoutAgreementRequest>): QueryGetTimedoutAgreementRequest;
};
export declare const QueryGetTimedoutAgreementResponse: {
    encode(message: QueryGetTimedoutAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutAgreementResponse;
    fromJSON(object: any): QueryGetTimedoutAgreementResponse;
    toJSON(message: QueryGetTimedoutAgreementResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTimedoutAgreementResponse>): QueryGetTimedoutAgreementResponse;
};
export declare const QueryAllTimedoutAgreementRequest: {
    encode(message: QueryAllTimedoutAgreementRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTimedoutAgreementRequest;
    fromJSON(object: any): QueryAllTimedoutAgreementRequest;
    toJSON(message: QueryAllTimedoutAgreementRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllTimedoutAgreementRequest>): QueryAllTimedoutAgreementRequest;
};
export declare const QueryAllTimedoutAgreementResponse: {
    encode(message: QueryAllTimedoutAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTimedoutAgreementResponse;
    fromJSON(object: any): QueryAllTimedoutAgreementResponse;
    toJSON(message: QueryAllTimedoutAgreementResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllTimedoutAgreementResponse>): QueryAllTimedoutAgreementResponse;
};
export declare const QueryGetAgreementRequest: {
    encode(message: QueryGetAgreementRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAgreementRequest;
    fromJSON(object: any): QueryGetAgreementRequest;
    toJSON(message: QueryGetAgreementRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetAgreementRequest>): QueryGetAgreementRequest;
};
export declare const QueryGetAgreementResponse: {
    encode(message: QueryGetAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetAgreementResponse;
    fromJSON(object: any): QueryGetAgreementResponse;
    toJSON(message: QueryGetAgreementResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetAgreementResponse>): QueryGetAgreementResponse;
};
export declare const QueryAllAgreementRequest: {
    encode(message: QueryAllAgreementRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAgreementRequest;
    fromJSON(object: any): QueryAllAgreementRequest;
    toJSON(message: QueryAllAgreementRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllAgreementRequest>): QueryAllAgreementRequest;
};
export declare const QueryAllAgreementResponse: {
    encode(message: QueryAllAgreementResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllAgreementResponse;
    fromJSON(object: any): QueryAllAgreementResponse;
    toJSON(message: QueryAllAgreementResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllAgreementResponse>): QueryAllAgreementResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a sentAgreement by id. */
    SentAgreement(request: QueryGetSentAgreementRequest): Promise<QueryGetSentAgreementResponse>;
    /** Queries a list of sentAgreement items. */
    SentAgreementAll(request: QueryAllSentAgreementRequest): Promise<QueryAllSentAgreementResponse>;
    /** Queries a timedoutAgreement by id. */
    TimedoutAgreement(request: QueryGetTimedoutAgreementRequest): Promise<QueryGetTimedoutAgreementResponse>;
    /** Queries a list of timedoutAgreement items. */
    TimedoutAgreementAll(request: QueryAllTimedoutAgreementRequest): Promise<QueryAllTimedoutAgreementResponse>;
    /** Queries a agreement by id. */
    Agreement(request: QueryGetAgreementRequest): Promise<QueryGetAgreementResponse>;
    /** Queries a list of agreement items. */
    AgreementAll(request: QueryAllAgreementRequest): Promise<QueryAllAgreementResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    SentAgreement(request: QueryGetSentAgreementRequest): Promise<QueryGetSentAgreementResponse>;
    SentAgreementAll(request: QueryAllSentAgreementRequest): Promise<QueryAllSentAgreementResponse>;
    TimedoutAgreement(request: QueryGetTimedoutAgreementRequest): Promise<QueryGetTimedoutAgreementResponse>;
    TimedoutAgreementAll(request: QueryAllTimedoutAgreementRequest): Promise<QueryAllTimedoutAgreementResponse>;
    Agreement(request: QueryGetAgreementRequest): Promise<QueryGetAgreementResponse>;
    AgreementAll(request: QueryAllAgreementRequest): Promise<QueryAllAgreementResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
