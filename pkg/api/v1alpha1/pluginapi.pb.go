// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/v1alpha1/pluginapi.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterPluginRequest struct {
	// Uniq name of the server
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Server socket address to register;
	// This server has to offers Plugin functionality
	// at this socket.
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPluginRequest) Reset()         { *m = RegisterPluginRequest{} }
func (m *RegisterPluginRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterPluginRequest) ProtoMessage()    {}
func (*RegisterPluginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f16fa2bf8e463b, []int{0}
}

func (m *RegisterPluginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPluginRequest.Unmarshal(m, b)
}
func (m *RegisterPluginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPluginRequest.Marshal(b, m, deterministic)
}
func (m *RegisterPluginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPluginRequest.Merge(m, src)
}
func (m *RegisterPluginRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterPluginRequest.Size(m)
}
func (m *RegisterPluginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPluginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPluginRequest proto.InternalMessageInfo

func (m *RegisterPluginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterPluginRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RegisterKeyServerReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterKeyServerReply) Reset()         { *m = RegisterKeyServerReply{} }
func (m *RegisterKeyServerReply) String() string { return proto.CompactTextString(m) }
func (*RegisterKeyServerReply) ProtoMessage()    {}
func (*RegisterKeyServerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f16fa2bf8e463b, []int{1}
}

func (m *RegisterKeyServerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterKeyServerReply.Unmarshal(m, b)
}
func (m *RegisterKeyServerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterKeyServerReply.Marshal(b, m, deterministic)
}
func (m *RegisterKeyServerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterKeyServerReply.Merge(m, src)
}
func (m *RegisterKeyServerReply) XXX_Size() int {
	return xxx_messageInfo_RegisterKeyServerReply.Size(m)
}
func (m *RegisterKeyServerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterKeyServerReply.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterKeyServerReply proto.InternalMessageInfo

type ValidateQuoteRequest struct {
	// CA signer name
	SignerName string `protobuf:"bytes,1,opt,name=signerName,proto3" json:"signerName,omitempty"`
	// base64 encoded public key used for generating the quote
	PublicKey []byte `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	// base64 encoded SGX Quote
	Quote []byte `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	// base64 encoded nonce used for generating the SGX quote
	Nonce                []byte   `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateQuoteRequest) Reset()         { *m = ValidateQuoteRequest{} }
func (m *ValidateQuoteRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateQuoteRequest) ProtoMessage()    {}
func (*ValidateQuoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f16fa2bf8e463b, []int{2}
}

func (m *ValidateQuoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateQuoteRequest.Unmarshal(m, b)
}
func (m *ValidateQuoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateQuoteRequest.Marshal(b, m, deterministic)
}
func (m *ValidateQuoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateQuoteRequest.Merge(m, src)
}
func (m *ValidateQuoteRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateQuoteRequest.Size(m)
}
func (m *ValidateQuoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateQuoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateQuoteRequest proto.InternalMessageInfo

func (m *ValidateQuoteRequest) GetSignerName() string {
	if m != nil {
		return m.SignerName
	}
	return ""
}

func (m *ValidateQuoteRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *ValidateQuoteRequest) GetQuote() []byte {
	if m != nil {
		return m.Quote
	}
	return nil
}

func (m *ValidateQuoteRequest) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

type ValidateQuoteReply struct {
	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	// Failure message in case of provided quote is invalid
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateQuoteReply) Reset()         { *m = ValidateQuoteReply{} }
func (m *ValidateQuoteReply) String() string { return proto.CompactTextString(m) }
func (*ValidateQuoteReply) ProtoMessage()    {}
func (*ValidateQuoteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f16fa2bf8e463b, []int{3}
}

func (m *ValidateQuoteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateQuoteReply.Unmarshal(m, b)
}
func (m *ValidateQuoteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateQuoteReply.Marshal(b, m, deterministic)
}
func (m *ValidateQuoteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateQuoteReply.Merge(m, src)
}
func (m *ValidateQuoteReply) XXX_Size() int {
	return xxx_messageInfo_ValidateQuoteReply.Size(m)
}
func (m *ValidateQuoteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateQuoteReply.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateQuoteReply proto.InternalMessageInfo

func (m *ValidateQuoteReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *ValidateQuoteReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetCAKeyAndCertificateRequest struct {
	// CA signer name
	SignerName string `protobuf:"bytes,1,opt,name=signerName,proto3" json:"signerName,omitempty"`
	// base64 encoded public key used for generating the quote
	PublicKey []byte `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	// base64 encoded SGX Quote
	Quote []byte `protobuf:"bytes,3,opt,name=quote,proto3" json:"quote,omitempty"`
	// base64 encoded nonce used for generating the SGX quote
	Nonce                []byte   `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCAKeyAndCertificateRequest) Reset()         { *m = GetCAKeyAndCertificateRequest{} }
func (m *GetCAKeyAndCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*GetCAKeyAndCertificateRequest) ProtoMessage()    {}
func (*GetCAKeyAndCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f16fa2bf8e463b, []int{4}
}

func (m *GetCAKeyAndCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCAKeyAndCertificateRequest.Unmarshal(m, b)
}
func (m *GetCAKeyAndCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCAKeyAndCertificateRequest.Marshal(b, m, deterministic)
}
func (m *GetCAKeyAndCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCAKeyAndCertificateRequest.Merge(m, src)
}
func (m *GetCAKeyAndCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_GetCAKeyAndCertificateRequest.Size(m)
}
func (m *GetCAKeyAndCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCAKeyAndCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCAKeyAndCertificateRequest proto.InternalMessageInfo

func (m *GetCAKeyAndCertificateRequest) GetSignerName() string {
	if m != nil {
		return m.SignerName
	}
	return ""
}

func (m *GetCAKeyAndCertificateRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *GetCAKeyAndCertificateRequest) GetQuote() []byte {
	if m != nil {
		return m.Quote
	}
	return nil
}

func (m *GetCAKeyAndCertificateRequest) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

type GetCAKeyAndCertificateReply struct {
	// The CA private key(PWK) is wrapped with a symmetric key(SWK)
	// that was wrapped with the given publicKey. Both the SWK and
	// PWK are concatenated and returned as single base64 encoded block.
	WrappedKey []byte `protobuf:"bytes,1,opt,name=wrappedKey,proto3" json:"wrappedKey,omitempty"`
	// base64 encoded PEM certificate
	Certificate          []byte   `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCAKeyAndCertificateReply) Reset()         { *m = GetCAKeyAndCertificateReply{} }
func (m *GetCAKeyAndCertificateReply) String() string { return proto.CompactTextString(m) }
func (*GetCAKeyAndCertificateReply) ProtoMessage()    {}
func (*GetCAKeyAndCertificateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_38f16fa2bf8e463b, []int{5}
}

func (m *GetCAKeyAndCertificateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCAKeyAndCertificateReply.Unmarshal(m, b)
}
func (m *GetCAKeyAndCertificateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCAKeyAndCertificateReply.Marshal(b, m, deterministic)
}
func (m *GetCAKeyAndCertificateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCAKeyAndCertificateReply.Merge(m, src)
}
func (m *GetCAKeyAndCertificateReply) XXX_Size() int {
	return xxx_messageInfo_GetCAKeyAndCertificateReply.Size(m)
}
func (m *GetCAKeyAndCertificateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCAKeyAndCertificateReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetCAKeyAndCertificateReply proto.InternalMessageInfo

func (m *GetCAKeyAndCertificateReply) GetWrappedKey() []byte {
	if m != nil {
		return m.WrappedKey
	}
	return nil
}

func (m *GetCAKeyAndCertificateReply) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterPluginRequest)(nil), "v1alpha1.RegisterPluginRequest")
	proto.RegisterType((*RegisterKeyServerReply)(nil), "v1alpha1.RegisterKeyServerReply")
	proto.RegisterType((*ValidateQuoteRequest)(nil), "v1alpha1.ValidateQuoteRequest")
	proto.RegisterType((*ValidateQuoteReply)(nil), "v1alpha1.ValidateQuoteReply")
	proto.RegisterType((*GetCAKeyAndCertificateRequest)(nil), "v1alpha1.GetCAKeyAndCertificateRequest")
	proto.RegisterType((*GetCAKeyAndCertificateReply)(nil), "v1alpha1.GetCAKeyAndCertificateReply")
}

