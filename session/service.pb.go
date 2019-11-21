// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: session/service.proto

package session

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type CreateRequest struct {
	// Message should be one of
	//
	// Types that are valid to be assigned to Message:
	//	*CreateRequest_Init
	//	*CreateRequest_Signed
	Message              isCreateRequest_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b329bee0fd1148e0, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

type isCreateRequest_Message interface {
	isCreateRequest_Message()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CreateRequest_Init struct {
	Init *Token `protobuf:"bytes,1,opt,name=Init,proto3,oneof" json:"Init,omitempty"`
}
type CreateRequest_Signed struct {
	Signed *Token `protobuf:"bytes,2,opt,name=Signed,proto3,oneof" json:"Signed,omitempty"`
}

func (*CreateRequest_Init) isCreateRequest_Message()   {}
func (*CreateRequest_Signed) isCreateRequest_Message() {}

func (m *CreateRequest) GetMessage() isCreateRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *CreateRequest) GetInit() *Token {
	if x, ok := m.GetMessage().(*CreateRequest_Init); ok {
		return x.Init
	}
	return nil
}

func (m *CreateRequest) GetSigned() *Token {
	if x, ok := m.GetMessage().(*CreateRequest_Signed); ok {
		return x.Signed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CreateRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CreateRequest_Init)(nil),
		(*CreateRequest_Signed)(nil),
	}
}

type CreateResponse struct {
	// Types that are valid to be assigned to Message:
	//	*CreateResponse_Unsigned
	//	*CreateResponse_Result
	Message              isCreateResponse_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b329bee0fd1148e0, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return m.Size()
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type isCreateResponse_Message interface {
	isCreateResponse_Message()
	MarshalTo([]byte) (int, error)
	Size() int
}

type CreateResponse_Unsigned struct {
	Unsigned *Token `protobuf:"bytes,1,opt,name=Unsigned,proto3,oneof" json:"Unsigned,omitempty"`
}
type CreateResponse_Result struct {
	Result *Token `protobuf:"bytes,2,opt,name=Result,proto3,oneof" json:"Result,omitempty"`
}

func (*CreateResponse_Unsigned) isCreateResponse_Message() {}
func (*CreateResponse_Result) isCreateResponse_Message()   {}

func (m *CreateResponse) GetMessage() isCreateResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *CreateResponse) GetUnsigned() *Token {
	if x, ok := m.GetMessage().(*CreateResponse_Unsigned); ok {
		return x.Unsigned
	}
	return nil
}

func (m *CreateResponse) GetResult() *Token {
	if x, ok := m.GetMessage().(*CreateResponse_Result); ok {
		return x.Result
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CreateResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CreateResponse_Unsigned)(nil),
		(*CreateResponse_Result)(nil),
	}
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "session.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "session.CreateResponse")
}

func init() { proto.RegisterFile("session/service.proto", fileDescriptor_b329bee0fd1148e0) }

var fileDescriptor_b329bee0fd1148e0 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0x6b, 0x84, 0x12, 0x30, 0xa2, 0x83, 0x11, 0x50, 0x65, 0xb0, 0x50, 0xc5, 0x90, 0x81,
	0x24, 0xa8, 0xcc, 0x30, 0x94, 0xa5, 0x0c, 0x2c, 0x29, 0x2c, 0x6c, 0x4d, 0x7a, 0x35, 0xe6, 0xc3,
	0x0e, 0x39, 0xa7, 0x12, 0x6f, 0xc2, 0x23, 0x31, 0x32, 0x32, 0xa2, 0xf0, 0x22, 0x08, 0x3b, 0xad,
	0x82, 0x50, 0x36, 0xff, 0x3f, 0x7c, 0x3f, 0xfb, 0xe8, 0x3e, 0x02, 0xa2, 0xd4, 0x2a, 0x41, 0x28,
	0x97, 0x32, 0x87, 0xb8, 0x28, 0xb5, 0xd1, 0xcc, 0x6f, 0xec, 0x60, 0x6f, 0x95, 0x9b, 0xd7, 0x02,
	0xd0, 0xa5, 0x41, 0x24, 0xa4, 0xb9, 0xaf, 0xb2, 0x38, 0xd7, 0xcf, 0x89, 0xd0, 0x42, 0x27, 0xd6,
	0xce, 0xaa, 0x85, 0x55, 0x56, 0xd8, 0x93, 0xab, 0x0f, 0x1f, 0xe8, 0xee, 0x65, 0x09, 0x33, 0x03,
	0x29, 0xbc, 0x54, 0x80, 0x86, 0x1d, 0xd3, 0xcd, 0x2b, 0x25, 0xcd, 0x80, 0x1c, 0x91, 0x70, 0x67,
	0xd4, 0x8f, 0x1b, 0x46, 0x7c, 0xa3, 0x1f, 0x41, 0x4d, 0x7a, 0xa9, 0x4d, 0x59, 0x48, 0xbd, 0xa9,
	0x14, 0x0a, 0xe6, 0x83, 0x8d, 0x8e, 0x5e, 0x93, 0x8f, 0xb7, 0xa9, 0x7f, 0x0d, 0x88, 0x33, 0x01,
	0x43, 0xa4, 0xfd, 0x15, 0x0b, 0x0b, 0xad, 0x10, 0xd8, 0x09, 0xdd, 0xba, 0x55, 0xe8, 0x06, 0x75,
	0x01, 0xd7, 0x8d, 0x5f, 0x68, 0x0a, 0x58, 0x3d, 0x99, 0x6e, 0xa8, 0xcb, 0x5b, 0xd0, 0xd1, 0x84,
	0xfa, 0x53, 0xd7, 0x62, 0xe7, 0xd4, 0x73, 0x7c, 0x76, 0xb0, 0xbe, 0xf9, 0xe7, 0xf3, 0xc1, 0xe1,
	0x3f, 0xdf, 0x3d, 0x34, 0x24, 0xa7, 0x64, 0x7c, 0xf1, 0x5e, 0x73, 0xf2, 0x51, 0x73, 0xf2, 0x59,
	0x73, 0xf2, 0x55, 0x73, 0xf2, 0xf6, 0xcd, 0x7b, 0x77, 0x61, 0x6b, 0xdf, 0x0a, 0x8b, 0x3c, 0x8f,
	0xe6, 0xb0, 0x4c, 0x14, 0xe8, 0x05, 0x46, 0x6e, 0xdb, 0xcd, 0xc8, 0xcc, 0xb3, 0xf2, 0xec, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x74, 0x3d, 0x2a, 0x06, 0xd7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SessionClient interface {
	// Create is a method that used to open a trusted session to manipulate
	// an object. In order to put or delete object client have to obtain session
	// token with trusted node. Trusted node will modify client's object
	// (add missing headers, checksums, homomorphic hash) and sign id with
	// session key. Session is established during 4-step handshake in one gRPC stream
	//
	// - First client stream message SHOULD BE type of `CreateRequest_Init`.
	// - First server stream message SHOULD BE type of `CreateResponse_Unsigned`.
	// - Second client stream message SHOULD BE type of `CreateRequest_Signed`.
	// - Second server stream message SHOULD BE type of `CreateResponse_Result`.
	Create(ctx context.Context, opts ...grpc.CallOption) (Session_CreateClient, error)
}

