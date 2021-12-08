// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: purchaseorder/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryGetPurchaseorderRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetPurchaseorderRequest) Reset()         { *m = QueryGetPurchaseorderRequest{} }
func (m *QueryGetPurchaseorderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetPurchaseorderRequest) ProtoMessage()    {}
func (*QueryGetPurchaseorderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e155db192b5ef7f, []int{0}
}
func (m *QueryGetPurchaseorderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPurchaseorderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPurchaseorderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPurchaseorderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPurchaseorderRequest.Merge(m, src)
}
func (m *QueryGetPurchaseorderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPurchaseorderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPurchaseorderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPurchaseorderRequest proto.InternalMessageInfo

func (m *QueryGetPurchaseorderRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetPurchaseorderResponse struct {
	Purchaseorder Purchaseorder `protobuf:"bytes,1,opt,name=Purchaseorder,proto3" json:"Purchaseorder"`
}

func (m *QueryGetPurchaseorderResponse) Reset()         { *m = QueryGetPurchaseorderResponse{} }
func (m *QueryGetPurchaseorderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetPurchaseorderResponse) ProtoMessage()    {}
func (*QueryGetPurchaseorderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e155db192b5ef7f, []int{1}
}
func (m *QueryGetPurchaseorderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPurchaseorderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPurchaseorderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPurchaseorderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPurchaseorderResponse.Merge(m, src)
}
func (m *QueryGetPurchaseorderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPurchaseorderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPurchaseorderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPurchaseorderResponse proto.InternalMessageInfo

func (m *QueryGetPurchaseorderResponse) GetPurchaseorder() Purchaseorder {
	if m != nil {
		return m.Purchaseorder
	}
	return Purchaseorder{}
}

type QueryAllPurchaseorderRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPurchaseorderRequest) Reset()         { *m = QueryAllPurchaseorderRequest{} }
func (m *QueryAllPurchaseorderRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPurchaseorderRequest) ProtoMessage()    {}
func (*QueryAllPurchaseorderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e155db192b5ef7f, []int{2}
}
func (m *QueryAllPurchaseorderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPurchaseorderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPurchaseorderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPurchaseorderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPurchaseorderRequest.Merge(m, src)
}
func (m *QueryAllPurchaseorderRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPurchaseorderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPurchaseorderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPurchaseorderRequest proto.InternalMessageInfo

func (m *QueryAllPurchaseorderRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPurchaseorderResponse struct {
	Purchaseorder []Purchaseorder     `protobuf:"bytes,1,rep,name=Purchaseorder,proto3" json:"Purchaseorder"`
	Pagination    *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPurchaseorderResponse) Reset()         { *m = QueryAllPurchaseorderResponse{} }
func (m *QueryAllPurchaseorderResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPurchaseorderResponse) ProtoMessage()    {}
func (*QueryAllPurchaseorderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e155db192b5ef7f, []int{3}
}
func (m *QueryAllPurchaseorderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPurchaseorderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPurchaseorderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPurchaseorderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPurchaseorderResponse.Merge(m, src)
}
func (m *QueryAllPurchaseorderResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPurchaseorderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPurchaseorderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPurchaseorderResponse proto.InternalMessageInfo

func (m *QueryAllPurchaseorderResponse) GetPurchaseorder() []Purchaseorder {
	if m != nil {
		return m.Purchaseorder
	}
	return nil
}

func (m *QueryAllPurchaseorderResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetPurchaseorderRequest)(nil), "stateset.core.purchaseorder.QueryGetPurchaseorderRequest")
	proto.RegisterType((*QueryGetPurchaseorderResponse)(nil), "stateset.core.purchaseorder.QueryGetPurchaseorderResponse")
	proto.RegisterType((*QueryAllPurchaseorderRequest)(nil), "stateset.core.purchaseorder.QueryAllPurchaseorderRequest")
	proto.RegisterType((*QueryAllPurchaseorderResponse)(nil), "stateset.core.purchaseorder.QueryAllPurchaseorderResponse")
}

func init() { proto.RegisterFile("purchaseorder/query.proto", fileDescriptor_3e155db192b5ef7f) }

var fileDescriptor_3e155db192b5ef7f = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbf, 0x4f, 0xdb, 0x40,
	0x1c, 0xc5, 0x7d, 0x4e, 0xda, 0xe1, 0xaa, 0x56, 0xd5, 0xa9, 0x43, 0x9b, 0xa6, 0x6e, 0xeb, 0xa1,
	0xad, 0xa2, 0xea, 0x4e, 0x49, 0x87, 0xaa, 0xdd, 0x92, 0xa1, 0xe9, 0x98, 0x7a, 0xe8, 0xc0, 0x76,
	0xb6, 0x0f, 0xc7, 0x92, 0xe3, 0x73, 0x7c, 0x67, 0x20, 0x42, 0x2c, 0x6c, 0x6c, 0x48, 0xfc, 0x3d,
	0x08, 0xc6, 0x8c, 0x91, 0x58, 0x98, 0x10, 0x4a, 0xf8, 0x43, 0x90, 0x7f, 0x20, 0x72, 0x01, 0x9b,
	0x08, 0xb1, 0x9d, 0xed, 0xef, 0x7b, 0xf7, 0x3e, 0xef, 0xce, 0xf0, 0x5d, 0x94, 0xc4, 0xce, 0x90,
	0x0a, 0xc6, 0x63, 0x97, 0xc5, 0x64, 0x9c, 0xb0, 0x78, 0x82, 0xa3, 0x98, 0x4b, 0x8e, 0xde, 0x0b,
	0x49, 0x25, 0x13, 0x4c, 0x62, 0x87, 0xc7, 0x0c, 0x2b, 0x83, 0x8d, 0xa6, 0xc7, 0xb9, 0x17, 0x30,
	0x42, 0x23, 0x9f, 0xd0, 0x30, 0xe4, 0x92, 0x4a, 0x9f, 0x87, 0x22, 0x97, 0x36, 0x5a, 0x0e, 0x17,
	0x23, 0x2e, 0x88, 0x4d, 0x05, 0xcb, 0x3d, 0xc9, 0x56, 0xdb, 0x66, 0x92, 0xb6, 0x49, 0x44, 0x3d,
	0x3f, 0xcc, 0x86, 0x8b, 0xd9, 0xcf, 0x6a, 0x02, 0xe5, 0xa9, 0x18, 0x79, 0xe3, 0x71, 0x8f, 0x67,
	0x4b, 0x92, 0xae, 0xf2, 0xb7, 0x26, 0x86, 0xcd, 0x7f, 0xa9, 0x75, 0x9f, 0xc9, 0xc1, 0xb2, 0xc8,
	0x62, 0xe3, 0x84, 0x09, 0x89, 0x5e, 0x41, 0xdd, 0x77, 0xdf, 0x82, 0x4f, 0xe0, 0x5b, 0xdd, 0xd2,
	0x7d, 0xd7, 0xdc, 0x86, 0x1f, 0x4a, 0xe6, 0x45, 0xc4, 0x43, 0xc1, 0xd0, 0x7f, 0xf8, 0x52, 0xf9,
	0x90, 0x69, 0x5f, 0x74, 0x5a, 0xb8, 0xa2, 0x08, 0xac, 0x28, 0x7a, 0xf5, 0xe9, 0xc5, 0x47, 0xcd,
	0x52, 0x6d, 0xcc, 0xcd, 0x22, 0x68, 0x37, 0x08, 0xee, 0x0d, 0xfa, 0x07, 0xc2, 0xdb, 0x56, 0x8a,
	0x4d, 0xbf, 0xe0, 0xbc, 0x42, 0x9c, 0x56, 0x88, 0xf3, 0x63, 0x29, 0x2a, 0xc4, 0x03, 0xea, 0xb1,
	0x42, 0x6b, 0x2d, 0x29, 0xcd, 0x53, 0x50, 0x10, 0xde, 0xdd, 0xa8, 0x9c, 0xb0, 0xf6, 0x04, 0x84,
	0xa8, 0xaf, 0x10, 0xe8, 0x19, 0xc1, 0xd7, 0x07, 0x09, 0xf2, 0x50, 0xcb, 0x08, 0x9d, 0x83, 0x1a,
	0x7c, 0x96, 0x21, 0xa0, 0x13, 0xb0, 0x92, 0x15, 0xfd, 0xaa, 0x4c, 0x59, 0x75, 0x15, 0x1a, 0xbf,
	0x1f, 0x23, 0xcd, 0xe3, 0x99, 0x3f, 0xf7, 0xcf, 0xae, 0x8e, 0xf4, 0x36, 0x22, 0xe4, 0xc6, 0x83,
	0xa4, 0x1e, 0xa4, 0xe2, 0xda, 0x92, 0x5d, 0xdf, 0xdd, 0x43, 0xc7, 0x00, 0xbe, 0x56, 0x2c, 0xbb,
	0x41, 0xb0, 0x0e, 0x44, 0xc9, 0x35, 0x59, 0x07, 0xa2, 0xec, 0xe0, 0xcd, 0x4e, 0x06, 0xf1, 0x1d,
	0xb5, 0xd6, 0x87, 0xe8, 0xfd, 0x9d, 0xce, 0x0d, 0x30, 0x9b, 0x1b, 0xe0, 0x72, 0x6e, 0x80, 0xc3,
	0x85, 0xa1, 0xcd, 0x16, 0x86, 0x76, 0xbe, 0x30, 0xb4, 0x0d, 0xec, 0xf9, 0x72, 0x98, 0xd8, 0xd8,
	0xe1, 0xa3, 0x15, 0xbf, 0x9d, 0x15, 0x47, 0x39, 0x89, 0x98, 0xb0, 0x9f, 0x67, 0x3f, 0xec, 0x8f,
	0xeb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x28, 0xfd, 0x1e, 0x6d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a purchaseorder by id.
	Purchaseorder(ctx context.Context, in *QueryGetPurchaseorderRequest, opts ...grpc.CallOption) (*QueryGetPurchaseorderResponse, error)
	// Queries a list of purchaseorder items.
	PurchaseorderAll(ctx context.Context, in *QueryAllPurchaseorderRequest, opts ...grpc.CallOption) (*QueryAllPurchaseorderResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Purchaseorder(ctx context.Context, in *QueryGetPurchaseorderRequest, opts ...grpc.CallOption) (*QueryGetPurchaseorderResponse, error) {
	out := new(QueryGetPurchaseorderResponse)
	err := c.cc.Invoke(ctx, "/stateset.core.purchaseorder.Query/Purchaseorder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PurchaseorderAll(ctx context.Context, in *QueryAllPurchaseorderRequest, opts ...grpc.CallOption) (*QueryAllPurchaseorderResponse, error) {
	out := new(QueryAllPurchaseorderResponse)
	err := c.cc.Invoke(ctx, "/stateset.core.purchaseorder.Query/PurchaseorderAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a purchaseorder by id.
	Purchaseorder(context.Context, *QueryGetPurchaseorderRequest) (*QueryGetPurchaseorderResponse, error)
	// Queries a list of purchaseorder items.
	PurchaseorderAll(context.Context, *QueryAllPurchaseorderRequest) (*QueryAllPurchaseorderResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Purchaseorder(ctx context.Context, req *QueryGetPurchaseorderRequest) (*QueryGetPurchaseorderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purchaseorder not implemented")
}
func (*UnimplementedQueryServer) PurchaseorderAll(ctx context.Context, req *QueryAllPurchaseorderRequest) (*QueryAllPurchaseorderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseorderAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Purchaseorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPurchaseorderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Purchaseorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.core.purchaseorder.Query/Purchaseorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Purchaseorder(ctx, req.(*QueryGetPurchaseorderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PurchaseorderAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPurchaseorderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PurchaseorderAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateset.core.purchaseorder.Query/PurchaseorderAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PurchaseorderAll(ctx, req.(*QueryAllPurchaseorderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stateset.core.purchaseorder.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Purchaseorder",
			Handler:    _Query_Purchaseorder_Handler,
		},
		{
			MethodName: "PurchaseorderAll",
			Handler:    _Query_PurchaseorderAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "purchaseorder/query.proto",
}

func (m *QueryGetPurchaseorderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPurchaseorderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPurchaseorderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetPurchaseorderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPurchaseorderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPurchaseorderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Purchaseorder.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllPurchaseorderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPurchaseorderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPurchaseorderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPurchaseorderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPurchaseorderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPurchaseorderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Purchaseorder) > 0 {
		for iNdEx := len(m.Purchaseorder) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Purchaseorder[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetPurchaseorderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetPurchaseorderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Purchaseorder.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllPurchaseorderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllPurchaseorderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Purchaseorder) > 0 {
		for _, e := range m.Purchaseorder {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetPurchaseorderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetPurchaseorderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPurchaseorderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetPurchaseorderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetPurchaseorderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPurchaseorderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Purchaseorder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Purchaseorder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllPurchaseorderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllPurchaseorderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPurchaseorderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllPurchaseorderResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllPurchaseorderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPurchaseorderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Purchaseorder", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Purchaseorder = append(m.Purchaseorder, Purchaseorder{})
			if err := m.Purchaseorder[len(m.Purchaseorder)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
