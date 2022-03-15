/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "cosmonaut.gofundme.gofundme";

export interface MsgCreateGofundme {
  creator: string;
  goal: number;
  start: string;
  end: string;
}

export interface MsgCreateGofundmeResponse {}

export interface MsgDonateFund {
  creator: string;
  id: number;
  donation: string;
}

export interface MsgDonateFundResponse {}

export interface MsgWithdrawDonation {
  creator: string;
  id: number;
}

export interface MsgWithdrawDonationResponse {}

const baseMsgCreateGofundme: object = {
  creator: "",
  goal: 0,
  start: "",
  end: "",
};

export const MsgCreateGofundme = {
  encode(message: MsgCreateGofundme, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.goal !== 0) {
      writer.uint32(16).uint64(message.goal);
    }
    if (message.start !== "") {
      writer.uint32(26).string(message.start);
    }
    if (message.end !== "") {
      writer.uint32(34).string(message.end);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateGofundme {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateGofundme } as MsgCreateGofundme;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.goal = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.start = reader.string();
          break;
        case 4:
          message.end = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateGofundme {
    const message = { ...baseMsgCreateGofundme } as MsgCreateGofundme;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.goal !== undefined && object.goal !== null) {
      message.goal = Number(object.goal);
    } else {
      message.goal = 0;
    }
    if (object.start !== undefined && object.start !== null) {
      message.start = String(object.start);
    } else {
      message.start = "";
    }
    if (object.end !== undefined && object.end !== null) {
      message.end = String(object.end);
    } else {
      message.end = "";
    }
    return message;
  },

  toJSON(message: MsgCreateGofundme): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.goal !== undefined && (obj.goal = message.goal);
    message.start !== undefined && (obj.start = message.start);
    message.end !== undefined && (obj.end = message.end);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateGofundme>): MsgCreateGofundme {
    const message = { ...baseMsgCreateGofundme } as MsgCreateGofundme;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.goal !== undefined && object.goal !== null) {
      message.goal = object.goal;
    } else {
      message.goal = 0;
    }
    if (object.start !== undefined && object.start !== null) {
      message.start = object.start;
    } else {
      message.start = "";
    }
    if (object.end !== undefined && object.end !== null) {
      message.end = object.end;
    } else {
      message.end = "";
    }
    return message;
  },
};

const baseMsgCreateGofundmeResponse: object = {};

export const MsgCreateGofundmeResponse = {
  encode(
    _: MsgCreateGofundmeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateGofundmeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateGofundmeResponse,
    } as MsgCreateGofundmeResponse;
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

  fromJSON(_: any): MsgCreateGofundmeResponse {
    const message = {
      ...baseMsgCreateGofundmeResponse,
    } as MsgCreateGofundmeResponse;
    return message;
  },

  toJSON(_: MsgCreateGofundmeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateGofundmeResponse>
  ): MsgCreateGofundmeResponse {
    const message = {
      ...baseMsgCreateGofundmeResponse,
    } as MsgCreateGofundmeResponse;
    return message;
  },
};

const baseMsgDonateFund: object = { creator: "", id: 0, donation: "" };

export const MsgDonateFund = {
  encode(message: MsgDonateFund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.donation !== "") {
      writer.uint32(26).string(message.donation);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDonateFund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDonateFund } as MsgDonateFund;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.donation = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDonateFund {
    const message = { ...baseMsgDonateFund } as MsgDonateFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.donation !== undefined && object.donation !== null) {
      message.donation = String(object.donation);
    } else {
      message.donation = "";
    }
    return message;
  },

  toJSON(message: MsgDonateFund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.donation !== undefined && (obj.donation = message.donation);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDonateFund>): MsgDonateFund {
    const message = { ...baseMsgDonateFund } as MsgDonateFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.donation !== undefined && object.donation !== null) {
      message.donation = object.donation;
    } else {
      message.donation = "";
    }
    return message;
  },
};

const baseMsgDonateFundResponse: object = {};

export const MsgDonateFundResponse = {
  encode(_: MsgDonateFundResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDonateFundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDonateFundResponse } as MsgDonateFundResponse;
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

  fromJSON(_: any): MsgDonateFundResponse {
    const message = { ...baseMsgDonateFundResponse } as MsgDonateFundResponse;
    return message;
  },

  toJSON(_: MsgDonateFundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDonateFundResponse>): MsgDonateFundResponse {
    const message = { ...baseMsgDonateFundResponse } as MsgDonateFundResponse;
    return message;
  },
};

const baseMsgWithdrawDonation: object = { creator: "", id: 0 };

export const MsgWithdrawDonation = {
  encode(
    message: MsgWithdrawDonation,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawDonation {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawDonation } as MsgWithdrawDonation;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawDonation {
    const message = { ...baseMsgWithdrawDonation } as MsgWithdrawDonation;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgWithdrawDonation): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdrawDonation>): MsgWithdrawDonation {
    const message = { ...baseMsgWithdrawDonation } as MsgWithdrawDonation;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgWithdrawDonationResponse: object = {};

export const MsgWithdrawDonationResponse = {
  encode(
    _: MsgWithdrawDonationResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgWithdrawDonationResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgWithdrawDonationResponse,
    } as MsgWithdrawDonationResponse;
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

  fromJSON(_: any): MsgWithdrawDonationResponse {
    const message = {
      ...baseMsgWithdrawDonationResponse,
    } as MsgWithdrawDonationResponse;
    return message;
  },

  toJSON(_: MsgWithdrawDonationResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgWithdrawDonationResponse>
  ): MsgWithdrawDonationResponse {
    const message = {
      ...baseMsgWithdrawDonationResponse,
    } as MsgWithdrawDonationResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateGofundme(
    request: MsgCreateGofundme
  ): Promise<MsgCreateGofundmeResponse>;
  DonateFund(request: MsgDonateFund): Promise<MsgDonateFundResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  WithdrawDonation(
    request: MsgWithdrawDonation
  ): Promise<MsgWithdrawDonationResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateGofundme(
    request: MsgCreateGofundme
  ): Promise<MsgCreateGofundmeResponse> {
    const data = MsgCreateGofundme.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.gofundme.gofundme.Msg",
      "CreateGofundme",
      data
    );
    return promise.then((data) =>
      MsgCreateGofundmeResponse.decode(new Reader(data))
    );
  }

  DonateFund(request: MsgDonateFund): Promise<MsgDonateFundResponse> {
    const data = MsgDonateFund.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.gofundme.gofundme.Msg",
      "DonateFund",
      data
    );
    return promise.then((data) =>
      MsgDonateFundResponse.decode(new Reader(data))
    );
  }

  WithdrawDonation(
    request: MsgWithdrawDonation
  ): Promise<MsgWithdrawDonationResponse> {
    const data = MsgWithdrawDonation.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.gofundme.gofundme.Msg",
      "WithdrawDonation",
      data
    );
    return promise.then((data) =>
      MsgWithdrawDonationResponse.decode(new Reader(data))
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
