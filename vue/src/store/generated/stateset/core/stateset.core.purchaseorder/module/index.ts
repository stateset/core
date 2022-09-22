// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCreateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCompletePurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteTimedoutPurchaseorder } from "./types/purchaseorder/tx";
import { MsgRequestPurchaseorder } from "./types/purchaseorder/tx";
import { MsgUpdateSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgDeleteSentPurchaseorder } from "./types/purchaseorder/tx";
import { MsgCancelPurchaseorder } from "./types/purchaseorder/tx";
import { MsgFinancePurchaseorder } from "./types/purchaseorder/tx";


const types = [
  ["/stateset.core.purchaseorder.MsgCreateSentPurchaseorder", MsgCreateSentPurchaseorder],
  ["/stateset.core.purchaseorder.MsgCreateTimedoutPurchaseorder", MsgCreateTimedoutPurchaseorder],
  ["/stateset.core.purchaseorder.MsgUpdateTimedoutPurchaseorder", MsgUpdateTimedoutPurchaseorder],
  ["/stateset.core.purchaseorder.MsgCompletePurchaseorder", MsgCompletePurchaseorder],
  ["/stateset.core.purchaseorder.MsgDeleteTimedoutPurchaseorder", MsgDeleteTimedoutPurchaseorder],
  ["/stateset.core.purchaseorder.MsgRequestPurchaseorder", MsgRequestPurchaseorder],
  ["/stateset.core.purchaseorder.MsgUpdateSentPurchaseorder", MsgUpdateSentPurchaseorder],
  ["/stateset.core.purchaseorder.MsgDeleteSentPurchaseorder", MsgDeleteSentPurchaseorder],
  ["/stateset.core.purchaseorder.MsgCancelPurchaseorder", MsgCancelPurchaseorder],
  ["/stateset.core.purchaseorder.MsgFinancePurchaseorder", MsgFinancePurchaseorder],
  
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
    msgCreateSentPurchaseorder: (data: MsgCreateSentPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCreateSentPurchaseorder", value: MsgCreateSentPurchaseorder.fromPartial( data ) }),
    msgCreateTimedoutPurchaseorder: (data: MsgCreateTimedoutPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCreateTimedoutPurchaseorder", value: MsgCreateTimedoutPurchaseorder.fromPartial( data ) }),
    msgUpdateTimedoutPurchaseorder: (data: MsgUpdateTimedoutPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgUpdateTimedoutPurchaseorder", value: MsgUpdateTimedoutPurchaseorder.fromPartial( data ) }),
    msgCompletePurchaseorder: (data: MsgCompletePurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCompletePurchaseorder", value: MsgCompletePurchaseorder.fromPartial( data ) }),
    msgDeleteTimedoutPurchaseorder: (data: MsgDeleteTimedoutPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgDeleteTimedoutPurchaseorder", value: MsgDeleteTimedoutPurchaseorder.fromPartial( data ) }),
    msgRequestPurchaseorder: (data: MsgRequestPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgRequestPurchaseorder", value: MsgRequestPurchaseorder.fromPartial( data ) }),
    msgUpdateSentPurchaseorder: (data: MsgUpdateSentPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgUpdateSentPurchaseorder", value: MsgUpdateSentPurchaseorder.fromPartial( data ) }),
    msgDeleteSentPurchaseorder: (data: MsgDeleteSentPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgDeleteSentPurchaseorder", value: MsgDeleteSentPurchaseorder.fromPartial( data ) }),
    msgCancelPurchaseorder: (data: MsgCancelPurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgCancelPurchaseorder", value: MsgCancelPurchaseorder.fromPartial( data ) }),
    msgFinancePurchaseorder: (data: MsgFinancePurchaseorder): EncodeObject => ({ typeUrl: "/stateset.core.purchaseorder.MsgFinancePurchaseorder", value: MsgFinancePurchaseorder.fromPartial( data ) }),
    
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
