import { Reader, Writer } from "protobufjs/minimal";
import { Invoice } from "../invoice/invoice";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
import { SentInvoice } from "../invoice/sent_invoice";
import { TimedoutInvoice } from "../invoice/timedout_invoice";
export declare const protobufPackage = "stateset.core.invoice";
export interface QueryGetInvoiceRequest {
    id: number;
}
export interface QueryGetInvoiceResponse {
    Invoice: Invoice | undefined;
}
export interface QueryAllInvoiceRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllInvoiceResponse {
    Invoice: Invoice[];
    pagination: PageResponse | undefined;
}
export interface QueryGetSentInvoiceRequest {
    id: number;
}
export interface QueryGetSentInvoiceResponse {
    SentInvoice: SentInvoice | undefined;
}
export interface QueryAllSentInvoiceRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllSentInvoiceResponse {
    SentInvoice: SentInvoice[];
    pagination: PageResponse | undefined;
}
export interface QueryGetTimedoutInvoiceRequest {
    id: number;
}
export interface QueryGetTimedoutInvoiceResponse {
    TimedoutInvoice: TimedoutInvoice | undefined;
}
export interface QueryAllTimedoutInvoiceRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllTimedoutInvoiceResponse {
    TimedoutInvoice: TimedoutInvoice[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetInvoiceRequest: {
    encode(message: QueryGetInvoiceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetInvoiceRequest;
    fromJSON(object: any): QueryGetInvoiceRequest;
    toJSON(message: QueryGetInvoiceRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetInvoiceRequest>): QueryGetInvoiceRequest;
};
export declare const QueryGetInvoiceResponse: {
    encode(message: QueryGetInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetInvoiceResponse;
    fromJSON(object: any): QueryGetInvoiceResponse;
    toJSON(message: QueryGetInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetInvoiceResponse>): QueryGetInvoiceResponse;
};
export declare const QueryAllInvoiceRequest: {
    encode(message: QueryAllInvoiceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllInvoiceRequest;
    fromJSON(object: any): QueryAllInvoiceRequest;
    toJSON(message: QueryAllInvoiceRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllInvoiceRequest>): QueryAllInvoiceRequest;
};
export declare const QueryAllInvoiceResponse: {
    encode(message: QueryAllInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllInvoiceResponse;
    fromJSON(object: any): QueryAllInvoiceResponse;
    toJSON(message: QueryAllInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllInvoiceResponse>): QueryAllInvoiceResponse;
};
export declare const QueryGetSentInvoiceRequest: {
    encode(message: QueryGetSentInvoiceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSentInvoiceRequest;
    fromJSON(object: any): QueryGetSentInvoiceRequest;
    toJSON(message: QueryGetSentInvoiceRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetSentInvoiceRequest>): QueryGetSentInvoiceRequest;
};
export declare const QueryGetSentInvoiceResponse: {
    encode(message: QueryGetSentInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetSentInvoiceResponse;
    fromJSON(object: any): QueryGetSentInvoiceResponse;
    toJSON(message: QueryGetSentInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetSentInvoiceResponse>): QueryGetSentInvoiceResponse;
};
export declare const QueryAllSentInvoiceRequest: {
    encode(message: QueryAllSentInvoiceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSentInvoiceRequest;
    fromJSON(object: any): QueryAllSentInvoiceRequest;
    toJSON(message: QueryAllSentInvoiceRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllSentInvoiceRequest>): QueryAllSentInvoiceRequest;
};
export declare const QueryAllSentInvoiceResponse: {
    encode(message: QueryAllSentInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllSentInvoiceResponse;
    fromJSON(object: any): QueryAllSentInvoiceResponse;
    toJSON(message: QueryAllSentInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllSentInvoiceResponse>): QueryAllSentInvoiceResponse;
};
export declare const QueryGetTimedoutInvoiceRequest: {
    encode(message: QueryGetTimedoutInvoiceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutInvoiceRequest;
    fromJSON(object: any): QueryGetTimedoutInvoiceRequest;
    toJSON(message: QueryGetTimedoutInvoiceRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetTimedoutInvoiceRequest>): QueryGetTimedoutInvoiceRequest;
};
export declare const QueryGetTimedoutInvoiceResponse: {
    encode(message: QueryGetTimedoutInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetTimedoutInvoiceResponse;
    fromJSON(object: any): QueryGetTimedoutInvoiceResponse;
    toJSON(message: QueryGetTimedoutInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetTimedoutInvoiceResponse>): QueryGetTimedoutInvoiceResponse;
};
export declare const QueryAllTimedoutInvoiceRequest: {
    encode(message: QueryAllTimedoutInvoiceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTimedoutInvoiceRequest;
    fromJSON(object: any): QueryAllTimedoutInvoiceRequest;
    toJSON(message: QueryAllTimedoutInvoiceRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllTimedoutInvoiceRequest>): QueryAllTimedoutInvoiceRequest;
};
export declare const QueryAllTimedoutInvoiceResponse: {
    encode(message: QueryAllTimedoutInvoiceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllTimedoutInvoiceResponse;
    fromJSON(object: any): QueryAllTimedoutInvoiceResponse;
    toJSON(message: QueryAllTimedoutInvoiceResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllTimedoutInvoiceResponse>): QueryAllTimedoutInvoiceResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a invoice by id. */
    Invoice(request: QueryGetInvoiceRequest): Promise<QueryGetInvoiceResponse>;
    /** Queries a list of invoice items. */
    InvoiceAll(request: QueryAllInvoiceRequest): Promise<QueryAllInvoiceResponse>;
    /** Queries a sentInvoice by id. */
    SentInvoice(request: QueryGetSentInvoiceRequest): Promise<QueryGetSentInvoiceResponse>;
    /** Queries a list of sentInvoice items. */
    SentInvoiceAll(request: QueryAllSentInvoiceRequest): Promise<QueryAllSentInvoiceResponse>;
    /** Queries a timedoutInvoice by id. */
    TimedoutInvoice(request: QueryGetTimedoutInvoiceRequest): Promise<QueryGetTimedoutInvoiceResponse>;
    /** Queries a list of timedoutInvoice items. */
    TimedoutInvoiceAll(request: QueryAllTimedoutInvoiceRequest): Promise<QueryAllTimedoutInvoiceResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Invoice(request: QueryGetInvoiceRequest): Promise<QueryGetInvoiceResponse>;
    InvoiceAll(request: QueryAllInvoiceRequest): Promise<QueryAllInvoiceResponse>;
    SentInvoice(request: QueryGetSentInvoiceRequest): Promise<QueryGetSentInvoiceResponse>;
    SentInvoiceAll(request: QueryAllSentInvoiceRequest): Promise<QueryAllSentInvoiceResponse>;
    TimedoutInvoice(request: QueryGetTimedoutInvoiceRequest): Promise<QueryGetTimedoutInvoiceResponse>;
    TimedoutInvoiceAll(request: QueryAllTimedoutInvoiceRequest): Promise<QueryAllTimedoutInvoiceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
