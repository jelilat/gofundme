import { Writer, Reader } from "protobufjs/minimal";
import { Params } from "../gofundme/params";
import { Gofundme } from "../gofundme/gofundme";
export declare const protobufPackage = "cosmonaut.gofundme.gofundme";
/** GenesisState defines the gofundme module's genesis state. */
export interface GenesisState {
    params: Params | undefined;
    gofundmeList: Gofundme[];
    /** this line is used by starport scaffolding # genesis/proto/state */
    gofundmeCount: number;
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
