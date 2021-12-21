///
//  Generated code. Do not modify.
//  source: cosmwasm/wasm/v1/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

const QueryContractInfoRequest$json = const {
  '1': 'QueryContractInfoRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
  ],
};

const QueryContractInfoResponse$json = const {
  '1': 'QueryContractInfoResponse',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'contract_info', '3': 2, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.ContractInfo', '8': const {}, '10': 'contractInfo'},
  ],
  '7': const {},
};

const QueryContractHistoryRequest$json = const {
  '1': 'QueryContractHistoryRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageRequest', '10': 'pagination'},
  ],
};

const QueryContractHistoryResponse$json = const {
  '1': 'QueryContractHistoryResponse',
  '2': const [
    const {'1': 'entries', '3': 1, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.ContractCodeHistoryEntry', '8': const {}, '10': 'entries'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageResponse', '10': 'pagination'},
  ],
};

const QueryContractsByCodeRequest$json = const {
  '1': 'QueryContractsByCodeRequest',
  '2': const [
    const {'1': 'code_id', '3': 1, '4': 1, '5': 4, '10': 'codeId'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageRequest', '10': 'pagination'},
  ],
};

const QueryContractsByCodeResponse$json = const {
  '1': 'QueryContractsByCodeResponse',
  '2': const [
    const {'1': 'contracts', '3': 1, '4': 3, '5': 9, '10': 'contracts'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageResponse', '10': 'pagination'},
  ],
};

const QueryAllContractStateRequest$json = const {
  '1': 'QueryAllContractStateRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageRequest', '10': 'pagination'},
  ],
};

const QueryAllContractStateResponse$json = const {
  '1': 'QueryAllContractStateResponse',
  '2': const [
    const {'1': 'models', '3': 1, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.Model', '8': const {}, '10': 'models'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageResponse', '10': 'pagination'},
  ],
};

const QueryRawContractStateRequest$json = const {
  '1': 'QueryRawContractStateRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'query_data', '3': 2, '4': 1, '5': 12, '10': 'queryData'},
  ],
};

const QueryRawContractStateResponse$json = const {
  '1': 'QueryRawContractStateResponse',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 12, '10': 'data'},
  ],
};

const QuerySmartContractStateRequest$json = const {
  '1': 'QuerySmartContractStateRequest',
  '2': const [
    const {'1': 'address', '3': 1, '4': 1, '5': 9, '10': 'address'},
    const {'1': 'query_data', '3': 2, '4': 1, '5': 12, '8': const {}, '10': 'queryData'},
  ],
};

const QuerySmartContractStateResponse$json = const {
  '1': 'QuerySmartContractStateResponse',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 12, '8': const {}, '10': 'data'},
  ],
};

const QueryCodeRequest$json = const {
  '1': 'QueryCodeRequest',
  '2': const [
    const {'1': 'code_id', '3': 1, '4': 1, '5': 4, '10': 'codeId'},
  ],
};

const CodeInfoResponse$json = const {
  '1': 'CodeInfoResponse',
  '2': const [
    const {'1': 'code_id', '3': 1, '4': 1, '5': 4, '8': const {}, '10': 'codeId'},
    const {'1': 'creator', '3': 2, '4': 1, '5': 9, '10': 'creator'},
    const {'1': 'data_hash', '3': 3, '4': 1, '5': 12, '8': const {}, '10': 'dataHash'},
  ],
  '7': const {},
  '9': const [
    const {'1': 4, '2': 5},
    const {'1': 5, '2': 6},
  ],
};

const QueryCodeResponse$json = const {
  '1': 'QueryCodeResponse',
  '2': const [
    const {'1': 'code_info', '3': 1, '4': 1, '5': 11, '6': '.cosmwasm.wasm.v1.CodeInfoResponse', '8': const {}, '10': 'codeInfo'},
    const {'1': 'data', '3': 2, '4': 1, '5': 12, '8': const {}, '10': 'data'},
  ],
  '7': const {},
};

const QueryCodesRequest$json = const {
  '1': 'QueryCodesRequest',
  '2': const [
    const {'1': 'pagination', '3': 1, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageRequest', '10': 'pagination'},
  ],
};

const QueryCodesResponse$json = const {
  '1': 'QueryCodesResponse',
  '2': const [
    const {'1': 'code_infos', '3': 1, '4': 3, '5': 11, '6': '.cosmwasm.wasm.v1.CodeInfoResponse', '8': const {}, '10': 'codeInfos'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageResponse', '10': 'pagination'},
  ],
};

const QueryPinnedCodesRequest$json = const {
  '1': 'QueryPinnedCodesRequest',
  '2': const [
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageRequest', '10': 'pagination'},
  ],
};

const QueryPinnedCodesResponse$json = const {
  '1': 'QueryPinnedCodesResponse',
  '2': const [
    const {'1': 'code_ids', '3': 1, '4': 3, '5': 4, '8': const {}, '10': 'codeIds'},
    const {'1': 'pagination', '3': 2, '4': 1, '5': 11, '6': '.cosmos.base.query.v1beta1.PageResponse', '10': 'pagination'},
  ],
};

