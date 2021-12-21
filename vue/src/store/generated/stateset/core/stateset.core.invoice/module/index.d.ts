import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSentInvoice } from "./types/invoice/tx";
import { MsgFactorInvoice } from "./types/invoice/tx";
import { MsgCreateTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteTimedoutInvoice } from "./types/invoice/tx";
import { MsgUpdateSentInvoice } from "./types/invoice/tx";
import { MsgUpdateTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteSentInvoice } from "./types/invoice/tx";
export declare const MissingWalletError: Error;
interface TxClientOptions {
    addr: string;
}
interface SignAndBroadcastOptions {
    fee: StdFee;
    memo?: string;
}
declare const txClient: (wallet: OfflineSigner, { addr: addr }?: TxClientOptions) => Promise<{
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }?: SignAndBroadcastOptions) => Promise<import("@cosmjs/stargate").BroadcastTxResponse>;
    msgCreateSentInvoice: (data: MsgCreateSentInvoice) => EncodeObject;
    msgFactorInvoice: (data: MsgFactorInvoice) => EncodeObject;
    msgCreateTimedoutInvoice: (data: MsgCreateTimedoutInvoice) => EncodeObject;
    msgDeleteTimedoutInvoice: (data: MsgDeleteTimedoutInvoice) => EncodeObject;
    msgUpdateSentInvoice: (data: MsgUpdateSentInvoice) => EncodeObject;
    msgUpdateTimedoutInvoice: (data: MsgUpdateTimedoutInvoice) => EncodeObject;
    msgDeleteSentInvoice: (data: MsgDeleteSentInvoice) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
