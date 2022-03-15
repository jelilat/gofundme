/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cosmonaut.gofundme.gofundme";

export interface Gofundme {
  id: number;
  creator: string;
  goal: number;
  start: string;
  end: string;
  totaldonations: number;
  claim: string;
  donation: string[];
  donor: string[];
}

const baseGofundme: object = {
  id: 0,
  creator: "",
  goal: 0,
  start: "",
  end: "",
  totaldonations: 0,
  claim: "",
  donation: "",
  donor: "",
};

export const Gofundme = {
  encode(message: Gofundme, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.goal !== 0) {
      writer.uint32(24).uint64(message.goal);
    }
    if (message.start !== "") {
      writer.uint32(34).string(message.start);
    }
    if (message.end !== "") {
      writer.uint32(42).string(message.end);
    }
    if (message.totaldonations !== 0) {
      writer.uint32(48).uint64(message.totaldonations);
    }
    if (message.claim !== "") {
      writer.uint32(58).string(message.claim);
    }
    for (const v of message.donation) {
      writer.uint32(66).string(v!);
    }
    for (const v of message.donor) {
      writer.uint32(74).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Gofundme {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGofundme } as Gofundme;
    message.donation = [];
    message.donor = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.goal = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.start = reader.string();
          break;
        case 5:
          message.end = reader.string();
          break;
        case 6:
          message.totaldonations = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.claim = reader.string();
          break;
        case 8:
          message.donation.push(reader.string());
          break;
        case 9:
          message.donor.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Gofundme {
    const message = { ...baseGofundme } as Gofundme;
    message.donation = [];
    message.donor = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
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
    if (object.totaldonations !== undefined && object.totaldonations !== null) {
      message.totaldonations = Number(object.totaldonations);
    } else {
      message.totaldonations = 0;
    }
    if (object.claim !== undefined && object.claim !== null) {
      message.claim = String(object.claim);
    } else {
      message.claim = "";
    }
    if (object.donation !== undefined && object.donation !== null) {
      for (const e of object.donation) {
        message.donation.push(String(e));
      }
    }
    if (object.donor !== undefined && object.donor !== null) {
      for (const e of object.donor) {
        message.donor.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Gofundme): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.creator !== undefined && (obj.creator = message.creator);
    message.goal !== undefined && (obj.goal = message.goal);
    message.start !== undefined && (obj.start = message.start);
    message.end !== undefined && (obj.end = message.end);
    message.totaldonations !== undefined &&
      (obj.totaldonations = message.totaldonations);
    message.claim !== undefined && (obj.claim = message.claim);
    if (message.donation) {
      obj.donation = message.donation.map((e) => e);
    } else {
      obj.donation = [];
    }
    if (message.donor) {
      obj.donor = message.donor.map((e) => e);
    } else {
      obj.donor = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Gofundme>): Gofundme {
    const message = { ...baseGofundme } as Gofundme;
    message.donation = [];
    message.donor = [];
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
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
    if (object.totaldonations !== undefined && object.totaldonations !== null) {
      message.totaldonations = object.totaldonations;
    } else {
      message.totaldonations = 0;
    }
    if (object.claim !== undefined && object.claim !== null) {
      message.claim = object.claim;
    } else {
      message.claim = "";
    }
    if (object.donation !== undefined && object.donation !== null) {
      for (const e of object.donation) {
        message.donation.push(e);
      }
    }
    if (object.donor !== undefined && object.donor !== null) {
      for (const e of object.donor) {
        message.donor.push(e);
      }
    }
    return message;
  },
};

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
