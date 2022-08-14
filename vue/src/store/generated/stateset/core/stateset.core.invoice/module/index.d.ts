import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgFactorInvoice } from "./types/invoice/tx";
import { MsgUpdateTimedoutInvoice } from "./types/invoice/tx";
import { MsgCreateTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteTimedoutInvoice } from "./types/invoice/tx";
import { MsgUpdateSentInvoice } from "./types/invoice/tx";
import { MsgDeleteSentInvoice } from "./types/invoice/tx";
import { MsgCreateSentInvoice } from "./types/invoice/tx";
import { MsgCreateInvoice } from "./types/invoice/tx";
export declare const MissingWalletError: Error;
export declare const registry: Registry;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => any;
    msgFactorInvoice: (data: MsgFactorInvoice) => EncodeObject;
    msgUpdateTimedoutInvoice: (data: MsgUpdateTimedoutInvoice) => EncodeObject;
    msgCreateTimedoutInvoice: (data: MsgCreateTimedoutInvoice) => EncodeObject;
    msgDeleteTimedoutInvoice: (data: MsgDeleteTimedoutInvoice) => EncodeObject;
    msgUpdateSentInvoice: (data: MsgUpdateSentInvoice) => EncodeObject;
    msgDeleteSentInvoice: (data: MsgDeleteSentInvoice) => EncodeObject;
    msgCreateSentInvoice: (data: MsgCreateSentInvoice) => EncodeObject;
    msgCreateInvoice: (data: MsgCreateInvoice) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
