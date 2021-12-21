import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgActivateAgreement } from "./types/agreement/tx";
import { MsgCreateTimedoutAgreement } from "./types/agreement/tx";
import { MsgExpireAgreement } from "./types/agreement/tx";
import { MsgUpdateSentAgreement } from "./types/agreement/tx";
import { MsgTerminateAgreement } from "./types/agreement/tx";
import { MsgUpdateTimedoutAgreement } from "./types/agreement/tx";
import { MsgDeleteTimedoutAgreement } from "./types/agreement/tx";
import { MsgCreateSentAgreement } from "./types/agreement/tx";
import { MsgRenewAgreement } from "./types/agreement/tx";
import { MsgDeleteSentAgreement } from "./types/agreement/tx";
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
    msgActivateAgreement: (data: MsgActivateAgreement) => EncodeObject;
    msgCreateTimedoutAgreement: (data: MsgCreateTimedoutAgreement) => EncodeObject;
    msgExpireAgreement: (data: MsgExpireAgreement) => EncodeObject;
    msgUpdateSentAgreement: (data: MsgUpdateSentAgreement) => EncodeObject;
    msgTerminateAgreement: (data: MsgTerminateAgreement) => EncodeObject;
    msgUpdateTimedoutAgreement: (data: MsgUpdateTimedoutAgreement) => EncodeObject;
    msgDeleteTimedoutAgreement: (data: MsgDeleteTimedoutAgreement) => EncodeObject;
    msgCreateSentAgreement: (data: MsgCreateSentAgreement) => EncodeObject;
    msgRenewAgreement: (data: MsgRenewAgreement) => EncodeObject;
    msgDeleteSentAgreement: (data: MsgDeleteSentAgreement) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
