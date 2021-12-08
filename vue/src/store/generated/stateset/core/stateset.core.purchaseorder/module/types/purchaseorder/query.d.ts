import { Reader, Writer } from "protobufjs/minimal";
import { Purchaseorder } from "../purchaseorder/purchaseorder";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
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
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a purchaseorder by id. */
    Purchaseorder(request: QueryGetPurchaseorderRequest): Promise<QueryGetPurchaseorderResponse>;
    /** Queries a list of purchaseorder items. */
    PurchaseorderAll(request: QueryAllPurchaseorderRequest): Promise<QueryAllPurchaseorderResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Purchaseorder(request: QueryGetPurchaseorderRequest): Promise<QueryGetPurchaseorderResponse>;
    PurchaseorderAll(request: QueryAllPurchaseorderRequest): Promise<QueryAllPurchaseorderResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
