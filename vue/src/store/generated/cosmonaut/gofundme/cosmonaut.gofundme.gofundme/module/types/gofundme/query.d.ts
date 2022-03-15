import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../gofundme/params";
import { Gofundme } from "../gofundme/gofundme";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "cosmonaut.gofundme.gofundme";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetGofundmeRequest {
    id: number;
}
export interface QueryGetGofundmeResponse {
    Gofundme: Gofundme | undefined;
}
export interface QueryAllGofundmeRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllGofundmeResponse {
    Gofundme: Gofundme[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetGofundmeRequest: {
    encode(message: QueryGetGofundmeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGofundmeRequest;
    fromJSON(object: any): QueryGetGofundmeRequest;
    toJSON(message: QueryGetGofundmeRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetGofundmeRequest>): QueryGetGofundmeRequest;
};
export declare const QueryGetGofundmeResponse: {
    encode(message: QueryGetGofundmeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetGofundmeResponse;
    fromJSON(object: any): QueryGetGofundmeResponse;
    toJSON(message: QueryGetGofundmeResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetGofundmeResponse>): QueryGetGofundmeResponse;
};
export declare const QueryAllGofundmeRequest: {
    encode(message: QueryAllGofundmeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGofundmeRequest;
    fromJSON(object: any): QueryAllGofundmeRequest;
    toJSON(message: QueryAllGofundmeRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllGofundmeRequest>): QueryAllGofundmeRequest;
};
export declare const QueryAllGofundmeResponse: {
    encode(message: QueryAllGofundmeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllGofundmeResponse;
    fromJSON(object: any): QueryAllGofundmeResponse;
    toJSON(message: QueryAllGofundmeResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllGofundmeResponse>): QueryAllGofundmeResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a Gofundme by id. */
    Gofundme(request: QueryGetGofundmeRequest): Promise<QueryGetGofundmeResponse>;
    /** Queries a list of Gofundme items. */
    GofundmeAll(request: QueryAllGofundmeRequest): Promise<QueryAllGofundmeResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Gofundme(request: QueryGetGofundmeRequest): Promise<QueryGetGofundmeResponse>;
    GofundmeAll(request: QueryAllGofundmeRequest): Promise<QueryAllGofundmeResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