func init() { proto.RegisterFile("pkg/api/v1alpha1/pluginapi.proto", fileDescriptor_38f16fa2bf8e463b) }

var fileDescriptor_38f16fa2bf8e463b = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x50, 0x42, 0x3a, 0x7c, 0x1c, 0x56, 0x25, 0xb2, 0x42, 0x09, 0x96, 0x25, 0x04, 0x97,
	0xda, 0x6a, 0xf9, 0x03, 0xb4, 0x15, 0x70, 0x88, 0xc4, 0x87, 0x11, 0x1c, 0xb8, 0xa0, 0x8d, 0x3d,
	0xb8, 0x0b, 0x9b, 0xdd, 0xcd, 0xee, 0x38, 0xc8, 0x37, 0xae, 0xfc, 0x30, 0xfe, 0x17, 0xf2, 0xc6,
	0x96, 0x93, 0x90, 0xc0, 0x91, 0x9b, 0xdf, 0x9b, 0x8f, 0x37, 0xe3, 0x79, 0x0b, 0x91, 0xf9, 0x56,
	0xa6, 0xdc, 0x88, 0x74, 0x79, 0xca, 0xa5, 0xb9, 0xe2, 0xa7, 0xa9, 0x91, 0x55, 0x29, 0x14, 0x37,
	0x22, 0x31, 0x56, 0x93, 0x66, 0xc3, 0x2e, 0x12, 0xbf, 0x80, 0xfb, 0x19, 0x96, 0xc2, 0x11, 0xda,
	0xb7, 0x3e, 0x29, 0xc3, 0x45, 0x85, 0x8e, 0x18, 0x83, 0x03, 0xc5, 0xe7, 0x18, 0x06, 0x51, 0xf0,
	0xf4, 0x30, 0xf3, 0xdf, 0x2c, 0x84, 0x5b, 0xbc, 0x28, 0x2c, 0x3a, 0x17, 0x5e, 0xf7, 0x74, 0x07,
	0xe3, 0x10, 0x46, 0x5d, 0x9b, 0x29, 0xd6, 0xef, 0xd1, 0x2e, 0xd1, 0x66, 0x68, 0x64, 0x1d, 0xff,
	0x08, 0xe0, 0xe8, 0x23, 0x97, 0xa2, 0xe0, 0x84, 0xef, 0x2a, 0x4d, 0xd8, 0x09, 0x4c, 0x00, 0x9c,
	0x28, 0x15, 0xda, 0xd7, 0xbd, 0xcc, 0x1a, 0xc3, 0x8e, 0xe1, 0xd0, 0x54, 0x33, 0x29, 0xf2, 0x29,
	0xd6, 0x5e, 0xee, 0x4e, 0xd6, 0x13, 0xec, 0x08, 0x6e, 0x2e, 0x9a, 0x6e, 0xe1, 0x0d, 0x1f, 0x59,
	0x81, 0x86, 0x55, 0x5a, 0xe5, 0x18, 0x1e, 0xac, 0x58, 0x0f, 0xe2, 0x97, 0xc0, 0xb6, 0x26, 0x30,
	0xb2, 0x66, 0x23, 0x18, 0x58, 0x74, 0x95, 0x24, 0xaf, 0x3d, 0xcc, 0x5a, 0xd4, 0x2c, 0x39, 0x47,
	0xe7, 0x78, 0x89, 0xdd, 0x92, 0x2d, 0x8c, 0x7f, 0x06, 0xf0, 0xf0, 0x15, 0xd2, 0xe5, 0xf9, 0x14,
	0xeb, 0x73, 0x55, 0x5c, 0xa2, 0x25, 0xf1, 0x45, 0xe4, 0xfc, 0x7f, 0xec, 0xf4, 0x19, 0x1e, 0xec,
	0x1b, 0xa5, 0x59, 0x6e, 0x02, 0xf0, 0xdd, 0x72, 0x63, 0xb0, 0x68, 0x94, 0x02, 0x5f, 0xb9, 0xc6,
	0xb0, 0x08, 0x6e, 0xe7, 0x7d, 0x4d, 0x3b, 0xca, 0x3a, 0x75, 0xc6, 0x61, 0xb8, 0xba, 0xa8, 0xad,
	0xd9, 0x07, 0xb8, 0xb7, 0x69, 0x12, 0xf6, 0x28, 0xe9, 0x1c, 0x94, 0xec, 0xb4, 0xcf, 0x38, 0xfa,
	0x33, 0x61, 0xcb, 0x18, 0xd7, 0xce, 0x7e, 0x05, 0x30, 0x68, 0xfb, 0xbd, 0x81, 0xbb, 0x1b, 0x27,
	0x62, 0x93, 0xbe, 0x7e, 0x97, 0x7b, 0xc6, 0xc7, 0x7b, 0xe3, 0xbe, 0x37, 0xfb, 0x0a, 0xa3, 0xdd,
	0xff, 0x87, 0x3d, 0xe9, 0x2b, 0xff, 0x7a, 0xcc, 0xf1, 0xe3, 0x7f, 0x27, 0x7a, 0xad, 0x8b, 0x8b,
	0x4f, 0xcf, 0x4b, 0x41, 0x57, 0xd5, 0x2c, 0xc9, 0xf5, 0x3c, 0x15, 0x8a, 0x50, 0xa6, 0x64, 0x2b,
	0x47, 0x58, 0x9c, 0x70, 0x22, 0x74, 0xc4, 0x49, 0x68, 0x75, 0x92, 0x6b, 0x45, 0x56, 0x4b, 0x89,
	0x36, 0xdd, 0x7e, 0xa1, 0xb3, 0x81, 0x7f, 0x98, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x54,
	0x1d, 0x3c, 0xf4, 0xbc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistryClient is the client API for Registry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistryClient interface {
	RegisterPlugin(ctx context.Context, in *RegisterPluginRequest, opts ...grpc.CallOption) (*RegisterKeyServerReply, error)
}

