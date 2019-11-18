// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service/verify_test.proto

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

type TestRequest struct {
	IntField                  int32  `protobuf:"varint,1,opt,name=IntField,proto3" json:"IntField,omitempty"`
	StringField               string `protobuf:"bytes,2,opt,name=StringField,proto3" json:"StringField,omitempty"`
	BytesField                []byte `protobuf:"bytes,3,opt,name=BytesField,proto3" json:"BytesField,omitempty"`
	RequestMetaHeader         `protobuf:"bytes,98,opt,name=Meta,proto3,embedded=Meta" json:"Meta"`
	RequestVerificationHeader `protobuf:"bytes,99,opt,name=Header,proto3,embedded=Header" json:"Header"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1effa83201a30d75, []int{0}
}
func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return m.Size()
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetIntField() int32 {
	if m != nil {
		return m.IntField
	}
	return 0
}

func (m *TestRequest) GetStringField() string {
	if m != nil {
		return m.StringField
	}
	return ""
}

func (m *TestRequest) GetBytesField() []byte {
	if m != nil {
		return m.BytesField
	}
	return nil
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "service.TestRequest")
}

func init() { proto.RegisterFile("service/verify_test.proto", fileDescriptor_1effa83201a30d75) }

var fileDescriptor_1effa83201a30d75 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0xc6, 0xeb, 0x7b, 0xa1, 0x14, 0x97, 0xc9, 0x62, 0x28, 0x19, 0xdc, 0xa8, 0x53, 0x96, 0x26,
	0x12, 0x2c, 0x4c, 0x0c, 0x15, 0x42, 0x30, 0xb0, 0x04, 0xc4, 0xc0, 0x82, 0x12, 0xe7, 0x24, 0x58,
	0xa2, 0x71, 0x89, 0x4f, 0x22, 0xf5, 0x4d, 0x78, 0xa4, 0x8e, 0x1d, 0x99, 0x2a, 0x14, 0x46, 0x5e,
	0x02, 0xd5, 0x36, 0x28, 0xc0, 0x96, 0xef, 0xcf, 0x2f, 0x47, 0xfe, 0xe8, 0x91, 0x86, 0xaa, 0x91,
	0x02, 0xa2, 0x06, 0x2a, 0x99, 0x2f, 0x1f, 0x10, 0x34, 0x86, 0x8b, 0x4a, 0xa1, 0x62, 0x7b, 0x2e,
	0xf2, 0xd8, 0x57, 0x67, 0x0e, 0x98, 0xd8, 0xd0, 0x3b, 0xfc, 0xc9, 0x39, 0x77, 0x5a, 0x48, 0x7c,
	0xac, 0xd3, 0x50, 0xa8, 0x79, 0x54, 0xa8, 0x42, 0x45, 0xc6, 0x4e, 0xeb, 0xdc, 0x28, 0x23, 0xcc,
	0x97, 0xad, 0x4f, 0x3e, 0x08, 0x1d, 0xde, 0x82, 0xc6, 0x18, 0x9e, 0x6b, 0xd0, 0xc8, 0x3c, 0x3a,
	0xb8, 0x2a, 0xf1, 0x42, 0xc2, 0x53, 0x36, 0x22, 0x3e, 0x09, 0x76, 0xe3, 0x6f, 0xcd, 0x7c, 0x3a,
	0xbc, 0xc1, 0x4a, 0x96, 0x85, 0x8d, 0xff, 0xf9, 0x24, 0xd8, 0x8f, 0xbb, 0x16, 0xe3, 0x94, 0xce,
	0x96, 0x08, 0xda, 0x16, 0xfe, 0xfb, 0x24, 0x38, 0x88, 0x3b, 0x0e, 0x3b, 0xa5, 0x3b, 0xd7, 0x80,
	0xc9, 0x28, 0xf5, 0x49, 0x30, 0x3c, 0xf6, 0x42, 0xf7, 0x82, 0xd0, 0x5d, 0xdf, 0x66, 0x97, 0x90,
	0x64, 0x50, 0xcd, 0x06, 0xab, 0xcd, 0xb8, 0xb7, 0xde, 0x8c, 0x49, 0x6c, 0x08, 0x76, 0x4e, 0xfb,
	0x36, 0x19, 0x09, 0xc3, 0x4e, 0x7e, 0xb3, 0x77, 0xdb, 0x11, 0xa4, 0x48, 0x50, 0xaa, 0xf2, 0xcf,
	0x3f, 0x1c, 0x3b, 0x3b, 0x5b, 0xb5, 0x9c, 0xac, 0x5b, 0x4e, 0x5e, 0x5b, 0x4e, 0xde, 0x5a, 0x4e,
	0x5e, 0xde, 0x79, 0xef, 0x3e, 0xe8, 0x4c, 0x56, 0xea, 0x85, 0x10, 0xd3, 0x0c, 0x9a, 0xa8, 0x04,
	0x95, 0xeb, 0xa9, 0x1d, 0xcc, 0xdd, 0x4a, 0xfb, 0x46, 0x9e, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0xdb, 0x5f, 0x7a, 0xb3, 0x01, 0x00, 0x00,
}

func (m *TestRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	{
		size, err := m.RequestVerificationHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVerifyTest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6
	i--
	dAtA[i] = 0x9a
	{
		size, err := m.RequestMetaHeader.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintVerifyTest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6
	i--
	dAtA[i] = 0x92
	if len(m.BytesField) > 0 {
		i -= len(m.BytesField)
		copy(dAtA[i:], m.BytesField)
		i = encodeVarintVerifyTest(dAtA, i, uint64(len(m.BytesField)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StringField) > 0 {
		i -= len(m.StringField)
		copy(dAtA[i:], m.StringField)
		i = encodeVarintVerifyTest(dAtA, i, uint64(len(m.StringField)))
		i--
		dAtA[i] = 0x12
	}
	if m.IntField != 0 {
		i = encodeVarintVerifyTest(dAtA, i, uint64(m.IntField))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVerifyTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerifyTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TestRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IntField != 0 {
		n += 1 + sovVerifyTest(uint64(m.IntField))
	}
	l = len(m.StringField)
	if l > 0 {
		n += 1 + l + sovVerifyTest(uint64(l))
	}
	l = len(m.BytesField)
	if l > 0 {
		n += 1 + l + sovVerifyTest(uint64(l))
	}
	l = m.RequestMetaHeader.Size()
	n += 2 + l + sovVerifyTest(uint64(l))
	l = m.RequestVerificationHeader.Size()
	n += 2 + l + sovVerifyTest(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVerifyTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerifyTest(x uint64) (n int) {
	return sovVerifyTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifyTest
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
			return fmt.Errorf("proto: TestRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntField", wireType)
			}
			m.IntField = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifyTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntField |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifyTest
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
				return ErrInvalidLengthVerifyTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifyTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesField", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifyTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthVerifyTest
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifyTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytesField = append(m.BytesField[:0], dAtA[iNdEx:postIndex]...)
			if m.BytesField == nil {
				m.BytesField = []byte{}
			}
			iNdEx = postIndex
		case 98:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifyTest
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
				return ErrInvalidLengthVerifyTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerifyTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RequestMetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 99:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestVerificationHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifyTest
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
				return ErrInvalidLengthVerifyTest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerifyTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RequestVerificationHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerifyTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifyTest
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifyTest
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
func skipVerifyTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerifyTest
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
					return 0, ErrIntOverflowVerifyTest
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
					return 0, ErrIntOverflowVerifyTest
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
				return 0, ErrInvalidLengthVerifyTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerifyTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerifyTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerifyTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerifyTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerifyTest = fmt.Errorf("proto: unexpected end of group")
)
