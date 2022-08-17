// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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


const types = [
  ["/stateset.core.agreement.MsgRenewAgreement", MsgRenewAgreement],
  ["/stateset.core.agreement.MsgActivateAgreement", MsgActivateAgreement],
  ["/stateset.core.agreement.MsgCreateSentAgreement", MsgCreateSentAgreement],
  ["/stateset.core.agreement.MsgDeleteSentAgreement", MsgDeleteSentAgreement],
  ["/stateset.core.agreement.MsgDeleteTimedoutAgreement", MsgDeleteTimedoutAgreement],
  ["/stateset.core.agreement.MsgExpireAgreement", MsgExpireAgreement],
  ["/stateset.core.agreement.MsgUpdateTimedoutAgreement", MsgUpdateTimedoutAgreement],
  ["/stateset.core.agreement.MsgTerminateAgreement", MsgTerminateAgreement],
  ["/stateset.core.agreement.MsgUpdateSentAgreement", MsgUpdateSentAgreement],
  ["/stateset.core.agreement.MsgCreateTimedoutAgreement", MsgCreateTimedoutAgreement],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRenewAgreement: (data: MsgRenewAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgRenewAgreement", value: MsgRenewAgreement.fromPartial( data ) }),
    msgActivateAgreement: (data: MsgActivateAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgActivateAgreement", value: MsgActivateAgreement.fromPartial( data ) }),
    msgCreateSentAgreement: (data: MsgCreateSentAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgCreateSentAgreement", value: MsgCreateSentAgreement.fromPartial( data ) }),
    msgDeleteSentAgreement: (data: MsgDeleteSentAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgDeleteSentAgreement", value: MsgDeleteSentAgreement.fromPartial( data ) }),
    msgDeleteTimedoutAgreement: (data: MsgDeleteTimedoutAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgDeleteTimedoutAgreement", value: MsgDeleteTimedoutAgreement.fromPartial( data ) }),
    msgExpireAgreement: (data: MsgExpireAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgExpireAgreement", value: MsgExpireAgreement.fromPartial( data ) }),
    msgUpdateTimedoutAgreement: (data: MsgUpdateTimedoutAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgUpdateTimedoutAgreement", value: MsgUpdateTimedoutAgreement.fromPartial( data ) }),
    msgTerminateAgreement: (data: MsgTerminateAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgTerminateAgreement", value: MsgTerminateAgreement.fromPartial( data ) }),
    msgUpdateSentAgreement: (data: MsgUpdateSentAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgUpdateSentAgreement", value: MsgUpdateSentAgreement.fromPartial( data ) }),
    msgCreateTimedoutAgreement: (data: MsgCreateTimedoutAgreement): EncodeObject => ({ typeUrl: "/stateset.core.agreement.MsgCreateTimedoutAgreement", value: MsgCreateTimedoutAgreement.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
