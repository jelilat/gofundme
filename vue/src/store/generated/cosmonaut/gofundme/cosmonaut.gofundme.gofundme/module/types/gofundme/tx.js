/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
export const protobufPackage = "cosmonaut.gofundme.gofundme";
const baseMsgCreateGofundme = {
    creator: "",
    goal: 0,
    start: "",
    end: "",
};
export const MsgCreateGofundme = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgCreateGofundme };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.goal = longToNumber(reader.uint64());
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
    fromJSON(object) {
        const message = { ...baseMsgCreateGofundme };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.goal !== undefined && object.goal !== null) {
            message.goal = Number(object.goal);
        }
        else {
            message.goal = 0;
        }
        if (object.start !== undefined && object.start !== null) {
            message.start = String(object.start);
        }
        else {
            message.start = "";
        }
        if (object.end !== undefined && object.end !== null) {
            message.end = String(object.end);
        }
        else {
            message.end = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.goal !== undefined && (obj.goal = message.goal);
        message.start !== undefined && (obj.start = message.start);
        message.end !== undefined && (obj.end = message.end);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgCreateGofundme };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.goal !== undefined && object.goal !== null) {
            message.goal = object.goal;
        }
        else {
            message.goal = 0;
        }
        if (object.start !== undefined && object.start !== null) {
            message.start = object.start;
        }
        else {
            message.start = "";
        }
        if (object.end !== undefined && object.end !== null) {
            message.end = object.end;
        }
        else {
            message.end = "";
        }
        return message;
    },
};
const baseMsgCreateGofundmeResponse = {};
export const MsgCreateGofundmeResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseMsgCreateGofundmeResponse,
        };
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
    fromJSON(_) {
        const message = {
            ...baseMsgCreateGofundmeResponse,
        };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = {
            ...baseMsgCreateGofundmeResponse,
        };
        return message;
    },
};
const baseMsgDonateFund = { creator: "", id: 0, donation: "" };
export const MsgDonateFund = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDonateFund };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.creator = reader.string();
                    break;
                case 2:
                    message.id = longToNumber(reader.uint64());
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
    fromJSON(object) {
        const message = { ...baseMsgDonateFund };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = String(object.creator);
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = Number(object.id);
        }
        else {
            message.id = 0;
        }
        if (object.donation !== undefined && object.donation !== null) {
            message.donation = String(object.donation);
        }
        else {
            message.donation = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.creator !== undefined && (obj.creator = message.creator);
        message.id !== undefined && (obj.id = message.id);
        message.donation !== undefined && (obj.donation = message.donation);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseMsgDonateFund };
        if (object.creator !== undefined && object.creator !== null) {
            message.creator = object.creator;
        }
        else {
            message.creator = "";
        }
        if (object.id !== undefined && object.id !== null) {
            message.id = object.id;
        }
        else {
            message.id = 0;
        }
        if (object.donation !== undefined && object.donation !== null) {
            message.donation = object.donation;
        }
        else {
            message.donation = "";
        }
        return message;
    },
};
const baseMsgDonateFundResponse = {};
export const MsgDonateFundResponse = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseMsgDonateFundResponse };
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
    fromJSON(_) {
        const message = { ...baseMsgDonateFundResponse };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseMsgDonateFundResponse };
        return message;
    },
};
export class MsgClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    CreateGofundme(request) {
        const data = MsgCreateGofundme.encode(request).finish();
        const promise = this.rpc.request("cosmonaut.gofundme.gofundme.Msg", "CreateGofundme", data);
        return promise.then((data) => MsgCreateGofundmeResponse.decode(new Reader(data)));
    }
    DonateFund(request) {
        const data = MsgDonateFund.encode(request).finish();
        const promise = this.rpc.request("cosmonaut.gofundme.gofundme.Msg", "DonateFund", data);
        return promise.then((data) => MsgDonateFundResponse.decode(new Reader(data)));
    }
}
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
