package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the payments Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the payments Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.payments.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Payment", Handler: _Query_Payment_Handler},
		{MethodName: "Payments", Handler: _Query_Payments_Handler},
		{MethodName: "PaymentsByPayer", Handler: _Query_PaymentsByPayer_Handler},
		{MethodName: "PaymentsByPayee", Handler: _Query_PaymentsByPayee_Handler},
		{MethodName: "PaymentsByStatus", Handler: _Query_PaymentsByStatus_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/payments",
}

func _Query_Payment_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryPaymentRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Payment(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.payments.Query/Payment"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Payment(ctx, request.(*QueryPaymentRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Payments_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryPaymentsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Payments(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.payments.Query/Payments"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Payments(ctx, request.(*QueryPaymentsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_PaymentsByPayer_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryPaymentsByPayerRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PaymentsByPayer(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.payments.Query/PaymentsByPayer"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).PaymentsByPayer(ctx, request.(*QueryPaymentsByPayerRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_PaymentsByPayee_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryPaymentsByPayeeRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PaymentsByPayee(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.payments.Query/PaymentsByPayee"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).PaymentsByPayee(ctx, request.(*QueryPaymentsByPayeeRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_PaymentsByStatus_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryPaymentsByStatusRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PaymentsByStatus(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.payments.Query/PaymentsByStatus"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).PaymentsByStatus(ctx, request.(*QueryPaymentsByStatusRequest))
	}
	return interceptor(ctx, req, info, handler)
}