type sessionClient struct {
	cc *grpc.ClientConn
}

func NewSessionClient(cc *grpc.ClientConn) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) Create(ctx context.Context, opts ...grpc.CallOption) (Session_CreateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Session_serviceDesc.Streams[0], "/session.Session/Create", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionCreateClient{stream}
	return x, nil
}

type Session_CreateClient interface {
	Send(*CreateRequest) error
	Recv() (*CreateResponse, error)
	grpc.ClientStream
}

type sessionCreateClient struct {
	grpc.ClientStream
}

func (x *sessionCreateClient) Send(m *CreateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionCreateClient) Recv() (*CreateResponse, error) {
	m := new(CreateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SessionServer is the server API for Session service.
type SessionServer interface {
	// Create is a method that used to open a trusted session to manipulate
	// an object. In order to put or delete object client have to obtain session
	// token with trusted node. Trusted node will modify client's object
	// (add missing headers, checksums, homomorphic hash) and sign id with
	// session key. Session is established during 4-step handshake in one gRPC stream
	//
	// - First client stream message SHOULD BE type of `CreateRequest_Init`.
	// - First server stream message SHOULD BE type of `CreateResponse_Unsigned`.
	// - Second client stream message SHOULD BE type of `CreateRequest_Signed`.
	// - Second server stream message SHOULD BE type of `CreateResponse_Result`.
	Create(Session_CreateServer) error
}

// UnimplementedSessionServer can be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (*UnimplementedSessionServer) Create(srv Session_CreateServer) error {
	return status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func RegisterSessionServer(s *grpc.Server, srv SessionServer) {
	s.RegisterService(&_Session_serviceDesc, srv)
}

func _Session_Create_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionServer).Create(&sessionCreateServer{stream})
}

type Session_CreateServer interface {
	Send(*CreateResponse) error
	Recv() (*CreateRequest, error)
	grpc.ServerStream
}

type sessionCreateServer struct {
	grpc.ServerStream
}

func (x *sessionCreateServer) Send(m *CreateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionCreateServer) Recv() (*CreateRequest, error) {
	m := new(CreateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Session_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.Session",
	HandlerType: (*SessionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Create",
			Handler:       _Session_Create_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "session/service.proto",
}

func (m *CreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Message != nil {
		{
			size := m.Message.Size()
			i -= size
			if _, err := m.Message.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateRequest_Init) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRequest_Init) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Init != nil {
		{
			size, err := m.Init.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *CreateRequest_Signed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateRequest_Signed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Signed != nil {
		{
			size, err := m.Signed.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *CreateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Message != nil {
		{
			size := m.Message.Size()
			i -= size
			if _, err := m.Message.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateResponse_Unsigned) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateResponse_Unsigned) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Unsigned != nil {
		{
			size, err := m.Unsigned.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *CreateResponse_Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateResponse_Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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
func (m *CreateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		n += m.Message.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CreateRequest_Init) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Init != nil {
		l = m.Init.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}
func (m *CreateRequest_Signed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Signed != nil {
		l = m.Signed.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}
func (m *CreateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		n += m.Message.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CreateResponse_Unsigned) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Unsigned != nil {
		l = m.Unsigned.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}
func (m *CreateResponse_Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Init", wireType)
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
			v := &Token{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &CreateRequest_Init{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signed", wireType)
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
			v := &Token{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &CreateRequest_Signed{v}
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
func (m *CreateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: CreateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unsigned", wireType)
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
			v := &Token{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &CreateResponse_Unsigned{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
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
			v := &Token{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Message = &CreateResponse_Result{v}
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
