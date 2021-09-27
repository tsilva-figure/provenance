// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/msgfees/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// ComputeMsgBasedRequest is the request type for the Msg.CalculateMsgBasedFees
// RPC method.
type CalculateMsgBasedRequest struct {
	// tx is the transaction to simulate.
	Tx *tx.Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *CalculateMsgBasedRequest) Reset()         { *m = CalculateMsgBasedRequest{} }
func (m *CalculateMsgBasedRequest) String() string { return proto.CompactTextString(m) }
func (*CalculateMsgBasedRequest) ProtoMessage()    {}
func (*CalculateMsgBasedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6bb65eaf858b5f, []int{0}
}
func (m *CalculateMsgBasedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalculateMsgBasedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalculateMsgBasedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalculateMsgBasedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateMsgBasedRequest.Merge(m, src)
}
func (m *CalculateMsgBasedRequest) XXX_Size() int {
	return m.Size()
}
func (m *CalculateMsgBasedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateMsgBasedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateMsgBasedRequest proto.InternalMessageInfo

func (m *CalculateMsgBasedRequest) GetTx() *tx.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

// CalculateMsgBasedFeesResponse is the response type for the Msg.CalculateMsgBasedFees
// RPC method.
type CalculateMsgBasedFeesResponse struct {
	// amount is the amount of coins to be paid as a fee
	FeeAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=fee_amount,json=feeAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_amount"`
}

func (m *CalculateMsgBasedFeesResponse) Reset()         { *m = CalculateMsgBasedFeesResponse{} }
func (m *CalculateMsgBasedFeesResponse) String() string { return proto.CompactTextString(m) }
func (*CalculateMsgBasedFeesResponse) ProtoMessage()    {}
func (*CalculateMsgBasedFeesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6bb65eaf858b5f, []int{1}
}
func (m *CalculateMsgBasedFeesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CalculateMsgBasedFeesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CalculateMsgBasedFeesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CalculateMsgBasedFeesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateMsgBasedFeesResponse.Merge(m, src)
}
func (m *CalculateMsgBasedFeesResponse) XXX_Size() int {
	return m.Size()
}
func (m *CalculateMsgBasedFeesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateMsgBasedFeesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateMsgBasedFeesResponse proto.InternalMessageInfo

func (m *CalculateMsgBasedFeesResponse) GetFeeAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.FeeAmount
	}
	return nil
}

// create fee for msg's (repeated)
type CreateFeeForMsgRequest struct {
	// msg to add Fee for.
	MsgFees []*MsgFees `protobuf:"bytes,1,rep,name=msg_fees,json=msgFees,proto3" json:"msg_fees,omitempty"`
}

