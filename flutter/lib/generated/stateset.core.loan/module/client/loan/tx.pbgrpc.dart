///
//  Generated code. Do not modify.
//  source: loan/tx.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'tx.pb.dart' as $1;
export 'tx.pb.dart';

class MsgClient extends $grpc.Client {
  static final _$requestLoan =
      $grpc.ClientMethod<$1.MsgRequestLoan, $1.MsgRequestLoanResponse>(
          '/stateset.core.loan.Msg/RequestLoan',
          ($1.MsgRequestLoan value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgRequestLoanResponse.fromBuffer(value));
  static final _$approveLoan =
      $grpc.ClientMethod<$1.MsgApproveLoan, $1.MsgApproveLoanResponse>(
          '/stateset.core.loan.Msg/ApproveLoan',
          ($1.MsgApproveLoan value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgApproveLoanResponse.fromBuffer(value));
  static final _$repayLoan =
      $grpc.ClientMethod<$1.MsgRepayLoan, $1.MsgRepayLoanResponse>(
          '/stateset.core.loan.Msg/RepayLoan',
          ($1.MsgRepayLoan value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgRepayLoanResponse.fromBuffer(value));
  static final _$liquidateLoan =
      $grpc.ClientMethod<$1.MsgLiquidateLoan, $1.MsgLiquidateLoanResponse>(
          '/stateset.core.loan.Msg/LiquidateLoan',
          ($1.MsgLiquidateLoan value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgLiquidateLoanResponse.fromBuffer(value));
  static final _$cancelLoan =
      $grpc.ClientMethod<$1.MsgCancelLoan, $1.MsgCancelLoanResponse>(
          '/stateset.core.loan.Msg/CancelLoan',
          ($1.MsgCancelLoan value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.MsgCancelLoanResponse.fromBuffer(value));

  MsgClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$1.MsgRequestLoanResponse> requestLoan(
      $1.MsgRequestLoan request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$requestLoan, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgApproveLoanResponse> approveLoan(
      $1.MsgApproveLoan request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$approveLoan, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgRepayLoanResponse> repayLoan(
      $1.MsgRepayLoan request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$repayLoan, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgLiquidateLoanResponse> liquidateLoan(
      $1.MsgLiquidateLoan request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$liquidateLoan, request, options: options);
  }

  $grpc.ResponseFuture<$1.MsgCancelLoanResponse> cancelLoan(
      $1.MsgCancelLoan request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$cancelLoan, request, options: options);
  }
}

abstract class MsgServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.loan.Msg';

  MsgServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$1.MsgRequestLoan, $1.MsgRequestLoanResponse>(
            'RequestLoan',
            requestLoan_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.MsgRequestLoan.fromBuffer(value),
            ($1.MsgRequestLoanResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$1.MsgApproveLoan, $1.MsgApproveLoanResponse>(
            'ApproveLoan',
            approveLoan_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.MsgApproveLoan.fromBuffer(value),
            ($1.MsgApproveLoanResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgRepayLoan, $1.MsgRepayLoanResponse>(
        'RepayLoan',
        repayLoan_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.MsgRepayLoan.fromBuffer(value),
        ($1.MsgRepayLoanResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$1.MsgLiquidateLoan, $1.MsgLiquidateLoanResponse>(
            'LiquidateLoan',
            liquidateLoan_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.MsgLiquidateLoan.fromBuffer(value),
            ($1.MsgLiquidateLoanResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.MsgCancelLoan, $1.MsgCancelLoanResponse>(
        'CancelLoan',
        cancelLoan_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.MsgCancelLoan.fromBuffer(value),
        ($1.MsgCancelLoanResponse value) => value.writeToBuffer()));
  }

  $async.Future<$1.MsgRequestLoanResponse> requestLoan_Pre(
      $grpc.ServiceCall call, $async.Future<$1.MsgRequestLoan> request) async {
    return requestLoan(call, await request);
  }

  $async.Future<$1.MsgApproveLoanResponse> approveLoan_Pre(
      $grpc.ServiceCall call, $async.Future<$1.MsgApproveLoan> request) async {
    return approveLoan(call, await request);
  }

  $async.Future<$1.MsgRepayLoanResponse> repayLoan_Pre(
      $grpc.ServiceCall call, $async.Future<$1.MsgRepayLoan> request) async {
    return repayLoan(call, await request);
  }

  $async.Future<$1.MsgLiquidateLoanResponse> liquidateLoan_Pre(
      $grpc.ServiceCall call,
      $async.Future<$1.MsgLiquidateLoan> request) async {
    return liquidateLoan(call, await request);
  }

  $async.Future<$1.MsgCancelLoanResponse> cancelLoan_Pre(
      $grpc.ServiceCall call, $async.Future<$1.MsgCancelLoan> request) async {
    return cancelLoan(call, await request);
  }

  $async.Future<$1.MsgRequestLoanResponse> requestLoan(
      $grpc.ServiceCall call, $1.MsgRequestLoan request);
  $async.Future<$1.MsgApproveLoanResponse> approveLoan(
      $grpc.ServiceCall call, $1.MsgApproveLoan request);
  $async.Future<$1.MsgRepayLoanResponse> repayLoan(
      $grpc.ServiceCall call, $1.MsgRepayLoan request);
  $async.Future<$1.MsgLiquidateLoanResponse> liquidateLoan(
      $grpc.ServiceCall call, $1.MsgLiquidateLoan request);
  $async.Future<$1.MsgCancelLoanResponse> cancelLoan(
      $grpc.ServiceCall call, $1.MsgCancelLoan request);
}
