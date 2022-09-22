// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: purchaseorder/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// EventCreatePurchaseOrder is an event emitted when an purchaseorder is created.
type EventCreatePurchaseOrder struct {
	// purchaseorder_id is the unique ID of purchaseorder
	PurchaseorderId string `protobuf:"bytes,1,opt,name=purchaseorder_id,json=purchaseorderId,proto3" json:"purchaseorder_id,omitempty" yaml:"purchaseorder_id"`
	// creator is the creator of the purchaseorder
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EventCreatePurchaseOrder) Reset()         { *m = EventCreatePurchaseOrder{} }
func (m *EventCreatePurchaseOrder) String() string { return proto.CompactTextString(m) }
func (*EventCreatePurchaseOrder) ProtoMessage()    {}
func (*EventCreatePurchaseOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_459c386467b7e936, []int{0}
}
func (m *EventCreatePurchaseOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreatePurchaseOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreatePurchaseOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreatePurchaseOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreatePurchaseOrder.Merge(m, src)
}
func (m *EventCreatePurchaseOrder) XXX_Size() int {
	return m.Size()
}
func (m *EventCreatePurchaseOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreatePurchaseOrder.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreatePurchaseOrder proto.InternalMessageInfo

func (m *EventCreatePurchaseOrder) GetPurchaseorderId() string {
	if m != nil {
		return m.PurchaseorderId
	}
	return ""
}

func (m *EventCreatePurchaseOrder) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// EventCompleted is an event emitted when an purchaseorder is completed.
type EventCompleted struct {
	// agreement_id is the unique ID of purchaseorder
	PurchaseorderId string `protobuf:"bytes,1,opt,name=purchaseorder_id,json=purchaseorderId,proto3" json:"purchaseorder_id,omitempty" yaml:"purchaseorder_id"`
	// creator is the creator of the purchaseorder
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EventCompleted) Reset()         { *m = EventCompleted{} }
func (m *EventCompleted) String() string { return proto.CompactTextString(m) }
func (*EventCompleted) ProtoMessage()    {}
func (*EventCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_459c386467b7e936, []int{1}
}
func (m *EventCompleted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCompleted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCompleted.Merge(m, src)
}
func (m *EventCompleted) XXX_Size() int {
	return m.Size()
}
func (m *EventCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_EventCompleted proto.InternalMessageInfo

func (m *EventCompleted) GetPurchaseorderId() string {
	if m != nil {
		return m.PurchaseorderId
	}
	return ""
}

func (m *EventCompleted) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// EventCancelled is an event emitted when an purchaseorder is cancelled.
type EventCancelled struct {
	// agreement_id is the unique ID of purchaseorder
	PurchaseorderId string `protobuf:"bytes,1,opt,name=purchaseorder_id,json=purchaseorderId,proto3" json:"purchaseorder_id,omitempty" yaml:"purchaseorder_id"`
	// creator is the creator of the purchaseorder
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EventCancelled) Reset()         { *m = EventCancelled{} }
func (m *EventCancelled) String() string { return proto.CompactTextString(m) }
func (*EventCancelled) ProtoMessage()    {}
func (*EventCancelled) Descriptor() ([]byte, []int) {
	return fileDescriptor_459c386467b7e936, []int{2}
}
func (m *EventCancelled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCancelled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCancelled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCancelled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCancelled.Merge(m, src)
}
func (m *EventCancelled) XXX_Size() int {
	return m.Size()
}
func (m *EventCancelled) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCancelled.DiscardUnknown(m)
}

var xxx_messageInfo_EventCancelled proto.InternalMessageInfo

func (m *EventCancelled) GetPurchaseorderId() string {
	if m != nil {
		return m.PurchaseorderId
	}
	return ""
}

func (m *EventCancelled) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// EventFinanced is an event emitted when an purchaseorder is financed.
type EventFinanced struct {
	// purchaseorder_id is the unique ID of purchaseorder
	PurchaseorderId string `protobuf:"bytes,1,opt,name=purchaseorder_id,json=purchaseorderId,proto3" json:"purchaseorder_id,omitempty" yaml:"purchaseorder_id"`
	// creator is the creator of the purchaseorder
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *EventFinanced) Reset()         { *m = EventFinanced{} }
func (m *EventFinanced) String() string { return proto.CompactTextString(m) }
func (*EventFinanced) ProtoMessage()    {}
func (*EventFinanced) Descriptor() ([]byte, []int) {
	return fileDescriptor_459c386467b7e936, []int{3}
}
func (m *EventFinanced) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFinanced) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFinanced.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFinanced) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFinanced.Merge(m, src)
}
func (m *EventFinanced) XXX_Size() int {
	return m.Size()
}
func (m *EventFinanced) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFinanced.DiscardUnknown(m)
}

