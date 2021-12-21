///
//  Generated code. Do not modify.
//  source: loan/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class MsgRequestLoan extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRequestLoan', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'amount')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'fee')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'collateral')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'deadline')
    ..hasRequiredFields = false
  ;

  MsgRequestLoan._() : super();
  factory MsgRequestLoan() => create();
  factory MsgRequestLoan.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRequestLoan.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRequestLoan clone() => MsgRequestLoan()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRequestLoan copyWith(void Function(MsgRequestLoan) updates) => super.copyWith((message) => updates(message as MsgRequestLoan)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRequestLoan create() => MsgRequestLoan._();
  MsgRequestLoan createEmptyInstance() => create();
  static $pb.PbList<MsgRequestLoan> createRepeated() => $pb.PbList<MsgRequestLoan>();
  @$core.pragma('dart2js:noInline')
  static MsgRequestLoan getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRequestLoan>(create);
  static MsgRequestLoan _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get creator => $_getSZ(0);
  @$pb.TagNumber(1)
  set creator($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCreator() => $_has(0);
  @$pb.TagNumber(1)
  void clearCreator() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get amount => $_getSZ(1);
  @$pb.TagNumber(2)
  set amount($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAmount() => $_has(1);
  @$pb.TagNumber(2)
  void clearAmount() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get fee => $_getSZ(2);
  @$pb.TagNumber(3)
  set fee($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasFee() => $_has(2);
  @$pb.TagNumber(3)
  void clearFee() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get collateral => $_getSZ(3);
  @$pb.TagNumber(4)
  set collateral($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasCollateral() => $_has(3);
  @$pb.TagNumber(4)
  void clearCollateral() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get deadline => $_getSZ(4);
  @$pb.TagNumber(5)
  set deadline($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasDeadline() => $_has(4);
  @$pb.TagNumber(5)
  void clearDeadline() => clearField(5);
}

class MsgRequestLoanResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRequestLoanResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgRequestLoanResponse._() : super();
  factory MsgRequestLoanResponse() => create();
  factory MsgRequestLoanResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRequestLoanResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRequestLoanResponse clone() => MsgRequestLoanResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRequestLoanResponse copyWith(void Function(MsgRequestLoanResponse) updates) => super.copyWith((message) => updates(message as MsgRequestLoanResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRequestLoanResponse create() => MsgRequestLoanResponse._();
  MsgRequestLoanResponse createEmptyInstance() => create();
  static $pb.PbList<MsgRequestLoanResponse> createRepeated() => $pb.PbList<MsgRequestLoanResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgRequestLoanResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRequestLoanResponse>(create);
  static MsgRequestLoanResponse _defaultInstance;
}

class MsgApproveLoan extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgApproveLoan', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgApproveLoan._() : super();
  factory MsgApproveLoan() => create();
  factory MsgApproveLoan.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgApproveLoan.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgApproveLoan clone() => MsgApproveLoan()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgApproveLoan copyWith(void Function(MsgApproveLoan) updates) => super.copyWith((message) => updates(message as MsgApproveLoan)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgApproveLoan create() => MsgApproveLoan._();
  MsgApproveLoan createEmptyInstance() => create();
  static $pb.PbList<MsgApproveLoan> createRepeated() => $pb.PbList<MsgApproveLoan>();
  @$core.pragma('dart2js:noInline')
  static MsgApproveLoan getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgApproveLoan>(create);
  static MsgApproveLoan _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get creator => $_getSZ(0);
  @$pb.TagNumber(1)
  set creator($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCreator() => $_has(0);
  @$pb.TagNumber(1)
  void clearCreator() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get id => $_getI64(1);
  @$pb.TagNumber(2)
  set id($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(1);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);
}

class MsgApproveLoanResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgApproveLoanResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgApproveLoanResponse._() : super();
  factory MsgApproveLoanResponse() => create();
  factory MsgApproveLoanResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgApproveLoanResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgApproveLoanResponse clone() => MsgApproveLoanResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgApproveLoanResponse copyWith(void Function(MsgApproveLoanResponse) updates) => super.copyWith((message) => updates(message as MsgApproveLoanResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgApproveLoanResponse create() => MsgApproveLoanResponse._();
  MsgApproveLoanResponse createEmptyInstance() => create();
  static $pb.PbList<MsgApproveLoanResponse> createRepeated() => $pb.PbList<MsgApproveLoanResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgApproveLoanResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgApproveLoanResponse>(create);
  static MsgApproveLoanResponse _defaultInstance;
}

class MsgRepayLoan extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRepayLoan', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgRepayLoan._() : super();
  factory MsgRepayLoan() => create();
  factory MsgRepayLoan.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRepayLoan.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRepayLoan clone() => MsgRepayLoan()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRepayLoan copyWith(void Function(MsgRepayLoan) updates) => super.copyWith((message) => updates(message as MsgRepayLoan)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRepayLoan create() => MsgRepayLoan._();
  MsgRepayLoan createEmptyInstance() => create();
  static $pb.PbList<MsgRepayLoan> createRepeated() => $pb.PbList<MsgRepayLoan>();
  @$core.pragma('dart2js:noInline')
  static MsgRepayLoan getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRepayLoan>(create);
  static MsgRepayLoan _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get creator => $_getSZ(0);
  @$pb.TagNumber(1)
  set creator($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCreator() => $_has(0);
  @$pb.TagNumber(1)
  void clearCreator() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get id => $_getI64(1);
  @$pb.TagNumber(2)
  set id($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(1);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);
}

class MsgRepayLoanResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRepayLoanResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgRepayLoanResponse._() : super();
  factory MsgRepayLoanResponse() => create();
  factory MsgRepayLoanResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRepayLoanResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRepayLoanResponse clone() => MsgRepayLoanResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRepayLoanResponse copyWith(void Function(MsgRepayLoanResponse) updates) => super.copyWith((message) => updates(message as MsgRepayLoanResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRepayLoanResponse create() => MsgRepayLoanResponse._();
  MsgRepayLoanResponse createEmptyInstance() => create();
  static $pb.PbList<MsgRepayLoanResponse> createRepeated() => $pb.PbList<MsgRepayLoanResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgRepayLoanResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRepayLoanResponse>(create);
  static MsgRepayLoanResponse _defaultInstance;
}

class MsgLiquidateLoan extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgLiquidateLoan', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgLiquidateLoan._() : super();
  factory MsgLiquidateLoan() => create();
  factory MsgLiquidateLoan.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgLiquidateLoan.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgLiquidateLoan clone() => MsgLiquidateLoan()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgLiquidateLoan copyWith(void Function(MsgLiquidateLoan) updates) => super.copyWith((message) => updates(message as MsgLiquidateLoan)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgLiquidateLoan create() => MsgLiquidateLoan._();
  MsgLiquidateLoan createEmptyInstance() => create();
  static $pb.PbList<MsgLiquidateLoan> createRepeated() => $pb.PbList<MsgLiquidateLoan>();
  @$core.pragma('dart2js:noInline')
  static MsgLiquidateLoan getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgLiquidateLoan>(create);
  static MsgLiquidateLoan _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get creator => $_getSZ(0);
  @$pb.TagNumber(1)
  set creator($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCreator() => $_has(0);
  @$pb.TagNumber(1)
  void clearCreator() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get id => $_getI64(1);
  @$pb.TagNumber(2)
  set id($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(1);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);
}

class MsgLiquidateLoanResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgLiquidateLoanResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgLiquidateLoanResponse._() : super();
  factory MsgLiquidateLoanResponse() => create();
  factory MsgLiquidateLoanResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgLiquidateLoanResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgLiquidateLoanResponse clone() => MsgLiquidateLoanResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgLiquidateLoanResponse copyWith(void Function(MsgLiquidateLoanResponse) updates) => super.copyWith((message) => updates(message as MsgLiquidateLoanResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgLiquidateLoanResponse create() => MsgLiquidateLoanResponse._();
  MsgLiquidateLoanResponse createEmptyInstance() => create();
  static $pb.PbList<MsgLiquidateLoanResponse> createRepeated() => $pb.PbList<MsgLiquidateLoanResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgLiquidateLoanResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgLiquidateLoanResponse>(create);
  static MsgLiquidateLoanResponse _defaultInstance;
}

class MsgCancelLoan extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCancelLoan', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCancelLoan._() : super();
  factory MsgCancelLoan() => create();
  factory MsgCancelLoan.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCancelLoan.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCancelLoan clone() => MsgCancelLoan()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCancelLoan copyWith(void Function(MsgCancelLoan) updates) => super.copyWith((message) => updates(message as MsgCancelLoan)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCancelLoan create() => MsgCancelLoan._();
  MsgCancelLoan createEmptyInstance() => create();
  static $pb.PbList<MsgCancelLoan> createRepeated() => $pb.PbList<MsgCancelLoan>();
  @$core.pragma('dart2js:noInline')
  static MsgCancelLoan getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCancelLoan>(create);
  static MsgCancelLoan _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get creator => $_getSZ(0);
  @$pb.TagNumber(1)
  set creator($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCreator() => $_has(0);
  @$pb.TagNumber(1)
  void clearCreator() => clearField(1);

  @$pb.TagNumber(2)
  $fixnum.Int64 get id => $_getI64(1);
  @$pb.TagNumber(2)
  set id($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasId() => $_has(1);
  @$pb.TagNumber(2)
  void clearId() => clearField(2);
}

class MsgCancelLoanResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCancelLoanResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.loan'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgCancelLoanResponse._() : super();
  factory MsgCancelLoanResponse() => create();
  factory MsgCancelLoanResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCancelLoanResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCancelLoanResponse clone() => MsgCancelLoanResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCancelLoanResponse copyWith(void Function(MsgCancelLoanResponse) updates) => super.copyWith((message) => updates(message as MsgCancelLoanResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCancelLoanResponse create() => MsgCancelLoanResponse._();
  MsgCancelLoanResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCancelLoanResponse> createRepeated() => $pb.PbList<MsgCancelLoanResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCancelLoanResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCancelLoanResponse>(create);
  static MsgCancelLoanResponse _defaultInstance;
}

