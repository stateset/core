///
//  Generated code. Do not modify.
//  source: agreement/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'query.pb.dart' as $0;
export 'query.pb.dart';

class QueryClient extends $grpc.Client {
  static final _$sentAgreement = $grpc.ClientMethod<
          $0.QueryGetSentAgreementRequest, $0.QueryGetSentAgreementResponse>(
      '/stateset.core.agreement.Query/SentAgreement',
      ($0.QueryGetSentAgreementRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetSentAgreementResponse.fromBuffer(value));
  static final _$sentAgreementAll = $grpc.ClientMethod<
          $0.QueryAllSentAgreementRequest, $0.QueryAllSentAgreementResponse>(
      '/stateset.core.agreement.Query/SentAgreementAll',
      ($0.QueryAllSentAgreementRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllSentAgreementResponse.fromBuffer(value));
  static final _$timedoutAgreement = $grpc.ClientMethod<
          $0.QueryGetTimedoutAgreementRequest,
          $0.QueryGetTimedoutAgreementResponse>(
      '/stateset.core.agreement.Query/TimedoutAgreement',
      ($0.QueryGetTimedoutAgreementRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetTimedoutAgreementResponse.fromBuffer(value));
  static final _$timedoutAgreementAll = $grpc.ClientMethod<
          $0.QueryAllTimedoutAgreementRequest,
          $0.QueryAllTimedoutAgreementResponse>(
      '/stateset.core.agreement.Query/TimedoutAgreementAll',
      ($0.QueryAllTimedoutAgreementRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllTimedoutAgreementResponse.fromBuffer(value));
  static final _$agreement = $grpc.ClientMethod<$0.QueryGetAgreementRequest,
          $0.QueryGetAgreementResponse>(
      '/stateset.core.agreement.Query/Agreement',
      ($0.QueryGetAgreementRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetAgreementResponse.fromBuffer(value));
  static final _$agreementAll = $grpc.ClientMethod<$0.QueryAllAgreementRequest,
          $0.QueryAllAgreementResponse>(
      '/stateset.core.agreement.Query/AgreementAll',
      ($0.QueryAllAgreementRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllAgreementResponse.fromBuffer(value));

  QueryClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.QueryGetSentAgreementResponse> sentAgreement(
      $0.QueryGetSentAgreementRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$sentAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllSentAgreementResponse> sentAgreementAll(
      $0.QueryAllSentAgreementRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$sentAgreementAll, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryGetTimedoutAgreementResponse> timedoutAgreement(
      $0.QueryGetTimedoutAgreementRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$timedoutAgreement, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllTimedoutAgreementResponse>
      timedoutAgreementAll($0.QueryAllTimedoutAgreementRequest request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$timedoutAgreementAll, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryGetAgreementResponse> agreement(
      $0.QueryGetAgreementRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$agreement, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllAgreementResponse> agreementAll(
      $0.QueryAllAgreementRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$agreementAll, request, options: options);
  }
}

abstract class QueryServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.agreement.Query';

  QueryServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.QueryGetSentAgreementRequest,
            $0.QueryGetSentAgreementResponse>(
        'SentAgreement',
        sentAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetSentAgreementRequest.fromBuffer(value),
        ($0.QueryGetSentAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllSentAgreementRequest,
            $0.QueryAllSentAgreementResponse>(
        'SentAgreementAll',
        sentAgreementAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllSentAgreementRequest.fromBuffer(value),
        ($0.QueryAllSentAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryGetTimedoutAgreementRequest,
            $0.QueryGetTimedoutAgreementResponse>(
        'TimedoutAgreement',
        timedoutAgreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetTimedoutAgreementRequest.fromBuffer(value),
        ($0.QueryGetTimedoutAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllTimedoutAgreementRequest,
            $0.QueryAllTimedoutAgreementResponse>(
        'TimedoutAgreementAll',
        timedoutAgreementAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllTimedoutAgreementRequest.fromBuffer(value),
        ($0.QueryAllTimedoutAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryGetAgreementRequest,
            $0.QueryGetAgreementResponse>(
        'Agreement',
        agreement_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetAgreementRequest.fromBuffer(value),
        ($0.QueryGetAgreementResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllAgreementRequest,
            $0.QueryAllAgreementResponse>(
        'AgreementAll',
        agreementAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllAgreementRequest.fromBuffer(value),
        ($0.QueryAllAgreementResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.QueryGetSentAgreementResponse> sentAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryGetSentAgreementRequest> request) async {
    return sentAgreement(call, await request);
  }

  $async.Future<$0.QueryAllSentAgreementResponse> sentAgreementAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllSentAgreementRequest> request) async {
    return sentAgreementAll(call, await request);
  }

  $async.Future<$0.QueryGetTimedoutAgreementResponse> timedoutAgreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryGetTimedoutAgreementRequest> request) async {
    return timedoutAgreement(call, await request);
  }

  $async.Future<$0.QueryAllTimedoutAgreementResponse> timedoutAgreementAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllTimedoutAgreementRequest> request) async {
    return timedoutAgreementAll(call, await request);
  }

  $async.Future<$0.QueryGetAgreementResponse> agreement_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryGetAgreementRequest> request) async {
    return agreement(call, await request);
  }

  $async.Future<$0.QueryAllAgreementResponse> agreementAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllAgreementRequest> request) async {
    return agreementAll(call, await request);
  }

  $async.Future<$0.QueryGetSentAgreementResponse> sentAgreement(
      $grpc.ServiceCall call, $0.QueryGetSentAgreementRequest request);
  $async.Future<$0.QueryAllSentAgreementResponse> sentAgreementAll(
      $grpc.ServiceCall call, $0.QueryAllSentAgreementRequest request);
  $async.Future<$0.QueryGetTimedoutAgreementResponse> timedoutAgreement(
      $grpc.ServiceCall call, $0.QueryGetTimedoutAgreementRequest request);
  $async.Future<$0.QueryAllTimedoutAgreementResponse> timedoutAgreementAll(
      $grpc.ServiceCall call, $0.QueryAllTimedoutAgreementRequest request);
  $async.Future<$0.QueryGetAgreementResponse> agreement(
      $grpc.ServiceCall call, $0.QueryGetAgreementRequest request);
  $async.Future<$0.QueryAllAgreementResponse> agreementAll(
      $grpc.ServiceCall call, $0.QueryAllAgreementRequest request);
}
