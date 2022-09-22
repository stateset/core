import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "stateset.core.loan";
/** EventCreateLoan is an event emitted when an loan is created. */
export interface EventCreateLoan {
    /** loan_id is the unique ID of loan */
    loanId: string;
    /** creator is the creator of the loan */
    creator: string;
}
/** EventLoanRequested is an event emitted when an loan is requested. */
export interface EventLoanRequested {
    /** loan_id is the unique ID of loan */
    loanId: string;
    /** creator is the creator of the loan */
    creator: string;
}
/** EventApproved is an event emitted when an loan is approved. */
export interface EventApproved {
    /** loan_id is the unique ID of loan */
    loanId: string;
    /** creator is the creator of the loan */
    creator: string;
}
/** EventRepaid is an event emitted when an loan is repaid. */
export interface EventRepaid {
    /** loan_id is the unique ID of loan */
    loanId: string;
    /** creator is the creator of the loan */
    creator: string;
}
/** EventLiquidated is an event emitted when an loan is liquidated. */
export interface EventLiquidated {
    /** loan_id is the unique ID of loan */
    loanId: string;
    /** creator is the creator of the loan */
    creator: string;
}
/** EventCancelled is an event emitted when an loan is cancelled. */
export interface EventCancelled {
    /** loan_id is the unique ID of loan */
    loanId: string;
    /** creator is the creator of the loan */
    creator: string;
}
export declare const EventCreateLoan: {
    encode(message: EventCreateLoan, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventCreateLoan;
    fromJSON(object: any): EventCreateLoan;
    toJSON(message: EventCreateLoan): unknown;
    fromPartial(object: DeepPartial<EventCreateLoan>): EventCreateLoan;
};
export declare const EventLoanRequested: {
    encode(message: EventLoanRequested, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventLoanRequested;
    fromJSON(object: any): EventLoanRequested;
    toJSON(message: EventLoanRequested): unknown;
    fromPartial(object: DeepPartial<EventLoanRequested>): EventLoanRequested;
};
export declare const EventApproved: {
    encode(message: EventApproved, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventApproved;
    fromJSON(object: any): EventApproved;
    toJSON(message: EventApproved): unknown;
    fromPartial(object: DeepPartial<EventApproved>): EventApproved;
};
export declare const EventRepaid: {
    encode(message: EventRepaid, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventRepaid;
    fromJSON(object: any): EventRepaid;
    toJSON(message: EventRepaid): unknown;
    fromPartial(object: DeepPartial<EventRepaid>): EventRepaid;
};
export declare const EventLiquidated: {
    encode(message: EventLiquidated, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventLiquidated;
    fromJSON(object: any): EventLiquidated;
    toJSON(message: EventLiquidated): unknown;
    fromPartial(object: DeepPartial<EventLiquidated>): EventLiquidated;
};
export declare const EventCancelled: {
    encode(message: EventCancelled, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EventCancelled;
    fromJSON(object: any): EventCancelled;
    toJSON(message: EventCancelled): unknown;
    fromPartial(object: DeepPartial<EventCancelled>): EventCancelled;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
