import { Reader, Writer } from "protobufjs/minimal";
import { Invoice } from "../invoice/invoice";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
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
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a invoice by id. */
    Invoice(request: QueryGetInvoiceRequest): Promise<QueryGetInvoiceResponse>;
    /** Queries a list of invoice items. */
    InvoiceAll(request: QueryAllInvoiceRequest): Promise<QueryAllInvoiceResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Invoice(request: QueryGetInvoiceRequest): Promise<QueryGetInvoiceResponse>;
    InvoiceAll(request: QueryAllInvoiceRequest): Promise<QueryAllInvoiceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
