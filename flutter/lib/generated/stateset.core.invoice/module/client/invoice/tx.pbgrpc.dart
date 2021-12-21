///
//  Generated code. Do not modify.
//  source: invoice/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'tx.pb.dart' as $1;
export 'tx.pb.dart';

class MsgClient extends $grpc.Client {
  static final _$factorInvoice =
      $grpc.ClientMethod<$1.MsgFactorInvoice, $1.MsgFactorInvoiceResponse>(
          '/stateset.core.invoice.Msg/FactorInvoice',
          ($1.MsgFactorInvoice value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgFactorInvoiceResponse.fromBuffer(value));
  static final _$createSentInvoice = $grpc.ClientMethod<$1.MsgCreateSentInvoice,
          $1.MsgCreateSentInvoiceResponse>(
      '/stateset.core.invoice.Msg/CreateSentInvoice',
      ($1.MsgCreateSentInvoice value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCreateSentInvoiceResponse.fromBuffer(value));
  static final _$updateSentInvoice = $grpc.ClientMethod<$1.MsgUpdateSentInvoice,
          $1.MsgUpdateSentInvoiceResponse>(
      '/stateset.core.invoice.Msg/UpdateSentInvoice',
      ($1.MsgUpdateSentInvoice value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgUpdateSentInvoiceResponse.fromBuffer(value));
  static final _$deleteSentInvoice = $grpc.ClientMethod<$1.MsgDeleteSentInvoice,
          $1.MsgDeleteSentInvoiceResponse>(
      '/stateset.core.invoice.Msg/DeleteSentInvoice',
      ($1.MsgDeleteSentInvoice value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgDeleteSentInvoiceResponse.fromBuffer(value));
  static final _$createTimedoutInvoice = $grpc.ClientMethod<
          $1.MsgCreateTimedoutInvoice, $1.MsgCreateTimedoutInvoiceResponse>(
      '/stateset.core.invoice.Msg/CreateTimedoutInvoice',
      ($1.MsgCreateTimedoutInvoice value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCreateTimedoutInvoiceResponse.fromBuffer(value));
  static final _$updateTimedoutInvoice = $grpc.ClientMethod<
          $1.MsgUpdateTimedoutInvoice, $1.MsgUpdateTimedoutInvoiceResponse>(
      '/stateset.core.invoice.Msg/UpdateTimedoutInvoice',
      ($1.MsgUpdateTimedoutInvoice value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgUpdateTimedoutInvoiceResponse.fromBuffer(value));
  static final _$deleteTimedoutInvoice = $grpc.ClientMethod<
          $1.MsgDeleteTimedoutInvoice, $1.MsgDeleteTimedoutInvoiceResponse>(
      '/stateset.core.invoice.Msg/DeleteTimedoutInvoice',
      ($1.MsgDeleteTimedoutInvoice value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgDeleteTimedoutInvoiceResponse.fromBuffer(value));

  MsgClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.MsgFactorInvoiceResponse> factorInvoice(
      $1.MsgFactorInvoice request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$factorInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCreateSentInvoiceResponse> createSentInvoice(
      $1.MsgCreateSentInvoice request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$createSentInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgUpdateSentInvoiceResponse> updateSentInvoice(
      $1.MsgUpdateSentInvoice request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$updateSentInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgDeleteSentInvoiceResponse> deleteSentInvoice(
      $1.MsgDeleteSentInvoice request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$deleteSentInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCreateTimedoutInvoiceResponse>
      createTimedoutInvoice($1.MsgCreateTimedoutInvoice request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$createTimedoutInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgUpdateTimedoutInvoiceResponse>
      updateTimedoutInvoice($1.MsgUpdateTimedoutInvoice request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$updateTimedoutInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgDeleteTimedoutInvoiceResponse>
      deleteTimedoutInvoice($1.MsgDeleteTimedoutInvoice request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$deleteTimedoutInvoice, request, options: options);
  }
}

abstract class MsgServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.invoice.Msg';

  MsgServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$1.MsgFactorInvoice, $1.MsgFactorInvoiceResponse>(
            'FactorInvoice',
            factorInvoice_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.MsgFactorInvoice.fromBuffer(value),
            ($1.MsgFactorInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCreateSentInvoice,
            $1.MsgCreateSentInvoiceResponse>(
        'CreateSentInvoice',
        createSentInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCreateSentInvoice.fromBuffer(value),
        ($1.MsgCreateSentInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgUpdateSentInvoice,
            $1.MsgUpdateSentInvoiceResponse>(
        'UpdateSentInvoice',
        updateSentInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgUpdateSentInvoice.fromBuffer(value),
        ($1.MsgUpdateSentInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgDeleteSentInvoice,
            $1.MsgDeleteSentInvoiceResponse>(
        'DeleteSentInvoice',
        deleteSentInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgDeleteSentInvoice.fromBuffer(value),
        ($1.MsgDeleteSentInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCreateTimedoutInvoice,
            $1.MsgCreateTimedoutInvoiceResponse>(
        'CreateTimedoutInvoice',
        createTimedoutInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCreateTimedoutInvoice.fromBuffer(value),
        ($1.MsgCreateTimedoutInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgUpdateTimedoutInvoice,
            $1.MsgUpdateTimedoutInvoiceResponse>(
        'UpdateTimedoutInvoice',
        updateTimedoutInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgUpdateTimedoutInvoice.fromBuffer(value),
        ($1.MsgUpdateTimedoutInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgDeleteTimedoutInvoice,
            $1.MsgDeleteTimedoutInvoiceResponse>(
        'DeleteTimedoutInvoice',
        deleteTimedoutInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgDeleteTimedoutInvoice.fromBuffer(value),
        ($1.MsgDeleteTimedoutInvoiceResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.MsgFactorInvoiceResponse> factorInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgFactorInvoice> request) async {
    return factorInvoice(call, await request);
  }

  $async.Future<$1.MsgCreateSentInvoiceResponse> createSentInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgCreateSentInvoice> request) async {
    return createSentInvoice(call, await request);
  }

  $async.Future<$1.MsgUpdateSentInvoiceResponse> updateSentInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgUpdateSentInvoice> request) async {
    return updateSentInvoice(call, await request);
  }

  $async.Future<$1.MsgDeleteSentInvoiceResponse> deleteSentInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgDeleteSentInvoice> request) async {
    return deleteSentInvoice(call, await request);
  }

  $async.Future<$1.MsgCreateTimedoutInvoiceResponse> createTimedoutInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgCreateTimedoutInvoice> request) async {
    return createTimedoutInvoice(call, await request);
  }

  $async.Future<$1.MsgUpdateTimedoutInvoiceResponse> updateTimedoutInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgUpdateTimedoutInvoice> request) async {
    return updateTimedoutInvoice(call, await request);
  }

  $async.Future<$1.MsgDeleteTimedoutInvoiceResponse> deleteTimedoutInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgDeleteTimedoutInvoice> request) async {
    return deleteTimedoutInvoice(call, await request);
  }

  $async.Future<$1.MsgFactorInvoiceResponse> factorInvoice(
      $grpc.ServiceCall call, $1.MsgFactorInvoice request);
  $async.Future<$1.MsgCreateSentInvoiceResponse> createSentInvoice(
      $grpc.ServiceCall call, $1.MsgCreateSentInvoice request);
  $async.Future<$1.MsgUpdateSentInvoiceResponse> updateSentInvoice(
      $grpc.ServiceCall call, $1.MsgUpdateSentInvoice request);
  $async.Future<$1.MsgDeleteSentInvoiceResponse> deleteSentInvoice(
      $grpc.ServiceCall call, $1.MsgDeleteSentInvoice request);
  $async.Future<$1.MsgCreateTimedoutInvoiceResponse> createTimedoutInvoice(
      $grpc.ServiceCall call, $1.MsgCreateTimedoutInvoice request);
  $async.Future<$1.MsgUpdateTimedoutInvoiceResponse> updateTimedoutInvoice(
      $grpc.ServiceCall call, $1.MsgUpdateTimedoutInvoice request);
  $async.Future<$1.MsgDeleteTimedoutInvoiceResponse> deleteTimedoutInvoice(
      $grpc.ServiceCall call, $1.MsgDeleteTimedoutInvoice request);
}
