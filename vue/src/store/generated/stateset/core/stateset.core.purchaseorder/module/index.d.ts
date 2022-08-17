import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgRequestPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCompletePurchaseorder } from "./types/purchaseorder/tx";
import { MsgCreateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgFinancePurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCancelPurchaseorder } from "./types/purchaseorder/tx";
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
    msgCreateSentPurchaseorder: (data: MsgCreateSentPurchaseorder) => EncodeObject;
    msgRequestPurchaseorder: (data: MsgRequestPurchaseorder) => EncodeObject;
    msgCompletePurchaseorder: (data: MsgCompletePurchaseorder) => EncodeObject;
    msgCreateTimedoutPurchaseorder: (data: MsgCreateTimedoutPurchaseorder) => EncodeObject;
    msgDeleteTimedoutPurchaseorder: (data: MsgDeleteTimedoutPurchaseorder) => EncodeObject;
    msgUpdateSentPurchaseorder: (data: MsgUpdateSentPurchaseorder) => EncodeObject;
    msgDeleteSentPurchaseorder: (data: MsgDeleteSentPurchaseorder) => EncodeObject;
    msgFinancePurchaseorder: (data: MsgFinancePurchaseorder) => EncodeObject;
    msgUpdateTimedoutPurchaseorder: (data: MsgUpdateTimedoutPurchaseorder) => EncodeObject;
    msgCancelPurchaseorder: (data: MsgCancelPurchaseorder) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
