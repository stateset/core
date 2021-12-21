///
//  Generated code. Do not modify.
//  source: invoice/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'query.pb.dart' as $0;
export 'query.pb.dart';

class QueryClient extends $grpc.Client {
  static final _$invoice =
      $grpc.ClientMethod<$0.QueryGetInvoiceRequest, $0.QueryGetInvoiceResponse>(
          '/stateset.core.invoice.Query/Invoice',
          ($0.QueryGetInvoiceRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.QueryGetInvoiceResponse.fromBuffer(value));
  static final _$invoiceAll =
      $grpc.ClientMethod<$0.QueryAllInvoiceRequest, $0.QueryAllInvoiceResponse>(
          '/stateset.core.invoice.Query/InvoiceAll',
          ($0.QueryAllInvoiceRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.QueryAllInvoiceResponse.fromBuffer(value));
  static final _$sentInvoice = $grpc.ClientMethod<$0.QueryGetSentInvoiceRequest,
          $0.QueryGetSentInvoiceResponse>(
      '/stateset.core.invoice.Query/SentInvoice',
      ($0.QueryGetSentInvoiceRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetSentInvoiceResponse.fromBuffer(value));
  static final _$sentInvoiceAll = $grpc.ClientMethod<
          $0.QueryAllSentInvoiceRequest, $0.QueryAllSentInvoiceResponse>(
      '/stateset.core.invoice.Query/SentInvoiceAll',
      ($0.QueryAllSentInvoiceRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllSentInvoiceResponse.fromBuffer(value));
  static final _$timedoutInvoice = $grpc.ClientMethod<
          $0.QueryGetTimedoutInvoiceRequest,
          $0.QueryGetTimedoutInvoiceResponse>(
      '/stateset.core.invoice.Query/TimedoutInvoice',
      ($0.QueryGetTimedoutInvoiceRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryGetTimedoutInvoiceResponse.fromBuffer(value));
  static final _$timedoutInvoiceAll = $grpc.ClientMethod<
          $0.QueryAllTimedoutInvoiceRequest,
          $0.QueryAllTimedoutInvoiceResponse>(
      '/stateset.core.invoice.Query/TimedoutInvoiceAll',
      ($0.QueryAllTimedoutInvoiceRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $0.QueryAllTimedoutInvoiceResponse.fromBuffer(value));

  QueryClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.QueryGetInvoiceResponse> invoice(
      $0.QueryGetInvoiceRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$invoice, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllInvoiceResponse> invoiceAll(
      $0.QueryAllInvoiceRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$invoiceAll, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryGetSentInvoiceResponse> sentInvoice(
      $0.QueryGetSentInvoiceRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$sentInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllSentInvoiceResponse> sentInvoiceAll(
      $0.QueryAllSentInvoiceRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$sentInvoiceAll, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryGetTimedoutInvoiceResponse> timedoutInvoice(
      $0.QueryGetTimedoutInvoiceRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$timedoutInvoice, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllTimedoutInvoiceResponse> timedoutInvoiceAll(
      $0.QueryAllTimedoutInvoiceRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$timedoutInvoiceAll, request, options: options);
  }
}

abstract class QueryServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.invoice.Query';

  QueryServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.QueryGetInvoiceRequest,
            $0.QueryGetInvoiceResponse>(
        'Invoice',
        invoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetInvoiceRequest.fromBuffer(value),
        ($0.QueryGetInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllInvoiceRequest,
            $0.QueryAllInvoiceResponse>(
        'InvoiceAll',
        invoiceAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllInvoiceRequest.fromBuffer(value),
        ($0.QueryAllInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryGetSentInvoiceRequest,
            $0.QueryGetSentInvoiceResponse>(
        'SentInvoice',
        sentInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetSentInvoiceRequest.fromBuffer(value),
        ($0.QueryGetSentInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllSentInvoiceRequest,
            $0.QueryAllSentInvoiceResponse>(
        'SentInvoiceAll',
        sentInvoiceAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllSentInvoiceRequest.fromBuffer(value),
        ($0.QueryAllSentInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryGetTimedoutInvoiceRequest,
            $0.QueryGetTimedoutInvoiceResponse>(
        'TimedoutInvoice',
        timedoutInvoice_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryGetTimedoutInvoiceRequest.fromBuffer(value),
        ($0.QueryGetTimedoutInvoiceResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.QueryAllTimedoutInvoiceRequest,
            $0.QueryAllTimedoutInvoiceResponse>(
        'TimedoutInvoiceAll',
        timedoutInvoiceAll_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $0.QueryAllTimedoutInvoiceRequest.fromBuffer(value),
        ($0.QueryAllTimedoutInvoiceResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.QueryGetInvoiceResponse> invoice_Pre($grpc.ServiceCall call,
      $async.Future<$0.QueryGetInvoiceRequest> request) async {
    return invoice(call, await request);
  }

  $async.Future<$0.QueryAllInvoiceResponse> invoiceAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllInvoiceRequest> request) async {
    return invoiceAll(call, await request);
  }

  $async.Future<$0.QueryGetSentInvoiceResponse> sentInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryGetSentInvoiceRequest> request) async {
    return sentInvoice(call, await request);
  }

  $async.Future<$0.QueryAllSentInvoiceResponse> sentInvoiceAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllSentInvoiceRequest> request) async {
    return sentInvoiceAll(call, await request);
  }

  $async.Future<$0.QueryGetTimedoutInvoiceResponse> timedoutInvoice_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryGetTimedoutInvoiceRequest> request) async {
    return timedoutInvoice(call, await request);
  }

  $async.Future<$0.QueryAllTimedoutInvoiceResponse> timedoutInvoiceAll_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.QueryAllTimedoutInvoiceRequest> request) async {
    return timedoutInvoiceAll(call, await request);
  }

  $async.Future<$0.QueryGetInvoiceResponse> invoice(
      $grpc.ServiceCall call, $0.QueryGetInvoiceRequest request);
  $async.Future<$0.QueryAllInvoiceResponse> invoiceAll(
      $grpc.ServiceCall call, $0.QueryAllInvoiceRequest request);
  $async.Future<$0.QueryGetSentInvoiceResponse> sentInvoice(
      $grpc.ServiceCall call, $0.QueryGetSentInvoiceRequest request);
  $async.Future<$0.QueryAllSentInvoiceResponse> sentInvoiceAll(
      $grpc.ServiceCall call, $0.QueryAllSentInvoiceRequest request);
  $async.Future<$0.QueryGetTimedoutInvoiceResponse> timedoutInvoice(
      $grpc.ServiceCall call, $0.QueryGetTimedoutInvoiceRequest request);
  $async.Future<$0.QueryAllTimedoutInvoiceResponse> timedoutInvoiceAll(
      $grpc.ServiceCall call, $0.QueryAllTimedoutInvoiceRequest request);
}
