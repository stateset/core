// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: purchaseorder/purchaseorder.proto

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

type Purchaseorder struct {
	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Did    string `protobuf:"bytes,2,opt,name=did,proto3" json:"did,omitempty"`
	Uri    string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Amount string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	State  string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *Purchaseorder) Reset()         { *m = Purchaseorder{} }
func (m *Purchaseorder) String() string { return proto.CompactTextString(m) }
func (*Purchaseorder) ProtoMessage()    {}
func (*Purchaseorder) Descriptor() ([]byte, []int) {
	return fileDescriptor_68cdaff5eff29fb6, []int{0}
}
func (m *Purchaseorder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Purchaseorder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Purchaseorder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Purchaseorder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Purchaseorder.Merge(m, src)
}
func (m *Purchaseorder) XXX_Size() int {
	return m.Size()
}
func (m *Purchaseorder) XXX_DiscardUnknown() {
	xxx_messageInfo_Purchaseorder.DiscardUnknown(m)
}

var xxx_messageInfo_Purchaseorder proto.InternalMessageInfo

func (m *Purchaseorder) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Purchaseorder) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *Purchaseorder) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Purchaseorder) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Purchaseorder) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*Purchaseorder)(nil), "stateset.core.purchaseorder.Purchaseorder")
}

func init() { proto.RegisterFile("purchaseorder/purchaseorder.proto", fileDescriptor_68cdaff5eff29fb6) }

var fileDescriptor_68cdaff5eff29fb6 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0x2d, 0x4a,
	0xce, 0x48, 0x2c, 0x4e, 0xcd, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x47, 0xe1, 0xe9, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x49, 0x17, 0x97, 0x24, 0x96, 0xa4, 0x16, 0xa7, 0x96, 0xe8, 0x25, 0xe7, 0x17,
	0xa5, 0xea, 0xa1, 0x28, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd3, 0x07, 0xb1, 0x20,
	0x5a, 0x94, 0x0a, 0xb9, 0x78, 0x03, 0x90, 0x95, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x70, 0x31, 0xa7, 0x64, 0xa6, 0x48,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x20, 0x91, 0xd2, 0xa2, 0x4c, 0x09, 0x66, 0x88,
	0x48, 0x69, 0x51, 0xa6, 0x90, 0x18, 0x17, 0x5b, 0x62, 0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x0b,
	0x58, 0x10, 0xca, 0x13, 0x12, 0xe1, 0x62, 0x05, 0xbb, 0x48, 0x82, 0x15, 0x2c, 0x0c, 0xe1, 0x38,
	0x79, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x5e, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xcc, 0x2b, 0xfa, 0x20, 0xaf, 0xe8, 0x57, 0xa0,
	0xfa, 0x57, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x07, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x63, 0x14, 0x37, 0xa0, 0x1b, 0x01, 0x00, 0x00,
}

func (m *Purchaseorder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Purchaseorder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Purchaseorder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintPurchaseorder(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintPurchaseorder(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintPurchaseorder(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintPurchaseorder(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPurchaseorder(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPurchaseorder(dAtA []byte, offset int, v uint64) int {
	offset -= sovPurchaseorder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Purchaseorder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPurchaseorder(uint64(m.Id))
	}
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovPurchaseorder(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovPurchaseorder(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovPurchaseorder(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovPurchaseorder(uint64(l))
	}
	return n
}

func sovPurchaseorder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPurchaseorder(x uint64) (n int) {
	return sovPurchaseorder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Purchaseorder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPurchaseorder
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
			return fmt.Errorf("proto: Purchaseorder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Purchaseorder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPurchaseorder
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPurchaseorder
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
				return ErrInvalidLengthPurchaseorder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPurchaseorder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPurchaseorder
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
				return ErrInvalidLengthPurchaseorder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPurchaseorder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPurchaseorder
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
				return ErrInvalidLengthPurchaseorder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPurchaseorder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPurchaseorder
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
				return ErrInvalidLengthPurchaseorder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPurchaseorder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPurchaseorder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPurchaseorder
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
func skipPurchaseorder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPurchaseorder
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
					return 0, ErrIntOverflowPurchaseorder
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
					return 0, ErrIntOverflowPurchaseorder
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
				return 0, ErrInvalidLengthPurchaseorder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPurchaseorder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPurchaseorder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPurchaseorder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPurchaseorder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPurchaseorder = fmt.Errorf("proto: unexpected end of group")
)
