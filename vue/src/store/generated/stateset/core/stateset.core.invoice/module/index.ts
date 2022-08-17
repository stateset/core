// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateInvoice } from "./types/invoice/tx";
import { MsgCreateSentInvoice } from "./types/invoice/tx";
import { MsgUpdateTimedoutInvoice } from "./types/invoice/tx";
import { MsgUpdateSentInvoice } from "./types/invoice/tx";
import { MsgDeleteSentInvoice } from "./types/invoice/tx";
import { MsgCreateTimedoutInvoice } from "./types/invoice/tx";
import { MsgFactorInvoice } from "./types/invoice/tx";
import { MsgDeleteTimedoutInvoice } from "./types/invoice/tx";


const types = [
  ["/stateset.core.invoice.MsgCreateInvoice", MsgCreateInvoice],
  ["/stateset.core.invoice.MsgCreateSentInvoice", MsgCreateSentInvoice],
  ["/stateset.core.invoice.MsgUpdateTimedoutInvoice", MsgUpdateTimedoutInvoice],
  ["/stateset.core.invoice.MsgUpdateSentInvoice", MsgUpdateSentInvoice],
  ["/stateset.core.invoice.MsgDeleteSentInvoice", MsgDeleteSentInvoice],
  ["/stateset.core.invoice.MsgCreateTimedoutInvoice", MsgCreateTimedoutInvoice],
  ["/stateset.core.invoice.MsgFactorInvoice", MsgFactorInvoice],
  ["/stateset.core.invoice.MsgDeleteTimedoutInvoice", MsgDeleteTimedoutInvoice],
  
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
    msgCreateInvoice: (data: MsgCreateInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgCreateInvoice", value: MsgCreateInvoice.fromPartial( data ) }),
    msgCreateSentInvoice: (data: MsgCreateSentInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgCreateSentInvoice", value: MsgCreateSentInvoice.fromPartial( data ) }),
    msgUpdateTimedoutInvoice: (data: MsgUpdateTimedoutInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgUpdateTimedoutInvoice", value: MsgUpdateTimedoutInvoice.fromPartial( data ) }),
    msgUpdateSentInvoice: (data: MsgUpdateSentInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgUpdateSentInvoice", value: MsgUpdateSentInvoice.fromPartial( data ) }),
    msgDeleteSentInvoice: (data: MsgDeleteSentInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgDeleteSentInvoice", value: MsgDeleteSentInvoice.fromPartial( data ) }),
    msgCreateTimedoutInvoice: (data: MsgCreateTimedoutInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgCreateTimedoutInvoice", value: MsgCreateTimedoutInvoice.fromPartial( data ) }),
    msgFactorInvoice: (data: MsgFactorInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgFactorInvoice", value: MsgFactorInvoice.fromPartial( data ) }),
    msgDeleteTimedoutInvoice: (data: MsgDeleteTimedoutInvoice): EncodeObject => ({ typeUrl: "/stateset.core.invoice.MsgDeleteTimedoutInvoice", value: MsgDeleteTimedoutInvoice.fromPartial( data ) }),
    
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
