import { Reader, Writer } from "protobufjs/minimal";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { SentPurchaseorder } from "../purchaseorder/sent_purchaseorder";
import { TimedoutPurchaseorder } from "../purchaseorder/timedout_purchaseorder";
export declare const protobufPackage = "stateset.core.purchaseorder";
export interface QueryGetPurchaseorderRequest {
    id: number;
}
export interface QueryGetPurchaseorderResponse {
    Purchaseorder: Purchaseorder | undefined;
}
export interface QueryAllPurchaseorderRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllPurchaseorderResponse {
    Purchaseorder: Purchaseorder[];
    pagination: PageResponse | undefined;
}
export interface QueryGetSentPurchaseorderRequest {
    id: number;
}
export interface QueryGetSentPurchaseorderResponse {
    SentPurchaseorder: SentPurchaseorder | undefined;
}
export interface QueryAllSentPurchaseorderRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSentPurchaseorderResponse {
    SentPurchaseorder: SentPurchaseorder[];
    pagination: PageResponse | undefined;
}
export interface QueryGetTimedoutPurchaseorderRequest {
    id: number;
}
export interface QueryGetTimedoutPurchaseorderResponse {
    TimedoutPurchaseorder: TimedoutPurchaseorder | undefined;
}
export interface QueryAllTimedoutPurchaseorderRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllTimedoutPurchaseorderResponse {
    TimedoutPurchaseorder: TimedoutPurchaseorder[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetPurchaseorderRequest: {
    encode(message: QueryGetPurchaseorderRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPurchaseorderRequest;
    fromJSON(object: any): QueryGetPurchaseorderRequest;
    toJSON(message: QueryGetPurchaseorderRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetPurchaseorderRequest>): QueryGetPurchaseorderRequest;
};
export declare const QueryGetPurchaseorderResponse: {
    encode(message: QueryGetPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetPurchaseorderResponse;
    fromJSON(object: any): QueryGetPurchaseorderResponse;
    toJSON(message: QueryGetPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetPurchaseorderResponse>): QueryGetPurchaseorderResponse;
};
export declare const QueryAllPurchaseorderRequest: {
    encode(message: QueryAllPurchaseorderRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPurchaseorderRequest;
    fromJSON(object: any): QueryAllPurchaseorderRequest;
    toJSON(message: QueryAllPurchaseorderRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllPurchaseorderRequest>): QueryAllPurchaseorderRequest;
};
export declare const QueryAllPurchaseorderResponse: {
    encode(message: QueryAllPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllPurchaseorderResponse;
    fromJSON(object: any): QueryAllPurchaseorderResponse;
    toJSON(message: QueryAllPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllPurchaseorderResponse>): QueryAllPurchaseorderResponse;
};
export declare const QueryGetSentPurchaseorderRequest: {
    encode(message: QueryGetSentPurchaseorderRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSentPurchaseorderRequest;
    fromJSON(object: any): QueryGetSentPurchaseorderRequest;
    toJSON(message: QueryGetSentPurchaseorderRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSentPurchaseorderRequest>): QueryGetSentPurchaseorderRequest;
};
export declare const QueryGetSentPurchaseorderResponse: {
    encode(message: QueryGetSentPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSentPurchaseorderResponse;
    fromJSON(object: any): QueryGetSentPurchaseorderResponse;
    toJSON(message: QueryGetSentPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSentPurchaseorderResponse>): QueryGetSentPurchaseorderResponse;
};
export declare const QueryAllSentPurchaseorderRequest: {
    encode(message: QueryAllSentPurchaseorderRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSentPurchaseorderRequest;
    fromJSON(object: any): QueryAllSentPurchaseorderRequest;
    toJSON(message: QueryAllSentPurchaseorderRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSentPurchaseorderRequest>): QueryAllSentPurchaseorderRequest;
};
export declare const QueryAllSentPurchaseorderResponse: {
    encode(message: QueryAllSentPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSentPurchaseorderResponse;
    fromJSON(object: any): QueryAllSentPurchaseorderResponse;
    toJSON(message: QueryAllSentPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSentPurchaseorderResponse>): QueryAllSentPurchaseorderResponse;
};
export declare const QueryGetTimedoutPurchaseorderRequest: {
    encode(message: QueryGetTimedoutPurchaseorderRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutPurchaseorderRequest;
    fromJSON(object: any): QueryGetTimedoutPurchaseorderRequest;
    toJSON(message: QueryGetTimedoutPurchaseorderRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTimedoutPurchaseorderRequest>): QueryGetTimedoutPurchaseorderRequest;
};
export declare const QueryGetTimedoutPurchaseorderResponse: {
    encode(message: QueryGetTimedoutPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutPurchaseorderResponse;
    fromJSON(object: any): QueryGetTimedoutPurchaseorderResponse;
    toJSON(message: QueryGetTimedoutPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTimedoutPurchaseorderResponse>): QueryGetTimedoutPurchaseorderResponse;
};
export declare const QueryAllTimedoutPurchaseorderRequest: {
    encode(message: QueryAllTimedoutPurchaseorderRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTimedoutPurchaseorderRequest;
    fromJSON(object: any): QueryAllTimedoutPurchaseorderRequest;
    toJSON(message: QueryAllTimedoutPurchaseorderRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllTimedoutPurchaseorderRequest>): QueryAllTimedoutPurchaseorderRequest;
};
export declare const QueryAllTimedoutPurchaseorderResponse: {
    encode(message: QueryAllTimedoutPurchaseorderResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTimedoutPurchaseorderResponse;
    fromJSON(object: any): QueryAllTimedoutPurchaseorderResponse;
    toJSON(message: QueryAllTimedoutPurchaseorderResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllTimedoutPurchaseorderResponse>): QueryAllTimedoutPurchaseorderResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a purchaseorder by id. */
    Purchaseorder(request: QueryGetPurchaseorderRequest): Promise<QueryGetPurchaseorderResponse>;
    /** Queries a list of purchaseorder items. */
    PurchaseorderAll(request: QueryAllPurchaseorderRequest): Promise<QueryAllPurchaseorderResponse>;
    /** Queries a sentPurchaseorder by id. */
    SentPurchaseorder(request: QueryGetSentPurchaseorderRequest): Promise<QueryGetSentPurchaseorderResponse>;
    /** Queries a list of sentPurchaseorder items. */
    SentPurchaseorderAll(request: QueryAllSentPurchaseorderRequest): Promise<QueryAllSentPurchaseorderResponse>;
    /** Queries a timedoutPurchaseorder by id. */
    TimedoutPurchaseorder(request: QueryGetTimedoutPurchaseorderRequest): Promise<QueryGetTimedoutPurchaseorderResponse>;
    /** Queries a list of timedoutPurchaseorder items. */
    TimedoutPurchaseorderAll(request: QueryAllTimedoutPurchaseorderRequest): Promise<QueryAllTimedoutPurchaseorderResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Purchaseorder(request: QueryGetPurchaseorderRequest): Promise<QueryGetPurchaseorderResponse>;
    PurchaseorderAll(request: QueryAllPurchaseorderRequest): Promise<QueryAllPurchaseorderResponse>;
    SentPurchaseorder(request: QueryGetSentPurchaseorderRequest): Promise<QueryGetSentPurchaseorderResponse>;
    SentPurchaseorderAll(request: QueryAllSentPurchaseorderRequest): Promise<QueryAllSentPurchaseorderResponse>;
    TimedoutPurchaseorder(request: QueryGetTimedoutPurchaseorderRequest): Promise<QueryGetTimedoutPurchaseorderResponse>;
    TimedoutPurchaseorderAll(request: QueryAllTimedoutPurchaseorderRequest): Promise<QueryAllTimedoutPurchaseorderResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
