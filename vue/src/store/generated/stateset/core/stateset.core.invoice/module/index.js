// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSentInvoice } from "./types/invoice/tx";
import { MsgUpdateTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteSentInvoice } from "./types/invoice/tx";
import { MsgCreateTimedoutInvoice } from "./types/invoice/tx";
import { MsgPayInvoice } from "./types/invoice/tx";
import { MsgUpdateSentInvoice } from "./types/invoice/tx";
import { MsgCreateInvoice } from "./types/invoice/tx";
import { MsgVoidInvoice } from "./types/invoice/tx";
import { MsgFactorInvoice } from "./types/invoice/tx";
const types = [
    ["/stateset.core.invoice.MsgCreateSentInvoice", MsgCreateSentInvoice],
    ["/stateset.core.invoice.MsgUpdateTimedoutInvoice", MsgUpdateTimedoutInvoice],
    ["/stateset.core.invoice.MsgDeleteTimedoutInvoice", MsgDeleteTimedoutInvoice],
    ["/stateset.core.invoice.MsgDeleteSentInvoice", MsgDeleteSentInvoice],
    ["/stateset.core.invoice.MsgCreateTimedoutInvoice", MsgCreateTimedoutInvoice],
    ["/stateset.core.invoice.MsgPayInvoice", MsgPayInvoice],
    ["/stateset.core.invoice.MsgUpdateSentInvoice", MsgUpdateSentInvoice],
    ["/stateset.core.invoice.MsgCreateInvoice", MsgCreateInvoice],
    ["/stateset.core.invoice.MsgVoidInvoice", MsgVoidInvoice],
    ["/stateset.core.invoice.MsgFactorInvoice", MsgFactorInvoice],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgCreateSentInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgCreateSentInvoice", value: MsgCreateSentInvoice.fromPartial(data) }),
        msgUpdateTimedoutInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgUpdateTimedoutInvoice", value: MsgUpdateTimedoutInvoice.fromPartial(data) }),
        msgDeleteTimedoutInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgDeleteTimedoutInvoice", value: MsgDeleteTimedoutInvoice.fromPartial(data) }),
        msgDeleteSentInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgDeleteSentInvoice", value: MsgDeleteSentInvoice.fromPartial(data) }),
        msgCreateTimedoutInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgCreateTimedoutInvoice", value: MsgCreateTimedoutInvoice.fromPartial(data) }),
        msgPayInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgPayInvoice", value: MsgPayInvoice.fromPartial(data) }),
        msgUpdateSentInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgUpdateSentInvoice", value: MsgUpdateSentInvoice.fromPartial(data) }),
        msgCreateInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgCreateInvoice", value: MsgCreateInvoice.fromPartial(data) }),
        msgVoidInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgVoidInvoice", value: MsgVoidInvoice.fromPartial(data) }),
        msgFactorInvoice: (data) => ({ typeUrl: "/stateset.core.invoice.MsgFactorInvoice", value: MsgFactorInvoice.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
