///
//  Generated code. Do not modify.
//  source: invoice/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class MsgFactorInvoice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgFactorInvoice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgFactorInvoice._() : super();
  factory MsgFactorInvoice() => create();
  factory MsgFactorInvoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgFactorInvoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgFactorInvoice clone() => MsgFactorInvoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgFactorInvoice copyWith(void Function(MsgFactorInvoice) updates) => super.copyWith((message) => updates(message as MsgFactorInvoice)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgFactorInvoice create() => MsgFactorInvoice._();
  MsgFactorInvoice createEmptyInstance() => create();
  static $pb.PbList<MsgFactorInvoice> createRepeated() => $pb.PbList<MsgFactorInvoice>();
  @$core.pragma('dart2js:noInline')
  static MsgFactorInvoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgFactorInvoice>(create);
  static MsgFactorInvoice _defaultInstance;

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

class MsgFactorInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgFactorInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgFactorInvoiceResponse._() : super();
  factory MsgFactorInvoiceResponse() => create();
  factory MsgFactorInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgFactorInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgFactorInvoiceResponse clone() => MsgFactorInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgFactorInvoiceResponse copyWith(void Function(MsgFactorInvoiceResponse) updates) => super.copyWith((message) => updates(message as MsgFactorInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgFactorInvoiceResponse create() => MsgFactorInvoiceResponse._();
  MsgFactorInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<MsgFactorInvoiceResponse> createRepeated() => $pb.PbList<MsgFactorInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgFactorInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgFactorInvoiceResponse>(create);
  static MsgFactorInvoiceResponse _defaultInstance;
}

class MsgCreateSentInvoice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateSentInvoice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgCreateSentInvoice._() : super();
  factory MsgCreateSentInvoice() => create();
  factory MsgCreateSentInvoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateSentInvoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateSentInvoice clone() => MsgCreateSentInvoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateSentInvoice copyWith(void Function(MsgCreateSentInvoice) updates) => super.copyWith((message) => updates(message as MsgCreateSentInvoice)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentInvoice create() => MsgCreateSentInvoice._();
  MsgCreateSentInvoice createEmptyInstance() => create();
  static $pb.PbList<MsgCreateSentInvoice> createRepeated() => $pb.PbList<MsgCreateSentInvoice>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentInvoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateSentInvoice>(create);
  static MsgCreateSentInvoice _defaultInstance;

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

class MsgCreateSentInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateSentInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCreateSentInvoiceResponse._() : super();
  factory MsgCreateSentInvoiceResponse() => create();
  factory MsgCreateSentInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateSentInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateSentInvoiceResponse clone() => MsgCreateSentInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateSentInvoiceResponse copyWith(void Function(MsgCreateSentInvoiceResponse) updates) => super.copyWith((message) => updates(message as MsgCreateSentInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentInvoiceResponse create() => MsgCreateSentInvoiceResponse._();
  MsgCreateSentInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCreateSentInvoiceResponse> createRepeated() => $pb.PbList<MsgCreateSentInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateSentInvoiceResponse>(create);
  static MsgCreateSentInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class MsgUpdateSentInvoice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateSentInvoice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgUpdateSentInvoice._() : super();
  factory MsgUpdateSentInvoice() => create();
  factory MsgUpdateSentInvoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateSentInvoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateSentInvoice clone() => MsgUpdateSentInvoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateSentInvoice copyWith(void Function(MsgUpdateSentInvoice) updates) => super.copyWith((message) => updates(message as MsgUpdateSentInvoice)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentInvoice create() => MsgUpdateSentInvoice._();
  MsgUpdateSentInvoice createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateSentInvoice> createRepeated() => $pb.PbList<MsgUpdateSentInvoice>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentInvoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateSentInvoice>(create);
  static MsgUpdateSentInvoice _defaultInstance;

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

class MsgUpdateSentInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateSentInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgUpdateSentInvoiceResponse._() : super();
  factory MsgUpdateSentInvoiceResponse() => create();
  factory MsgUpdateSentInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateSentInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateSentInvoiceResponse clone() => MsgUpdateSentInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateSentInvoiceResponse copyWith(void Function(MsgUpdateSentInvoiceResponse) updates) => super.copyWith((message) => updates(message as MsgUpdateSentInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentInvoiceResponse create() => MsgUpdateSentInvoiceResponse._();
  MsgUpdateSentInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateSentInvoiceResponse> createRepeated() => $pb.PbList<MsgUpdateSentInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateSentInvoiceResponse>(create);
  static MsgUpdateSentInvoiceResponse _defaultInstance;
}

class MsgDeleteSentInvoice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteSentInvoice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgDeleteSentInvoice._() : super();
  factory MsgDeleteSentInvoice() => create();
  factory MsgDeleteSentInvoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteSentInvoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteSentInvoice clone() => MsgDeleteSentInvoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteSentInvoice copyWith(void Function(MsgDeleteSentInvoice) updates) => super.copyWith((message) => updates(message as MsgDeleteSentInvoice)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentInvoice create() => MsgDeleteSentInvoice._();
  MsgDeleteSentInvoice createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteSentInvoice> createRepeated() => $pb.PbList<MsgDeleteSentInvoice>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentInvoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteSentInvoice>(create);
  static MsgDeleteSentInvoice _defaultInstance;

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

class MsgDeleteSentInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteSentInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgDeleteSentInvoiceResponse._() : super();
  factory MsgDeleteSentInvoiceResponse() => create();
  factory MsgDeleteSentInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteSentInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteSentInvoiceResponse clone() => MsgDeleteSentInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteSentInvoiceResponse copyWith(void Function(MsgDeleteSentInvoiceResponse) updates) => super.copyWith((message) => updates(message as MsgDeleteSentInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentInvoiceResponse create() => MsgDeleteSentInvoiceResponse._();
  MsgDeleteSentInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteSentInvoiceResponse> createRepeated() => $pb.PbList<MsgDeleteSentInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteSentInvoiceResponse>(create);
  static MsgDeleteSentInvoiceResponse _defaultInstance;
}

class MsgCreateTimedoutInvoice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateTimedoutInvoice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgCreateTimedoutInvoice._() : super();
  factory MsgCreateTimedoutInvoice() => create();
  factory MsgCreateTimedoutInvoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateTimedoutInvoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutInvoice clone() => MsgCreateTimedoutInvoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutInvoice copyWith(void Function(MsgCreateTimedoutInvoice) updates) => super.copyWith((message) => updates(message as MsgCreateTimedoutInvoice)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutInvoice create() => MsgCreateTimedoutInvoice._();
  MsgCreateTimedoutInvoice createEmptyInstance() => create();
  static $pb.PbList<MsgCreateTimedoutInvoice> createRepeated() => $pb.PbList<MsgCreateTimedoutInvoice>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutInvoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateTimedoutInvoice>(create);
  static MsgCreateTimedoutInvoice _defaultInstance;

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

class MsgCreateTimedoutInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateTimedoutInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCreateTimedoutInvoiceResponse._() : super();
  factory MsgCreateTimedoutInvoiceResponse() => create();
  factory MsgCreateTimedoutInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateTimedoutInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutInvoiceResponse clone() => MsgCreateTimedoutInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutInvoiceResponse copyWith(void Function(MsgCreateTimedoutInvoiceResponse) updates) => super.copyWith((message) => updates(message as MsgCreateTimedoutInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutInvoiceResponse create() => MsgCreateTimedoutInvoiceResponse._();
  MsgCreateTimedoutInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCreateTimedoutInvoiceResponse> createRepeated() => $pb.PbList<MsgCreateTimedoutInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateTimedoutInvoiceResponse>(create);
  static MsgCreateTimedoutInvoiceResponse _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class MsgUpdateTimedoutInvoice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateTimedoutInvoice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgUpdateTimedoutInvoice._() : super();
  factory MsgUpdateTimedoutInvoice() => create();
  factory MsgUpdateTimedoutInvoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateTimedoutInvoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutInvoice clone() => MsgUpdateTimedoutInvoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutInvoice copyWith(void Function(MsgUpdateTimedoutInvoice) updates) => super.copyWith((message) => updates(message as MsgUpdateTimedoutInvoice)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutInvoice create() => MsgUpdateTimedoutInvoice._();
  MsgUpdateTimedoutInvoice createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateTimedoutInvoice> createRepeated() => $pb.PbList<MsgUpdateTimedoutInvoice>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutInvoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateTimedoutInvoice>(create);
  static MsgUpdateTimedoutInvoice _defaultInstance;

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

class MsgUpdateTimedoutInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateTimedoutInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgUpdateTimedoutInvoiceResponse._() : super();
  factory MsgUpdateTimedoutInvoiceResponse() => create();
  factory MsgUpdateTimedoutInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateTimedoutInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutInvoiceResponse clone() => MsgUpdateTimedoutInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutInvoiceResponse copyWith(void Function(MsgUpdateTimedoutInvoiceResponse) updates) => super.copyWith((message) => updates(message as MsgUpdateTimedoutInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutInvoiceResponse create() => MsgUpdateTimedoutInvoiceResponse._();
  MsgUpdateTimedoutInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateTimedoutInvoiceResponse> createRepeated() => $pb.PbList<MsgUpdateTimedoutInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateTimedoutInvoiceResponse>(create);
  static MsgUpdateTimedoutInvoiceResponse _defaultInstance;
}

class MsgDeleteTimedoutInvoice extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteTimedoutInvoice', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgDeleteTimedoutInvoice._() : super();
  factory MsgDeleteTimedoutInvoice() => create();
  factory MsgDeleteTimedoutInvoice.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteTimedoutInvoice.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutInvoice clone() => MsgDeleteTimedoutInvoice()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutInvoice copyWith(void Function(MsgDeleteTimedoutInvoice) updates) => super.copyWith((message) => updates(message as MsgDeleteTimedoutInvoice)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutInvoice create() => MsgDeleteTimedoutInvoice._();
  MsgDeleteTimedoutInvoice createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteTimedoutInvoice> createRepeated() => $pb.PbList<MsgDeleteTimedoutInvoice>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutInvoice getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteTimedoutInvoice>(create);
  static MsgDeleteTimedoutInvoice _defaultInstance;

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

class MsgDeleteTimedoutInvoiceResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteTimedoutInvoiceResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.invoice'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgDeleteTimedoutInvoiceResponse._() : super();
  factory MsgDeleteTimedoutInvoiceResponse() => create();
  factory MsgDeleteTimedoutInvoiceResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteTimedoutInvoiceResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutInvoiceResponse clone() => MsgDeleteTimedoutInvoiceResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutInvoiceResponse copyWith(void Function(MsgDeleteTimedoutInvoiceResponse) updates) => super.copyWith((message) => updates(message as MsgDeleteTimedoutInvoiceResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutInvoiceResponse create() => MsgDeleteTimedoutInvoiceResponse._();
  MsgDeleteTimedoutInvoiceResponse createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteTimedoutInvoiceResponse> createRepeated() => $pb.PbList<MsgDeleteTimedoutInvoiceResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutInvoiceResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteTimedoutInvoiceResponse>(create);
  static MsgDeleteTimedoutInvoiceResponse _defaultInstance;
}