func (m *CreateFeeForMsgRequest) Reset()         { *m = CreateFeeForMsgRequest{} }
func (m *CreateFeeForMsgRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFeeForMsgRequest) ProtoMessage()    {}
func (*CreateFeeForMsgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6bb65eaf858b5f, []int{2}
}
func (m *CreateFeeForMsgRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateFeeForMsgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateFeeForMsgRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateFeeForMsgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFeeForMsgRequest.Merge(m, src)
}
func (m *CreateFeeForMsgRequest) XXX_Size() int {
	return m.Size()
}
func (m *CreateFeeForMsgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFeeForMsgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFeeForMsgRequest proto.InternalMessageInfo

func (m *CreateFeeForMsgRequest) GetMsgFees() []*MsgFees {
	if m != nil {
		return m.MsgFees
	}
	return nil
}

// response for CreateFeeForMsg
type CreateFeeForMsgResponse struct {
	// msg to add Fee for.
	MsgFees []*MsgFees `protobuf:"bytes,1,rep,name=msg_fees,json=msgFees,proto3" json:"msg_fees,omitempty"`
}

func (m *CreateFeeForMsgResponse) Reset()         { *m = CreateFeeForMsgResponse{} }
func (m *CreateFeeForMsgResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFeeForMsgResponse) ProtoMessage()    {}
func (*CreateFeeForMsgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c6bb65eaf858b5f, []int{3}
}
func (m *CreateFeeForMsgResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateFeeForMsgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateFeeForMsgResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateFeeForMsgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFeeForMsgResponse.Merge(m, src)
}
func (m *CreateFeeForMsgResponse) XXX_Size() int {
	return m.Size()
}
func (m *CreateFeeForMsgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFeeForMsgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFeeForMsgResponse proto.InternalMessageInfo

func (m *CreateFeeForMsgResponse) GetMsgFees() []*MsgFees {
	if m != nil {
		return m.MsgFees
	}
	return nil
}

func init() {
	proto.RegisterType((*CalculateMsgBasedRequest)(nil), "provenance.msgfees.v1.CalculateMsgBasedRequest")
	proto.RegisterType((*CalculateMsgBasedFeesResponse)(nil), "provenance.msgfees.v1.CalculateMsgBasedFeesResponse")
	proto.RegisterType((*CreateFeeForMsgRequest)(nil), "provenance.msgfees.v1.CreateFeeForMsgRequest")
	proto.RegisterType((*CreateFeeForMsgResponse)(nil), "provenance.msgfees.v1.CreateFeeForMsgResponse")
}

func init() { proto.RegisterFile("provenance/msgfees/v1/tx.proto", fileDescriptor_4c6bb65eaf858b5f) }

var fileDescriptor_4c6bb65eaf858b5f = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x33, 0x5d, 0xf0, 0xc7, 0xec, 0x41, 0x08, 0x56, 0xdb, 0xa0, 0xd9, 0x25, 0x2a, 0x54,
	0xa1, 0x33, 0xb6, 0xeb, 0x45, 0x6f, 0xdb, 0x42, 0x6f, 0x05, 0xa9, 0x7b, 0xf2, 0x52, 0x26, 0xd9,
	0xd7, 0x31, 0x6e, 0x33, 0x13, 0x33, 0x93, 0x92, 0xbd, 0x7a, 0xf5, 0x22, 0xf8, 0x5f, 0xe8, 0xc1,
	0x7f, 0x63, 0x8f, 0x0b, 0x5e, 0x3c, 0xa9, 0xb4, 0xfe, 0x21, 0xd2, 0x64, 0xd2, 0x2d, 0x36, 0x85,
	0xb2, 0xa7, 0x64, 0xf8, 0xbe, 0xf7, 0x79, 0xdf, 0xf7, 0xe6, 0x0d, 0x76, 0xe3, 0x44, 0xce, 0x40,
	0x30, 0x11, 0x00, 0x8d, 0x14, 0x9f, 0x00, 0x28, 0x3a, 0xeb, 0x50, 0x9d, 0x91, 0x38, 0x91, 0x5a,
	0xda, 0xf5, 0x2b, 0x9d, 0x18, 0x9d, 0xcc, 0x3a, 0xce, 0x5d, 0x2e, 0xb9, 0xcc, 0x23, 0xe8, 0xf2,
	0xaf, 0x08, 0x76, 0xdc, 0x40, 0xaa, 0x48, 0x2a, 0xea, 0x33, 0x05, 0x74, 0xd6, 0xf1, 0x41, 0xb3,
	0x0e, 0x0d, 0x64, 0x28, 0x36, 0x74, 0x71, 0xb6, 0xd2, 0x97, 0x07, 0xa3, 0x3f, 0xaa, 0x36, 0x53,
	0xd6, 0x2d, 0x82, 0x9a, 0x05, 0x64, 0x5c, 0x54, 0x2f, 0x0e, 0x46, 0x7a, 0xc0, 0xa5, 0xe4, 0x53,
	0xa0, 0x2c, 0x0e, 0x29, 0x13, 0x42, 0x6a, 0xa6, 0x43, 0x29, 0x4a, 0xd5, 0x31, 0xd5, 0x75, 0xb6,
	0xaa, 0x5d, 0xb6, 0xe9, 0x34, 0x4d, 0x66, 0x7e, 0xf2, 0xd3, 0x09, 0x65, 0xe2, 0xbc, 0x90, 0xbc,
	0x63, 0xdc, 0xe8, 0xb3, 0x69, 0x90, 0x4e, 0x99, 0x86, 0xa1, 0xe2, 0x3d, 0xa6, 0xe0, 0x74, 0x04,
	0x1f, 0x52, 0x50, 0xda, 0x7e, 0x82, 0x6b, 0x3a, 0x6b, 0xa0, 0x43, 0xd4, 0xda, 0xef, 0xd6, 0x89,
	0xf1, 0xa2, 0x33, 0x62, 0xf8, 0xe4, 0x24, 0x1b, 0xd5, 0x74, 0xe6, 0x7d, 0x42, 0xf8, 0xe1, 0x06,
	0x63, 0x00, 0xa0, 0x46, 0xa0, 0x62, 0x29, 0x14, 0xd8, 0xef, 0x31, 0x9e, 0x00, 0x8c, 0x59, 0x24,
	0x53, 0xa1, 0x1b, 0xe8, 0x70, 0xaf, 0xb5, 0xdf, 0x6d, 0x96, 0xc0, 0xe5, 0x38, 0x57, 0xc8, 0xbe,
	0x0c, 0x45, 0xef, 0xf9, 0xc5, 0xaf, 0x03, 0xeb, 0xeb, 0xef, 0x83, 0x16, 0x0f, 0xf5, 0xbb, 0xd4,
	0x27, 0x81, 0x8c, 0xcc, 0x24, 0xcc, 0xa7, 0xad, 0x4e, 0xcf, 0xa8, 0x3e, 0x8f, 0x41, 0xe5, 0x09,
	0x6a, 0x74, 0x7b, 0x02, 0x70, 0x9c, 0xd3, 0xbd, 0x37, 0xf8, 0x5e, 0x3f, 0x01, 0xa6, 0x61, 0x00,
	0x30, 0x90, 0xc9, 0x50, 0xf1, 0xb2, 0x9d, 0x97, 0xf8, 0x56, 0xa4, 0xf8, 0x78, 0x39, 0x6c, 0xe3,
	0xc1, 0x25, 0x95, 0xf7, 0x4f, 0x86, 0x8a, 0xe7, 0xfe, 0x6f, 0x46, 0xc5, 0x8f, 0x77, 0x82, 0xef,
	0x6f, 0x40, 0x4d, 0x6f, 0xd7, 0xa7, 0x76, 0xbf, 0xd5, 0xf0, 0xde, 0x50, 0x71, 0x3b, 0xc1, 0x77,
	0xfe, 0xa3, 0xdb, 0xed, 0x2d, 0x8c, 0xea, 0xd6, 0x1c, 0xb2, 0x6b, 0x78, 0x61, 0xda, 0xb3, 0xec,
	0xef, 0x08, 0xd7, 0x2b, 0x2f, 0xcd, 0xa6, 0xdb, 0x58, 0x5b, 0xd6, 0xc4, 0x79, 0xb1, 0x6b, 0xc2,
	0xfa, 0x4e, 0x78, 0xf4, 0xe3, 0x8f, 0xbf, 0x5f, 0x6a, 0x4f, 0xbd, 0xc7, 0x74, 0xed, 0x59, 0xe4,
	0xcb, 0x4b, 0x83, 0x32, 0x71, 0x5c, 0x4e, 0xf6, 0x15, 0x7a, 0xd6, 0x0b, 0x2f, 0xe6, 0x2e, 0xba,
	0x9c, 0xbb, 0xe8, 0xcf, 0xdc, 0x45, 0x9f, 0x17, 0xae, 0x75, 0xb9, 0x70, 0xad, 0x9f, 0x0b, 0xd7,
	0xc2, 0x8d, 0x50, 0x56, 0x5b, 0x78, 0x8d, 0xde, 0x1e, 0xad, 0xed, 0xd0, 0x55, 0x4c, 0x3b, 0x94,
	0xeb, 0x65, 0xb3, 0xd5, 0x7b, 0xcc, 0x97, 0xca, 0xbf, 0x91, 0xbf, 0x8d, 0xa3, 0x7f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0x5d, 0x98, 0x08, 0x3f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// create fee for an associated Msg (repeated)
	CreateFeeForMsg(ctx context.Context, in *CreateFeeForMsgRequest, opts ...grpc.CallOption) (*CreateFeeForMsgResponse, error)
	// CalculateMsgBasedFees simulates executing a transaction for estimating gas usage.
	CalculateMsgBasedFees(ctx context.Context, in *CalculateMsgBasedRequest, opts ...grpc.CallOption) (*CalculateMsgBasedFeesResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateFeeForMsg(ctx context.Context, in *CreateFeeForMsgRequest, opts ...grpc.CallOption) (*CreateFeeForMsgResponse, error) {
	out := new(CreateFeeForMsgResponse)
	err := c.cc.Invoke(ctx, "/provenance.msgfees.v1.Msg/CreateFeeForMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CalculateMsgBasedFees(ctx context.Context, in *CalculateMsgBasedRequest, opts ...grpc.CallOption) (*CalculateMsgBasedFeesResponse, error) {
	out := new(CalculateMsgBasedFeesResponse)
	err := c.cc.Invoke(ctx, "/provenance.msgfees.v1.Msg/CalculateMsgBasedFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// create fee for an associated Msg (repeated)
	CreateFeeForMsg(context.Context, *CreateFeeForMsgRequest) (*CreateFeeForMsgResponse, error)
	// CalculateMsgBasedFees simulates executing a transaction for estimating gas usage.
	CalculateMsgBasedFees(context.Context, *CalculateMsgBasedRequest) (*CalculateMsgBasedFeesResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateFeeForMsg(ctx context.Context, req *CreateFeeForMsgRequest) (*CreateFeeForMsgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeeForMsg not implemented")
}
func (*UnimplementedMsgServer) CalculateMsgBasedFees(ctx context.Context, req *CalculateMsgBasedRequest) (*CalculateMsgBasedFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateMsgBasedFees not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateFeeForMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeeForMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateFeeForMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.msgfees.v1.Msg/CreateFeeForMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateFeeForMsg(ctx, req.(*CreateFeeForMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CalculateMsgBasedFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateMsgBasedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CalculateMsgBasedFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.msgfees.v1.Msg/CalculateMsgBasedFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CalculateMsgBasedFees(ctx, req.(*CalculateMsgBasedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.msgfees.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeeForMsg",
			Handler:    _Msg_CreateFeeForMsg_Handler,
		},
		{
			MethodName: "CalculateMsgBasedFees",
			Handler:    _Msg_CalculateMsgBasedFees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provenance/msgfees/v1/tx.proto",
}

func (m *CalculateMsgBasedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalculateMsgBasedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalculateMsgBasedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CalculateMsgBasedFeesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CalculateMsgBasedFeesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CalculateMsgBasedFeesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeAmount) > 0 {
		for iNdEx := len(m.FeeAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateFeeForMsgRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateFeeForMsgRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateFeeForMsgRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgFees) > 0 {
		for iNdEx := len(m.MsgFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MsgFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CreateFeeForMsgResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateFeeForMsgResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateFeeForMsgResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgFees) > 0 {
		for iNdEx := len(m.MsgFees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MsgFees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CalculateMsgBasedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *CalculateMsgBasedFeesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeeAmount) > 0 {
		for _, e := range m.FeeAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *CreateFeeForMsgRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MsgFees) > 0 {
		for _, e := range m.MsgFees {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *CreateFeeForMsgResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MsgFees) > 0 {
		for _, e := range m.MsgFees {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CalculateMsgBasedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: CalculateMsgBasedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalculateMsgBasedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &tx.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *CalculateMsgBasedFeesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: CalculateMsgBasedFeesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CalculateMsgBasedFeesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAmount = append(m.FeeAmount, types.Coin{})
			if err := m.FeeAmount[len(m.FeeAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *CreateFeeForMsgRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: CreateFeeForMsgRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateFeeForMsgRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgFees = append(m.MsgFees, &MsgFees{})
			if err := m.MsgFees[len(m.MsgFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *CreateFeeForMsgResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: CreateFeeForMsgResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateFeeForMsgResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgFees = append(m.MsgFees, &MsgFees{})
			if err := m.MsgFees[len(m.MsgFees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)