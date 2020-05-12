// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/meta.proto

package service

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// RequestMetaHeader contains information about request meta headers
// (should be embedded into message)
type RequestMetaHeader struct {
	// TTL must be larger than zero, it decreased in every NeoFS Node
	TTL uint32 `protobuf:"varint,1,opt,name=TTL,proto3" json:"TTL,omitempty"`
	// Epoch for user can be empty, because node sets epoch to the actual value
	Epoch uint64 `protobuf:"varint,2,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	// Version defines protocol version
	// TODO: not used for now, should be implemented in future
	Version uint32 `protobuf:"varint,3,opt,name=Version,proto3" json:"Version,omitempty"`
	// Raw determines whether the request is raw or not
	Raw                  bool     `protobuf:"varint,4,opt,name=Raw,proto3" json:"Raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMetaHeader) Reset()         { *m = RequestMetaHeader{} }
func (m *RequestMetaHeader) String() string { return proto.CompactTextString(m) }
func (*RequestMetaHeader) ProtoMessage()    {}
func (*RequestMetaHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638867e7b43457c, []int{0}
}
func (m *RequestMetaHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestMetaHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RequestMetaHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMetaHeader.Merge(m, src)
}
func (m *RequestMetaHeader) XXX_Size() int {
	return m.Size()
}
func (m *RequestMetaHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMetaHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMetaHeader proto.InternalMessageInfo

func (m *RequestMetaHeader) GetTTL() uint32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func (m *RequestMetaHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *RequestMetaHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RequestMetaHeader) GetRaw() bool {
	if m != nil {
		return m.Raw
	}
	return false
}

// ResponseMetaHeader contains meta information based on request processing by server
// (should be embedded into message)
type ResponseMetaHeader struct {
	// Current NeoFS epoch on server
	Epoch uint64 `protobuf:"varint,1,opt,name=Epoch,proto3" json:"Epoch,omitempty"`
	// Version defines protocol version
	// TODO: not used for now, should be implemented in future
	Version              uint32   `protobuf:"varint,2,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMetaHeader) Reset()         { *m = ResponseMetaHeader{} }
func (m *ResponseMetaHeader) String() string { return proto.CompactTextString(m) }
func (*ResponseMetaHeader) ProtoMessage()    {}
func (*ResponseMetaHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_a638867e7b43457c, []int{1}
}
func (m *ResponseMetaHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseMetaHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ResponseMetaHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMetaHeader.Merge(m, src)
}
func (m *ResponseMetaHeader) XXX_Size() int {
	return m.Size()
}
func (m *ResponseMetaHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMetaHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMetaHeader proto.InternalMessageInfo

func (m *ResponseMetaHeader) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ResponseMetaHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*RequestMetaHeader)(nil), "service.RequestMetaHeader")
	proto.RegisterType((*ResponseMetaHeader)(nil), "service.ResponseMetaHeader")
}

func init() { proto.RegisterFile("service/meta.proto", fileDescriptor_a638867e7b43457c) }

var fileDescriptor_a638867e7b43457c = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0xcf, 0x4d, 0x2d, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x87, 0x8a, 0x49, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7,
	0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0xe5, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0xe8,
	0x53, 0x4a, 0xe7, 0x12, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xf1, 0x4d, 0x2d, 0x49, 0xf4,
	0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x12, 0x12, 0xe0, 0x62, 0x0e, 0x09, 0xf1, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0d, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0x5d, 0x0b, 0xf2, 0x93, 0x33, 0x24, 0x98,
	0x14, 0x18, 0x35, 0x58, 0x82, 0x20, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0xb0, 0xd4, 0xa2, 0xe2, 0xcc,
	0xfc, 0x3c, 0x09, 0x66, 0xb0, 0x5a, 0x18, 0x17, 0x64, 0x42, 0x50, 0x62, 0xb9, 0x04, 0x8b, 0x02,
	0xa3, 0x06, 0x47, 0x10, 0x88, 0xa9, 0xe4, 0xc2, 0x25, 0x14, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0x8a, 0x64, 0x13, 0xdc, 0x5c, 0x46, 0x1c, 0xe6, 0x32, 0xa1, 0x98, 0xeb, 0x14, 0x7c, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x37, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x26, 0x92, 0x9f, 0xf3, 0x8a, 0x0b, 0x92, 0x93, 0x75, 0x53, 0x52,
	0xcb, 0xf4, 0xf3, 0x52, 0xf3, 0xd3, 0x8a, 0x75, 0x13, 0x0b, 0x32, 0x75, 0xd3, 0xf3, 0xf5, 0xa1,
	0xc1, 0xb3, 0x8a, 0x49, 0xd0, 0x2f, 0x35, 0xdf, 0x2d, 0x58, 0xcf, 0x31, 0xc0, 0x53, 0x2f, 0x18,
	0x22, 0x96, 0xc4, 0x06, 0x0e, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x12, 0x93,
	0x5e, 0x58, 0x01, 0x00, 0x00,
}

func (m *RequestMetaHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestMetaHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestMetaHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Raw {
		i--
		if m.Raw {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Version != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if m.Epoch != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x10
	}
	if m.TTL != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.TTL))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ResponseMetaHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseMetaHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseMetaHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Version != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x10
	}
	if m.Epoch != 0 {
		i = encodeVarintMeta(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMeta(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeta(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestMetaHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TTL != 0 {
		n += 1 + sovMeta(uint64(m.TTL))
	}
	if m.Epoch != 0 {
		n += 1 + sovMeta(uint64(m.Epoch))
	}
	if m.Version != 0 {
		n += 1 + sovMeta(uint64(m.Version))
	}
	if m.Raw {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ResponseMetaHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovMeta(uint64(m.Epoch))
	}
	if m.Version != 0 {
		n += 1 + sovMeta(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMeta(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeta(x uint64) (n int) {
	return sovMeta(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestMetaHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: RequestMetaHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestMetaHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTL", wireType)
			}
			m.TTL = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TTL |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Raw = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseMetaHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeta
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
			return fmt.Errorf("proto: ResponseMetaHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseMetaHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeta
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMeta(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMeta
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMeta(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
					return 0, ErrIntOverflowMeta
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
				return 0, ErrInvalidLengthMeta
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMeta
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMeta
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMeta        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeta          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMeta = fmt.Errorf("proto: unexpected end of group")
)
