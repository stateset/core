///
//  Generated code. Do not modify.
//  source: agreement/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class MsgActivateAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgActivateAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgActivateAgreement._() : super();
  factory MsgActivateAgreement() => create();
  factory MsgActivateAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgActivateAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgActivateAgreement clone() => MsgActivateAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgActivateAgreement copyWith(void Function(MsgActivateAgreement) updates) => super.copyWith((message) => updates(message as MsgActivateAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgActivateAgreement create() => MsgActivateAgreement._();
  MsgActivateAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgActivateAgreement> createRepeated() => $pb.PbList<MsgActivateAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgActivateAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgActivateAgreement>(create);
  static MsgActivateAgreement _defaultInstance;

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

class MsgActivateAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgActivateAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgActivateAgreementResponse._() : super();
  factory MsgActivateAgreementResponse() => create();
  factory MsgActivateAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgActivateAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgActivateAgreementResponse clone() => MsgActivateAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgActivateAgreementResponse copyWith(void Function(MsgActivateAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgActivateAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgActivateAgreementResponse create() => MsgActivateAgreementResponse._();
  MsgActivateAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgActivateAgreementResponse> createRepeated() => $pb.PbList<MsgActivateAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgActivateAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgActivateAgreementResponse>(create);
  static MsgActivateAgreementResponse _defaultInstance;
}

class MsgExpireAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgExpireAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgExpireAgreement._() : super();
  factory MsgExpireAgreement() => create();
  factory MsgExpireAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgExpireAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgExpireAgreement clone() => MsgExpireAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgExpireAgreement copyWith(void Function(MsgExpireAgreement) updates) => super.copyWith((message) => updates(message as MsgExpireAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgExpireAgreement create() => MsgExpireAgreement._();
  MsgExpireAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgExpireAgreement> createRepeated() => $pb.PbList<MsgExpireAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgExpireAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgExpireAgreement>(create);
  static MsgExpireAgreement _defaultInstance;

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

class MsgExpireAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgExpireAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgExpireAgreementResponse._() : super();
  factory MsgExpireAgreementResponse() => create();
  factory MsgExpireAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgExpireAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgExpireAgreementResponse clone() => MsgExpireAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgExpireAgreementResponse copyWith(void Function(MsgExpireAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgExpireAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgExpireAgreementResponse create() => MsgExpireAgreementResponse._();
  MsgExpireAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgExpireAgreementResponse> createRepeated() => $pb.PbList<MsgExpireAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgExpireAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgExpireAgreementResponse>(create);
  static MsgExpireAgreementResponse _defaultInstance;
}

class MsgRenewAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRenewAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgRenewAgreement._() : super();
  factory MsgRenewAgreement() => create();
  factory MsgRenewAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRenewAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRenewAgreement clone() => MsgRenewAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRenewAgreement copyWith(void Function(MsgRenewAgreement) updates) => super.copyWith((message) => updates(message as MsgRenewAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRenewAgreement create() => MsgRenewAgreement._();
  MsgRenewAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgRenewAgreement> createRepeated() => $pb.PbList<MsgRenewAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgRenewAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRenewAgreement>(create);
  static MsgRenewAgreement _defaultInstance;

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

class MsgRenewAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRenewAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgRenewAgreementResponse._() : super();
  factory MsgRenewAgreementResponse() => create();
  factory MsgRenewAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRenewAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRenewAgreementResponse clone() => MsgRenewAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRenewAgreementResponse copyWith(void Function(MsgRenewAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgRenewAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRenewAgreementResponse create() => MsgRenewAgreementResponse._();
  MsgRenewAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgRenewAgreementResponse> createRepeated() => $pb.PbList<MsgRenewAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgRenewAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRenewAgreementResponse>(create);
  static MsgRenewAgreementResponse _defaultInstance;
}

class MsgTerminateAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgTerminateAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgTerminateAgreement._() : super();
  factory MsgTerminateAgreement() => create();
  factory MsgTerminateAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgTerminateAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgTerminateAgreement clone() => MsgTerminateAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgTerminateAgreement copyWith(void Function(MsgTerminateAgreement) updates) => super.copyWith((message) => updates(message as MsgTerminateAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgTerminateAgreement create() => MsgTerminateAgreement._();
  MsgTerminateAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgTerminateAgreement> createRepeated() => $pb.PbList<MsgTerminateAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgTerminateAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgTerminateAgreement>(create);
  static MsgTerminateAgreement _defaultInstance;

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

class MsgTerminateAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgTerminateAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgTerminateAgreementResponse._() : super();
  factory MsgTerminateAgreementResponse() => create();
  factory MsgTerminateAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgTerminateAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgTerminateAgreementResponse clone() => MsgTerminateAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgTerminateAgreementResponse copyWith(void Function(MsgTerminateAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgTerminateAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgTerminateAgreementResponse create() => MsgTerminateAgreementResponse._();
  MsgTerminateAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgTerminateAgreementResponse> createRepeated() => $pb.PbList<MsgTerminateAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgTerminateAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgTerminateAgreementResponse>(create);
  static MsgTerminateAgreementResponse _defaultInstance;
}

class MsgCreateSentAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateSentAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgCreateSentAgreement._() : super();
  factory MsgCreateSentAgreement() => create();
  factory MsgCreateSentAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateSentAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateSentAgreement clone() => MsgCreateSentAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateSentAgreement copyWith(void Function(MsgCreateSentAgreement) updates) => super.copyWith((message) => updates(message as MsgCreateSentAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentAgreement create() => MsgCreateSentAgreement._();
  MsgCreateSentAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgCreateSentAgreement> createRepeated() => $pb.PbList<MsgCreateSentAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateSentAgreement>(create);
  static MsgCreateSentAgreement _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get creator => $_getSZ(0);
  @$pb.TagNumber(1)
  set creator($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCreator() => $_has(0);
  @$pb.TagNumber(1)
  void clearCreator() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get did => $_getSZ(1);
  @$pb.TagNumber(2)
  set did($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDid() => $_has(1);
  @$pb.TagNumber(2)
  void clearDid() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get chain => $_getSZ(2);
  @$pb.TagNumber(3)
  set chain($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasChain() => $_has(2);
  @$pb.TagNumber(3)
  void clearChain() => clearField(3);
}

class MsgCreateSentAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateSentAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCreateSentAgreementResponse._() : super();
  factory MsgCreateSentAgreementResponse() => create();
  factory MsgCreateSentAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateSentAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateSentAgreementResponse clone() => MsgCreateSentAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateSentAgreementResponse copyWith(void Function(MsgCreateSentAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgCreateSentAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentAgreementResponse create() => MsgCreateSentAgreementResponse._();
  MsgCreateSentAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCreateSentAgreementResponse> createRepeated() => $pb.PbList<MsgCreateSentAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateSentAgreementResponse>(create);
  static MsgCreateSentAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class MsgUpdateSentAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateSentAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgUpdateSentAgreement._() : super();
  factory MsgUpdateSentAgreement() => create();
  factory MsgUpdateSentAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateSentAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateSentAgreement clone() => MsgUpdateSentAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateSentAgreement copyWith(void Function(MsgUpdateSentAgreement) updates) => super.copyWith((message) => updates(message as MsgUpdateSentAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentAgreement create() => MsgUpdateSentAgreement._();
  MsgUpdateSentAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateSentAgreement> createRepeated() => $pb.PbList<MsgUpdateSentAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateSentAgreement>(create);
  static MsgUpdateSentAgreement _defaultInstance;

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

  @$pb.TagNumber(3)
  $core.String get did => $_getSZ(2);
  @$pb.TagNumber(3)
  set did($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDid() => $_has(2);
  @$pb.TagNumber(3)
  void clearDid() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get chain => $_getSZ(3);
  @$pb.TagNumber(4)
  set chain($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasChain() => $_has(3);
  @$pb.TagNumber(4)
  void clearChain() => clearField(4);
}

class MsgUpdateSentAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateSentAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgUpdateSentAgreementResponse._() : super();
  factory MsgUpdateSentAgreementResponse() => create();
  factory MsgUpdateSentAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateSentAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateSentAgreementResponse clone() => MsgUpdateSentAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateSentAgreementResponse copyWith(void Function(MsgUpdateSentAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgUpdateSentAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentAgreementResponse create() => MsgUpdateSentAgreementResponse._();
  MsgUpdateSentAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateSentAgreementResponse> createRepeated() => $pb.PbList<MsgUpdateSentAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateSentAgreementResponse>(create);
  static MsgUpdateSentAgreementResponse _defaultInstance;
}

class MsgDeleteSentAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteSentAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgDeleteSentAgreement._() : super();
  factory MsgDeleteSentAgreement() => create();
  factory MsgDeleteSentAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteSentAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteSentAgreement clone() => MsgDeleteSentAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteSentAgreement copyWith(void Function(MsgDeleteSentAgreement) updates) => super.copyWith((message) => updates(message as MsgDeleteSentAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentAgreement create() => MsgDeleteSentAgreement._();
  MsgDeleteSentAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteSentAgreement> createRepeated() => $pb.PbList<MsgDeleteSentAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteSentAgreement>(create);
  static MsgDeleteSentAgreement _defaultInstance;

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

class MsgDeleteSentAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteSentAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgDeleteSentAgreementResponse._() : super();
  factory MsgDeleteSentAgreementResponse() => create();
  factory MsgDeleteSentAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteSentAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteSentAgreementResponse clone() => MsgDeleteSentAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteSentAgreementResponse copyWith(void Function(MsgDeleteSentAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgDeleteSentAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentAgreementResponse create() => MsgDeleteSentAgreementResponse._();
  MsgDeleteSentAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteSentAgreementResponse> createRepeated() => $pb.PbList<MsgDeleteSentAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteSentAgreementResponse>(create);
  static MsgDeleteSentAgreementResponse _defaultInstance;
}

class MsgCreateTimedoutAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateTimedoutAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgCreateTimedoutAgreement._() : super();
  factory MsgCreateTimedoutAgreement() => create();
  factory MsgCreateTimedoutAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateTimedoutAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutAgreement clone() => MsgCreateTimedoutAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutAgreement copyWith(void Function(MsgCreateTimedoutAgreement) updates) => super.copyWith((message) => updates(message as MsgCreateTimedoutAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutAgreement create() => MsgCreateTimedoutAgreement._();
  MsgCreateTimedoutAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgCreateTimedoutAgreement> createRepeated() => $pb.PbList<MsgCreateTimedoutAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateTimedoutAgreement>(create);
  static MsgCreateTimedoutAgreement _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get creator => $_getSZ(0);
  @$pb.TagNumber(1)
  set creator($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCreator() => $_has(0);
  @$pb.TagNumber(1)
  void clearCreator() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get did => $_getSZ(1);
  @$pb.TagNumber(2)
  set did($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDid() => $_has(1);
  @$pb.TagNumber(2)
  void clearDid() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get chain => $_getSZ(2);
  @$pb.TagNumber(3)
  set chain($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasChain() => $_has(2);
  @$pb.TagNumber(3)
  void clearChain() => clearField(3);
}

class MsgCreateTimedoutAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateTimedoutAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCreateTimedoutAgreementResponse._() : super();
  factory MsgCreateTimedoutAgreementResponse() => create();
  factory MsgCreateTimedoutAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateTimedoutAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutAgreementResponse clone() => MsgCreateTimedoutAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutAgreementResponse copyWith(void Function(MsgCreateTimedoutAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgCreateTimedoutAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutAgreementResponse create() => MsgCreateTimedoutAgreementResponse._();
  MsgCreateTimedoutAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCreateTimedoutAgreementResponse> createRepeated() => $pb.PbList<MsgCreateTimedoutAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateTimedoutAgreementResponse>(create);
  static MsgCreateTimedoutAgreementResponse _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class MsgUpdateTimedoutAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateTimedoutAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgUpdateTimedoutAgreement._() : super();
  factory MsgUpdateTimedoutAgreement() => create();
  factory MsgUpdateTimedoutAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateTimedoutAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutAgreement clone() => MsgUpdateTimedoutAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutAgreement copyWith(void Function(MsgUpdateTimedoutAgreement) updates) => super.copyWith((message) => updates(message as MsgUpdateTimedoutAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutAgreement create() => MsgUpdateTimedoutAgreement._();
  MsgUpdateTimedoutAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateTimedoutAgreement> createRepeated() => $pb.PbList<MsgUpdateTimedoutAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateTimedoutAgreement>(create);
  static MsgUpdateTimedoutAgreement _defaultInstance;

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

  @$pb.TagNumber(3)
  $core.String get did => $_getSZ(2);
  @$pb.TagNumber(3)
  set did($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasDid() => $_has(2);
  @$pb.TagNumber(3)
  void clearDid() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get chain => $_getSZ(3);
  @$pb.TagNumber(4)
  set chain($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasChain() => $_has(3);
  @$pb.TagNumber(4)
  void clearChain() => clearField(4);
}

class MsgUpdateTimedoutAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateTimedoutAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgUpdateTimedoutAgreementResponse._() : super();
  factory MsgUpdateTimedoutAgreementResponse() => create();
  factory MsgUpdateTimedoutAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateTimedoutAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutAgreementResponse clone() => MsgUpdateTimedoutAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutAgreementResponse copyWith(void Function(MsgUpdateTimedoutAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgUpdateTimedoutAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutAgreementResponse create() => MsgUpdateTimedoutAgreementResponse._();
  MsgUpdateTimedoutAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateTimedoutAgreementResponse> createRepeated() => $pb.PbList<MsgUpdateTimedoutAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateTimedoutAgreementResponse>(create);
  static MsgUpdateTimedoutAgreementResponse _defaultInstance;
}

class MsgDeleteTimedoutAgreement extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteTimedoutAgreement', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgDeleteTimedoutAgreement._() : super();
  factory MsgDeleteTimedoutAgreement() => create();
  factory MsgDeleteTimedoutAgreement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteTimedoutAgreement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutAgreement clone() => MsgDeleteTimedoutAgreement()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutAgreement copyWith(void Function(MsgDeleteTimedoutAgreement) updates) => super.copyWith((message) => updates(message as MsgDeleteTimedoutAgreement)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutAgreement create() => MsgDeleteTimedoutAgreement._();
  MsgDeleteTimedoutAgreement createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteTimedoutAgreement> createRepeated() => $pb.PbList<MsgDeleteTimedoutAgreement>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutAgreement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteTimedoutAgreement>(create);
  static MsgDeleteTimedoutAgreement _defaultInstance;

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

class MsgDeleteTimedoutAgreementResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteTimedoutAgreementResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.agreement'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgDeleteTimedoutAgreementResponse._() : super();
  factory MsgDeleteTimedoutAgreementResponse() => create();
  factory MsgDeleteTimedoutAgreementResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteTimedoutAgreementResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutAgreementResponse clone() => MsgDeleteTimedoutAgreementResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutAgreementResponse copyWith(void Function(MsgDeleteTimedoutAgreementResponse) updates) => super.copyWith((message) => updates(message as MsgDeleteTimedoutAgreementResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutAgreementResponse create() => MsgDeleteTimedoutAgreementResponse._();
  MsgDeleteTimedoutAgreementResponse createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteTimedoutAgreementResponse> createRepeated() => $pb.PbList<MsgDeleteTimedoutAgreementResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutAgreementResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteTimedoutAgreementResponse>(create);
  static MsgDeleteTimedoutAgreementResponse _defaultInstance;
}

