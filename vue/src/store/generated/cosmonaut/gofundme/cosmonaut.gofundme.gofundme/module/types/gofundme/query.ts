/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../gofundme/params";
import { Gofundme } from "../gofundme/gofundme";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "cosmonaut.gofundme.gofundme";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetGofundmeRequest: object = { id: 0 };

export const QueryGetGofundmeRequest = {
  encode(
    message: QueryGetGofundmeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetGofundmeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetGofundmeRequest,
    } as QueryGetGofundmeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGofundmeRequest {
    const message = {
      ...baseQueryGetGofundmeRequest,
    } as QueryGetGofundmeRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetGofundmeRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGofundmeRequest>
  ): QueryGetGofundmeRequest {
    const message = {
      ...baseQueryGetGofundmeRequest,
    } as QueryGetGofundmeRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetGofundmeResponse: object = {};

export const QueryGetGofundmeResponse = {
  encode(
    message: QueryGetGofundmeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Gofundme !== undefined) {
      Gofundme.encode(message.Gofundme, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetGofundmeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetGofundmeResponse,
    } as QueryGetGofundmeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Gofundme = Gofundme.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetGofundmeResponse {
    const message = {
      ...baseQueryGetGofundmeResponse,
    } as QueryGetGofundmeResponse;
    if (object.Gofundme !== undefined && object.Gofundme !== null) {
      message.Gofundme = Gofundme.fromJSON(object.Gofundme);
    } else {
      message.Gofundme = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetGofundmeResponse): unknown {
    const obj: any = {};
    message.Gofundme !== undefined &&
      (obj.Gofundme = message.Gofundme
        ? Gofundme.toJSON(message.Gofundme)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetGofundmeResponse>
  ): QueryGetGofundmeResponse {
    const message = {
      ...baseQueryGetGofundmeResponse,
    } as QueryGetGofundmeResponse;
    if (object.Gofundme !== undefined && object.Gofundme !== null) {
      message.Gofundme = Gofundme.fromPartial(object.Gofundme);
    } else {
      message.Gofundme = undefined;
    }
    return message;
  },
};

const baseQueryAllGofundmeRequest: object = {};

export const QueryAllGofundmeRequest = {
  encode(
    message: QueryAllGofundmeRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllGofundmeRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllGofundmeRequest,
    } as QueryAllGofundmeRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGofundmeRequest {
    const message = {
      ...baseQueryAllGofundmeRequest,
    } as QueryAllGofundmeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGofundmeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllGofundmeRequest>
  ): QueryAllGofundmeRequest {
    const message = {
      ...baseQueryAllGofundmeRequest,
    } as QueryAllGofundmeRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllGofundmeResponse: object = {};

export const QueryAllGofundmeResponse = {
  encode(
    message: QueryAllGofundmeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Gofundme) {
      Gofundme.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllGofundmeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllGofundmeResponse,
    } as QueryAllGofundmeResponse;
    message.Gofundme = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Gofundme.push(Gofundme.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllGofundmeResponse {
    const message = {
      ...baseQueryAllGofundmeResponse,
    } as QueryAllGofundmeResponse;
    message.Gofundme = [];
    if (object.Gofundme !== undefined && object.Gofundme !== null) {
      for (const e of object.Gofundme) {
        message.Gofundme.push(Gofundme.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllGofundmeResponse): unknown {
    const obj: any = {};
    if (message.Gofundme) {
      obj.Gofundme = message.Gofundme.map((e) =>
        e ? Gofundme.toJSON(e) : undefined
      );
    } else {
      obj.Gofundme = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllGofundmeResponse>
  ): QueryAllGofundmeResponse {
    const message = {
      ...baseQueryAllGofundmeResponse,
    } as QueryAllGofundmeResponse;
    message.Gofundme = [];
    if (object.Gofundme !== undefined && object.Gofundme !== null) {
      for (const e of object.Gofundme) {
        message.Gofundme.push(Gofundme.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Gofundme by id. */
  Gofundme(request: QueryGetGofundmeRequest): Promise<QueryGetGofundmeResponse>;
  /** Queries a list of Gofundme items. */
  GofundmeAll(
    request: QueryAllGofundmeRequest
  ): Promise<QueryAllGofundmeResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.gofundme.gofundme.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Gofundme(
    request: QueryGetGofundmeRequest
  ): Promise<QueryGetGofundmeResponse> {
    const data = QueryGetGofundmeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.gofundme.gofundme.Query",
      "Gofundme",
      data
    );
    return promise.then((data) =>
      QueryGetGofundmeResponse.decode(new Reader(data))
    );
  }

  GofundmeAll(
    request: QueryAllGofundmeRequest
  ): Promise<QueryAllGofundmeResponse> {
    const data = QueryAllGofundmeRequest.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.gofundme.gofundme.Query",
      "GofundmeAll",
      data
    );
    return promise.then((data) =>
      QueryAllGofundmeResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
