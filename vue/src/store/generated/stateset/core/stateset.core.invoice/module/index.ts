// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSentInvoice } from "./types/invoice/tx";
import { MsgFactorInvoice } from "./types/invoice/tx";
import { MsgCreateTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteTimedoutInvoice } from "./types/invoice/tx";
import { MsgUpdateSentInvoice } from "./types/invoice/tx";
import { MsgUpdateTimedoutInvoice } from "./types/invoice/tx";
import { MsgDeleteSentInvoice } from "./types/invoice/tx";


const types = [
  ["/stateset.core.invoice.MsgCreateSentInvoice", MsgCreateSentInvoice],
  ["/stateset.core.invoice.MsgFactorInvoice", MsgFactorInvoice],
  ["/stateset.core.invoice.MsgCreateTimedoutInvoice", MsgCreateTimedoutInvoice],
  ["/stateset.core.invoice.MsgDeleteTimedoutInvoice", MsgDeleteTimedoutInvoice],
  ["/stateset.core.invoice.MsgUpdateSentInvoice", MsgUpdateSentInvoice],
  ["/stateset.core.invoice.MsgUpdateTimedoutInvoice", MsgUpdateTimedoutInvoice],
  ["/stateset.core.invoice.MsgDeleteSentInvoice", MsgDeleteSentInvoice],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

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

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateSentInvoice: (data: MsgCreateSentInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgCreateSentInvoice", value: data }),
    msgFactorInvoice: (data: MsgFactorInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgFactorInvoice", value: data }),
    msgCreateTimedoutInvoice: (data: MsgCreateTimedoutInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgCreateTimedoutInvoice", value: data }),
    msgDeleteTimedoutInvoice: (data: MsgDeleteTimedoutInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgDeleteTimedoutInvoice", value: data }),
    msgUpdateSentInvoice: (data: MsgUpdateSentInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgUpdateSentInvoice", value: data }),
    msgUpdateTimedoutInvoice: (data: MsgUpdateTimedoutInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgUpdateTimedoutInvoice", value: data }),
    msgDeleteSentInvoice: (data: MsgDeleteSentInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgDeleteSentInvoice", value: data }),
    
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
