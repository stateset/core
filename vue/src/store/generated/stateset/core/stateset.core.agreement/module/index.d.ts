import { StdFee } from "@cosmjs/launchpad";
import { Registry, OfflineSigner, EncodeObject } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRenewAgreement } from "./types/agreement/tx";
import { MsgActivateAgreement } from "./types/agreement/tx";
import { MsgCreateSentAgreement } from "./types/agreement/tx";
import { MsgDeleteSentAgreement } from "./types/agreement/tx";
import { MsgDeleteTimedoutAgreement } from "./types/agreement/tx";
import { MsgExpireAgreement } from "./types/agreement/tx";
import { MsgUpdateTimedoutAgreement } from "./types/agreement/tx";
import { MsgTerminateAgreement } from "./types/agreement/tx";
import { MsgUpdateSentAgreement } from "./types/agreement/tx";
import { MsgCreateTimedoutAgreement } from "./types/agreement/tx";
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
    msgRenewAgreement: (data: MsgRenewAgreement) => EncodeObject;
    msgActivateAgreement: (data: MsgActivateAgreement) => EncodeObject;
    msgCreateSentAgreement: (data: MsgCreateSentAgreement) => EncodeObject;
    msgDeleteSentAgreement: (data: MsgDeleteSentAgreement) => EncodeObject;
    msgDeleteTimedoutAgreement: (data: MsgDeleteTimedoutAgreement) => EncodeObject;
    msgExpireAgreement: (data: MsgExpireAgreement) => EncodeObject;
    msgUpdateTimedoutAgreement: (data: MsgUpdateTimedoutAgreement) => EncodeObject;
    msgTerminateAgreement: (data: MsgTerminateAgreement) => EncodeObject;
    msgUpdateSentAgreement: (data: MsgUpdateSentAgreement) => EncodeObject;
    msgCreateTimedoutAgreement: (data: MsgCreateTimedoutAgreement) => EncodeObject;
}>;
interface QueryClientOptions {
    addr: string;
}
declare const queryClient: ({ addr: addr }?: QueryClientOptions) => Promise<Api<unknown>>;
export { txClient, queryClient, };
