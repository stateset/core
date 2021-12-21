///
//  Generated code. Do not modify.
//  source: loan/query.proto
//
// @dart = 2.3
// ignore_for_file: annotate_overrides,camel_case_types,unnecessary_const,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type,unnecessary_this,prefer_final_fields

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'query.pb.dart' as $0;
export 'query.pb.dart';

class QueryClient extends $grpc.Client {
  static final _$loan =
      $grpc.ClientMethod<$0.QueryGetLoanRequest, $0.QueryGetLoanResponse>(
          '/stateset.core.loan.Query/Loan',
          ($0.QueryGetLoanRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.QueryGetLoanResponse.fromBuffer(value));
  static final _$loanAll =
      $grpc.ClientMethod<$0.QueryAllLoanRequest, $0.QueryAllLoanResponse>(
          '/stateset.core.loan.Query/LoanAll',
          ($0.QueryAllLoanRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.QueryAllLoanResponse.fromBuffer(value));

  QueryClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options,
      $core.Iterable<$grpc.ClientInterceptor> interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.QueryGetLoanResponse> loan(
      $0.QueryGetLoanRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$loan, request, options: options);
  }

  $grpc.ResponseFuture<$0.QueryAllLoanResponse> loanAll(
      $0.QueryAllLoanRequest request,
      {$grpc.CallOptions options}) {
    return $createUnaryCall(_$loanAll, request, options: options);
  }
}

abstract class QueryServiceBase extends $grpc.Service {
  $core.String get $name => 'stateset.core.loan.Query';

  QueryServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.QueryGetLoanRequest, $0.QueryGetLoanResponse>(
            'Loan',
            loan_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.QueryGetLoanRequest.fromBuffer(value),
            ($0.QueryGetLoanResponse value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$0.QueryAllLoanRequest, $0.QueryAllLoanResponse>(
            'LoanAll',
            loanAll_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.QueryAllLoanRequest.fromBuffer(value),
            ($0.QueryAllLoanResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.QueryGetLoanResponse> loan_Pre($grpc.ServiceCall call,
      $async.Future<$0.QueryGetLoanRequest> request) async {
    return loan(call, await request);
  }

  $async.Future<$0.QueryAllLoanResponse> loanAll_Pre($grpc.ServiceCall call,
      $async.Future<$0.QueryAllLoanRequest> request) async {
    return loanAll(call, await request);
  }

  $async.Future<$0.QueryGetLoanResponse> loan(
      $grpc.ServiceCall call, $0.QueryGetLoanRequest request);
  $async.Future<$0.QueryAllLoanResponse> loanAll(
      $grpc.ServiceCall call, $0.QueryAllLoanRequest request);
}
