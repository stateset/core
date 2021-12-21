///
//  Generated code. Do not modify.
//  source: loan/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'loan.pb.dart' as $2;
import '../cosmos/base/query/v1beta1/pagination.pb.dart' as $4;

class QueryGetLoanRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetLoanRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetLoanRequest._() : super();
  factory QueryGetLoanRequest() => create();
  factory QueryGetLoanRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetLoanRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetLoanRequest clone() => QueryGetLoanRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetLoanRequest copyWith(void Function(QueryGetLoanRequest) updates) => super.copyWith((message) => updates(message as QueryGetLoanRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetLoanRequest create() => QueryGetLoanRequest._();
  QueryGetLoanRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetLoanRequest> createRepeated() => $pb.PbList<QueryGetLoanRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetLoanRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetLoanRequest>(create);
  static QueryGetLoanRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetLoanResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetLoanResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..aOM<$2.Loan>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Loan', protoName: 'Loan', subBuilder: $2.Loan.create)
    ..hasRequiredFields = false
  ;

  QueryGetLoanResponse._() : super();
  factory QueryGetLoanResponse() => create();
  factory QueryGetLoanResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetLoanResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetLoanResponse clone() => QueryGetLoanResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetLoanResponse copyWith(void Function(QueryGetLoanResponse) updates) => super.copyWith((message) => updates(message as QueryGetLoanResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetLoanResponse create() => QueryGetLoanResponse._();
  QueryGetLoanResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetLoanResponse> createRepeated() => $pb.PbList<QueryGetLoanResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetLoanResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetLoanResponse>(create);
  static QueryGetLoanResponse _defaultInstance;

  @$pb.TagNumber(1)
  $2.Loan get loan => $_getN(0);
  @$pb.TagNumber(1)
  set loan($2.Loan v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasLoan() => $_has(0);
  @$pb.TagNumber(1)
  void clearLoan() => clearField(1);
  @$pb.TagNumber(1)
  $2.Loan ensureLoan() => $_ensure(0);
}

class QueryAllLoanRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllLoanRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..aOM<$4.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $4.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllLoanRequest._() : super();
  factory QueryAllLoanRequest() => create();
  factory QueryAllLoanRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllLoanRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllLoanRequest clone() => QueryAllLoanRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllLoanRequest copyWith(void Function(QueryAllLoanRequest) updates) => super.copyWith((message) => updates(message as QueryAllLoanRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllLoanRequest create() => QueryAllLoanRequest._();
  QueryAllLoanRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllLoanRequest> createRepeated() => $pb.PbList<QueryAllLoanRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllLoanRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllLoanRequest>(create);
  static QueryAllLoanRequest _defaultInstance;

  @$pb.TagNumber(1)
  $4.PageRequest get pagination => $_getN(0);
  @$pb.TagNumber(1)
  set pagination($4.PageRequest v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasPagination() => $_has(0);
  @$pb.TagNumber(1)
  void clearPagination() => clearField(1);
  @$pb.TagNumber(1)
  $4.PageRequest ensurePagination() => $_ensure(0);
}

class QueryAllLoanResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllLoanResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..pc<$2.Loan>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Loan', $pb.PbFieldType.PM, protoName: 'Loan', subBuilder: $2.Loan.create)
    ..aOM<$4.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $4.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllLoanResponse._() : super();
  factory QueryAllLoanResponse() => create();
  factory QueryAllLoanResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllLoanResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllLoanResponse clone() => QueryAllLoanResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllLoanResponse copyWith(void Function(QueryAllLoanResponse) updates) => super.copyWith((message) => updates(message as QueryAllLoanResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllLoanResponse create() => QueryAllLoanResponse._();
  QueryAllLoanResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllLoanResponse> createRepeated() => $pb.PbList<QueryAllLoanResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllLoanResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllLoanResponse>(create);
  static QueryAllLoanResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$2.Loan> get loan => $_getList(0);

  @$pb.TagNumber(2)
  $4.PageResponse get pagination => $_getN(1);
  @$pb.TagNumber(2)
  set pagination($4.PageResponse v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPagination() => $_has(1);
  @$pb.TagNumber(2)
  void clearPagination() => clearField(2);
  @$pb.TagNumber(2)
  $4.PageResponse ensurePagination() => $_ensure(1);
}

