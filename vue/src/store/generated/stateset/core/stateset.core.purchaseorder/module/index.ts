// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgRequestPurchaseorder } from "./types/purchaseorder/tx";
import { MsgFinancePurchaseorder } from "./types/purchaseorder/tx";
import { MsgCancelPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCreateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCompletePurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteSentPurchaseorder } from "./types/purchaseorder/tx";


const types = [
  ["/stateset.core.purchaseorder.MsgCreateTimedoutPurchaseorder", MsgCreateTimedoutPurchaseorder],
  ["/stateset.core.purchaseorder.MsgDeleteTimedoutPurchaseorder", MsgDeleteTimedoutPurchaseorder],
  ["/stateset.core.purchaseorder.MsgRequestPurchaseorder", MsgRequestPurchaseorder],
  ["/stateset.core.purchaseorder.MsgFinancePurchaseorder", MsgFinancePurchaseorder],
  ["/stateset.core.purchaseorder.MsgCancelPurchaseorder", MsgCancelPurchaseorder],
  ["/stateset.core.purchaseorder.MsgCreateSentPurchaseorder", MsgCreateSentPurchaseorder],
  ["/stateset.core.purchaseorder.MsgCompletePurchaseorder", MsgCompletePurchaseorder],
  ["/stateset.core.purchaseorder.MsgUpdateTimedoutPurchaseorder", MsgUpdateTimedoutPurchaseorder],
  ["/stateset.core.purchaseorder.MsgUpdateSentPurchaseorder", MsgUpdateSentPurchaseorder],
  ["/stateset.core.purchaseorder.MsgDeleteSentPurchaseorder", MsgDeleteSentPurchaseorder],
  
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
    msgCreateTimedoutPurchaseorder: (data: MsgCreateTimedoutPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCreateTimedoutPurchaseorder", value: data }),
    msgDeleteTimedoutPurchaseorder: (data: MsgDeleteTimedoutPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgDeleteTimedoutPurchaseorder", value: data }),
    msgRequestPurchaseorder: (data: MsgRequestPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgRequestPurchaseorder", value: data }),
    msgFinancePurchaseorder: (data: MsgFinancePurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgFinancePurchaseorder", value: data }),
    msgCancelPurchaseorder: (data: MsgCancelPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCancelPurchaseorder", value: data }),
    msgCreateSentPurchaseorder: (data: MsgCreateSentPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCreateSentPurchaseorder", value: data }),
    msgCompletePurchaseorder: (data: MsgCompletePurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCompletePurchaseorder", value: data }),
    msgUpdateTimedoutPurchaseorder: (data: MsgUpdateTimedoutPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgUpdateTimedoutPurchaseorder", value: data }),
    msgUpdateSentPurchaseorder: (data: MsgUpdateSentPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgUpdateSentPurchaseorder", value: data }),
    msgDeleteSentPurchaseorder: (data: MsgDeleteSentPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgDeleteSentPurchaseorder", value: data }),
    
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
