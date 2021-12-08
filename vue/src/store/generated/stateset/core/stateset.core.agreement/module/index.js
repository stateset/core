// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgExpireAgreement } from "./types/agreement/tx";
import { MsgRenewAgreement } from "./types/agreement/tx";
import { MsgCreateTimedoutAgreement } from "./types/agreement/tx";
import { MsgUpdateTimedoutAgreement } from "./types/agreement/tx";
import { MsgTerminateAgreement } from "./types/agreement/tx";
import { MsgCreateSentAgreement } from "./types/agreement/tx";
import { MsgDeleteTimedoutAgreement } from "./types/agreement/tx";
import { MsgDeleteSentAgreement } from "./types/agreement/tx";
import { MsgUpdateSentAgreement } from "./types/agreement/tx";
import { MsgActivateAgreement } from "./types/agreement/tx";
const types = [
    ["/stateset.core.agreement.MsgExpireAgreement", MsgExpireAgreement],
    ["/stateset.core.agreement.MsgRenewAgreement", MsgRenewAgreement],
    ["/stateset.core.agreement.MsgCreateTimedoutAgreement", MsgCreateTimedoutAgreement],
    ["/stateset.core.agreement.MsgUpdateTimedoutAgreement", MsgUpdateTimedoutAgreement],
    ["/stateset.core.agreement.MsgTerminateAgreement", MsgTerminateAgreement],
    ["/stateset.core.agreement.MsgCreateSentAgreement", MsgCreateSentAgreement],
    ["/stateset.core.agreement.MsgDeleteTimedoutAgreement", MsgDeleteTimedoutAgreement],
    ["/stateset.core.agreement.MsgDeleteSentAgreement", MsgDeleteSentAgreement],
    ["/stateset.core.agreement.MsgUpdateSentAgreement", MsgUpdateSentAgreement],
    ["/stateset.core.agreement.MsgActivateAgreement", MsgActivateAgreement],
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
        msgExpireAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgExpireAgreement", value: data }),
        msgRenewAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgRenewAgreement", value: data }),
        msgCreateTimedoutAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgCreateTimedoutAgreement", value: data }),
        msgUpdateTimedoutAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgUpdateTimedoutAgreement", value: data }),
        msgTerminateAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgTerminateAgreement", value: data }),
        msgCreateSentAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgCreateSentAgreement", value: data }),
        msgDeleteTimedoutAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgDeleteTimedoutAgreement", value: data }),
        msgDeleteSentAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgDeleteSentAgreement", value: data }),
        msgUpdateSentAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgUpdateSentAgreement", value: data }),
        msgActivateAgreement: (data) => ({ typeUrl: "/stateset.core.agreement.MsgActivateAgreement", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
