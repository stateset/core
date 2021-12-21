///
//  Generated code. Do not modify.
//  source: cosmwasm/wasm/v1/proposal.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

const StoreCodeProposal$json = const {
  '1': 'StoreCodeProposal',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'run_as', '3': 3, '4': 1, '5': 9, '10': 'runAs'},
    const {'1': 'wasm_byte_code', '3': 4, '4': 1, '5': 12, '8': const {}, '10': 'wasmByteCode'},
    const {'1': 'instantiate_permission', '3': 7, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.AccessConfig', '10': 'instantiatePermission'},
  ],
  '9': const [
    const {'1': 5, '2': 6},
    const {'1': 6, '2': 7},
  ],
};

const InstantiateContractProposal$json = const {
  '1': 'InstantiateContractProposal',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'run_as', '3': 3, '4': 1, '5': 9, '10': 'runAs'},
    const {'1': 'admin', '3': 4, '4': 1, '5': 9, '10': 'admin'},
    const {'1': 'code_id', '3': 5, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'label', '3': 6, '4': 1, '5': 9, '10': 'label'},
    const {'1': 'msg', '3': 7, '4': 1, '5': 12, '8': const {}, '10': 'msg'},
    const {'1': 'funds', '3': 8, '4': 3, '5': 11, '6': '.cosmos.base.v1beta1.Coin', '8': const {}, '10': 'funds'},
  ],
};

const MigrateContractProposal$json = const {
  '1': 'MigrateContractProposal',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'run_as', '3': 3, '4': 1, '5': 9, '10': 'runAs'},
    const {'1': 'contract', '3': 4, '4': 1, '5': 9, '10': 'contract'},
    const {'1': 'code_id', '3': 5, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'msg', '3': 6, '4': 1, '5': 12, '8': const {}, '10': 'msg'},
  ],
};

const UpdateAdminProposal$json = const {
  '1': 'UpdateAdminProposal',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'new_admin', '3': 3, '4': 1, '5': 9, '8': const {}, '10': 'newAdmin'},
    const {'1': 'contract', '3': 4, '4': 1, '5': 9, '10': 'contract'},
  ],
};

const ClearAdminProposal$json = const {
  '1': 'ClearAdminProposal',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '10': 'title'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'contract', '3': 3, '4': 1, '5': 9, '10': 'contract'},
  ],
};

const PinCodesProposal$json = const {
  '1': 'PinCodesProposal',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'title'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'description'},
    const {'1': 'code_ids', '3': 3, '4': 3, '5': 4, '8': const {}, '10': 'codeIds'},
  ],
};

const UnpinCodesProposal$json = const {
  '1': 'UnpinCodesProposal',
  '2': const [
    const {'1': 'title', '3': 1, '4': 1, '5': 9, '8': const {}, '10': 'title'},
    const {'1': 'description', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'description'},
    const {'1': 'code_ids', '3': 3, '4': 3, '5': 4, '8': const {}, '10': 'codeIds'},
  ],
};

