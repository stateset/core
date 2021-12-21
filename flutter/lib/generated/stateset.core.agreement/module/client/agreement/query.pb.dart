///
//  Generated code. Do not modify.
//  source: agreement/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'sent_agreement.pb.dart' as $2;
import '../cosmos/base/query/v1beta1/pagination.pb.dart' as $6;
import 'timedout_agreement.pb.dart' as $3;
import 'agreement.pb.dart' as $4;

class QueryGetSentAgreementRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetSentAgreementRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetSentAgreementRequest._() : super();
  factory QueryGetSentAgreementRequest() => create();
  factory QueryGetSentAgreementRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetSentAgreementRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetSentAgreementRequest clone() => QueryGetSentAgreementRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetSentAgreementRequest copyWith(void Function(QueryGetSentAgreementRequest) updates) => super.copyWith((message) => updates(message as QueryGetSentAgreementRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetSentAgreementRequest create() => QueryGetSentAgreementRequest._();
  QueryGetSentAgreementRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetSentAgreementRequest> createRepeated() => $pb.PbList<QueryGetSentAgreementRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetSentAgreementRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetSentAgreementRequest>(create);
  static QueryGetSentAgreementRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetSentAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetSentAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOM<$2.SentAgreement>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SentAgreement', protoName: 'SentAgreement', subBuilder: $2.SentAgreement.create)
    ..hasRequiredFields = false
  ;

  QueryGetSentAgreementResponse._() : super();
  factory QueryGetSentAgreementResponse() => create();
  factory QueryGetSentAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetSentAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetSentAgreementResponse clone() => QueryGetSentAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetSentAgreementResponse copyWith(void Function(QueryGetSentAgreementResponse) updates) => super.copyWith((message) => updates(message as QueryGetSentAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetSentAgreementResponse create() => QueryGetSentAgreementResponse._();
  QueryGetSentAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetSentAgreementResponse> createRepeated() => $pb.PbList<QueryGetSentAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetSentAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetSentAgreementResponse>(create);
  static QueryGetSentAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $2.SentAgreement get sentAgreement => $_getN(0);
  @$pb.TagNumber(1)
  set sentAgreement($2.SentAgreement v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSentAgreement() => $_has(0);
  @$pb.TagNumber(1)
  void clearSentAgreement() => clearField(1);
  @$pb.TagNumber(1)
  $2.SentAgreement ensureSentAgreement() => $_ensure(0);
}

class QueryAllSentAgreementRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllSentAgreementRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllSentAgreementRequest._() : super();
  factory QueryAllSentAgreementRequest() => create();
  factory QueryAllSentAgreementRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllSentAgreementRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllSentAgreementRequest clone() => QueryAllSentAgreementRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllSentAgreementRequest copyWith(void Function(QueryAllSentAgreementRequest) updates) => super.copyWith((message) => updates(message as QueryAllSentAgreementRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllSentAgreementRequest create() => QueryAllSentAgreementRequest._();
  QueryAllSentAgreementRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllSentAgreementRequest> createRepeated() => $pb.PbList<QueryAllSentAgreementRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllSentAgreementRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllSentAgreementRequest>(create);
  static QueryAllSentAgreementRequest _defaultInstance;

  @$pb.TagNumber(1)
  $6.PageRequest get pagination => $_getN(0);
  @$pb.TagNumber(1)
  set pagination($6.PageRequest v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasPagination() => $_has(0);
  @$pb.TagNumber(1)
  void clearPagination() => clearField(1);
  @$pb.TagNumber(1)
  $6.PageRequest ensurePagination() => $_ensure(0);
}

class QueryAllSentAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllSentAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..pc<$2.SentAgreement>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SentAgreement', $pb.PbFieldType.PM, protoName: 'SentAgreement', subBuilder: $2.SentAgreement.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllSentAgreementResponse._() : super();
  factory QueryAllSentAgreementResponse() => create();
  factory QueryAllSentAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllSentAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllSentAgreementResponse clone() => QueryAllSentAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllSentAgreementResponse copyWith(void Function(QueryAllSentAgreementResponse) updates) => super.copyWith((message) => updates(message as QueryAllSentAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllSentAgreementResponse create() => QueryAllSentAgreementResponse._();
  QueryAllSentAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllSentAgreementResponse> createRepeated() => $pb.PbList<QueryAllSentAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllSentAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllSentAgreementResponse>(create);
  static QueryAllSentAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$2.SentAgreement> get sentAgreement => $_getList(0);

  @$pb.TagNumber(2)
  $6.PageResponse get pagination => $_getN(1);
  @$pb.TagNumber(2)
  set pagination($6.PageResponse v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPagination() => $_has(1);
  @$pb.TagNumber(2)
  void clearPagination() => clearField(2);
  @$pb.TagNumber(2)
  $6.PageResponse ensurePagination() => $_ensure(1);
}

class QueryGetTimedoutAgreementRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetTimedoutAgreementRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetTimedoutAgreementRequest._() : super();
  factory QueryGetTimedoutAgreementRequest() => create();
  factory QueryGetTimedoutAgreementRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTimedoutAgreementRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutAgreementRequest clone() => QueryGetTimedoutAgreementRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutAgreementRequest copyWith(void Function(QueryGetTimedoutAgreementRequest) updates) => super.copyWith((message) => updates(message as QueryGetTimedoutAgreementRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutAgreementRequest create() => QueryGetTimedoutAgreementRequest._();
  QueryGetTimedoutAgreementRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetTimedoutAgreementRequest> createRepeated() => $pb.PbList<QueryGetTimedoutAgreementRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutAgreementRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTimedoutAgreementRequest>(create);
  static QueryGetTimedoutAgreementRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetTimedoutAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetTimedoutAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOM<$3.TimedoutAgreement>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TimedoutAgreement', protoName: 'TimedoutAgreement', subBuilder: $3.TimedoutAgreement.create)
    ..hasRequiredFields = false
  ;

  QueryGetTimedoutAgreementResponse._() : super();
  factory QueryGetTimedoutAgreementResponse() => create();
  factory QueryGetTimedoutAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTimedoutAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutAgreementResponse clone() => QueryGetTimedoutAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutAgreementResponse copyWith(void Function(QueryGetTimedoutAgreementResponse) updates) => super.copyWith((message) => updates(message as QueryGetTimedoutAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutAgreementResponse create() => QueryGetTimedoutAgreementResponse._();
  QueryGetTimedoutAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetTimedoutAgreementResponse> createRepeated() => $pb.PbList<QueryGetTimedoutAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTimedoutAgreementResponse>(create);
  static QueryGetTimedoutAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $3.TimedoutAgreement get timedoutAgreement => $_getN(0);
  @$pb.TagNumber(1)
  set timedoutAgreement($3.TimedoutAgreement v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimedoutAgreement() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimedoutAgreement() => clearField(1);
  @$pb.TagNumber(1)
  $3.TimedoutAgreement ensureTimedoutAgreement() => $_ensure(0);
}

class QueryAllTimedoutAgreementRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllTimedoutAgreementRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllTimedoutAgreementRequest._() : super();
  factory QueryAllTimedoutAgreementRequest() => create();
  factory QueryAllTimedoutAgreementRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllTimedoutAgreementRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutAgreementRequest clone() => QueryAllTimedoutAgreementRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutAgreementRequest copyWith(void Function(QueryAllTimedoutAgreementRequest) updates) => super.copyWith((message) => updates(message as QueryAllTimedoutAgreementRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutAgreementRequest create() => QueryAllTimedoutAgreementRequest._();
  QueryAllTimedoutAgreementRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllTimedoutAgreementRequest> createRepeated() => $pb.PbList<QueryAllTimedoutAgreementRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutAgreementRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllTimedoutAgreementRequest>(create);
  static QueryAllTimedoutAgreementRequest _defaultInstance;

  @$pb.TagNumber(1)
  $6.PageRequest get pagination => $_getN(0);
  @$pb.TagNumber(1)
  set pagination($6.PageRequest v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasPagination() => $_has(0);
  @$pb.TagNumber(1)
  void clearPagination() => clearField(1);
  @$pb.TagNumber(1)
  $6.PageRequest ensurePagination() => $_ensure(0);
}

class QueryAllTimedoutAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllTimedoutAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..pc<$3.TimedoutAgreement>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TimedoutAgreement', $pb.PbFieldType.PM, protoName: 'TimedoutAgreement', subBuilder: $3.TimedoutAgreement.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllTimedoutAgreementResponse._() : super();
  factory QueryAllTimedoutAgreementResponse() => create();
  factory QueryAllTimedoutAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllTimedoutAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutAgreementResponse clone() => QueryAllTimedoutAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutAgreementResponse copyWith(void Function(QueryAllTimedoutAgreementResponse) updates) => super.copyWith((message) => updates(message as QueryAllTimedoutAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutAgreementResponse create() => QueryAllTimedoutAgreementResponse._();
  QueryAllTimedoutAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllTimedoutAgreementResponse> createRepeated() => $pb.PbList<QueryAllTimedoutAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllTimedoutAgreementResponse>(create);
  static QueryAllTimedoutAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.TimedoutAgreement> get timedoutAgreement => $_getList(0);

  @$pb.TagNumber(2)
  $6.PageResponse get pagination => $_getN(1);
  @$pb.TagNumber(2)
  set pagination($6.PageResponse v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPagination() => $_has(1);
  @$pb.TagNumber(2)
  void clearPagination() => clearField(2);
  @$pb.TagNumber(2)
  $6.PageResponse ensurePagination() => $_ensure(1);
}

class QueryGetAgreementRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetAgreementRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetAgreementRequest._() : super();
  factory QueryGetAgreementRequest() => create();
  factory QueryGetAgreementRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetAgreementRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetAgreementRequest clone() => QueryGetAgreementRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetAgreementRequest copyWith(void Function(QueryGetAgreementRequest) updates) => super.copyWith((message) => updates(message as QueryGetAgreementRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetAgreementRequest create() => QueryGetAgreementRequest._();
  QueryGetAgreementRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetAgreementRequest> createRepeated() => $pb.PbList<QueryGetAgreementRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetAgreementRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetAgreementRequest>(create);
  static QueryGetAgreementRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOM<$4.Agreement>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Agreement', protoName: 'Agreement', subBuilder: $4.Agreement.create)
    ..hasRequiredFields = false
  ;

  QueryGetAgreementResponse._() : super();
  factory QueryGetAgreementResponse() => create();
  factory QueryGetAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetAgreementResponse clone() => QueryGetAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetAgreementResponse copyWith(void Function(QueryGetAgreementResponse) updates) => super.copyWith((message) => updates(message as QueryGetAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetAgreementResponse create() => QueryGetAgreementResponse._();
  QueryGetAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetAgreementResponse> createRepeated() => $pb.PbList<QueryGetAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetAgreementResponse>(create);
  static QueryGetAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $4.Agreement get agreement => $_getN(0);
  @$pb.TagNumber(1)
  set agreement($4.Agreement v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasAgreement() => $_has(0);
  @$pb.TagNumber(1)
  void clearAgreement() => clearField(1);
  @$pb.TagNumber(1)
  $4.Agreement ensureAgreement() => $_ensure(0);
}

class QueryAllAgreementRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllAgreementRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllAgreementRequest._() : super();
  factory QueryAllAgreementRequest() => create();
  factory QueryAllAgreementRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllAgreementRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllAgreementRequest clone() => QueryAllAgreementRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllAgreementRequest copyWith(void Function(QueryAllAgreementRequest) updates) => super.copyWith((message) => updates(message as QueryAllAgreementRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllAgreementRequest create() => QueryAllAgreementRequest._();
  QueryAllAgreementRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllAgreementRequest> createRepeated() => $pb.PbList<QueryAllAgreementRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllAgreementRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllAgreementRequest>(create);
  static QueryAllAgreementRequest _defaultInstance;

  @$pb.TagNumber(1)
  $6.PageRequest get pagination => $_getN(0);
  @$pb.TagNumber(1)
  set pagination($6.PageRequest v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasPagination() => $_has(0);
  @$pb.TagNumber(1)
  void clearPagination() => clearField(1);
  @$pb.TagNumber(1)
  $6.PageRequest ensurePagination() => $_ensure(0);
}

class QueryAllAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..pc<$4.Agreement>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Agreement', $pb.PbFieldType.PM, protoName: 'Agreement', subBuilder: $4.Agreement.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllAgreementResponse._() : super();
  factory QueryAllAgreementResponse() => create();
  factory QueryAllAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllAgreementResponse clone() => QueryAllAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllAgreementResponse copyWith(void Function(QueryAllAgreementResponse) updates) => super.copyWith((message) => updates(message as QueryAllAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllAgreementResponse create() => QueryAllAgreementResponse._();
  QueryAllAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllAgreementResponse> createRepeated() => $pb.PbList<QueryAllAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllAgreementResponse>(create);
  static QueryAllAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$4.Agreement> get agreement => $_getList(0);

  @$pb.TagNumber(2)
  $6.PageResponse get pagination => $_getN(1);
  @$pb.TagNumber(2)
  set pagination($6.PageResponse v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasPagination() => $_has(1);
  @$pb.TagNumber(2)
  void clearPagination() => clearField(2);
  @$pb.TagNumber(2)
  $6.PageResponse ensurePagination() => $_ensure(1);
}

