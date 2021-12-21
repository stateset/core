///
//  Generated code. Do not modify.
//  source: cosmwasm/wasm/v1/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

const MsgStoreCode$json = const {
  '1': 'MsgStoreCode',
  '2': const [
    const {'1': 'sender', '3': 1, '4': 1, '5': 9, '10': 'sender'},
    const {'1': 'wasm_byte_code', '3': 2, '4': 1, '5': 12, '8': const {}, '10': 'wasmByteCode'},
    const {'1': 'instantiate_permission', '3': 5, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.AccessConfig', '10': 'instantiatePermission'},
  ],
  '9': const [
    const {'1': 3, '2': 4},
    const {'1': 4, '2': 5},
  ],
};

const MsgStoreCodeResponse$json = const {
  '1': 'MsgStoreCodeResponse',
  '2': const [
    const {'1': 'code_id', '3': 1, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
  ],
};

const MsgInstantiateContract$json = const {
  '1': 'MsgInstantiateContract',
  '2': const [
    const {'1': 'sender', '3': 1, '4': 1, '5': 9, '10': 'sender'},
    const {'1': 'admin', '3': 2, '4': 1, '5': 9, '10': 'admin'},
    const {'1': 'code_id', '3': 3, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'label', '3': 4, '4': 1, '5': 9, '10': 'label'},
    const {'1': 'msg', '3': 5, '4': 1, '5': 12, '8': const {}, '10': 'msg'},
    const {'1': 'funds', '3': 6, '4': 3, '5': 11, '6': '.cosmos.base.v1beta1.Coin', '8': const {}, '10': 'funds'},
  ],
};

const MsgInstantiateContractResponse$json = const {
  '1': 'MsgInstantiateContractResponse',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'data', '3': 2, '4': 1, '5': 12, '10': 'data'},
  ],
};

const MsgExecuteContract$json = const {
  '1': 'MsgExecuteContract',
  '2': const [
    const {'1': 'sender', '3': 1, '4': 1, '5': 9, '10': 'sender'},
    const {'1': 'contract', '3': 2, '4': 1, '5': 9, '10': 'contract'},
    const {'1': 'msg', '3': 3, '4': 1, '5': 12, '8': const {}, '10': 'msg'},
    const {'1': 'funds', '3': 5, '4': 3, '5': 11, '6': '.cosmos.base.v1beta1.Coin', '8': const {}, '10': 'funds'},
  ],
};

const MsgExecuteContractResponse$json = const {
  '1': 'MsgExecuteContractResponse',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 12, '10': 'data'},
  ],
};

const MsgMigrateContract$json = const {
  '1': 'MsgMigrateContract',
  '2': const [
    const {'1': 'sender', '3': 1, '4': 1, '5': 9, '10': 'sender'},
    const {'1': 'contract', '3': 2, '4': 1, '5': 9, '10': 'contract'},
    const {'1': 'code_id', '3': 3, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'msg', '3': 4, '4': 1, '5': 12, '8': const {}, '10': 'msg'},
  ],
};

const MsgMigrateContractResponse$json = const {
  '1': 'MsgMigrateContractResponse',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 12, '10': 'data'},
  ],
};

const MsgUpdateAdmin$json = const {
  '1': 'MsgUpdateAdmin',
  '2': const [
    const {'1': 'sender', '3': 1, '4': 1, '5': 9, '10': 'sender'},
    const {'1': 'new_admin', '3': 2, '4': 1, '5': 9, '10': 'newAdmin'},
    const {'1': 'contract', '3': 3, '4': 1, '5': 9, '10': 'contract'},
  ],
};

const MsgUpdateAdminResponse$json = const {
  '1': 'MsgUpdateAdminResponse',
};

const MsgClearAdmin$json = const {
  '1': 'MsgClearAdmin',
  '2': const [
    const {'1': 'sender', '3': 1, '4': 1, '5': 9, '10': 'sender'},
    const {'1': 'contract', '3': 3, '4': 1, '5': 9, '10': 'contract'},
  ],
};

const MsgClearAdminResponse$json = const {
  '1': 'MsgClearAdminResponse',
};

