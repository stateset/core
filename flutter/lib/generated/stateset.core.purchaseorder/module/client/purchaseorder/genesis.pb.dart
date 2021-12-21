///
//  Generated code. Do not modify.
//  source: purchaseorder/genesis.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'purchaseorder.pb.dart' as $2;
import 'sent_purchaseorder.pb.dart' as $3;
import 'timedout_purchaseorder.pb.dart' as $4;

class GenesisState extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'GenesisState', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'stateset.core.purchaseorder'), createEmptyInstance: create)
    ..pc<$2.Purchaseorder>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'purchaseorderList', $pb.PbFieldType.PM, protoName: 'purchaseorderList', subBuilder: $2.Purchaseorder.create)
    ..a<$fixnum.Int64>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'purchaseorderCount', $pb.PbFieldType.OU6, protoName: 'purchaseorderCount', defaultOrMaker: $fixnum.Int64.ZERO)
    ..pc<$3.SentPurchaseorder>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sentPurchaseorderList', $pb.PbFieldType.PM, protoName: 'sentPurchaseorderList', subBuilder: $3.SentPurchaseorder.create)
    ..a<$fixnum.Int64>(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'sentPurchaseorderCount', $pb.PbFieldType.OU6, protoName: 'sentPurchaseorderCount', defaultOrMaker: $fixnum.Int64.ZERO)
    ..pc<$4.TimedoutPurchaseorder>(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timedoutPurchaseorderList', $pb.PbFieldType.PM, protoName: 'timedoutPurchaseorderList', subBuilder: $4.TimedoutPurchaseorder.create)
    ..a<$fixnum.Int64>(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'timedoutPurchaseorderCount', $pb.PbFieldType.OU6, protoName: 'timedoutPurchaseorderCount', defaultOrMaker: $fixnum.Int64.ZERO)
    ..hasRequiredFields = false
  ;

  GenesisState._() : super();
  factory GenesisState() => create();
  factory GenesisState.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GenesisState.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  GenesisState clone() => GenesisState()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  GenesisState copyWith(void Function(GenesisState) updates) => super.copyWith((message) => updates(message as GenesisState)); // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GenesisState create() => GenesisState._();
  GenesisState createEmptyInstance() => create();
  static $pb.PbList<GenesisState> createRepeated() => $pb.PbList<GenesisState>();
  @$core.pragma('dart2js:noInline')
  static GenesisState getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GenesisState>(create);
  static GenesisState _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$2.Purchaseorder> get purchaseorderList => $_getList(0);

  @$pb.TagNumber(2)
  $fixnum.Int64 get purchaseorderCount => $_getI64(1);
  @$pb.TagNumber(2)
  set purchaseorderCount($fixnum.Int64 v) { $_setInt64(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPurchaseorderCount() => $_has(1);
  @$pb.TagNumber(2)
  void clearPurchaseorderCount() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$3.SentPurchaseorder> get sentPurchaseorderList => $_getList(2);

  @$pb.TagNumber(4)
  $fixnum.Int64 get sentPurchaseorderCount => $_getI64(3);
  @$pb.TagNumber(4)
  set sentPurchaseorderCount($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasSentPurchaseorderCount() => $_has(3);
  @$pb.TagNumber(4)
  void clearSentPurchaseorderCount() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$4.TimedoutPurchaseorder> get timedoutPurchaseorderList => $_getList(4);

  @$pb.TagNumber(6)
  $fixnum.Int64 get timedoutPurchaseorderCount => $_getI64(5);
  @$pb.TagNumber(6)
  set timedoutPurchaseorderCount($fixnum.Int64 v) { $_setInt64(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasTimedoutPurchaseorderCount() => $_has(5);
  @$pb.TagNumber(6)
  void clearTimedoutPurchaseorderCount() => clearField(6);
}

