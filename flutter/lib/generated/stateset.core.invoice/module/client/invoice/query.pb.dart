///
//  Generated code. Do not modify.
//  source: invoice/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'invoice.pb.dart' as $2;
import '../cosmos/base/query/v1beta1/pagination.pb.dart' as $6;
import 'sent_invoice.pb.dart' as $3;
import 'timedout_invoice.pb.dart' as $4;

class QueryGetInvoiceRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetInvoiceRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetInvoiceRequest._() : super();
  factory QueryGetInvoiceRequest() => create();
  factory QueryGetInvoiceRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetInvoiceRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetInvoiceRequest clone() => QueryGetInvoiceRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetInvoiceRequest copyWith(void Function(QueryGetInvoiceRequest) updates) => super.copyWith((message) => updates(message as QueryGetInvoiceRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetInvoiceRequest create() => QueryGetInvoiceRequest._();
  QueryGetInvoiceRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetInvoiceRequest> createRepeated() => $pb.PbList<QueryGetInvoiceRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetInvoiceRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetInvoiceRequest>(create);
  static QueryGetInvoiceRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOM<$2.Invoice>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Invoice', protoName: 'Invoice', subBuilder: $2.Invoice.create)
    ..hasRequiredFields = false
  ;

  QueryGetInvoiceResponse._() : super();
  factory QueryGetInvoiceResponse() => create();
  factory QueryGetInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetInvoiceResponse clone() => QueryGetInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetInvoiceResponse copyWith(void Function(QueryGetInvoiceResponse) updates) => super.copyWith((message) => updates(message as QueryGetInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetInvoiceResponse create() => QueryGetInvoiceResponse._();
  QueryGetInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetInvoiceResponse> createRepeated() => $pb.PbList<QueryGetInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetInvoiceResponse>(create);
  static QueryGetInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $2.Invoice get invoice => $_getN(0);
  @$pb.TagNumber(1)
  set invoice($2.Invoice v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasInvoice() => $_has(0);
  @$pb.TagNumber(1)
  void clearInvoice() => clearField(1);
  @$pb.TagNumber(1)
  $2.Invoice ensureInvoice() => $_ensure(0);
}

class QueryAllInvoiceRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllInvoiceRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllInvoiceRequest._() : super();
  factory QueryAllInvoiceRequest() => create();
  factory QueryAllInvoiceRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllInvoiceRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllInvoiceRequest clone() => QueryAllInvoiceRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllInvoiceRequest copyWith(void Function(QueryAllInvoiceRequest) updates) => super.copyWith((message) => updates(message as QueryAllInvoiceRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllInvoiceRequest create() => QueryAllInvoiceRequest._();
  QueryAllInvoiceRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllInvoiceRequest> createRepeated() => $pb.PbList<QueryAllInvoiceRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllInvoiceRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllInvoiceRequest>(create);
  static QueryAllInvoiceRequest _defaultInstance;

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

class QueryAllInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..pc<$2.Invoice>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Invoice', $pb.PbFieldType.PM, protoName: 'Invoice', subBuilder: $2.Invoice.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllInvoiceResponse._() : super();
  factory QueryAllInvoiceResponse() => create();
  factory QueryAllInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllInvoiceResponse clone() => QueryAllInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllInvoiceResponse copyWith(void Function(QueryAllInvoiceResponse) updates) => super.copyWith((message) => updates(message as QueryAllInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllInvoiceResponse create() => QueryAllInvoiceResponse._();
  QueryAllInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllInvoiceResponse> createRepeated() => $pb.PbList<QueryAllInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllInvoiceResponse>(create);
  static QueryAllInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$2.Invoice> get invoice => $_getList(0);

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

class QueryGetSentInvoiceRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetSentInvoiceRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetSentInvoiceRequest._() : super();
  factory QueryGetSentInvoiceRequest() => create();
  factory QueryGetSentInvoiceRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetSentInvoiceRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetSentInvoiceRequest clone() => QueryGetSentInvoiceRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetSentInvoiceRequest copyWith(void Function(QueryGetSentInvoiceRequest) updates) => super.copyWith((message) => updates(message as QueryGetSentInvoiceRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetSentInvoiceRequest create() => QueryGetSentInvoiceRequest._();
  QueryGetSentInvoiceRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetSentInvoiceRequest> createRepeated() => $pb.PbList<QueryGetSentInvoiceRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetSentInvoiceRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetSentInvoiceRequest>(create);
  static QueryGetSentInvoiceRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetSentInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetSentInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOM<$3.SentInvoice>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SentInvoice', protoName: 'SentInvoice', subBuilder: $3.SentInvoice.create)
    ..hasRequiredFields = false
  ;

  QueryGetSentInvoiceResponse._() : super();
  factory QueryGetSentInvoiceResponse() => create();
  factory QueryGetSentInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetSentInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetSentInvoiceResponse clone() => QueryGetSentInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetSentInvoiceResponse copyWith(void Function(QueryGetSentInvoiceResponse) updates) => super.copyWith((message) => updates(message as QueryGetSentInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetSentInvoiceResponse create() => QueryGetSentInvoiceResponse._();
  QueryGetSentInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetSentInvoiceResponse> createRepeated() => $pb.PbList<QueryGetSentInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetSentInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetSentInvoiceResponse>(create);
  static QueryGetSentInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $3.SentInvoice get sentInvoice => $_getN(0);
  @$pb.TagNumber(1)
  set sentInvoice($3.SentInvoice v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasSentInvoice() => $_has(0);
  @$pb.TagNumber(1)
  void clearSentInvoice() => clearField(1);
  @$pb.TagNumber(1)
  $3.SentInvoice ensureSentInvoice() => $_ensure(0);
}

class QueryAllSentInvoiceRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllSentInvoiceRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllSentInvoiceRequest._() : super();
  factory QueryAllSentInvoiceRequest() => create();
  factory QueryAllSentInvoiceRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllSentInvoiceRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllSentInvoiceRequest clone() => QueryAllSentInvoiceRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllSentInvoiceRequest copyWith(void Function(QueryAllSentInvoiceRequest) updates) => super.copyWith((message) => updates(message as QueryAllSentInvoiceRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllSentInvoiceRequest create() => QueryAllSentInvoiceRequest._();
  QueryAllSentInvoiceRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllSentInvoiceRequest> createRepeated() => $pb.PbList<QueryAllSentInvoiceRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllSentInvoiceRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllSentInvoiceRequest>(create);
  static QueryAllSentInvoiceRequest _defaultInstance;

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

class QueryAllSentInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllSentInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..pc<$3.SentInvoice>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'SentInvoice', $pb.PbFieldType.PM, protoName: 'SentInvoice', subBuilder: $3.SentInvoice.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllSentInvoiceResponse._() : super();
  factory QueryAllSentInvoiceResponse() => create();
  factory QueryAllSentInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllSentInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllSentInvoiceResponse clone() => QueryAllSentInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllSentInvoiceResponse copyWith(void Function(QueryAllSentInvoiceResponse) updates) => super.copyWith((message) => updates(message as QueryAllSentInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllSentInvoiceResponse create() => QueryAllSentInvoiceResponse._();
  QueryAllSentInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllSentInvoiceResponse> createRepeated() => $pb.PbList<QueryAllSentInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllSentInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllSentInvoiceResponse>(create);
  static QueryAllSentInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.SentInvoice> get sentInvoice => $_getList(0);

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

class QueryGetTimedoutInvoiceRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetTimedoutInvoiceRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  QueryGetTimedoutInvoiceRequest._() : super();
  factory QueryGetTimedoutInvoiceRequest() => create();
  factory QueryGetTimedoutInvoiceRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTimedoutInvoiceRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutInvoiceRequest clone() => QueryGetTimedoutInvoiceRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutInvoiceRequest copyWith(void Function(QueryGetTimedoutInvoiceRequest) updates) => super.copyWith((message) => updates(message as QueryGetTimedoutInvoiceRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutInvoiceRequest create() => QueryGetTimedoutInvoiceRequest._();
  QueryGetTimedoutInvoiceRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetTimedoutInvoiceRequest> createRepeated() => $pb.PbList<QueryGetTimedoutInvoiceRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutInvoiceRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTimedoutInvoiceRequest>(create);
  static QueryGetTimedoutInvoiceRequest _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetTimedoutInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryGetTimedoutInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOM<$4.TimedoutInvoice>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TimedoutInvoice', protoName: 'TimedoutInvoice', subBuilder: $4.TimedoutInvoice.create)
    ..hasRequiredFields = false
  ;

  QueryGetTimedoutInvoiceResponse._() : super();
  factory QueryGetTimedoutInvoiceResponse() => create();
  factory QueryGetTimedoutInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTimedoutInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutInvoiceResponse clone() => QueryGetTimedoutInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTimedoutInvoiceResponse copyWith(void Function(QueryGetTimedoutInvoiceResponse) updates) => super.copyWith((message) => updates(message as QueryGetTimedoutInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutInvoiceResponse create() => QueryGetTimedoutInvoiceResponse._();
  QueryGetTimedoutInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetTimedoutInvoiceResponse> createRepeated() => $pb.PbList<QueryGetTimedoutInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTimedoutInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTimedoutInvoiceResponse>(create);
  static QueryGetTimedoutInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $4.TimedoutInvoice get timedoutInvoice => $_getN(0);
  @$pb.TagNumber(1)
  set timedoutInvoice($4.TimedoutInvoice v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTimedoutInvoice() => $_has(0);
  @$pb.TagNumber(1)
  void clearTimedoutInvoice() => clearField(1);
  @$pb.TagNumber(1)
  $4.TimedoutInvoice ensureTimedoutInvoice() => $_ensure(0);
}

class QueryAllTimedoutInvoiceRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllTimedoutInvoiceRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOM<$6.PageRequest>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageRequest.create)
    ..hasRequiredFields = false
  ;

  QueryAllTimedoutInvoiceRequest._() : super();
  factory QueryAllTimedoutInvoiceRequest() => create();
  factory QueryAllTimedoutInvoiceRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllTimedoutInvoiceRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutInvoiceRequest clone() => QueryAllTimedoutInvoiceRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutInvoiceRequest copyWith(void Function(QueryAllTimedoutInvoiceRequest) updates) => super.copyWith((message) => updates(message as QueryAllTimedoutInvoiceRequest)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutInvoiceRequest create() => QueryAllTimedoutInvoiceRequest._();
  QueryAllTimedoutInvoiceRequest createEmptyInstance() => create();
  static $pb.PbList<QueryAllTimedoutInvoiceRequest> createRepeated() => $pb.PbList<QueryAllTimedoutInvoiceRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutInvoiceRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllTimedoutInvoiceRequest>(create);
  static QueryAllTimedoutInvoiceRequest _defaultInstance;

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

class QueryAllTimedoutInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'QueryAllTimedoutInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..pc<$4.TimedoutInvoice>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'TimedoutInvoice', $pb.PbFieldType.PM, protoName: 'TimedoutInvoice', subBuilder: $4.TimedoutInvoice.create)
    ..aOM<$6.PageResponse>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'pagination', subBuilder: $6.PageResponse.create)
    ..hasRequiredFields = false
  ;

  QueryAllTimedoutInvoiceResponse._() : super();
  factory QueryAllTimedoutInvoiceResponse() => create();
  factory QueryAllTimedoutInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryAllTimedoutInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutInvoiceResponse clone() => QueryAllTimedoutInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryAllTimedoutInvoiceResponse copyWith(void Function(QueryAllTimedoutInvoiceResponse) updates) => super.copyWith((message) => updates(message as QueryAllTimedoutInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutInvoiceResponse create() => QueryAllTimedoutInvoiceResponse._();
  QueryAllTimedoutInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<QueryAllTimedoutInvoiceResponse> createRepeated() => $pb.PbList<QueryAllTimedoutInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryAllTimedoutInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryAllTimedoutInvoiceResponse>(create);
  static QueryAllTimedoutInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$4.TimedoutInvoice> get timedoutInvoice => $_getList(0);

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

