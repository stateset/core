// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteTimedoutInvoice } from "./types/invoice/tx";
import { MsgCreateSentInvoice } from "./types/invoice/tx";
import { MsgCreateTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteSentInvoice } from "./types/invoice/tx";
import { MsgFactorInvoice } from "./types/invoice/tx";
import { MsgUpdateTimedoutInvoice } from "./types/invoice/tx";
import { MsgUpdateSentInvoice } from "./types/invoice/tx";
const types = [
    ["/stateset.core.invoice.MsgDeleteTimedoutInvoice", MsgDeleteTimedoutInvoice],
    ["/stateset.core.invoice.MsgCreateSentInvoice", MsgCreateSentInvoice],
    ["/stateset.core.invoice.MsgCreateTimedoutInvoice", MsgCreateTimedoutInvoice],
    ["/stateset.core.invoice.MsgDeleteSentInvoice", MsgDeleteSentInvoice],
    ["/stateset.core.invoice.MsgFactorInvoice", MsgFactorInvoice],
    ["/stateset.core.invoice.MsgUpdateTimedoutInvoice", MsgUpdateTimedoutInvoice],
    ["/stateset.core.invoice.MsgUpdateSentInvoice", MsgUpdateSentInvoice],
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
        msgDeleteTimedoutInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgDeleteTimedoutInvoice", value: data }),
        msgCreateSentInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgCreateSentInvoice", value: data }),
        msgCreateTimedoutInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgCreateTimedoutInvoice", value: data }),
        msgDeleteSentInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgDeleteSentInvoice", value: data }),
        msgFactorInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgFactorInvoice", value: data }),
        msgUpdateTimedoutInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgUpdateTimedoutInvoice", value: data }),
        msgUpdateSentInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgUpdateSentInvoice", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
