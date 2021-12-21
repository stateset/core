///
//  Generated code. Do not modify.
//  source: cosmwasm/wasm/v1/genesis.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

const GenesisState$json = const {
  '1': 'GenesisState',
  '2': const [
    const {'1': 'params', '3': 1, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.Params', '8': const {}, '10': 'params'},
    const {'1': 'codes', '3': 2, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.Code', '8': const {}, '10': 'codes'},
    const {'1': 'contracts', '3': 3, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.Contract', '8': const {}, '10': 'contracts'},
    const {'1': 'sequences', '3': 4, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.Sequence', '8': const {}, '10': 'sequences'},
    const {'1': 'gen_msgs', '3': 5, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.GenesisState.GenMsgs', '8': const {}, '10': 'genMsgs'},
  ],
  '3': const [GenesisState_GenMsgs$json],
};

const GenesisState_GenMsgs$json = const {
  '1': 'GenMsgs',
  '2': const [
    const {'1': 'store_code', '3': 1, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.MsgStoreCode', '9': 0, '10': 'storeCode'},
    const {'1': 'instantiate_contract', '3': 2, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.MsgInstantiateContract', '9': 0, '10': 'instantiateContract'},
    const {'1': 'execute_contract', '3': 3, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.MsgExecuteContract', '9': 0, '10': 'executeContract'},
  ],
  '8': const [
    const {'1': 'sum'},
  ],
};

const Code$json = const {
  '1': 'Code',
  '2': const [
    const {'1': 'code_id', '3': 1, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'code_info', '3': 2, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.CodeInfo', '8': const {}, '10': 'codeInfo'},
    const {'1': 'code_bytes', '3': 3, '4': 1, '5': 12, '10': 'codeBytes'},
    const {'1': 'pinned', '3': 4, '4': 1, '5': 8, '10': 'pinned'},
  ],
};

const Contract$json = const {
  '1': 'Contract',
  '2': const [
    const {'1': 'contract_address', '3': 1, '4': 1, '5': 9, '10': 'contractAddress'},
    const {'1': 'contract_info', '3': 2, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.ContractInfo', '8': const {}, '10': 'contractInfo'},
    const {'1': 'contract_state', '3': 3, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.Model', '8': const {}, '10': 'contractState'},
  ],
};

const Sequence$json = const {
  '1': 'Sequence',
  '2': const [
    const {'1': 'id_key', '3': 1, '4': 1, '5': 12, '8': const {}, '10': 'idKey'},
    const {'1': 'value', '3': 2, '4': 1, '5': 4, '10': 'value'},
  ],
};