var xxx_messageInfo_EventFinanced proto.InternalMessageInfo

func (m *EventFinanced) GetPurchaseorderId() string {
	if m != nil {
		return m.PurchaseorderId
	}
	return ""
}

func (m *EventFinanced) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreatePurchaseOrder)(nil), "stateset.core.purchaseorder.EventCreatePurchaseOrder")
	proto.RegisterType((*EventCompleted)(nil), "stateset.core.purchaseorder.EventCompleted")
	proto.RegisterType((*EventCancelled)(nil), "stateset.core.purchaseorder.EventCancelled")
	proto.RegisterType((*EventFinanced)(nil), "stateset.core.purchaseorder.EventFinanced")
}

func init() { proto.RegisterFile("purchaseorder/events.proto", fileDescriptor_459c386467b7e936) }

var fileDescriptor_459c386467b7e936 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0x2d, 0x4a,
	0xce, 0x48, 0x2c, 0x4e, 0xcd, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0x29,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2e, 0x2e, 0x49, 0x2c, 0x49, 0x2d, 0x4e, 0x2d,
	0xd1, 0x4b, 0xce, 0x2f, 0x4a, 0xd5, 0x43, 0x51, 0x29, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56,
	0xa7, 0x0f, 0x62, 0x41, 0xb4, 0x28, 0xd5, 0x70, 0x49, 0xb8, 0x82, 0x8c, 0x70, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x0d, 0x80, 0xea, 0xf0, 0x07, 0xe9, 0x10, 0x72, 0xe3, 0x12, 0x40, 0x31, 0x22, 0x3e,
	0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xd3, 0x49, 0xfa, 0xd3, 0x3d, 0x79, 0xf1, 0xca, 0xc4,
	0xdc, 0x1c, 0x2b, 0x25, 0x74, 0x15, 0x4a, 0x41, 0xfc, 0x28, 0x42, 0x9e, 0x29, 0x42, 0x12, 0x5c,
	0xec, 0xc9, 0x20, 0xe3, 0xf3, 0x8b, 0x24, 0x98, 0x40, 0xda, 0x83, 0x60, 0x5c, 0xa5, 0x22, 0x2e,
	0x3e, 0x88, 0xed, 0xf9, 0xb9, 0x05, 0x39, 0xa9, 0x25, 0xa9, 0x29, 0xf4, 0xb4, 0x33, 0x31, 0x2f,
	0x39, 0x35, 0x27, 0x87, 0x2e, 0x76, 0x16, 0x72, 0xf1, 0x82, 0xed, 0x74, 0xcb, 0xcc, 0x03, 0xd9,
	0x4a, 0x07, 0x2b, 0x9d, 0x3c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a,
	0x2f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x96, 0x60, 0xf4, 0x41,
	0x09, 0x46, 0xbf, 0x42, 0x1f, 0x35, 0x71, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x53,
	0x8a, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x8a, 0xb6, 0x12, 0x7a, 0x02, 0x00, 0x00,
}

func (m *EventCreatePurchaseOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreatePurchaseOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreatePurchaseOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PurchaseorderId) > 0 {
		i -= len(m.PurchaseorderId)
		copy(dAtA[i:], m.PurchaseorderId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PurchaseorderId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCompleted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCompleted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCompleted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PurchaseorderId) > 0 {
		i -= len(m.PurchaseorderId)
		copy(dAtA[i:], m.PurchaseorderId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PurchaseorderId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCancelled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCancelled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCancelled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PurchaseorderId) > 0 {
		i -= len(m.PurchaseorderId)
		copy(dAtA[i:], m.PurchaseorderId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PurchaseorderId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventFinanced) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFinanced) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFinanced) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PurchaseorderId) > 0 {
		i -= len(m.PurchaseorderId)
		copy(dAtA[i:], m.PurchaseorderId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PurchaseorderId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreatePurchaseOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PurchaseorderId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCompleted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PurchaseorderId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCancelled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PurchaseorderId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventFinanced) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PurchaseorderId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreatePurchaseOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreatePurchaseOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreatePurchaseOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseorderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PurchaseorderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCompleted) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCompleted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCompleted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseorderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PurchaseorderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCancelled) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCancelled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCancelled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseorderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PurchaseorderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventFinanced) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventFinanced: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFinanced: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PurchaseorderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PurchaseorderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
