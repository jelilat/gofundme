import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "cosmonaut.gofundme.gofundme";
export interface MsgCreateGofundme {
    creator: string;
    goal: number;
    start: string;
    end: string;
}
export interface MsgCreateGofundmeResponse {
}
export interface MsgDonateFund {
    creator: string;
    id: number;
    donation: string;
}
export interface MsgDonateFundResponse {
}
export declare const MsgCreateGofundme: {
    encode(message: MsgCreateGofundme, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGofundme;
    fromJSON(object: any): MsgCreateGofundme;
    toJSON(message: MsgCreateGofundme): unknown;
    fromPartial(object: DeepPartial<MsgCreateGofundme>): MsgCreateGofundme;
};
export declare const MsgCreateGofundmeResponse: {
    encode(_: MsgCreateGofundmeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateGofundmeResponse;
    fromJSON(_: any): MsgCreateGofundmeResponse;
    toJSON(_: MsgCreateGofundmeResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateGofundmeResponse>): MsgCreateGofundmeResponse;
};
export declare const MsgDonateFund: {
    encode(message: MsgDonateFund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDonateFund;
    fromJSON(object: any): MsgDonateFund;
    toJSON(message: MsgDonateFund): unknown;
    fromPartial(object: DeepPartial<MsgDonateFund>): MsgDonateFund;
};
export declare const MsgDonateFundResponse: {
    encode(_: MsgDonateFundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDonateFundResponse;
    fromJSON(_: any): MsgDonateFundResponse;
    toJSON(_: MsgDonateFundResponse): unknown;
    fromPartial(_: DeepPartial<MsgDonateFundResponse>): MsgDonateFundResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateGofundme(request: MsgCreateGofundme): Promise<MsgCreateGofundmeResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DonateFund(request: MsgDonateFund): Promise<MsgDonateFundResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateGofundme(request: MsgCreateGofundme): Promise<MsgCreateGofundmeResponse>;
    DonateFund(request: MsgDonateFund): Promise<MsgDonateFundResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
