///
//  Generated code. Do not modify.
//  source: purchaseorder/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class MsgFinancePurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgFinancePurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgFinancePurchaseorder._() : super();
  factory MsgFinancePurchaseorder() => create();
  factory MsgFinancePurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgFinancePurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgFinancePurchaseorder clone() => MsgFinancePurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgFinancePurchaseorder copyWith(void Function(MsgFinancePurchaseorder) updates) => super.copyWith((message) => updates(message as MsgFinancePurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgFinancePurchaseorder create() => MsgFinancePurchaseorder._();
  MsgFinancePurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgFinancePurchaseorder> createRepeated() => $pb.PbList<MsgFinancePurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgFinancePurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgFinancePurchaseorder>(create);
  static MsgFinancePurchaseorder _defaultInstance;

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

class MsgFinancePurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgFinancePurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgFinancePurchaseorderResponse._() : super();
  factory MsgFinancePurchaseorderResponse() => create();
  factory MsgFinancePurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgFinancePurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgFinancePurchaseorderResponse clone() => MsgFinancePurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgFinancePurchaseorderResponse copyWith(void Function(MsgFinancePurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgFinancePurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgFinancePurchaseorderResponse create() => MsgFinancePurchaseorderResponse._();
  MsgFinancePurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgFinancePurchaseorderResponse> createRepeated() => $pb.PbList<MsgFinancePurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgFinancePurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgFinancePurchaseorderResponse>(create);
  static MsgFinancePurchaseorderResponse _defaultInstance;
}

class MsgCancelPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCancelPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCancelPurchaseorder._() : super();
  factory MsgCancelPurchaseorder() => create();
  factory MsgCancelPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCancelPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCancelPurchaseorder clone() => MsgCancelPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCancelPurchaseorder copyWith(void Function(MsgCancelPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgCancelPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCancelPurchaseorder create() => MsgCancelPurchaseorder._();
  MsgCancelPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgCancelPurchaseorder> createRepeated() => $pb.PbList<MsgCancelPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgCancelPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCancelPurchaseorder>(create);
  static MsgCancelPurchaseorder _defaultInstance;

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

class MsgCancelPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCancelPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgCancelPurchaseorderResponse._() : super();
  factory MsgCancelPurchaseorderResponse() => create();
  factory MsgCancelPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCancelPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCancelPurchaseorderResponse clone() => MsgCancelPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCancelPurchaseorderResponse copyWith(void Function(MsgCancelPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgCancelPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCancelPurchaseorderResponse create() => MsgCancelPurchaseorderResponse._();
  MsgCancelPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCancelPurchaseorderResponse> createRepeated() => $pb.PbList<MsgCancelPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCancelPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCancelPurchaseorderResponse>(create);
  static MsgCancelPurchaseorderResponse _defaultInstance;
}

class MsgCompletePurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCompletePurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCompletePurchaseorder._() : super();
  factory MsgCompletePurchaseorder() => create();
  factory MsgCompletePurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCompletePurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCompletePurchaseorder clone() => MsgCompletePurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCompletePurchaseorder copyWith(void Function(MsgCompletePurchaseorder) updates) => super.copyWith((message) => updates(message as MsgCompletePurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCompletePurchaseorder create() => MsgCompletePurchaseorder._();
  MsgCompletePurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgCompletePurchaseorder> createRepeated() => $pb.PbList<MsgCompletePurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgCompletePurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCompletePurchaseorder>(create);
  static MsgCompletePurchaseorder _defaultInstance;

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

class MsgCompletePurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCompletePurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgCompletePurchaseorderResponse._() : super();
  factory MsgCompletePurchaseorderResponse() => create();
  factory MsgCompletePurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCompletePurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCompletePurchaseorderResponse clone() => MsgCompletePurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCompletePurchaseorderResponse copyWith(void Function(MsgCompletePurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgCompletePurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCompletePurchaseorderResponse create() => MsgCompletePurchaseorderResponse._();
  MsgCompletePurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCompletePurchaseorderResponse> createRepeated() => $pb.PbList<MsgCompletePurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCompletePurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCompletePurchaseorderResponse>(create);
  static MsgCompletePurchaseorderResponse _defaultInstance;
}

class MsgCreateSentPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateSentPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgCreateSentPurchaseorder._() : super();
  factory MsgCreateSentPurchaseorder() => create();
  factory MsgCreateSentPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateSentPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateSentPurchaseorder clone() => MsgCreateSentPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateSentPurchaseorder copyWith(void Function(MsgCreateSentPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgCreateSentPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentPurchaseorder create() => MsgCreateSentPurchaseorder._();
  MsgCreateSentPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgCreateSentPurchaseorder> createRepeated() => $pb.PbList<MsgCreateSentPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateSentPurchaseorder>(create);
  static MsgCreateSentPurchaseorder _defaultInstance;

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

class MsgCreateSentPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateSentPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCreateSentPurchaseorderResponse._() : super();
  factory MsgCreateSentPurchaseorderResponse() => create();
  factory MsgCreateSentPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateSentPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateSentPurchaseorderResponse clone() => MsgCreateSentPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateSentPurchaseorderResponse copyWith(void Function(MsgCreateSentPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgCreateSentPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentPurchaseorderResponse create() => MsgCreateSentPurchaseorderResponse._();
  MsgCreateSentPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCreateSentPurchaseorderResponse> createRepeated() => $pb.PbList<MsgCreateSentPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateSentPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateSentPurchaseorderResponse>(create);
  static MsgCreateSentPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class MsgUpdateSentPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateSentPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgUpdateSentPurchaseorder._() : super();
  factory MsgUpdateSentPurchaseorder() => create();
  factory MsgUpdateSentPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateSentPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateSentPurchaseorder clone() => MsgUpdateSentPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateSentPurchaseorder copyWith(void Function(MsgUpdateSentPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgUpdateSentPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentPurchaseorder create() => MsgUpdateSentPurchaseorder._();
  MsgUpdateSentPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateSentPurchaseorder> createRepeated() => $pb.PbList<MsgUpdateSentPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateSentPurchaseorder>(create);
  static MsgUpdateSentPurchaseorder _defaultInstance;

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

class MsgUpdateSentPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateSentPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgUpdateSentPurchaseorderResponse._() : super();
  factory MsgUpdateSentPurchaseorderResponse() => create();
  factory MsgUpdateSentPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateSentPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateSentPurchaseorderResponse clone() => MsgUpdateSentPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateSentPurchaseorderResponse copyWith(void Function(MsgUpdateSentPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgUpdateSentPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentPurchaseorderResponse create() => MsgUpdateSentPurchaseorderResponse._();
  MsgUpdateSentPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateSentPurchaseorderResponse> createRepeated() => $pb.PbList<MsgUpdateSentPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateSentPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateSentPurchaseorderResponse>(create);
  static MsgUpdateSentPurchaseorderResponse _defaultInstance;
}

class MsgDeleteSentPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteSentPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgDeleteSentPurchaseorder._() : super();
  factory MsgDeleteSentPurchaseorder() => create();
  factory MsgDeleteSentPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteSentPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteSentPurchaseorder clone() => MsgDeleteSentPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteSentPurchaseorder copyWith(void Function(MsgDeleteSentPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgDeleteSentPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentPurchaseorder create() => MsgDeleteSentPurchaseorder._();
  MsgDeleteSentPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteSentPurchaseorder> createRepeated() => $pb.PbList<MsgDeleteSentPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteSentPurchaseorder>(create);
  static MsgDeleteSentPurchaseorder _defaultInstance;

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

class MsgDeleteSentPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteSentPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgDeleteSentPurchaseorderResponse._() : super();
  factory MsgDeleteSentPurchaseorderResponse() => create();
  factory MsgDeleteSentPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteSentPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteSentPurchaseorderResponse clone() => MsgDeleteSentPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteSentPurchaseorderResponse copyWith(void Function(MsgDeleteSentPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgDeleteSentPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentPurchaseorderResponse create() => MsgDeleteSentPurchaseorderResponse._();
  MsgDeleteSentPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteSentPurchaseorderResponse> createRepeated() => $pb.PbList<MsgDeleteSentPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteSentPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteSentPurchaseorderResponse>(create);
  static MsgDeleteSentPurchaseorderResponse _defaultInstance;
}

class MsgCreateTimedoutPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateTimedoutPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgCreateTimedoutPurchaseorder._() : super();
  factory MsgCreateTimedoutPurchaseorder() => create();
  factory MsgCreateTimedoutPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateTimedoutPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutPurchaseorder clone() => MsgCreateTimedoutPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutPurchaseorder copyWith(void Function(MsgCreateTimedoutPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgCreateTimedoutPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutPurchaseorder create() => MsgCreateTimedoutPurchaseorder._();
  MsgCreateTimedoutPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgCreateTimedoutPurchaseorder> createRepeated() => $pb.PbList<MsgCreateTimedoutPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateTimedoutPurchaseorder>(create);
  static MsgCreateTimedoutPurchaseorder _defaultInstance;

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

class MsgCreateTimedoutPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgCreateTimedoutPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgCreateTimedoutPurchaseorderResponse._() : super();
  factory MsgCreateTimedoutPurchaseorderResponse() => create();
  factory MsgCreateTimedoutPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgCreateTimedoutPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutPurchaseorderResponse clone() => MsgCreateTimedoutPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgCreateTimedoutPurchaseorderResponse copyWith(void Function(MsgCreateTimedoutPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgCreateTimedoutPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutPurchaseorderResponse create() => MsgCreateTimedoutPurchaseorderResponse._();
  MsgCreateTimedoutPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgCreateTimedoutPurchaseorderResponse> createRepeated() => $pb.PbList<MsgCreateTimedoutPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgCreateTimedoutPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgCreateTimedoutPurchaseorderResponse>(create);
  static MsgCreateTimedoutPurchaseorderResponse _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class MsgUpdateTimedoutPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateTimedoutPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'chain')
    ..hasRequiredFields = false
  ;

  MsgUpdateTimedoutPurchaseorder._() : super();
  factory MsgUpdateTimedoutPurchaseorder() => create();
  factory MsgUpdateTimedoutPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateTimedoutPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutPurchaseorder clone() => MsgUpdateTimedoutPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutPurchaseorder copyWith(void Function(MsgUpdateTimedoutPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgUpdateTimedoutPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutPurchaseorder create() => MsgUpdateTimedoutPurchaseorder._();
  MsgUpdateTimedoutPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateTimedoutPurchaseorder> createRepeated() => $pb.PbList<MsgUpdateTimedoutPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateTimedoutPurchaseorder>(create);
  static MsgUpdateTimedoutPurchaseorder _defaultInstance;

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

class MsgUpdateTimedoutPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgUpdateTimedoutPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgUpdateTimedoutPurchaseorderResponse._() : super();
  factory MsgUpdateTimedoutPurchaseorderResponse() => create();
  factory MsgUpdateTimedoutPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgUpdateTimedoutPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutPurchaseorderResponse clone() => MsgUpdateTimedoutPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgUpdateTimedoutPurchaseorderResponse copyWith(void Function(MsgUpdateTimedoutPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgUpdateTimedoutPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutPurchaseorderResponse create() => MsgUpdateTimedoutPurchaseorderResponse._();
  MsgUpdateTimedoutPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgUpdateTimedoutPurchaseorderResponse> createRepeated() => $pb.PbList<MsgUpdateTimedoutPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgUpdateTimedoutPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgUpdateTimedoutPurchaseorderResponse>(create);
  static MsgUpdateTimedoutPurchaseorderResponse _defaultInstance;
}

class MsgDeleteTimedoutPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteTimedoutPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'id', $pb.PbFieldType.OU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  MsgDeleteTimedoutPurchaseorder._() : super();
  factory MsgDeleteTimedoutPurchaseorder() => create();
  factory MsgDeleteTimedoutPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteTimedoutPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutPurchaseorder clone() => MsgDeleteTimedoutPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutPurchaseorder copyWith(void Function(MsgDeleteTimedoutPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgDeleteTimedoutPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutPurchaseorder create() => MsgDeleteTimedoutPurchaseorder._();
  MsgDeleteTimedoutPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteTimedoutPurchaseorder> createRepeated() => $pb.PbList<MsgDeleteTimedoutPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteTimedoutPurchaseorder>(create);
  static MsgDeleteTimedoutPurchaseorder _defaultInstance;

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

class MsgDeleteTimedoutPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgDeleteTimedoutPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgDeleteTimedoutPurchaseorderResponse._() : super();
  factory MsgDeleteTimedoutPurchaseorderResponse() => create();
  factory MsgDeleteTimedoutPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgDeleteTimedoutPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutPurchaseorderResponse clone() => MsgDeleteTimedoutPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgDeleteTimedoutPurchaseorderResponse copyWith(void Function(MsgDeleteTimedoutPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgDeleteTimedoutPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutPurchaseorderResponse create() => MsgDeleteTimedoutPurchaseorderResponse._();
  MsgDeleteTimedoutPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgDeleteTimedoutPurchaseorderResponse> createRepeated() => $pb.PbList<MsgDeleteTimedoutPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgDeleteTimedoutPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgDeleteTimedoutPurchaseorderResponse>(create);
  static MsgDeleteTimedoutPurchaseorderResponse _defaultInstance;
}

class MsgRequestPurchaseorder extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRequestPurchaseorder', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'creator')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'did')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'uri')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'amount')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'state')
    ..hasRequiredFields = false
  ;

  MsgRequestPurchaseorder._() : super();
  factory MsgRequestPurchaseorder() => create();
  factory MsgRequestPurchaseorder.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRequestPurchaseorder.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRequestPurchaseorder clone() => MsgRequestPurchaseorder()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRequestPurchaseorder copyWith(void Function(MsgRequestPurchaseorder) updates) => super.copyWith((message) => updates(message as MsgRequestPurchaseorder)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRequestPurchaseorder create() => MsgRequestPurchaseorder._();
  MsgRequestPurchaseorder createEmptyInstance() => create();
  static $pb.PbList<MsgRequestPurchaseorder> createRepeated() => $pb.PbList<MsgRequestPurchaseorder>();
  @$core.pragma('dart2js:noInline')
  static MsgRequestPurchaseorder getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRequestPurchaseorder>(create);
  static MsgRequestPurchaseorder _defaultInstance;

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
  $core.String get uri => $_getSZ(2);
  @$pb.TagNumber(3)
  set uri($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasUri() => $_has(2);
  @$pb.TagNumber(3)
  void clearUri() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get amount => $_getSZ(3);
  @$pb.TagNumber(4)
  set amount($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasAmount() => $_has(3);
  @$pb.TagNumber(4)
  void clearAmount() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get state => $_getSZ(4);
  @$pb.TagNumber(5)
  set state($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasState() => $_has(4);
  @$pb.TagNumber(5)
  void clearState() => clearField(5);
}

class MsgRequestPurchaseorderResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'MsgRequestPurchaseorderResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  MsgRequestPurchaseorderResponse._() : super();
  factory MsgRequestPurchaseorderResponse() => create();
  factory MsgRequestPurchaseorderResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory MsgRequestPurchaseorderResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  MsgRequestPurchaseorderResponse clone() => MsgRequestPurchaseorderResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  MsgRequestPurchaseorderResponse copyWith(void Function(MsgRequestPurchaseorderResponse) updates) => super.copyWith((message) => updates(message as MsgRequestPurchaseorderResponse)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static MsgRequestPurchaseorderResponse create() => MsgRequestPurchaseorderResponse._();
  MsgRequestPurchaseorderResponse createEmptyInstance() => create();
  static $pb.PbList<MsgRequestPurchaseorderResponse> createRepeated() => $pb.PbList<MsgRequestPurchaseorderResponse>();
  @$core.pragma('dart2js:noInline')
  static MsgRequestPurchaseorderResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<MsgRequestPurchaseorderResponse>(create);
  static MsgRequestPurchaseorderResponse _defaultInstance;
}

