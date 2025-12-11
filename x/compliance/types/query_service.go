package types

import (
	"context"

	grpc "google.golang.org/grpc"
)

// RegisterQueryServer registers the compliance Query service with the router.
func RegisterQueryServer(router grpc.ServiceRegistrar, srv QueryServer) {
	router.RegisterService(&Query_ServiceDesc, srv)
}

// Query_ServiceDesc describes the compliance Query service for the legacy router.
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.compliance.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Profile", Handler: _Query_Profile_Handler},
		{MethodName: "Profiles", Handler: _Query_Profiles_Handler},
		{MethodName: "ProfilesByStatus", Handler: _Query_ProfilesByStatus_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateset/compliance",
}

func _Query_Profile_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryProfileRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Profile(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.compliance.Query/Profile"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Profile(ctx, request.(*QueryProfileRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_Profiles_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryProfilesRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Profiles(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.compliance.Query/Profiles"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).Profiles(ctx, request.(*QueryProfilesRequest))
	}
	return interceptor(ctx, req, info, handler)
}

func _Query_ProfilesByStatus_Handler(srv interface{}, ctx context.Context, decode func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	req := new(QueryProfilesByStatusRequest)
	if err := decode(req); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProfilesByStatus(ctx, req)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/stateset.compliance.Query/ProfilesByStatus"}
	handler := func(ctx context.Context, request interface{}) (interface{}, error) {
		return srv.(QueryServer).ProfilesByStatus(ctx, request.(*QueryProfilesByStatusRequest))
	}
	return interceptor(ctx, req, info, handler)
}
