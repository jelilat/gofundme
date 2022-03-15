import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "cosmonaut.gofundme.gofundme";
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
export declare const Gofundme: {
    encode(message: Gofundme, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Gofundme;
    fromJSON(object: any): Gofundme;
    toJSON(message: Gofundme): unknown;
    fromPartial(object: DeepPartial<Gofundme>): Gofundme;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
