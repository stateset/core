///
//  Generated code. Do not modify.
//  source: purchaseorder/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'tx.pb.dart' as $1;
export 'tx.pb.dart';

class MsgClient extends $grpc.Client {
  static final _$financePurchaseorder = $grpc.ClientMethod<
          $1.MsgFinancePurchaseorder, $1.MsgFinancePurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/FinancePurchaseorder',
      ($1.MsgFinancePurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgFinancePurchaseorderResponse.fromBuffer(value));
  static final _$cancelPurchaseorder = $grpc.ClientMethod<
          $1.MsgCancelPurchaseorder, $1.MsgCancelPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/CancelPurchaseorder',
      ($1.MsgCancelPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCancelPurchaseorderResponse.fromBuffer(value));
  static final _$completePurchaseorder = $grpc.ClientMethod<
          $1.MsgCompletePurchaseorder, $1.MsgCompletePurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/CompletePurchaseorder',
      ($1.MsgCompletePurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCompletePurchaseorderResponse.fromBuffer(value));
  static final _$createSentPurchaseorder = $grpc.ClientMethod<
          $1.MsgCreateSentPurchaseorder, $1.MsgCreateSentPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/CreateSentPurchaseorder',
      ($1.MsgCreateSentPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCreateSentPurchaseorderResponse.fromBuffer(value));
  static final _$updateSentPurchaseorder = $grpc.ClientMethod<
          $1.MsgUpdateSentPurchaseorder, $1.MsgUpdateSentPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/UpdateSentPurchaseorder',
      ($1.MsgUpdateSentPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgUpdateSentPurchaseorderResponse.fromBuffer(value));
  static final _$deleteSentPurchaseorder = $grpc.ClientMethod<
          $1.MsgDeleteSentPurchaseorder, $1.MsgDeleteSentPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/DeleteSentPurchaseorder',
      ($1.MsgDeleteSentPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgDeleteSentPurchaseorderResponse.fromBuffer(value));
  static final _$createTimedoutPurchaseorder = $grpc.ClientMethod<
          $1.MsgCreateTimedoutPurchaseorder,
          $1.MsgCreateTimedoutPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/CreateTimedoutPurchaseorder',
      ($1.MsgCreateTimedoutPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCreateTimedoutPurchaseorderResponse.fromBuffer(value));
  static final _$updateTimedoutPurchaseorder = $grpc.ClientMethod<
          $1.MsgUpdateTimedoutPurchaseorder,
          $1.MsgUpdateTimedoutPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/UpdateTimedoutPurchaseorder',
      ($1.MsgUpdateTimedoutPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgUpdateTimedoutPurchaseorderResponse.fromBuffer(value));
  static final _$deleteTimedoutPurchaseorder = $grpc.ClientMethod<
          $1.MsgDeleteTimedoutPurchaseorder,
          $1.MsgDeleteTimedoutPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/DeleteTimedoutPurchaseorder',
      ($1.MsgDeleteTimedoutPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgDeleteTimedoutPurchaseorderResponse.fromBuffer(value));
  static final _$requestPurchaseorder = $grpc.ClientMethod<
          $1.MsgRequestPurchaseorder, $1.MsgRequestPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Msg/RequestPurchaseorder',
      ($1.MsgRequestPurchaseorder value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgRequestPurchaseorderResponse.fromBuffer(value));

  MsgClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.MsgFinancePurchaseorderResponse> financePurchaseorder(
      $1.MsgFinancePurchaseorder request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$financePurchaseorder, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCancelPurchaseorderResponse> cancelPurchaseorder(
      $1.MsgCancelPurchaseorder request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$cancelPurchaseorder, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCompletePurchaseorderResponse>
      completePurchaseorder($1.MsgCompletePurchaseorder request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$completePurchaseorder, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCreateSentPurchaseorderResponse>
      createSentPurchaseorder($1.MsgCreateSentPurchaseorder request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$createSentPurchaseorder, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgUpdateSentPurchaseorderResponse>
      updateSentPurchaseorder($1.MsgUpdateSentPurchaseorder request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$updateSentPurchaseorder, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgDeleteSentPurchaseorderResponse>
      deleteSentPurchaseorder($1.MsgDeleteSentPurchaseorder request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$deleteSentPurchaseorder, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgCreateTimedoutPurchaseorderResponse>
      createTimedoutPurchaseorder($1.MsgCreateTimedoutPurchaseorder request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$createTimedoutPurchaseorder, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgUpdateTimedoutPurchaseorderResponse>
      updateTimedoutPurchaseorder($1.MsgUpdateTimedoutPurchaseorder request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$updateTimedoutPurchaseorder, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgDeleteTimedoutPurchaseorderResponse>
      deleteTimedoutPurchaseorder($1.MsgDeleteTimedoutPurchaseorder request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$deleteTimedoutPurchaseorder, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgRequestPurchaseorderResponse> requestPurchaseorder(
      $1.MsgRequestPurchaseorder request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$requestPurchaseorder, request, options: options);
  }
}

abstract class MsgServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.purchaseorder.Msg';

  MsgServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.MsgFinancePurchaseorder,
            $1.MsgFinancePurchaseorderResponse>(
        'FinancePurchaseorder',
        financePurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgFinancePurchaseorder.fromBuffer(value),
        ($1.MsgFinancePurchaseorderResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCancelPurchaseorder,
            $1.MsgCancelPurchaseorderResponse>(
        'CancelPurchaseorder',
        cancelPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCancelPurchaseorder.fromBuffer(value),
        ($1.MsgCancelPurchaseorderResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCompletePurchaseorder,
            $1.MsgCompletePurchaseorderResponse>(
        'CompletePurchaseorder',
        completePurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCompletePurchaseorder.fromBuffer(value),
        ($1.MsgCompletePurchaseorderResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCreateSentPurchaseorder,
            $1.MsgCreateSentPurchaseorderResponse>(
        'CreateSentPurchaseorder',
        createSentPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCreateSentPurchaseorder.fromBuffer(value),
        ($1.MsgCreateSentPurchaseorderResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgUpdateSentPurchaseorder,
            $1.MsgUpdateSentPurchaseorderResponse>(
        'UpdateSentPurchaseorder',
        updateSentPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgUpdateSentPurchaseorder.fromBuffer(value),
        ($1.MsgUpdateSentPurchaseorderResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgDeleteSentPurchaseorder,
            $1.MsgDeleteSentPurchaseorderResponse>(
        'DeleteSentPurchaseorder',
        deleteSentPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgDeleteSentPurchaseorder.fromBuffer(value),
        ($1.MsgDeleteSentPurchaseorderResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCreateTimedoutPurchaseorder,
            $1.MsgCreateTimedoutPurchaseorderResponse>(
        'CreateTimedoutPurchaseorder',
        createTimedoutPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCreateTimedoutPurchaseorder.fromBuffer(value),
        ($1.MsgCreateTimedoutPurchaseorderResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgUpdateTimedoutPurchaseorder,
            $1.MsgUpdateTimedoutPurchaseorderResponse>(
        'UpdateTimedoutPurchaseorder',
        updateTimedoutPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgUpdateTimedoutPurchaseorder.fromBuffer(value),
        ($1.MsgUpdateTimedoutPurchaseorderResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgDeleteTimedoutPurchaseorder,
            $1.MsgDeleteTimedoutPurchaseorderResponse>(
        'DeleteTimedoutPurchaseorder',
        deleteTimedoutPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgDeleteTimedoutPurchaseorder.fromBuffer(value),
        ($1.MsgDeleteTimedoutPurchaseorderResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgRequestPurchaseorder,
            $1.MsgRequestPurchaseorderResponse>(
        'RequestPurchaseorder',
        requestPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgRequestPurchaseorder.fromBuffer(value),
        ($1.MsgRequestPurchaseorderResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.MsgFinancePurchaseorderResponse> financePurchaseorder_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgFinancePurchaseorder> request) async {
    return financePurchaseorder(call, await request);
  }

  $async.Future<$1.MsgCancelPurchaseorderResponse> cancelPurchaseorder_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgCancelPurchaseorder> request) async {
    return cancelPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgCompletePurchaseorderResponse> completePurchaseorder_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgCompletePurchaseorder> request) async {
    return completePurchaseorder(call, await request);
  }

  $async.Future<$1.MsgCreateSentPurchaseorderResponse>
      createSentPurchaseorder_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgCreateSentPurchaseorder> request) async {
    return createSentPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgUpdateSentPurchaseorderResponse>
      updateSentPurchaseorder_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgUpdateSentPurchaseorder> request) async {
    return updateSentPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgDeleteSentPurchaseorderResponse>
      deleteSentPurchaseorder_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgDeleteSentPurchaseorder> request) async {
    return deleteSentPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgCreateTimedoutPurchaseorderResponse>
      createTimedoutPurchaseorder_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgCreateTimedoutPurchaseorder> request) async {
    return createTimedoutPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgUpdateTimedoutPurchaseorderResponse>
      updateTimedoutPurchaseorder_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgUpdateTimedoutPurchaseorder> request) async {
    return updateTimedoutPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgDeleteTimedoutPurchaseorderResponse>
      deleteTimedoutPurchaseorder_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgDeleteTimedoutPurchaseorder> request) async {
    return deleteTimedoutPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgRequestPurchaseorderResponse> requestPurchaseorder_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgRequestPurchaseorder> request) async {
    return requestPurchaseorder(call, await request);
  }

  $async.Future<$1.MsgFinancePurchaseorderResponse> financePurchaseorder(
      $grpc.ServiceCall call, $1.MsgFinancePurchaseorder request);
  $async.Future<$1.MsgCancelPurchaseorderResponse> cancelPurchaseorder(
      $grpc.ServiceCall call, $1.MsgCancelPurchaseorder request);
  $async.Future<$1.MsgCompletePurchaseorderResponse> completePurchaseorder(
      $grpc.ServiceCall call, $1.MsgCompletePurchaseorder request);
  $async.Future<$1.MsgCreateSentPurchaseorderResponse> createSentPurchaseorder(
      $grpc.ServiceCall call, $1.MsgCreateSentPurchaseorder request);
  $async.Future<$1.MsgUpdateSentPurchaseorderResponse> updateSentPurchaseorder(
      $grpc.ServiceCall call, $1.MsgUpdateSentPurchaseorder request);
  $async.Future<$1.MsgDeleteSentPurchaseorderResponse> deleteSentPurchaseorder(
      $grpc.ServiceCall call, $1.MsgDeleteSentPurchaseorder request);
  $async.Future<$1.MsgCreateTimedoutPurchaseorderResponse>
      createTimedoutPurchaseorder(
          $grpc.ServiceCall call, $1.MsgCreateTimedoutPurchaseorder request);
  $async.Future<$1.MsgUpdateTimedoutPurchaseorderResponse>
      updateTimedoutPurchaseorder(
          $grpc.ServiceCall call, $1.MsgUpdateTimedoutPurchaseorder request);
  $async.Future<$1.MsgDeleteTimedoutPurchaseorderResponse>
      deleteTimedoutPurchaseorder(
          $grpc.ServiceCall call, $1.MsgDeleteTimedoutPurchaseorder request);
  $async.Future<$1.MsgRequestPurchaseorderResponse> requestPurchaseorder(
      $grpc.ServiceCall call, $1.MsgRequestPurchaseorder request);
}
