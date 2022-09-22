// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRequestRefund } from "./types/refund/tx";
import { MsgRejectRefund } from "./types/refund/tx";
import { MsgApproveRefund } from "./types/refund/tx";
const types = [
    ["/stateset.core.refund.MsgRequestRefund", MsgRequestRefund],
    ["/stateset.core.refund.MsgRejectRefund", MsgRejectRefund],
    ["/stateset.core.refund.MsgApproveRefund", MsgApproveRefund],
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
        msgRequestRefund: (data) => ({ typeUrl: "/stateset.core.refund.MsgRequestRefund", value: MsgRequestRefund.fromPartial(data) }),
        msgRejectRefund: (data) => ({ typeUrl: "/stateset.core.refund.MsgRejectRefund", value: MsgRejectRefund.fromPartial(data) }),
        msgApproveRefund: (data) => ({ typeUrl: "/stateset.core.refund.MsgApproveRefund", value: MsgApproveRefund.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
