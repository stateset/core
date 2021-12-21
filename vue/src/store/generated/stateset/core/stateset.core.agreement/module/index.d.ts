import { StdFee } from "@cosmjs/launchpad";
import { OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteTimedoutAgreement } from "./types/agreement/tx";
import { MsgCreateSentAgreement } from "./types/agreement/tx";
import { MsgDeleteSentAgreement } from "./types/agreement/tx";
import { MsgExpireAgreement } from "./types/agreement/tx";
import { MsgRenewAgreement } from "./types/agreement/tx";
import { MsgUpdateSentAgreement } from "./types/agreement/tx";
import { MsgCreateTimedoutAgreement } from "./types/agreement/tx";
import { MsgActivateAgreement } from "./types/agreement/tx";
import { MsgTerminateAgreement } from "./types/agreement/tx";
import { MsgUpdateTimedoutAgreement } from "./types/agreement/tx";
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
    msgDeleteTimedoutAgreement: (data: MsgDeleteTimedoutAgreement) => EncodeObject;
    msgCreateSentAgreement: (data: MsgCreateSentAgreement) => EncodeObject;
    msgDeleteSentAgreement: (data: MsgDeleteSentAgreement) => EncodeObject;
    msgExpireAgreement: (data: MsgExpireAgreement) => EncodeObject;
    msgRenewAgreement: (data: MsgRenewAgreement) => EncodeObject;
    msgUpdateSentAgreement: (data: MsgUpdateSentAgreement) => EncodeObject;
    msgCreateTimedoutAgreement: (data: MsgCreateTimedoutAgreement) => EncodeObject;
    msgActivateAgreement: (data: MsgActivateAgreement) => EncodeObject;
    msgTerminateAgreement: (data: MsgTerminateAgreement) => EncodeObject;
    msgUpdateTimedoutAgreement: (data: MsgUpdateTimedoutAgreement) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