type registryClient struct {
	cc *grpc.ClientConn
}

func NewRegistryClient(cc *grpc.ClientConn) RegistryClient {
	return &registryClient{cc}
}

func (c *registryClient) RegisterPlugin(ctx context.Context, in *RegisterPluginRequest, opts ...grpc.CallOption) (*RegisterKeyServerReply, error) {
	out := new(RegisterKeyServerReply)
	err := c.cc.Invoke(ctx, "/v1alpha1.Registry/RegisterPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistryServer is the server API for Registry service.
type RegistryServer interface {
	RegisterPlugin(context.Context, *RegisterPluginRequest) (*RegisterKeyServerReply, error)
}

func RegisterRegistryServer(s *grpc.Server, srv RegistryServer) {
	s.RegisterService(&_Registry_serviceDesc, srv)
}

func _Registry_RegisterPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServer).RegisterPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Registry/RegisterPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServer).RegisterPlugin(ctx, req.(*RegisterPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Registry",
	HandlerType: (*RegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPlugin",
			Handler:    _Registry_RegisterPlugin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/v1alpha1/pluginapi.proto",
}

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PluginClient interface {
	// ValidateQuote validates the given SGX quote
	ValidateQuote(ctx context.Context, in *ValidateQuoteRequest, opts ...grpc.CallOption) (*ValidateQuoteReply, error)
	// GetCAKeyCertificate retrieves the stored CA key and certificate at the key-manager
	// for given signer signerName.
	// On success, returns the key and certificate.
	// Otherwise, appropriate error gets returned.
	GetCAKeyAndCertificate(ctx context.Context, in *GetCAKeyAndCertificateRequest, opts ...grpc.CallOption) (*GetCAKeyAndCertificateReply, error)
}

type pluginClient struct {
	cc *grpc.ClientConn
}

func NewPluginClient(cc *grpc.ClientConn) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) ValidateQuote(ctx context.Context, in *ValidateQuoteRequest, opts ...grpc.CallOption) (*ValidateQuoteReply, error) {
	out := new(ValidateQuoteReply)
	err := c.cc.Invoke(ctx, "/v1alpha1.Plugin/ValidateQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) GetCAKeyAndCertificate(ctx context.Context, in *GetCAKeyAndCertificateRequest, opts ...grpc.CallOption) (*GetCAKeyAndCertificateReply, error) {
	out := new(GetCAKeyAndCertificateReply)
	err := c.cc.Invoke(ctx, "/v1alpha1.Plugin/GetCAKeyAndCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
type PluginServer interface {
	// ValidateQuote validates the given SGX quote
	ValidateQuote(context.Context, *ValidateQuoteRequest) (*ValidateQuoteReply, error)
	// GetCAKeyCertificate retrieves the stored CA key and certificate at the key-manager
	// for given signer signerName.
	// On success, returns the key and certificate.
	// Otherwise, appropriate error gets returned.
	GetCAKeyAndCertificate(context.Context, *GetCAKeyAndCertificateRequest) (*GetCAKeyAndCertificateReply, error)
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_ValidateQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ValidateQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Plugin/ValidateQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ValidateQuote(ctx, req.(*ValidateQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_GetCAKeyAndCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCAKeyAndCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetCAKeyAndCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Plugin/GetCAKeyAndCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetCAKeyAndCertificate(ctx, req.(*GetCAKeyAndCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateQuote",
			Handler:    _Plugin_ValidateQuote_Handler,
		},
		{
			MethodName: "GetCAKeyAndCertificate",
			Handler:    _Plugin_GetCAKeyAndCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/v1alpha1/pluginapi.proto",
}
