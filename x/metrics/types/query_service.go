package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the metrics Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the metrics Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.metrics.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "SystemMetrics", Handler: _Query_SystemMetrics_Handler},
		{MethodName: "Counter", Handler: _Query_Counter_Handler},
		{MethodName: "Gauge", Handler: _Query_Gauge_Handler},
		{MethodName: "ModuleHealth", Handler: _Query_ModuleHealth_Handler},
		{MethodName: "Alerts", Handler: _Query_Alerts_Handler},
		{MethodName: "PrometheusMetrics", Handler: _Query_PrometheusMetrics_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/metrics",
}

func _Query_SystemMetrics_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QuerySystemMetricsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SystemMetrics(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.metrics.Query/SystemMetrics"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).SystemMetrics(ctx, request.(*QuerySystemMetricsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Counter_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryCounterRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Counter(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.metrics.Query/Counter"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Counter(ctx, request.(*QueryCounterRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Gauge_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryGaugeRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Gauge(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.metrics.Query/Gauge"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Gauge(ctx, request.(*QueryGaugeRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_ModuleHealth_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryModuleHealthRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModuleHealth(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.metrics.Query/ModuleHealth"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).ModuleHealth(ctx, request.(*QueryModuleHealthRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Alerts_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryAlertsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Alerts(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.metrics.Query/Alerts"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Alerts(ctx, request.(*QueryAlertsRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_PrometheusMetrics_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryPrometheusMetricsRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PrometheusMetrics(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.metrics.Query/PrometheusMetrics"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).PrometheusMetrics(ctx, request.(*QueryPrometheusMetricsRequest))
	}
	return interceptor(ctx, req, info, handler)
}
