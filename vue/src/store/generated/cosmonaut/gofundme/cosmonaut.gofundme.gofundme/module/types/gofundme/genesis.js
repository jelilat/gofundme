/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../gofundme/params";
import { Gofundme } from "../gofundme/gofundme";
export const protobufPackage = "cosmonaut.gofundme.gofundme";
const baseGenesisState = { gofundmeCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.gofundmeList) {
            Gofundme.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.gofundmeCount !== 0) {
            writer.uint32(24).uint64(message.gofundmeCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.gofundmeList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.gofundmeList.push(Gofundme.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.gofundmeCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.gofundmeList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.gofundmeList !== undefined && object.gofundmeList !== null) {
            for (const e of object.gofundmeList) {
                message.gofundmeList.push(Gofundme.fromJSON(e));
            }
        }
        if (object.gofundmeCount !== undefined && object.gofundmeCount !== null) {
            message.gofundmeCount = Number(object.gofundmeCount);
        }
        else {
            message.gofundmeCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.gofundmeList) {
            obj.gofundmeList = message.gofundmeList.map((e) => e ? Gofundme.toJSON(e) : undefined);
        }
        else {
            obj.gofundmeList = [];
        }
        message.gofundmeCount !== undefined &&
            (obj.gofundmeCount = message.gofundmeCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.gofundmeList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.gofundmeList !== undefined && object.gofundmeList !== null) {
            for (const e of object.gofundmeList) {
                message.gofundmeList.push(Gofundme.fromPartial(e));
            }
        }
        if (object.gofundmeCount !== undefined && object.gofundmeCount !== null) {
            message.gofundmeCount = object.gofundmeCount;
        }
        else {
            message.gofundmeCount = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
