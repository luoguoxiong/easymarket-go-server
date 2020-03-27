// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GoodsReq struct {
	IsHot                int32    `protobuf:"varint,1,opt,name=isHot,proto3" json:"isHot,omitempty" form:"isHot"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	Size_                int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty" form:"size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsReq) Reset()         { *m = GoodsReq{} }
func (m *GoodsReq) String() string { return proto.CompactTextString(m) }
func (*GoodsReq) ProtoMessage()    {}
func (*GoodsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *GoodsReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsReq.Merge(m, src)
}
func (m *GoodsReq) XXX_Size() int {
	return m.Size()
}
func (m *GoodsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsReq proto.InternalMessageInfo

type GoodsListRes struct {
	GoodsList            []*GoodsRes `protobuf:"bytes,1,rep,name=goodsList,proto3" json:"goodsList,omitempty" form:"goodsList"`
	Total                int32       `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty" form:"total"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GoodsListRes) Reset()         { *m = GoodsListRes{} }
func (m *GoodsListRes) String() string { return proto.CompactTextString(m) }
func (*GoodsListRes) ProtoMessage()    {}
func (*GoodsListRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *GoodsListRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsListRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsListRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsListRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsListRes.Merge(m, src)
}
func (m *GoodsListRes) XXX_Size() int {
	return m.Size()
}
func (m *GoodsListRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsListRes.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsListRes proto.InternalMessageInfo

type GoodsRes struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RetailPrice          float64  `protobuf:"fixed64,3,opt,name=retail_price,json=retailPrice,proto3" json:"retail_price,omitempty"`
	GoodsBrief           string   `protobuf:"bytes,4,opt,name=goods_brief,json=goodsBrief,proto3" json:"goods_brief,omitempty"`
	ListPicUrl           string   `protobuf:"bytes,5,opt,name=list_pic_url,json=listPicUrl,proto3" json:"list_pic_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodsRes) Reset()         { *m = GoodsRes{} }
func (m *GoodsRes) String() string { return proto.CompactTextString(m) }
func (*GoodsRes) ProtoMessage()    {}
func (*GoodsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}
func (m *GoodsRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsRes.Merge(m, src)
}
func (m *GoodsRes) XXX_Size() int {
	return m.Size()
}
func (m *GoodsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsRes.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsRes proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GoodsReq)(nil), "goods.service.v1.GoodsReq")
	proto.RegisterType((*GoodsListRes)(nil), "goods.service.v1.GoodsListRes")
	proto.RegisterType((*GoodsRes)(nil), "goods.service.v1.GoodsRes")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x6f, 0x13, 0x3f,
	0x10, 0xad, 0xf3, 0xe7, 0xa7, 0x5f, 0x9c, 0x55, 0x89, 0x0c, 0x15, 0x4b, 0x0a, 0x9b, 0x60, 0x24,
	0xd4, 0x0b, 0xbb, 0xa2, 0xdc, 0x38, 0x46, 0x48, 0xe5, 0x80, 0x50, 0x65, 0x89, 0x0b, 0x97, 0xc8,
	0xd9, 0x38, 0x66, 0xa4, 0x4d, 0xbc, 0x5d, 0x3b, 0x95, 0x40, 0x9c, 0x38, 0xf0, 0x05, 0xe0, 0xc0,
	0x47, 0xea, 0x11, 0x89, 0x13, 0x97, 0x08, 0x02, 0x9f, 0x20, 0x9f, 0x00, 0x79, 0xbc, 0x6d, 0x21,
	0x52, 0x7b, 0x9b, 0x79, 0xef, 0x79, 0xde, 0xcc, 0xdb, 0xa5, 0x1d, 0x59, 0x42, 0x5a, 0x56, 0xc6,
	0x19, 0xd6, 0xd3, 0xc6, 0x4c, 0x6d, 0x6a, 0x55, 0x75, 0x0a, 0xb9, 0x4a, 0x4f, 0x1f, 0xf7, 0x1f,
	0x69, 0x70, 0x6f, 0x96, 0x93, 0x34, 0x37, 0xf3, 0x4c, 0x1b, 0x6d, 0x32, 0x14, 0x4e, 0x96, 0x33,
	0xec, 0xb0, 0xc1, 0x2a, 0x0c, 0xe8, 0xef, 0x6b, 0x63, 0x74, 0xa1, 0x2e, 0x55, 0x6a, 0x5e, 0xba,
	0xb7, 0x35, 0x79, 0xb7, 0x26, 0x65, 0x09, 0x99, 0x5c, 0x2c, 0x8c, 0x93, 0x0e, 0xcc, 0xc2, 0x06,
	0x96, 0xbf, 0xa7, 0xff, 0x1f, 0x79, 0x77, 0xa1, 0x4e, 0xd8, 0x43, 0xda, 0x06, 0xfb, 0xdc, 0xb8,
	0x98, 0x0c, 0xc9, 0x41, 0x7b, 0xd4, 0xdb, 0xac, 0x06, 0xd1, 0xcc, 0x54, 0xf3, 0xa7, 0x1c, 0x61,
	0x2e, 0x02, 0xcd, 0x1e, 0xd0, 0x56, 0x29, 0xb5, 0x8a, 0x1b, 0x28, 0xbb, 0xb1, 0x59, 0x0d, 0xba,
	0x41, 0xe6, 0x51, 0x2e, 0x90, 0xf4, 0x22, 0x0b, 0xef, 0x54, 0xdc, 0xdc, 0x16, 0x79, 0x94, 0x0b,
	0x24, 0xf9, 0x47, 0x42, 0x23, 0xb4, 0x7f, 0x01, 0xd6, 0x09, 0x65, 0xd9, 0x4b, 0xda, 0xd1, 0xe7,
	0x7d, 0x4c, 0x86, 0xcd, 0x83, 0xee, 0x61, 0x3f, 0xdd, 0x8e, 0x27, 0xad, 0x37, 0xb6, 0xa3, 0x5b,
	0x9b, 0xd5, 0xa0, 0x17, 0xc6, 0x5e, 0x3c, 0xe3, 0xe2, 0x72, 0x84, 0x3f, 0xc9, 0x19, 0x27, 0x8b,
	0x7a, 0xd7, 0xbf, 0x4e, 0x42, 0x98, 0x8b, 0x40, 0xf3, 0xcf, 0xe4, 0x22, 0x07, 0xcb, 0x76, 0x69,
	0x03, 0xa6, 0x21, 0x04, 0xd1, 0x80, 0x29, 0x63, 0xb4, 0xb5, 0x90, 0xf3, 0x70, 0x6f, 0x47, 0x60,
	0xcd, 0xee, 0xd3, 0xa8, 0x52, 0x4e, 0x42, 0x31, 0x2e, 0x2b, 0xc8, 0xc3, 0x99, 0x44, 0x74, 0x03,
	0x76, 0xec, 0x21, 0x36, 0xa0, 0x5d, 0x5c, 0x64, 0x3c, 0xa9, 0x40, 0xcd, 0xe2, 0x16, 0xbe, 0xa6,
	0x08, 0x8d, 0x3c, 0xc2, 0x86, 0x34, 0x2a, 0xc0, 0xba, 0x71, 0x09, 0xf9, 0x78, 0x59, 0x15, 0x71,
	0x3b, 0x28, 0x3c, 0x76, 0x0c, 0xf9, 0xab, 0xaa, 0x38, 0xfc, 0x4e, 0x68, 0x1b, 0xd7, 0x62, 0x9a,
	0xee, 0x1e, 0x29, 0x87, 0xf5, 0x33, 0xf4, 0x60, 0x57, 0xe7, 0x72, 0xd2, 0xbf, 0x26, 0x33, 0x7e,
	0xef, 0xc3, 0xb7, 0xdf, 0x9f, 0x1a, 0xb7, 0xd9, 0x5e, 0x86, 0x9a, 0x4c, 0xff, 0x3b, 0x56, 0xd3,
	0xe8, 0xdc, 0x08, 0x13, 0xbc, 0xce, 0x26, 0xb9, 0x82, 0xab, 0xbf, 0x26, 0xdf, 0x47, 0xab, 0x3d,
	0x76, 0x73, 0xcb, 0xca, 0xf3, 0xa3, 0x3b, 0x67, 0x3f, 0x93, 0x9d, 0xb3, 0x75, 0x42, 0xbe, 0xae,
	0x13, 0xf2, 0x63, 0x9d, 0x90, 0x2f, 0xbf, 0x92, 0x9d, 0xd7, 0x4d, 0x59, 0xc2, 0xe4, 0x3f, 0xfc,
	0x37, 0x9f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xac, 0x33, 0xc3, 0x24, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoodsClient interface {
	GetGoodsDetail(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsRes, error)
	GetGoodsList(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsListRes, error)
}

type goodsClient struct {
	cc *grpc.ClientConn
}

func NewGoodsClient(cc *grpc.ClientConn) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetGoodsDetail(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsRes, error) {
	out := new(GoodsRes)
	err := c.cc.Invoke(ctx, "/goods.service.v1.Goods/GetGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetGoodsList(ctx context.Context, in *GoodsReq, opts ...grpc.CallOption) (*GoodsListRes, error) {
	out := new(GoodsListRes)
	err := c.cc.Invoke(ctx, "/goods.service.v1.Goods/GetGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the server API for Goods service.
type GoodsServer interface {
	GetGoodsDetail(context.Context, *GoodsReq) (*GoodsRes, error)
	GetGoodsList(context.Context, *GoodsReq) (*GoodsListRes, error)
}

// UnimplementedGoodsServer can be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (*UnimplementedGoodsServer) GetGoodsDetail(ctx context.Context, req *GoodsReq) (*GoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}
func (*UnimplementedGoodsServer) GetGoodsList(ctx context.Context, req *GoodsReq) (*GoodsListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsList not implemented")
}

func RegisterGoodsServer(s *grpc.Server, srv GoodsServer) {
	s.RegisterService(&_Goods_serviceDesc, srv)
}

func _Goods_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.service.v1.Goods/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsDetail(ctx, req.(*GoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.service.v1.Goods/GetGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetGoodsList(ctx, req.(*GoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Goods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "goods.service.v1.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGoodsDetail",
			Handler:    _Goods_GetGoodsDetail_Handler,
		},
		{
			MethodName: "GetGoodsList",
			Handler:    _Goods_GetGoodsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *GoodsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Size_ != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if m.Page != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x10
	}
	if m.IsHot != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.IsHot))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GoodsListRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsListRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsListRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Total != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.GoodsList) > 0 {
		for iNdEx := len(m.GoodsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GoodsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GoodsRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ListPicUrl) > 0 {
		i -= len(m.ListPicUrl)
		copy(dAtA[i:], m.ListPicUrl)
		i = encodeVarintApi(dAtA, i, uint64(len(m.ListPicUrl)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.GoodsBrief) > 0 {
		i -= len(m.GoodsBrief)
		copy(dAtA[i:], m.GoodsBrief)
		i = encodeVarintApi(dAtA, i, uint64(len(m.GoodsBrief)))
		i--
		dAtA[i] = 0x22
	}
	if m.RetailPrice != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.RetailPrice))))
		i--
		dAtA[i] = 0x19
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintApi(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GoodsReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IsHot != 0 {
		n += 1 + sovApi(uint64(m.IsHot))
	}
	if m.Page != 0 {
		n += 1 + sovApi(uint64(m.Page))
	}
	if m.Size_ != 0 {
		n += 1 + sovApi(uint64(m.Size_))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoodsListRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GoodsList) > 0 {
		for _, e := range m.GoodsList {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if m.Total != 0 {
		n += 1 + sovApi(uint64(m.Total))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GoodsRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovApi(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.RetailPrice != 0 {
		n += 9
	}
	l = len(m.GoodsBrief)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.ListPicUrl)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoodsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: GoodsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsHot", wireType)
			}
			m.IsHot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsHot |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func (m *GoodsListRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: GoodsListRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsListRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoodsList = append(m.GoodsList, &GoodsRes{})
			if err := m.GoodsList[len(m.GoodsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func (m *GoodsRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: GoodsRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field RetailPrice", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.RetailPrice = float64(math.Float64frombits(v))
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodsBrief", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GoodsBrief = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListPicUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListPicUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
