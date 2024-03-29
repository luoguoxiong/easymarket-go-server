// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	api "easymarket-go-server/common/goods/api"
	api2 "easymarket-go-server/common/topic/api"
	api1 "easymarket-go-server/common/wechat/api"
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

type ActivityServiceErrCode int32

const (
	ActivityServiceErrCode_Undefined    ActivityServiceErrCode = 0
	ActivityServiceErrCode_TokenTimeOut ActivityServiceErrCode = 1001
	ActivityServiceErrCode_TokenIsError ActivityServiceErrCode = 1002
	ActivityServiceErrCode_TokenIsNew   ActivityServiceErrCode = 1003
)

var ActivityServiceErrCode_name = map[int32]string{
	0:    "Undefined",
	1001: "TokenTimeOut",
	1002: "TokenIsError",
	1003: "TokenIsNew",
}

var ActivityServiceErrCode_value = map[string]int32{
	"Undefined":    0,
	"TokenTimeOut": 1001,
	"TokenIsError": 1002,
	"TokenIsNew":   1003,
}

func (x ActivityServiceErrCode) String() string {
	return proto.EnumName(ActivityServiceErrCode_name, int32(x))
}

func (ActivityServiceErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type GoodsSellRes struct {
	// 商品规格
	GoodsSize []*api.GoodsSize `protobuf:"bytes,1,rep,name=goodsSize,proto3" json:"goodsSize" form:"goodsSize"`
	// 商品轮播图
	GoodsGallery []*api.GoodsGallery `protobuf:"bytes,2,rep,name=goodsGallery,proto3" json:"goodsGallery" form:"goodsGallery"`
	// 商品特点
	Attribute []*api.Attribute `protobuf:"bytes,3,rep,name=attribute,proto3" json:"attribute" form:"attribute"`
	// 商品疑问列表
	Issue []*api.GoodsIssue `protobuf:"bytes,4,rep,name=issue,proto3" json:"issue" form:"issue"`
	// 产品列表
	ProductList          []*api.Product `protobuf:"bytes,5,rep,name=productList,proto3" json:"productList" form:"productList"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GoodsSellRes) Reset()         { *m = GoodsSellRes{} }
func (m *GoodsSellRes) String() string { return proto.CompactTextString(m) }
func (*GoodsSellRes) ProtoMessage()    {}
func (*GoodsSellRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *GoodsSellRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GoodsSellRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GoodsSellRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GoodsSellRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodsSellRes.Merge(m, src)
}
func (m *GoodsSellRes) XXX_Size() int {
	return m.Size()
}
func (m *GoodsSellRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodsSellRes.DiscardUnknown(m)
}

var xxx_messageInfo_GoodsSellRes proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("app.service.v1.ActivityServiceErrCode", ActivityServiceErrCode_name, ActivityServiceErrCode_value)
	proto.RegisterType((*GoodsSellRes)(nil), "app.service.v1.GoodsSellRes")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x9b, 0x0d, 0x4b, 0xe9, 0x34, 0xbb, 0x49, 0x07, 0x76, 0x49, 0xd3, 0x10, 0xb7, 0x96,
	0x10, 0x68, 0xa5, 0xda, 0xb0, 0xdc, 0xf6, 0xd6, 0x74, 0x57, 0x51, 0xa5, 0x85, 0x45, 0x6e, 0x01,
	0x09, 0x21, 0xa4, 0x89, 0xfd, 0xea, 0x0c, 0x75, 0x3c, 0x66, 0x3c, 0x69, 0x15, 0x8e, 0x7c, 0x05,
	0x2e, 0x7c, 0xa4, 0x3d, 0x22, 0x71, 0xb7, 0xa0, 0x70, 0x21, 0x5c, 0x50, 0x3f, 0x01, 0x9a, 0x37,
	0x4e, 0xe2, 0x34, 0x71, 0xb6, 0xb7, 0x78, 0xfe, 0xff, 0xf7, 0xff, 0xcd, 0x7b, 0x1e, 0xdb, 0x21,
	0x5b, 0x2c, 0xe1, 0x4e, 0x22, 0x85, 0x12, 0xf4, 0x21, 0x4b, 0x12, 0x27, 0x05, 0x79, 0xc9, 0x7d,
	0x70, 0x2e, 0x3f, 0x6d, 0x1d, 0x86, 0x5c, 0x0d, 0x46, 0x7d, 0xc7, 0x17, 0x43, 0x37, 0x14, 0xa1,
	0x70, 0xd1, 0xd6, 0x1f, 0x9d, 0xe3, 0x15, 0x5e, 0xe0, 0x2f, 0x53, 0xde, 0xda, 0x0b, 0x85, 0x08,
	0x23, 0x98, 0xbb, 0x60, 0x98, 0xa8, 0x71, 0x2e, 0xb6, 0x73, 0x91, 0x25, 0xdc, 0x65, 0x71, 0x2c,
	0x14, 0x53, 0x5c, 0xc4, 0x69, 0xae, 0xba, 0xc0, 0xd2, 0xf1, 0x90, 0xc9, 0x0b, 0x50, 0x87, 0xa1,
	0x38, 0xd4, 0x7b, 0x00, 0xe9, 0xfa, 0x62, 0x38, 0x14, 0xb1, 0x1b, 0x0a, 0x11, 0xa4, 0xa6, 0x70,
	0xba, 0xd5, 0xf5, 0x05, 0x4a, 0x24, 0xdc, 0xbf, 0x55, 0xf0, 0xc9, 0xba, 0x82, 0x2b, 0xf0, 0x07,
	0x4c, 0x2d, 0x56, 0xd8, 0x93, 0x2a, 0xa9, 0xf5, 0x34, 0xfa, 0x14, 0xa2, 0xc8, 0x83, 0x94, 0x7e,
	0x47, 0xb6, 0x70, 0x2b, 0xa7, 0xfc, 0x27, 0x68, 0x56, 0xf6, 0xab, 0x1f, 0x6f, 0x3f, 0xdd, 0x73,
	0x70, 0xa5, 0x30, 0x34, 0xa7, 0x37, 0xb5, 0x74, 0x0f, 0x26, 0x99, 0x35, 0xaf, 0xb8, 0xc9, 0xac,
	0xc6, 0xb9, 0x90, 0xc3, 0x67, 0xf6, 0x6c, 0xc9, 0xf6, 0xe6, 0x32, 0xbd, 0x20, 0x35, 0xbc, 0xe8,
	0xb1, 0x28, 0x02, 0x39, 0x6e, 0xde, 0x43, 0x40, 0xa7, 0x04, 0x90, 0xbb, 0xba, 0x1f, 0x4d, 0x32,
	0x6b, 0xa1, 0xee, 0x26, 0xb3, 0xde, 0x2d, 0x60, 0xf2, 0x55, 0xdb, 0x5b, 0x30, 0xe9, 0x56, 0x98,
	0x52, 0x92, 0xf7, 0x47, 0x0a, 0x9a, 0xd5, 0xb2, 0x56, 0x8e, 0xa6, 0x16, 0xd3, 0xca, 0xac, 0x62,
	0xde, 0xca, 0x6c, 0xc9, 0xf6, 0xe6, 0x32, 0xfd, 0x9c, 0xdc, 0xe7, 0x69, 0x3a, 0x82, 0xe6, 0x5b,
	0x98, 0xdc, 0x2e, 0xe9, 0xe1, 0x44, 0x7b, 0xba, 0xbb, 0x93, 0xcc, 0x32, 0xf6, 0x9b, 0xcc, 0xaa,
	0x99, 0x58, 0xbc, 0xb4, 0x3d, 0xb3, 0x4c, 0xfb, 0x64, 0x3b, 0x91, 0x22, 0x18, 0xf9, 0xea, 0x25,
	0x4f, 0x55, 0xf3, 0x3e, 0x86, 0xee, 0x2e, 0x87, 0x7e, 0x69, 0x4c, 0xdd, 0x0f, 0x27, 0x99, 0x55,
	0xac, 0xb8, 0xc9, 0x2c, 0x6a, 0x72, 0x0b, 0x8b, 0xb6, 0x57, 0xb4, 0x3c, 0xf9, 0x9e, 0x3c, 0x3e,
	0xf2, 0x15, 0xbf, 0xe4, 0x6a, 0x7c, 0x6a, 0x12, 0x5f, 0x48, 0x79, 0x2c, 0x02, 0xa0, 0x0f, 0xc8,
	0xd6, 0x57, 0x71, 0x00, 0xe7, 0x3c, 0x86, 0xa0, 0xb1, 0x41, 0x77, 0x48, 0xed, 0x4c, 0x5c, 0x40,
	0x7c, 0xc6, 0x87, 0xf0, 0x6a, 0xa4, 0x1a, 0xff, 0x6c, 0xce, 0x96, 0x4e, 0xd2, 0x17, 0x52, 0x0a,
	0xd9, 0x98, 0x6c, 0xd2, 0x3a, 0x21, 0xf9, 0xd2, 0x17, 0x70, 0xd5, 0xf8, 0x77, 0xf3, 0xe9, 0x7f,
	0x5b, 0xa4, 0x7a, 0x94, 0x24, 0x74, 0x48, 0xea, 0xdf, 0xc0, 0xf1, 0x80, 0xa9, 0x1e, 0xa8, 0x57,
	0x09, 0xc4, 0x27, 0xcf, 0x69, 0xcb, 0x31, 0xc7, 0xaf, 0xd8, 0x8a, 0x26, 0x7b, 0xf0, 0x63, 0xab,
	0xbd, 0x42, 0xc3, 0xb2, 0xc0, 0x83, 0xd4, 0xde, 0xff, 0xf9, 0xf7, 0xbf, 0x7f, 0xb9, 0xd7, 0xb2,
	0x1f, 0xb9, 0x2c, 0x49, 0xdc, 0x2b, 0xcc, 0x75, 0xc3, 0x3c, 0x38, 0x78, 0x56, 0x79, 0x42, 0x81,
	0x6c, 0x1b, 0xdc, 0x4b, 0x11, 0xf2, 0x98, 0xee, 0xad, 0x88, 0x43, 0x45, 0xb3, 0xd6, 0x88, 0xa9,
	0xdd, 0x46, 0xd4, 0x63, 0x7b, 0xa7, 0x88, 0x8a, 0xb4, 0xaa, 0x31, 0x3e, 0xa9, 0xf5, 0x40, 0xe1,
	0x4d, 0xd5, 0xd3, 0xa4, 0xad, 0x92, 0x3b, 0xae, 0x31, 0x65, 0x27, 0x5a, 0x17, 0x6a, 0xd2, 0xfb,
	0x48, 0xda, 0xa1, 0x75, 0x24, 0x99, 0x67, 0x3f, 0xd2, 0xa1, 0x7d, 0xf2, 0x70, 0x0a, 0x79, 0x0e,
	0x8a, 0xf1, 0x88, 0xee, 0x97, 0x44, 0x19, 0x59, 0xc3, 0xca, 0x37, 0x92, 0xda, 0x14, 0x41, 0x35,
	0x4a, 0xe6, 0x20, 0xfa, 0x03, 0x32, 0xba, 0x92, 0xc5, 0x41, 0xce, 0x38, 0x58, 0x4e, 0x40, 0xb9,
	0x00, 0x79, 0xa3, 0xe5, 0x36, 0xab, 0xaf, 0x55, 0x3a, 0xc0, 0xa1, 0xa1, 0x13, 0x87, 0x66, 0x95,
	0xc5, 0x98, 0xc9, 0xac, 0x9c, 0xdc, 0xac, 0x7a, 0x79, 0x72, 0x08, 0x31, 0x93, 0xfb, 0x9a, 0xbc,
	0xd3, 0x03, 0x75, 0xa6, 0xdf, 0x8a, 0xb4, 0xe5, 0xe0, 0xdb, 0xb1, 0x18, 0x82, 0x82, 0x99, 0x56,
	0x99, 0x76, 0xbb, 0x03, 0xf4, 0xd1, 0x10, 0x3b, 0x40, 0x0b, 0x76, 0xd0, 0x29, 0xa9, 0x9f, 0x37,
	0xb0, 0x4e, 0xbf, 0xdd, 0x80, 0x79, 0x8b, 0x63, 0x03, 0x09, 0xa9, 0x4f, 0x41, 0x1e, 0x44, 0x4c,
	0x41, 0x40, 0x0f, 0x4a, 0xf7, 0x8a, 0xfa, 0x5d, 0x70, 0x2d, 0xc4, 0xbd, 0x47, 0x69, 0x01, 0x27,
	0xf3, 0x78, 0x85, 0xc4, 0x63, 0xa6, 0x20, 0x14, 0x72, 0x8c, 0xdd, 0xd9, 0xcb, 0xe3, 0x9f, 0xea,
	0xc7, 0x03, 0x1e, 0x05, 0x25, 0x47, 0x41, 0x7b, 0x7a, 0x79, 0xc6, 0x32, 0xd5, 0xcf, 0x13, 0xa6,
	0x47, 0x7c, 0xbb, 0x40, 0xa5, 0x1f, 0x94, 0x13, 0x35, 0x6c, 0xad, 0x9c, 0xda, 0x8f, 0x10, 0x54,
	0xa7, 0x0f, 0x16, 0x40, 0xf9, 0x4d, 0x9b, 0x7d, 0xd8, 0xee, 0xf0, 0x10, 0xb5, 0x9d, 0xc5, 0xff,
	0x05, 0x4e, 0xf1, 0xab, 0xb8, 0xf2, 0x79, 0x4d, 0x21, 0x8a, 0xba, 0xbb, 0xaf, 0xff, 0xec, 0x6c,
	0xbc, 0xbe, 0xee, 0x54, 0x7e, 0xbb, 0xee, 0x54, 0xfe, 0xb8, 0xee, 0x54, 0x7e, 0xfd, 0xab, 0xb3,
	0xf1, 0x6d, 0x95, 0x25, 0xbc, 0xff, 0x36, 0x7e, 0x61, 0x3f, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x96, 0xa8, 0xe6, 0x7c, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppClient interface {
	// 微信小程序登录
	WeChatGetOpenID(ctx context.Context, in *api1.CodeReq, opts ...grpc.CallOption) (*api1.OpenIdRes, error)
	// 微信小程序登录
	WeChatLogin(ctx context.Context, in *api1.LoginReq, opts ...grpc.CallOption) (*api1.LoginRes, error)
	// 获取商品列表
	GetGoodsList(ctx context.Context, in *api.GoodsReq, opts ...grpc.CallOption) (*api.GoodsListRes, error)
	// 获取商品详情
	GetGoodsDetail(ctx context.Context, in *api.GoodsDetailReq, opts ...grpc.CallOption) (*api.GoodsRes, error)
	// 获取制造商详情
	GetBrandDetail(ctx context.Context, in *api.BrandsDetailReq, opts ...grpc.CallOption) (*api.BrandsDetailRes, error)
	// 获取制造商列表
	GetBrandList(ctx context.Context, in *api.BrandsListReq, opts ...grpc.CallOption) (*api.BrandListRes, error)
	// 获取专题详情
	GetTopic(ctx context.Context, in *api2.TopicReq, opts ...grpc.CallOption) (*api2.TopicRes, error)
	// 获取专题列表
	GetTopicList(ctx context.Context, in *api2.TopicListReq, opts ...grpc.CallOption) (*api2.TopicListRes, error)
	// 获取相似专题列表
	GetTopicRelated(ctx context.Context, in *api2.TopicRelatedReq, opts ...grpc.CallOption) (*api2.TopicListRes, error)
	// 获取子商品分类列表
	GetCategoryList(ctx context.Context, in *api.CategoryChildReq, opts ...grpc.CallOption) (*api.CateGoryListRes, error)
	// 获取商品分类详情
	GetCategory(ctx context.Context, in *api.CategoryReq, opts ...grpc.CallOption) (*api.CategoryRes, error)
	// 获取商品售卖信息
	GetGoodsSell(ctx context.Context, in *api.GoodsDetailReq, opts ...grpc.CallOption) (*GoodsSellRes, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) WeChatGetOpenID(ctx context.Context, in *api1.CodeReq, opts ...grpc.CallOption) (*api1.OpenIdRes, error) {
	out := new(api1.OpenIdRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/WeChatGetOpenID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) WeChatLogin(ctx context.Context, in *api1.LoginReq, opts ...grpc.CallOption) (*api1.LoginRes, error) {
	out := new(api1.LoginRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/WeChatLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetGoodsList(ctx context.Context, in *api.GoodsReq, opts ...grpc.CallOption) (*api.GoodsListRes, error) {
	out := new(api.GoodsListRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetGoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetGoodsDetail(ctx context.Context, in *api.GoodsDetailReq, opts ...grpc.CallOption) (*api.GoodsRes, error) {
	out := new(api.GoodsRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetGoodsDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetBrandDetail(ctx context.Context, in *api.BrandsDetailReq, opts ...grpc.CallOption) (*api.BrandsDetailRes, error) {
	out := new(api.BrandsDetailRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetBrandDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetBrandList(ctx context.Context, in *api.BrandsListReq, opts ...grpc.CallOption) (*api.BrandListRes, error) {
	out := new(api.BrandListRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetBrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetTopic(ctx context.Context, in *api2.TopicReq, opts ...grpc.CallOption) (*api2.TopicRes, error) {
	out := new(api2.TopicRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetTopicList(ctx context.Context, in *api2.TopicListReq, opts ...grpc.CallOption) (*api2.TopicListRes, error) {
	out := new(api2.TopicListRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetTopicList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetTopicRelated(ctx context.Context, in *api2.TopicRelatedReq, opts ...grpc.CallOption) (*api2.TopicListRes, error) {
	out := new(api2.TopicListRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetTopicRelated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetCategoryList(ctx context.Context, in *api.CategoryChildReq, opts ...grpc.CallOption) (*api.CateGoryListRes, error) {
	out := new(api.CateGoryListRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetCategory(ctx context.Context, in *api.CategoryReq, opts ...grpc.CallOption) (*api.CategoryRes, error) {
	out := new(api.CategoryRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) GetGoodsSell(ctx context.Context, in *api.GoodsDetailReq, opts ...grpc.CallOption) (*GoodsSellRes, error) {
	out := new(GoodsSellRes)
	err := c.cc.Invoke(ctx, "/app.service.v1.App/GetGoodsSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
type AppServer interface {
	// 微信小程序登录
	WeChatGetOpenID(context.Context, *api1.CodeReq) (*api1.OpenIdRes, error)
	// 微信小程序登录
	WeChatLogin(context.Context, *api1.LoginReq) (*api1.LoginRes, error)
	// 获取商品列表
	GetGoodsList(context.Context, *api.GoodsReq) (*api.GoodsListRes, error)
	// 获取商品详情
	GetGoodsDetail(context.Context, *api.GoodsDetailReq) (*api.GoodsRes, error)
	// 获取制造商详情
	GetBrandDetail(context.Context, *api.BrandsDetailReq) (*api.BrandsDetailRes, error)
	// 获取制造商列表
	GetBrandList(context.Context, *api.BrandsListReq) (*api.BrandListRes, error)
	// 获取专题详情
	GetTopic(context.Context, *api2.TopicReq) (*api2.TopicRes, error)
	// 获取专题列表
	GetTopicList(context.Context, *api2.TopicListReq) (*api2.TopicListRes, error)
	// 获取相似专题列表
	GetTopicRelated(context.Context, *api2.TopicRelatedReq) (*api2.TopicListRes, error)
	// 获取子商品分类列表
	GetCategoryList(context.Context, *api.CategoryChildReq) (*api.CateGoryListRes, error)
	// 获取商品分类详情
	GetCategory(context.Context, *api.CategoryReq) (*api.CategoryRes, error)
	// 获取商品售卖信息
	GetGoodsSell(context.Context, *api.GoodsDetailReq) (*GoodsSellRes, error)
}

// UnimplementedAppServer can be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (*UnimplementedAppServer) WeChatGetOpenID(ctx context.Context, req *api1.CodeReq) (*api1.OpenIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeChatGetOpenID not implemented")
}
func (*UnimplementedAppServer) WeChatLogin(ctx context.Context, req *api1.LoginReq) (*api1.LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeChatLogin not implemented")
}
func (*UnimplementedAppServer) GetGoodsList(ctx context.Context, req *api.GoodsReq) (*api.GoodsListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsList not implemented")
}
func (*UnimplementedAppServer) GetGoodsDetail(ctx context.Context, req *api.GoodsDetailReq) (*api.GoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsDetail not implemented")
}
func (*UnimplementedAppServer) GetBrandDetail(ctx context.Context, req *api.BrandsDetailReq) (*api.BrandsDetailRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandDetail not implemented")
}
func (*UnimplementedAppServer) GetBrandList(ctx context.Context, req *api.BrandsListReq) (*api.BrandListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrandList not implemented")
}
func (*UnimplementedAppServer) GetTopic(ctx context.Context, req *api2.TopicReq) (*api2.TopicRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (*UnimplementedAppServer) GetTopicList(ctx context.Context, req *api2.TopicListReq) (*api2.TopicListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicList not implemented")
}
func (*UnimplementedAppServer) GetTopicRelated(ctx context.Context, req *api2.TopicRelatedReq) (*api2.TopicListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicRelated not implemented")
}
func (*UnimplementedAppServer) GetCategoryList(ctx context.Context, req *api.CategoryChildReq) (*api.CateGoryListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryList not implemented")
}
func (*UnimplementedAppServer) GetCategory(ctx context.Context, req *api.CategoryReq) (*api.CategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (*UnimplementedAppServer) GetGoodsSell(ctx context.Context, req *api.GoodsDetailReq) (*GoodsSellRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoodsSell not implemented")
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_WeChatGetOpenID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api1.CodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).WeChatGetOpenID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/WeChatGetOpenID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).WeChatGetOpenID(ctx, req.(*api1.CodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_WeChatLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api1.LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).WeChatLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/WeChatLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).WeChatLogin(ctx, req.(*api1.LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetGoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetGoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetGoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetGoodsList(ctx, req.(*api.GoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetGoodsDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GoodsDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetGoodsDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetGoodsDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetGoodsDetail(ctx, req.(*api.GoodsDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetBrandDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.BrandsDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetBrandDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetBrandDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetBrandDetail(ctx, req.(*api.BrandsDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetBrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.BrandsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetBrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetBrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetBrandList(ctx, req.(*api.BrandsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api2.TopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetTopic(ctx, req.(*api2.TopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetTopicList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api2.TopicListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetTopicList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetTopicList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetTopicList(ctx, req.(*api2.TopicListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetTopicRelated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api2.TopicRelatedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetTopicRelated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetTopicRelated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetTopicRelated(ctx, req.(*api2.TopicRelatedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CategoryChildReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetCategoryList(ctx, req.(*api.CategoryChildReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.CategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetCategory(ctx, req.(*api.CategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_GetGoodsSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(api.GoodsDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetGoodsSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.service.v1.App/GetGoodsSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetGoodsSell(ctx, req.(*api.GoodsDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.service.v1.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WeChatGetOpenID",
			Handler:    _App_WeChatGetOpenID_Handler,
		},
		{
			MethodName: "WeChatLogin",
			Handler:    _App_WeChatLogin_Handler,
		},
		{
			MethodName: "GetGoodsList",
			Handler:    _App_GetGoodsList_Handler,
		},
		{
			MethodName: "GetGoodsDetail",
			Handler:    _App_GetGoodsDetail_Handler,
		},
		{
			MethodName: "GetBrandDetail",
			Handler:    _App_GetBrandDetail_Handler,
		},
		{
			MethodName: "GetBrandList",
			Handler:    _App_GetBrandList_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _App_GetTopic_Handler,
		},
		{
			MethodName: "GetTopicList",
			Handler:    _App_GetTopicList_Handler,
		},
		{
			MethodName: "GetTopicRelated",
			Handler:    _App_GetTopicRelated_Handler,
		},
		{
			MethodName: "GetCategoryList",
			Handler:    _App_GetCategoryList_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _App_GetCategory_Handler,
		},
		{
			MethodName: "GetGoodsSell",
			Handler:    _App_GetGoodsSell_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func (m *GoodsSellRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoodsSellRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GoodsSellRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ProductList) > 0 {
		for iNdEx := len(m.ProductList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProductList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Issue) > 0 {
		for iNdEx := len(m.Issue) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issue[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Attribute) > 0 {
		for iNdEx := len(m.Attribute) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attribute[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.GoodsGallery) > 0 {
		for iNdEx := len(m.GoodsGallery) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GoodsGallery[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintApi(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.GoodsSize) > 0 {
		for iNdEx := len(m.GoodsSize) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GoodsSize[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *GoodsSellRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GoodsSize) > 0 {
		for _, e := range m.GoodsSize {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if len(m.GoodsGallery) > 0 {
		for _, e := range m.GoodsGallery {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if len(m.Attribute) > 0 {
		for _, e := range m.Attribute {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if len(m.Issue) > 0 {
		for _, e := range m.Issue {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
	}
	if len(m.ProductList) > 0 {
		for _, e := range m.ProductList {
			l = e.Size()
			n += 1 + l + sovApi(uint64(l))
		}
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
func (m *GoodsSellRes) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GoodsSellRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoodsSellRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodsSize", wireType)
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
			m.GoodsSize = append(m.GoodsSize, &api.GoodsSize{})
			if err := m.GoodsSize[len(m.GoodsSize)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoodsGallery", wireType)
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
			m.GoodsGallery = append(m.GoodsGallery, &api.GoodsGallery{})
			if err := m.GoodsGallery[len(m.GoodsGallery)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attribute", wireType)
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
			m.Attribute = append(m.Attribute, &api.Attribute{})
			if err := m.Attribute[len(m.Attribute)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issue", wireType)
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
			m.Issue = append(m.Issue, &api.GoodsIssue{})
			if err := m.Issue[len(m.Issue)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProductList", wireType)
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
			m.ProductList = append(m.ProductList, &api.Product{})
			if err := m.ProductList[len(m.ProductList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
