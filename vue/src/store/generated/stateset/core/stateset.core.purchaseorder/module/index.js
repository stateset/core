// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgFinancePurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCancelPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCompletePurchaseorder } from "./types/purchaseorder/tx";
import { MsgRequestPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCreateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCreateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteTimedoutPurchaseorder } from "./types/purchaseorder/tx";
const types = [
    ["/stateset.core.purchaseorder.MsgFinancePurchaseorder", MsgFinancePurchaseorder],
    ["/stateset.core.purchaseorder.MsgUpdateSentPurchaseorder", MsgUpdateSentPurchaseorder],
    ["/stateset.core.purchaseorder.MsgUpdateTimedoutPurchaseorder", MsgUpdateTimedoutPurchaseorder],
    ["/stateset.core.purchaseorder.MsgCancelPurchaseorder", MsgCancelPurchaseorder],
    ["/stateset.core.purchaseorder.MsgCompletePurchaseorder", MsgCompletePurchaseorder],
    ["/stateset.core.purchaseorder.MsgRequestPurchaseorder", MsgRequestPurchaseorder],
    ["/stateset.core.purchaseorder.MsgCreateSentPurchaseorder", MsgCreateSentPurchaseorder],
    ["/stateset.core.purchaseorder.MsgDeleteSentPurchaseorder", MsgDeleteSentPurchaseorder],
    ["/stateset.core.purchaseorder.MsgCreateTimedoutPurchaseorder", MsgCreateTimedoutPurchaseorder],
    ["/stateset.core.purchaseorder.MsgDeleteTimedoutPurchaseorder", MsgDeleteTimedoutPurchaseorder],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgFinancePurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgFinancePurchaseorder", value: data }),
        msgUpdateSentPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgUpdateSentPurchaseorder", value: data }),
        msgUpdateTimedoutPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgUpdateTimedoutPurchaseorder", value: data }),
        msgCancelPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgCancelPurchaseorder", value: data }),
        msgCompletePurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgCompletePurchaseorder", value: data }),
        msgRequestPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgRequestPurchaseorder", value: data }),
        msgCreateSentPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgCreateSentPurchaseorder", value: data }),
        msgDeleteSentPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgDeleteSentPurchaseorder", value: data }),
        msgCreateTimedoutPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgCreateTimedoutPurchaseorder", value: data }),
        msgDeleteTimedoutPurchaseorder: (data) => ({ typeUrl: "/stateset.core.purchaseorder.MsgDeleteTimedoutPurchaseorder", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
