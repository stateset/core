///
//  Generated code. Do not modify.
//  source: cosmwasm/wasm/v1/types.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

const AccessType$json = const {
  '1': 'AccessType',
  '2': const [
    const {'1': 'ACCESS_TYPE_UNSPECIFIED', '2': 0, '3': const {}},
    const {'1': 'ACCESS_TYPE_NOBODY', '2': 1, '3': const {}},
    const {'1': 'ACCESS_TYPE_ONLY_ADDRESS', '2': 2, '3': const {}},
    const {'1': 'ACCESS_TYPE_EVERYBODY', '2': 3, '3': const {}},
  ],
  '3': const {},
};

const ContractCodeHistoryOperationType$json = const {
  '1': 'ContractCodeHistoryOperationType',
  '2': const [
    const {'1': 'CONTRACT_CODE_HISTORY_OPERATION_TYPE_UNSPECIFIED', '2': 0, '3': const {}},
    const {'1': 'CONTRACT_CODE_HISTORY_OPERATION_TYPE_INIT', '2': 1, '3': const {}},
    const {'1': 'CONTRACT_CODE_HISTORY_OPERATION_TYPE_MIGRATE', '2': 2, '3': const {}},
    const {'1': 'CONTRACT_CODE_HISTORY_OPERATION_TYPE_GENESIS', '2': 3, '3': const {}},
  ],
  '3': const {},
};

const AccessTypeParam$json = const {
  '1': 'AccessTypeParam',
  '2': const [
    const {'1': 'value', '3': 1, '4': 1, '5': 14, '6': '.cosmwasm.wasm.v1.AccessType', '8': const {}, '10': 'value'},
  ],
  '7': const {},
};

const AccessConfig$json = const {
  '1': 'AccessConfig',
  '2': const [
    const {'1': 'permission', '3': 1, '4': 1, '5': 14, '6': '.cosmwasm.wasm.v1.AccessType', '8': const {}, '10': 'permission'},
    const {'1': 'address', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'address'},
  ],
  '7': const {},
};

const Params$json = const {
  '1': 'Params',
  '2': const [
    const {'1': 'code_upload_access', '3': 1, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.AccessConfig', '8': const {}, '10': 'codeUploadAccess'},
    const {'1': 'instantiate_default_permission', '3': 2, '4': 1, '5': 14, '6': '.cosmwasm.wasm.v1.AccessType', '8': const {}, '10': 'instantiateDefaultPermission'},
    const {'1': 'max_wasm_code_size', '3': 3, '4': 1, '5': 4, '8': const {}, '10': 'maxWasmCodeSize'},
  ],
  '7': const {},
};

const CodeInfo$json = const {
  '1': 'CodeInfo',
  '2': const [
    const {'1': 'code_hash', '3': 1, '4': 1, '5': 12, '10': 'codeHash'},
    const {'1': 'creator', '3': 2, '4': 1, '5': 9, '10': 'creator'},
    const {'1': 'instantiate_config', '3': 5, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.AccessConfig', '8': const {}, '10': 'instantiateConfig'},
  ],
  '9': const [
    const {'1': 3, '2': 4},
    const {'1': 4, '2': 5},
  ],
};

const ContractInfo$json = const {
  '1': 'ContractInfo',
  '2': const [
    const {'1': 'code_id', '3': 1, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'creator', '3': 2, '4': 1, '5': 9, '10': 'creator'},
    const {'1': 'admin', '3': 3, '4': 1, '5': 9, '10': 'admin'},
    const {'1': 'label', '3': 4, '4': 1, '5': 9, '10': 'label'},
    const {'1': 'created', '3': 5, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.AbsoluteTxPosition', '10': 'created'},
    const {'1': 'ibc_port_id', '3': 6, '4': 1, '5': 9, '8': const {}, '10': 'ibcPortId'},
    const {'1': 'extension', '3': 7, '4': 1, '5': 11, '6': '.google.protobuf.Any', '8': const {}, '10': 'extension'},
  ],
  '7': const {},
};

const ContractCodeHistoryEntry$json = const {
  '1': 'ContractCodeHistoryEntry',
  '2': const [
    const {'1': 'operation', '3': 1, '4': 1, '5': 14, '6': '.cosmwasm.wasm.v1.ContractCodeHistoryOperationType', '10': 'operation'},
    const {'1': 'code_id', '3': 2, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'updated', '3': 3, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.AbsoluteTxPosition', '10': 'updated'},
    const {'1': 'msg', '3': 4, '4': 1, '5': 12, '8': const {}, '10': 'msg'},
  ],
};

const AbsoluteTxPosition$json = const {
  '1': 'AbsoluteTxPosition',
  '2': const [
    const {'1': 'block_height', '3': 1, '4': 1, '5': 4, '10': 'blockHeight'},
    const {'1': 'tx_index', '3': 2, '4': 1, '5': 4, '10': 'txIndex'},
  ],
};

const Model$json = const {
  '1': 'Model',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 12, '8': const {}, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 12, '10': 'value'},
  ],
};

