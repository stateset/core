///
//  Generated code. Do not modify.
//  source: agreement/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'tx.pb.dart' as $1;
export 'tx.pb.dart';

class MsgClient extends $grpc.Client {
  static final _$activateAgreement = $grpc.ClientMethod<$1.MsgActivateAgreement,
          $1.MsgActivateAgreementResponse>(
      '/stateset.core.agreement.Msg/ActivateAgreement',
      ($1.MsgActivateAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgActivateAgreementResponse.fromBuffer(value));
  static final _$expireAgreement =
      $grpc.ClientMethod<$1.MsgExpireAgreement, $1.MsgExpireAgreementResponse>(
          '/stateset.core.agreement.Msg/ExpireAgreement',
          ($1.MsgExpireAgreement value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgExpireAgreementResponse.fromBuffer(value));
  static final _$renewAgreement =
      $grpc.ClientMethod<$1.MsgRenewAgreement, $1.MsgRenewAgreementResponse>(
          '/stateset.core.agreement.Msg/RenewAgreement',
          ($1.MsgRenewAgreement value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgRenewAgreementResponse.fromBuffer(value));
  static final _$terminateAgreement = $grpc.ClientMethod<
          $1.MsgTerminateAgreement, $1.MsgTerminateAgreementResponse>(
      '/stateset.core.agreement.Msg/TerminateAgreement',
      ($1.MsgTerminateAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgTerminateAgreementResponse.fromBuffer(value));
  static final _$createSentAgreement = $grpc.ClientMethod<
          $1.MsgCreateSentAgreement, $1.MsgCreateSentAgreementResponse>(
      '/stateset.core.agreement.Msg/CreateSentAgreement',
      ($1.MsgCreateSentAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCreateSentAgreementResponse.fromBuffer(value));
  static final _$updateSentAgreement = $grpc.ClientMethod<
          $1.MsgUpdateSentAgreement, $1.MsgUpdateSentAgreementResponse>(
      '/stateset.core.agreement.Msg/UpdateSentAgreement',
      ($1.MsgUpdateSentAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgUpdateSentAgreementResponse.fromBuffer(value));
  static final _$deleteSentAgreement = $grpc.ClientMethod<
          $1.MsgDeleteSentAgreement, $1.MsgDeleteSentAgreementResponse>(
      '/stateset.core.agreement.Msg/DeleteSentAgreement',
      ($1.MsgDeleteSentAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgDeleteSentAgreementResponse.fromBuffer(value));
  static final _$createTimedoutAgreement = $grpc.ClientMethod<
          $1.MsgCreateTimedoutAgreement, $1.MsgCreateTimedoutAgreementResponse>(
      '/stateset.core.agreement.Msg/CreateTimedoutAgreement',
      ($1.MsgCreateTimedoutAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgCreateTimedoutAgreementResponse.fromBuffer(value));
  static final _$updateTimedoutAgreement = $grpc.ClientMethod<
          $1.MsgUpdateTimedoutAgreement, $1.MsgUpdateTimedoutAgreementResponse>(
      '/stateset.core.agreement.Msg/UpdateTimedoutAgreement',
      ($1.MsgUpdateTimedoutAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgUpdateTimedoutAgreementResponse.fromBuffer(value));
  static final _$deleteTimedoutAgreement = $grpc.ClientMethod<
          $1.MsgDeleteTimedoutAgreement, $1.MsgDeleteTimedoutAgreementResponse>(
      '/stateset.core.agreement.Msg/DeleteTimedoutAgreement',
      ($1.MsgDeleteTimedoutAgreement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.MsgDeleteTimedoutAgreementResponse.fromBuffer(value));

  MsgClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.MsgActivateAgreementResponse> activateAgreement(
      $1.MsgActivateAgreement request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$activateAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgExpireAgreementResponse> expireAgreement(
      $1.MsgExpireAgreement request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$expireAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgRenewAgreementResponse> renewAgreement(
      $1.MsgRenewAgreement request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$renewAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgTerminateAgreementResponse> terminateAgreement(
      $1.MsgTerminateAgreement request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$terminateAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCreateSentAgreementResponse> createSentAgreement(
      $1.MsgCreateSentAgreement request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$createSentAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgUpdateSentAgreementResponse> updateSentAgreement(
      $1.MsgUpdateSentAgreement request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$updateSentAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgDeleteSentAgreementResponse> deleteSentAgreement(
      $1.MsgDeleteSentAgreement request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$deleteSentAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCreateTimedoutAgreementResponse>
      createTimedoutAgreement($1.MsgCreateTimedoutAgreement request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$createTimedoutAgreement, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgUpdateTimedoutAgreementResponse>
      updateTimedoutAgreement($1.MsgUpdateTimedoutAgreement request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$updateTimedoutAgreement, request,
        options: options);
  }

  $grpc.ResponseFuture<$1.MsgDeleteTimedoutAgreementResponse>
      deleteTimedoutAgreement($1.MsgDeleteTimedoutAgreement request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$deleteTimedoutAgreement, request,
        options: options);
  }
}

abstract class MsgServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.agreement.Msg';

  MsgServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.MsgActivateAgreement,
            $1.MsgActivateAgreementResponse>(
        'ActivateAgreement',
        activateAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgActivateAgreement.fromBuffer(value),
        ($1.MsgActivateAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgExpireAgreement,
            $1.MsgExpireAgreementResponse>(
        'ExpireAgreement',
        expireAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgExpireAgreement.fromBuffer(value),
        ($1.MsgExpireAgreementResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$1.MsgRenewAgreement, $1.MsgRenewAgreementResponse>(
            'RenewAgreement',
            renewAgreement_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.MsgRenewAgreement.fromBuffer(value),
            ($1.MsgRenewAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgTerminateAgreement,
            $1.MsgTerminateAgreementResponse>(
        'TerminateAgreement',
        terminateAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgTerminateAgreement.fromBuffer(value),
        ($1.MsgTerminateAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCreateSentAgreement,
            $1.MsgCreateSentAgreementResponse>(
        'CreateSentAgreement',
        createSentAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCreateSentAgreement.fromBuffer(value),
        ($1.MsgCreateSentAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgUpdateSentAgreement,
            $1.MsgUpdateSentAgreementResponse>(
        'UpdateSentAgreement',
        updateSentAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgUpdateSentAgreement.fromBuffer(value),
        ($1.MsgUpdateSentAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgDeleteSentAgreement,
            $1.MsgDeleteSentAgreementResponse>(
        'DeleteSentAgreement',
        deleteSentAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgDeleteSentAgreement.fromBuffer(value),
        ($1.MsgDeleteSentAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCreateTimedoutAgreement,
            $1.MsgCreateTimedoutAgreementResponse>(
        'CreateTimedoutAgreement',
        createTimedoutAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgCreateTimedoutAgreement.fromBuffer(value),
        ($1.MsgCreateTimedoutAgreementResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgUpdateTimedoutAgreement,
            $1.MsgUpdateTimedoutAgreementResponse>(
        'UpdateTimedoutAgreement',
        updateTimedoutAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgUpdateTimedoutAgreement.fromBuffer(value),
        ($1.MsgUpdateTimedoutAgreementResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgDeleteTimedoutAgreement,
            $1.MsgDeleteTimedoutAgreementResponse>(
        'DeleteTimedoutAgreement',
        deleteTimedoutAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.MsgDeleteTimedoutAgreement.fromBuffer(value),
        ($1.MsgDeleteTimedoutAgreementResponse value) =>
            value.writeToBuffer()));
  }

  $async.Future<$1.MsgActivateAgreementResponse> activateAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgActivateAgreement> request) async {
    return activateAgreement(call, await request);
  }

  $async.Future<$1.MsgExpireAgreementResponse> expireAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgExpireAgreement> request) async {
    return expireAgreement(call, await request);
  }

  $async.Future<$1.MsgRenewAgreementResponse> renewAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgRenewAgreement> request) async {
    return renewAgreement(call, await request);
  }

  $async.Future<$1.MsgTerminateAgreementResponse> terminateAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgTerminateAgreement> request) async {
    return terminateAgreement(call, await request);
  }

  $async.Future<$1.MsgCreateSentAgreementResponse> createSentAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgCreateSentAgreement> request) async {
    return createSentAgreement(call, await request);
  }

  $async.Future<$1.MsgUpdateSentAgreementResponse> updateSentAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgUpdateSentAgreement> request) async {
    return updateSentAgreement(call, await request);
  }

  $async.Future<$1.MsgDeleteSentAgreementResponse> deleteSentAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgDeleteSentAgreement> request) async {
    return deleteSentAgreement(call, await request);
  }

  $async.Future<$1.MsgCreateTimedoutAgreementResponse>
      createTimedoutAgreement_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgCreateTimedoutAgreement> request) async {
    return createTimedoutAgreement(call, await request);
  }

  $async.Future<$1.MsgUpdateTimedoutAgreementResponse>
      updateTimedoutAgreement_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgUpdateTimedoutAgreement> request) async {
    return updateTimedoutAgreement(call, await request);
  }

  $async.Future<$1.MsgDeleteTimedoutAgreementResponse>
      deleteTimedoutAgreement_Pre($grpc.ServiceCall call,
          $async.Future<$1.MsgDeleteTimedoutAgreement> request) async {
    return deleteTimedoutAgreement(call, await request);
  }

  $async.Future<$1.MsgActivateAgreementResponse> activateAgreement(
      $grpc.ServiceCall call, $1.MsgActivateAgreement request);
  $async.Future<$1.MsgExpireAgreementResponse> expireAgreement(
      $grpc.ServiceCall call, $1.MsgExpireAgreement request);
  $async.Future<$1.MsgRenewAgreementResponse> renewAgreement(
      $grpc.ServiceCall call, $1.MsgRenewAgreement request);
  $async.Future<$1.MsgTerminateAgreementResponse> terminateAgreement(
      $grpc.ServiceCall call, $1.MsgTerminateAgreement request);
  $async.Future<$1.MsgCreateSentAgreementResponse> createSentAgreement(
      $grpc.ServiceCall call, $1.MsgCreateSentAgreement request);
  $async.Future<$1.MsgUpdateSentAgreementResponse> updateSentAgreement(
      $grpc.ServiceCall call, $1.MsgUpdateSentAgreement request);
  $async.Future<$1.MsgDeleteSentAgreementResponse> deleteSentAgreement(
      $grpc.ServiceCall call, $1.MsgDeleteSentAgreement request);
  $async.Future<$1.MsgCreateTimedoutAgreementResponse> createTimedoutAgreement(
      $grpc.ServiceCall call, $1.MsgCreateTimedoutAgreement request);
  $async.Future<$1.MsgUpdateTimedoutAgreementResponse> updateTimedoutAgreement(
      $grpc.ServiceCall call, $1.MsgUpdateTimedoutAgreement request);
  $async.Future<$1.MsgDeleteTimedoutAgreementResponse> deleteTimedoutAgreement(
      $grpc.ServiceCall call, $1.MsgDeleteTimedoutAgreement request);
}
