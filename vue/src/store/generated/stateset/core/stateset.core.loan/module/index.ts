// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRequestLoan } from "./types/loan/tx";
import { MsgCancelLoan } from "./types/loan/tx";
import { MsgLiquidateLoan } from "./types/loan/tx";
import { MsgApproveLoan } from "./types/loan/tx";
import { MsgRepayLoan } from "./types/loan/tx";


const types = [
  ["/stateset.core.loan.MsgRequestLoan", MsgRequestLoan],
  ["/stateset.core.loan.MsgCancelLoan", MsgCancelLoan],
  ["/stateset.core.loan.MsgLiquidateLoan", MsgLiquidateLoan],
  ["/stateset.core.loan.MsgApproveLoan", MsgApproveLoan],
  ["/stateset.core.loan.MsgRepayLoan", MsgRepayLoan],
  
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
    msgRequestLoan: (data: MsgRequestLoan): EncodeObject => ({ typeUrl: "/stateset.core.loan.MsgRequestLoan", value: data }),
    msgCancelLoan: (data: MsgCancelLoan): EncodeObject => ({ typeUrl: "/stateset.core.loan.MsgCancelLoan", value: data }),
    msgLiquidateLoan: (data: MsgLiquidateLoan): EncodeObject => ({ typeUrl: "/stateset.core.loan.MsgLiquidateLoan", value: data }),
    msgApproveLoan: (data: MsgApproveLoan): EncodeObject => ({ typeUrl: "/stateset.core.loan.MsgApproveLoan", value: data }),
    msgRepayLoan: (data: MsgRepayLoan): EncodeObject => ({ typeUrl: "/stateset.core.loan.MsgRepayLoan", value: data }),
    
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
