// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bootstrap/service.proto

package bootstrap

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	service "github.com/nspcc-dev/neofs-api-go/service"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Node state
type Request_State int32

const (
	// used by default
	Request_Unknown Request_State = 0
	// used to inform that node online
	Request_Online Request_State = 1
	// used to inform that node offline
	Request_Offline Request_State = 2
)

var Request_State_name = map[int32]string{
	0: "Unknown",
	1: "Online",
	2: "Offline",
}

var Request_State_value = map[string]int32{
	"Unknown": 0,
	"Online":  1,
	"Offline": 2,
}

func (x Request_State) String() string {
	return proto.EnumName(Request_State_name, int32(x))
}

func (Request_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21bce759c9d8eb63, []int{0, 0}
}

type Request struct {
	// Type is NodeType, can be InnerRingNode (type=1) or StorageNode (type=2)
	Type NodeType `protobuf:"varint,1,opt,name=type,proto3,customtype=NodeType" json:"type"`
	// Info contains information about node
	Info NodeInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
	// State contains node status
	State Request_State `protobuf:"varint,3,opt,name=state,proto3,enum=bootstrap.Request_State" json:"state,omitempty"`
	// RequestMetaHeader contains information about request meta headers (should be embedded into message)
	service.RequestMetaHeader `protobuf:"bytes,98,opt,name=Meta,proto3,embedded=Meta" json:"Meta"`
	// RequestVerificationHeader is a set of signatures of every NeoFS Node that processed request (should be embedded into message)
	service.RequestVerificationHeader `protobuf:"bytes,99,opt,name=Verify,proto3,embedded=Verify" json:"Verify"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_21bce759c9d8eb63, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetInfo() NodeInfo {
	if m != nil {
		return m.Info
	}
	return NodeInfo{}
}

func (m *Request) GetState() Request_State {
	if m != nil {
		return m.State
	}
	return Request_Unknown
}

func init() {
	proto.RegisterEnum("bootstrap.Request_State", Request_State_name, Request_State_value)
	proto.RegisterType((*Request)(nil), "bootstrap.Request")
}

func init() { proto.RegisterFile("bootstrap/service.proto", fileDescriptor_21bce759c9d8eb63) }

var fileDescriptor_21bce759c9d8eb63 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xee, 0xd4, 0xb4, 0xdd, 0x9d, 0x05, 0x59, 0x66, 0x57, 0x0c, 0x3d, 0xa4, 0xa5, 0xa7, 0x82,
	0x66, 0x02, 0xdd, 0x8b, 0x47, 0x0d, 0x22, 0xee, 0x61, 0x3f, 0x48, 0xfd, 0x00, 0x6f, 0x93, 0xf4,
	0x4d, 0x1d, 0x74, 0xe7, 0x1d, 0x33, 0xd3, 0x4a, 0xff, 0x89, 0xbf, 0xc1, 0xff, 0x21, 0xec, 0x71,
	0x8f, 0xe2, 0xa1, 0x48, 0xfc, 0x23, 0x32, 0xd3, 0xb4, 0x16, 0xf7, 0x94, 0xbc, 0xcf, 0xd7, 0xbc,
	0x1f, 0xf4, 0x71, 0x8e, 0x68, 0x8d, 0xad, 0x84, 0x4e, 0x0c, 0x54, 0x4b, 0x59, 0x00, 0xd7, 0x15,
	0x5a, 0x64, 0x87, 0x3b, 0xa2, 0xcf, 0x1a, 0x26, 0xb9, 0x01, 0x2b, 0x36, 0x74, 0xff, 0x74, 0x8b,
	0x2d, 0xa1, 0x92, 0xe5, 0xaa, 0x41, 0x1f, 0xfd, 0x4b, 0xb3, 0x2b, 0x0d, 0xa6, 0x81, 0xe3, 0xb9,
	0xb4, 0x1f, 0x17, 0x39, 0x2f, 0xf0, 0x26, 0x99, 0xe3, 0x1c, 0x13, 0x0f, 0xe7, 0x8b, 0xd2, 0x57,
	0xbe, 0xf0, 0x7f, 0x1b, 0xf9, 0xe8, 0x47, 0x9b, 0xf6, 0x32, 0xf8, 0xb2, 0x00, 0x63, 0xd9, 0x53,
	0x1a, 0xb8, 0xa4, 0x90, 0x0c, 0xc9, 0xb8, 0x93, 0x86, 0xb7, 0xeb, 0x41, 0xeb, 0xd7, 0x7a, 0x70,
	0x70, 0x89, 0x33, 0x78, 0xb3, 0xd2, 0x50, 0xaf, 0x07, 0x81, 0xfb, 0x66, 0x5e, 0xc5, 0x62, 0x1a,
	0x48, 0x55, 0x62, 0xd8, 0x1e, 0x92, 0xf1, 0xd1, 0xe4, 0x84, 0xef, 0xda, 0xe1, 0xce, 0x70, 0xae,
	0x4a, 0x4c, 0x03, 0x17, 0x91, 0x79, 0x19, 0xe3, 0xb4, 0x63, 0xac, 0xb0, 0x10, 0x3e, 0x18, 0x92,
	0xf1, 0xc3, 0x49, 0xb8, 0xa7, 0x6f, 0xde, 0xe7, 0x53, 0xc7, 0x67, 0x1b, 0x19, 0x7b, 0x46, 0x83,
	0x0b, 0xb0, 0x22, 0xcc, 0x7d, 0x7c, 0x9f, 0x6f, 0x37, 0xd6, 0x88, 0x1d, 0xf7, 0x1a, 0xc4, 0x0c,
	0xaa, 0xf4, 0xc0, 0xbd, 0x72, 0xb7, 0x1e, 0x90, 0xcc, 0x3b, 0xd8, 0x4b, 0xda, 0x7d, 0xe7, 0x17,
	0x15, 0x16, 0xde, 0x3b, 0xfa, 0xdf, 0xeb, 0x59, 0x59, 0x08, 0x2b, 0x51, 0xdd, 0xcb, 0x68, 0xbc,
	0xa3, 0x98, 0x76, 0x7c, 0x3f, 0xec, 0x88, 0xf6, 0xde, 0xaa, 0x4f, 0x0a, 0xbf, 0xaa, 0xe3, 0x16,
	0xa3, 0xb4, 0x7b, 0xa5, 0x3e, 0x4b, 0x05, 0xc7, 0xc4, 0x11, 0x57, 0x65, 0xe9, 0x8b, 0xf6, 0xe4,
	0x39, 0x3d, 0x4c, 0xb7, 0x03, 0xb1, 0x33, 0xda, 0xbb, 0xae, 0xb0, 0x00, 0x63, 0x18, 0xbb, 0x3f,
	0x67, 0xff, 0x74, 0x0f, 0x9b, 0xea, 0x0a, 0xc4, 0xec, 0x42, 0xe8, 0xf4, 0xfd, 0x6d, 0x1d, 0x91,
	0xbb, 0x3a, 0x22, 0x3f, 0xeb, 0x88, 0xfc, 0xae, 0x23, 0xf2, 0xed, 0x4f, 0xd4, 0xfa, 0xf0, 0x64,
	0xef, 0x9c, 0xca, 0xe8, 0xa2, 0x88, 0x67, 0xb0, 0x4c, 0x14, 0x60, 0x69, 0x62, 0xa1, 0x65, 0x3c,
	0xc7, 0x64, 0x17, 0xf6, 0xbd, 0x7d, 0x72, 0x09, 0xf8, 0x6a, 0xca, 0x5f, 0x5c, 0x9f, 0xf3, 0x5d,
	0x37, 0x79, 0xd7, 0x5f, 0xfa, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xc6, 0xb6, 0xb6,
	0x7f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BootstrapClient is the client API for Bootstrap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BootstrapClient interface {
	// Process is method that allows to register node in the network and receive actual netmap
	Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*SpreadMap, error)
}

type bootstrapClient struct {
	cc *grpc.ClientConn
}

func NewBootstrapClient(cc *grpc.ClientConn) BootstrapClient {
	return &bootstrapClient{cc}
}

func (c *bootstrapClient) Process(ctx context.Context, in *Request, opts ...grpc.CallOption) (*SpreadMap, error) {
	out := new(SpreadMap)
	err := c.cc.Invoke(ctx, "/bootstrap.Bootstrap/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstrapServer is the server API for Bootstrap service.
type BootstrapServer interface {
	// Process is method that allows to register node in the network and receive actual netmap
	Process(context.Context, *Request) (*SpreadMap, error)
}

// UnimplementedBootstrapServer can be embedded to have forward compatible implementations.
type UnimplementedBootstrapServer struct {
}

func (*UnimplementedBootstrapServer) Process(ctx context.Context, req *Request) (*SpreadMap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterBootstrapServer(s *grpc.Server, srv BootstrapServer) {
	s.RegisterService(&_Bootstrap_serviceDesc, srv)
}

func _Bootstrap_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bootstrap.Bootstrap/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapServer).Process(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bootstrap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bootstrap.Bootstrap",
	HandlerType: (*BootstrapServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _Bootstrap_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bootstrap/service.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintService(dAtA, i, uint64(size))
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
		i = encodeVarintService(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6
	i--
	dAtA[i] = 0x92
	if m.State != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Info.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintService(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Type != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovService(uint64(m.Type))
	}
	l = m.Info.Size()
	n += 1 + l + sovService(uint64(l))
	if m.State != 0 {
		n += 1 + sovService(uint64(m.State))
	}
	l = m.RequestMetaHeader.Size()
	n += 2 + l + sovService(uint64(l))
	l = m.RequestVerificationHeader.Size()
	n += 2 + l + sovService(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= NodeType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Request_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 98:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestMetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
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
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
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
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
