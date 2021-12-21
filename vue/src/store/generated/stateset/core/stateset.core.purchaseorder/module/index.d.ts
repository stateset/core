import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCancelPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCreateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCreateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgRequestPurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCompletePurchaseorder } from "./types/purchaseorder/tx";
import { MsgFinancePurchaseorder } from "./types/purchaseorder/tx";
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
    msgCancelPurchaseorder: (data: MsgCancelPurchaseorder) => EncodeObject;
    msgCreateTimedoutPurchaseorder: (data: MsgCreateTimedoutPurchaseorder) => EncodeObject;
    msgUpdateTimedoutPurchaseorder: (data: MsgUpdateTimedoutPurchaseorder) => EncodeObject;
    msgDeleteTimedoutPurchaseorder: (data: MsgDeleteTimedoutPurchaseorder) => EncodeObject;
    msgCreateSentPurchaseorder: (data: MsgCreateSentPurchaseorder) => EncodeObject;
    msgRequestPurchaseorder: (data: MsgRequestPurchaseorder) => EncodeObject;
    msgUpdateSentPurchaseorder: (data: MsgUpdateSentPurchaseorder) => EncodeObject;
    msgDeleteSentPurchaseorder: (data: MsgDeleteSentPurchaseorder) => EncodeObject;
    msgCompletePurchaseorder: (data: MsgCompletePurchaseorder) => EncodeObject;
    msgFinancePurchaseorder: (data: MsgFinancePurchaseorder) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
