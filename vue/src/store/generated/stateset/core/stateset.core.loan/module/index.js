// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgApproveLoan } from "./types/loan/tx";
import { MsgLiquidateLoan } from "./types/loan/tx";
import { MsgCancelLoan } from "./types/loan/tx";
import { MsgRepayLoan } from "./types/loan/tx";
import { MsgRequestLoan } from "./types/loan/tx";
const types = [
    ["/stateset.core.loan.MsgApproveLoan", MsgApproveLoan],
    ["/stateset.core.loan.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/stateset.core.loan.MsgCancelLoan", MsgCancelLoan],
    ["/stateset.core.loan.MsgRepayLoan", MsgRepayLoan],
    ["/stateset.core.loan.MsgRequestLoan", MsgRequestLoan],
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
        msgApproveLoan: (data) => ({ typeUrl: "/stateset.core.loan.MsgApproveLoan", value: MsgApproveLoan.fromPartial(data) }),
        msgLiquidateLoan: (data) => ({ typeUrl: "/stateset.core.loan.MsgLiquidateLoan", value: MsgLiquidateLoan.fromPartial(data) }),
        msgCancelLoan: (data) => ({ typeUrl: "/stateset.core.loan.MsgCancelLoan", value: MsgCancelLoan.fromPartial(data) }),
        msgRepayLoan: (data) => ({ typeUrl: "/stateset.core.loan.MsgRepayLoan", value: MsgRepayLoan.fromPartial(data) }),
        msgRequestLoan: (data) => ({ typeUrl: "/stateset.core.loan.MsgRequestLoan", value: MsgRequestLoan.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
