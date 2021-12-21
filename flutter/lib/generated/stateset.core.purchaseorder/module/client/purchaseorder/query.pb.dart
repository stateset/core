///
//  Generated code. Do not modify.
//  source: purchaseorder/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'purchaseorder.pb.dart' as $2;
import '../cosmos/base/query/v1beta1/pagination.pb.dart' as $6;
import 'sent_purchaseorder.pb.dart' as $3;
import 'timedout_purchaseorder.pb.dart' as $4;

class QueryGetPurchaseorderRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetPurchaseorderRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetPurchaseorderRequest._() : super();
  factory QueryGetPurchaseorderRequest() => create();
  factory QueryGetPurchaseorderRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetPurchaseorderRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetPurchaseorderRequest clone() => QueryGetPurchaseorderRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetPurchaseorderRequest copyWith(void Function(QueryGetPurchaseorderRequest) updates) => super.copyWith((message) => updates(message as QueryGetPurchaseorderRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetPurchaseorderRequest create() => QueryGetPurchaseorderRequest._();
  QueryGetPurchaseorderRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetPurchaseorderRequest> createRepeated() => $pb.PbList<QueryGetPurchaseorderRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetPurchaseorderRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetPurchaseorderRequest>(create);
  static QueryGetPurchaseorderRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOM<$2.Purchaseorder>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Purchaseorder', protoName: 'Purchaseorder', subBuilder: $2.Purchaseorder.create)
    ..hasRequiredFields = false
  ;

  QueryGetPurchaseorderResponse._() : super();
  factory QueryGetPurchaseorderResponse() => create();
  factory QueryGetPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetPurchaseorderResponse clone() => QueryGetPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetPurchaseorderResponse copyWith(void Function(QueryGetPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as QueryGetPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetPurchaseorderResponse create() => QueryGetPurchaseorderResponse._();
  QueryGetPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetPurchaseorderResponse> createRepeated() => $pb.PbList<QueryGetPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetPurchaseorderResponse>(create);
  static QueryGetPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $2.Purchaseorder get purchaseorder => $_getN(0);
  @$pb.TagNumber(1)
  set purchaseorder($2.Purchaseorder v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasPurchaseorder() => $_has(0);
  @$pb.TagNumber(1)
  void clearPurchaseorder() => clearField(1);
  @$pb.TagNumber(1)
  $2.Purchaseorder ensurePurchaseorder() => $_ensure(0);
}

class QueryAllPurchaseorderRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllPurchaseorderRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllPurchaseorderRequest._() : super();
  factory QueryAllPurchaseorderRequest() => create();
  factory QueryAllPurchaseorderRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllPurchaseorderRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllPurchaseorderRequest clone() => QueryAllPurchaseorderRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllPurchaseorderRequest copyWith(void Function(QueryAllPurchaseorderRequest) updates) => super.copyWith((message) => updates(message as QueryAllPurchaseorderRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllPurchaseorderRequest create() => QueryAllPurchaseorderRequest._();
  QueryAllPurchaseorderRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllPurchaseorderRequest> createRepeated() => $pb.PbList<QueryAllPurchaseorderRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllPurchaseorderRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllPurchaseorderRequest>(create);
  static QueryAllPurchaseorderRequest _defaultInstance;

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

class QueryAllPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..pc<$2.Purchaseorder>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Purchaseorder', $pb.PbFieldType.PM, protoName: 'Purchaseorder', subBuilder: $2.Purchaseorder.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllPurchaseorderResponse._() : super();
  factory QueryAllPurchaseorderResponse() => create();
  factory QueryAllPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllPurchaseorderResponse clone() => QueryAllPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllPurchaseorderResponse copyWith(void Function(QueryAllPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as QueryAllPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllPurchaseorderResponse create() => QueryAllPurchaseorderResponse._();
  QueryAllPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllPurchaseorderResponse> createRepeated() => $pb.PbList<QueryAllPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllPurchaseorderResponse>(create);
  static QueryAllPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$2.Purchaseorder> get purchaseorder => $_getList(0);

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

class QueryGetSentPurchaseorderRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetSentPurchaseorderRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetSentPurchaseorderRequest._() : super();
  factory QueryGetSentPurchaseorderRequest() => create();
  factory QueryGetSentPurchaseorderRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetSentPurchaseorderRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetSentPurchaseorderRequest clone() => QueryGetSentPurchaseorderRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetSentPurchaseorderRequest copyWith(void Function(QueryGetSentPurchaseorderRequest) updates) => super.copyWith((message) => updates(message as QueryGetSentPurchaseorderRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetSentPurchaseorderRequest create() => QueryGetSentPurchaseorderRequest._();
  QueryGetSentPurchaseorderRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetSentPurchaseorderRequest> createRepeated() => $pb.PbList<QueryGetSentPurchaseorderRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetSentPurchaseorderRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetSentPurchaseorderRequest>(create);
  static QueryGetSentPurchaseorderRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetSentPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetSentPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOM<$3.SentPurchaseorder>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SentPurchaseorder', protoName: 'SentPurchaseorder', subBuilder: $3.SentPurchaseorder.create)
    ..hasRequiredFields = false
  ;

  QueryGetSentPurchaseorderResponse._() : super();
  factory QueryGetSentPurchaseorderResponse() => create();
  factory QueryGetSentPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetSentPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetSentPurchaseorderResponse clone() => QueryGetSentPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetSentPurchaseorderResponse copyWith(void Function(QueryGetSentPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as QueryGetSentPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetSentPurchaseorderResponse create() => QueryGetSentPurchaseorderResponse._();
  QueryGetSentPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetSentPurchaseorderResponse> createRepeated() => $pb.PbList<QueryGetSentPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetSentPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetSentPurchaseorderResponse>(create);
  static QueryGetSentPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $3.SentPurchaseorder get sentPurchaseorder => $_getN(0);
  @$pb.TagNumber(1)
  set sentPurchaseorder($3.SentPurchaseorder v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSentPurchaseorder() => $_has(0);
  @$pb.TagNumber(1)
  void clearSentPurchaseorder() => clearField(1);
  @$pb.TagNumber(1)
  $3.SentPurchaseorder ensureSentPurchaseorder() => $_ensure(0);
}

class QueryAllSentPurchaseorderRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllSentPurchaseorderRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllSentPurchaseorderRequest._() : super();
  factory QueryAllSentPurchaseorderRequest() => create();
  factory QueryAllSentPurchaseorderRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllSentPurchaseorderRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllSentPurchaseorderRequest clone() => QueryAllSentPurchaseorderRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllSentPurchaseorderRequest copyWith(void Function(QueryAllSentPurchaseorderRequest) updates) => super.copyWith((message) => updates(message as QueryAllSentPurchaseorderRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllSentPurchaseorderRequest create() => QueryAllSentPurchaseorderRequest._();
  QueryAllSentPurchaseorderRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllSentPurchaseorderRequest> createRepeated() => $pb.PbList<QueryAllSentPurchaseorderRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllSentPurchaseorderRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllSentPurchaseorderRequest>(create);
  static QueryAllSentPurchaseorderRequest _defaultInstance;

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

class QueryAllSentPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllSentPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..pc<$3.SentPurchaseorder>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SentPurchaseorder', $pb.PbFieldType.PM, protoName: 'SentPurchaseorder', subBuilder: $3.SentPurchaseorder.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllSentPurchaseorderResponse._() : super();
  factory QueryAllSentPurchaseorderResponse() => create();
  factory QueryAllSentPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllSentPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllSentPurchaseorderResponse clone() => QueryAllSentPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllSentPurchaseorderResponse copyWith(void Function(QueryAllSentPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as QueryAllSentPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllSentPurchaseorderResponse create() => QueryAllSentPurchaseorderResponse._();
  QueryAllSentPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllSentPurchaseorderResponse> createRepeated() => $pb.PbList<QueryAllSentPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllSentPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllSentPurchaseorderResponse>(create);
  static QueryAllSentPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.SentPurchaseorder> get sentPurchaseorder => $_getList(0);

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

class QueryGetTimedoutPurchaseorderRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetTimedoutPurchaseorderRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetTimedoutPurchaseorderRequest._() : super();
  factory QueryGetTimedoutPurchaseorderRequest() => create();
  factory QueryGetTimedoutPurchaseorderRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTimedoutPurchaseorderRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutPurchaseorderRequest clone() => QueryGetTimedoutPurchaseorderRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutPurchaseorderRequest copyWith(void Function(QueryGetTimedoutPurchaseorderRequest) updates) => super.copyWith((message) => updates(message as QueryGetTimedoutPurchaseorderRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutPurchaseorderRequest create() => QueryGetTimedoutPurchaseorderRequest._();
  QueryGetTimedoutPurchaseorderRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetTimedoutPurchaseorderRequest> createRepeated() => $pb.PbList<QueryGetTimedoutPurchaseorderRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutPurchaseorderRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTimedoutPurchaseorderRequest>(create);
  static QueryGetTimedoutPurchaseorderRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetTimedoutPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetTimedoutPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOM<$4.TimedoutPurchaseorder>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TimedoutPurchaseorder', protoName: 'TimedoutPurchaseorder', subBuilder: $4.TimedoutPurchaseorder.create)
    ..hasRequiredFields = false
  ;

  QueryGetTimedoutPurchaseorderResponse._() : super();
  factory QueryGetTimedoutPurchaseorderResponse() => create();
  factory QueryGetTimedoutPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTimedoutPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutPurchaseorderResponse clone() => QueryGetTimedoutPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutPurchaseorderResponse copyWith(void Function(QueryGetTimedoutPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as QueryGetTimedoutPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutPurchaseorderResponse create() => QueryGetTimedoutPurchaseorderResponse._();
  QueryGetTimedoutPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetTimedoutPurchaseorderResponse> createRepeated() => $pb.PbList<QueryGetTimedoutPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTimedoutPurchaseorderResponse>(create);
  static QueryGetTimedoutPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $4.TimedoutPurchaseorder get timedoutPurchaseorder => $_getN(0);
  @$pb.TagNumber(1)
  set timedoutPurchaseorder($4.TimedoutPurchaseorder v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimedoutPurchaseorder() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimedoutPurchaseorder() => clearField(1);
  @$pb.TagNumber(1)
  $4.TimedoutPurchaseorder ensureTimedoutPurchaseorder() => $_ensure(0);
}

class QueryAllTimedoutPurchaseorderRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllTimedoutPurchaseorderRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllTimedoutPurchaseorderRequest._() : super();
  factory QueryAllTimedoutPurchaseorderRequest() => create();
  factory QueryAllTimedoutPurchaseorderRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllTimedoutPurchaseorderRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutPurchaseorderRequest clone() => QueryAllTimedoutPurchaseorderRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutPurchaseorderRequest copyWith(void Function(QueryAllTimedoutPurchaseorderRequest) updates) => super.copyWith((message) => updates(message as QueryAllTimedoutPurchaseorderRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutPurchaseorderRequest create() => QueryAllTimedoutPurchaseorderRequest._();
  QueryAllTimedoutPurchaseorderRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllTimedoutPurchaseorderRequest> createRepeated() => $pb.PbList<QueryAllTimedoutPurchaseorderRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutPurchaseorderRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllTimedoutPurchaseorderRequest>(create);
  static QueryAllTimedoutPurchaseorderRequest _defaultInstance;

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

class QueryAllTimedoutPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllTimedoutPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..pc<$4.TimedoutPurchaseorder>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TimedoutPurchaseorder', $pb.PbFieldType.PM, protoName: 'TimedoutPurchaseorder', subBuilder: $4.TimedoutPurchaseorder.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllTimedoutPurchaseorderResponse._() : super();
  factory QueryAllTimedoutPurchaseorderResponse() => create();
  factory QueryAllTimedoutPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllTimedoutPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutPurchaseorderResponse clone() => QueryAllTimedoutPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutPurchaseorderResponse copyWith(void Function(QueryAllTimedoutPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as QueryAllTimedoutPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutPurchaseorderResponse create() => QueryAllTimedoutPurchaseorderResponse._();
  QueryAllTimedoutPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllTimedoutPurchaseorderResponse> createRepeated() => $pb.PbList<QueryAllTimedoutPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllTimedoutPurchaseorderResponse>(create);
  static QueryAllTimedoutPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$4.TimedoutPurchaseorder> get timedoutPurchaseorder => $_getList(0);

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

