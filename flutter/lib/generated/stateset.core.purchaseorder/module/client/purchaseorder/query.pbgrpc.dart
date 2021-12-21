///
//  Generated code. Do not modify.
//  source: purchaseorder/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'query.pb.dart' as $0;
export 'query.pb.dart';

class QueryClient extends $grpc.Client {
  static final _$purchaseorder = $grpc.ClientMethod<
          $0.QueryGetPurchaseorderRequest, $0.QueryGetPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Query/Purchaseorder',
      ($0.QueryGetPurchaseorderRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetPurchaseorderResponse.fromBuffer(value));
  static final _$purchaseorderAll = $grpc.ClientMethod<
          $0.QueryAllPurchaseorderRequest, $0.QueryAllPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Query/PurchaseorderAll',
      ($0.QueryAllPurchaseorderRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllPurchaseorderResponse.fromBuffer(value));
  static final _$sentPurchaseorder = $grpc.ClientMethod<
          $0.QueryGetSentPurchaseorderRequest,
          $0.QueryGetSentPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Query/SentPurchaseorder',
      ($0.QueryGetSentPurchaseorderRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetSentPurchaseorderResponse.fromBuffer(value));
  static final _$sentPurchaseorderAll = $grpc.ClientMethod<
          $0.QueryAllSentPurchaseorderRequest,
          $0.QueryAllSentPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Query/SentPurchaseorderAll',
      ($0.QueryAllSentPurchaseorderRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllSentPurchaseorderResponse.fromBuffer(value));
  static final _$timedoutPurchaseorder = $grpc.ClientMethod<
          $0.QueryGetTimedoutPurchaseorderRequest,
          $0.QueryGetTimedoutPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Query/TimedoutPurchaseorder',
      ($0.QueryGetTimedoutPurchaseorderRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetTimedoutPurchaseorderResponse.fromBuffer(value));
  static final _$timedoutPurchaseorderAll = $grpc.ClientMethod<
          $0.QueryAllTimedoutPurchaseorderRequest,
          $0.QueryAllTimedoutPurchaseorderResponse>(
      '/stateset.core.purchaseorder.Query/TimedoutPurchaseorderAll',
      ($0.QueryAllTimedoutPurchaseorderRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllTimedoutPurchaseorderResponse.fromBuffer(value));

  QueryClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.QueryGetPurchaseorderResponse> purchaseorder(
      $0.QueryGetPurchaseorderRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$purchaseorder, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllPurchaseorderResponse> purchaseorderAll(
      $0.QueryAllPurchaseorderRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$purchaseorderAll, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryGetSentPurchaseorderResponse> sentPurchaseorder(
      $0.QueryGetSentPurchaseorderRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$sentPurchaseorder, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllSentPurchaseorderResponse>
      sentPurchaseorderAll($0.QueryAllSentPurchaseorderRequest request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$sentPurchaseorderAll, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryGetTimedoutPurchaseorderResponse>
      timedoutPurchaseorder($0.QueryGetTimedoutPurchaseorderRequest request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$timedoutPurchaseorder, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllTimedoutPurchaseorderResponse>
      timedoutPurchaseorderAll($0.QueryAllTimedoutPurchaseorderRequest request,
          {$grpc.CallOptions options}) {
    return $createUnaryCall(_$timedoutPurchaseorderAll, request,
        options: options);
  }
}

abstract class QueryServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.purchaseorder.Query';

  QueryServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.QueryGetPurchaseorderRequest,
            $0.QueryGetPurchaseorderResponse>(
        'Purchaseorder',
        purchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetPurchaseorderRequest.fromBuffer(value),
        ($0.QueryGetPurchaseorderResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllPurchaseorderRequest,
            $0.QueryAllPurchaseorderResponse>(
        'PurchaseorderAll',
        purchaseorderAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllPurchaseorderRequest.fromBuffer(value),
        ($0.QueryAllPurchaseorderResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryGetSentPurchaseorderRequest,
            $0.QueryGetSentPurchaseorderResponse>(
        'SentPurchaseorder',
        sentPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetSentPurchaseorderRequest.fromBuffer(value),
        ($0.QueryGetSentPurchaseorderResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllSentPurchaseorderRequest,
            $0.QueryAllSentPurchaseorderResponse>(
        'SentPurchaseorderAll',
        sentPurchaseorderAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllSentPurchaseorderRequest.fromBuffer(value),
        ($0.QueryAllSentPurchaseorderResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryGetTimedoutPurchaseorderRequest,
            $0.QueryGetTimedoutPurchaseorderResponse>(
        'TimedoutPurchaseorder',
        timedoutPurchaseorder_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetTimedoutPurchaseorderRequest.fromBuffer(value),
        ($0.QueryGetTimedoutPurchaseorderResponse value) =>
            value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllTimedoutPurchaseorderRequest,
            $0.QueryAllTimedoutPurchaseorderResponse>(
        'TimedoutPurchaseorderAll',
        timedoutPurchaseorderAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllTimedoutPurchaseorderRequest.fromBuffer(value),
        ($0.QueryAllTimedoutPurchaseorderResponse value) =>
            value.writeToBuffer()));
  }

  $async.Future<$0.QueryGetPurchaseorderResponse> purchaseorder_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryGetPurchaseorderRequest> request) async {
    return purchaseorder(call, await request);
  }

  $async.Future<$0.QueryAllPurchaseorderResponse> purchaseorderAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllPurchaseorderRequest> request) async {
    return purchaseorderAll(call, await request);
  }

  $async.Future<$0.QueryGetSentPurchaseorderResponse> sentPurchaseorder_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryGetSentPurchaseorderRequest> request) async {
    return sentPurchaseorder(call, await request);
  }

  $async.Future<$0.QueryAllSentPurchaseorderResponse> sentPurchaseorderAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllSentPurchaseorderRequest> request) async {
    return sentPurchaseorderAll(call, await request);
  }

  $async.Future<$0.QueryGetTimedoutPurchaseorderResponse>
      timedoutPurchaseorder_Pre(
          $grpc.ServiceCall call,
          $async.Future<$0.QueryGetTimedoutPurchaseorderRequest>
              request) async {
    return timedoutPurchaseorder(call, await request);
  }

  $async.Future<$0.QueryAllTimedoutPurchaseorderResponse>
      timedoutPurchaseorderAll_Pre(
          $grpc.ServiceCall call,
          $async.Future<$0.QueryAllTimedoutPurchaseorderRequest>
              request) async {
    return timedoutPurchaseorderAll(call, await request);
  }

  $async.Future<$0.QueryGetPurchaseorderResponse> purchaseorder(
      $grpc.ServiceCall call, $0.QueryGetPurchaseorderRequest request);
  $async.Future<$0.QueryAllPurchaseorderResponse> purchaseorderAll(
      $grpc.ServiceCall call, $0.QueryAllPurchaseorderRequest request);
  $async.Future<$0.QueryGetSentPurchaseorderResponse> sentPurchaseorder(
      $grpc.ServiceCall call, $0.QueryGetSentPurchaseorderRequest request);
  $async.Future<$0.QueryAllSentPurchaseorderResponse> sentPurchaseorderAll(
      $grpc.ServiceCall call, $0.QueryAllSentPurchaseorderRequest request);
  $async.Future<$0.QueryGetTimedoutPurchaseorderResponse> timedoutPurchaseorder(
      $grpc.ServiceCall call, $0.QueryGetTimedoutPurchaseorderRequest request);
  $async.Future<$0.QueryAllTimedoutPurchaseorderResponse>
      timedoutPurchaseorderAll($grpc.ServiceCall call,
          $0.QueryAllTimedoutPurchaseorderRequest request);
}
