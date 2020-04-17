// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

package api

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathGoodsGetGoodsDetail = "/goods"
var PathGoodsGetGoodsList = "/goods/list"
var PathGoodsGetBrandDetail = "/brand"
var PathGoodsGetBrandList = "/brand/list"
var PathGoodsGetCategoryList = "/category/list"
var PathGoodsGetCategory = "/category"
var PathGoodsGetGoodsSize = "/goods/sizeList"
var PathGoodsGetGoodsGallery = "/goods/gallery"
var PathGoodsGetGoodsAttribute = "/goods/attribute"
var PathGoodsGetGoodsIssueList = "/goods/issue"
var PathGoodsGetGoodsProductList = "/goods/product"

// GoodsBMServer is the server API for Goods service.
type GoodsBMServer interface {
	// 获取商品详情
	GetGoodsDetail(ctx context.Context, req *GoodsDetailReq) (resp *GoodsRes, err error)

	// 获取商品列表
	GetGoodsList(ctx context.Context, req *GoodsReq) (resp *GoodsListRes, err error)

	// 获取制造商详情
	GetBrandDetail(ctx context.Context, req *BrandsDetailReq) (resp *BrandsDetailRes, err error)

	// 获取制造商列表
	GetBrandList(ctx context.Context, req *BrandsListReq) (resp *BrandListRes, err error)

	// 获取子商品分类列表
	GetCategoryList(ctx context.Context, req *CategoryChildReq) (resp *CateGoryListRes, err error)

	// 获取商品分类详情
	GetCategory(ctx context.Context, req *CategoryReq) (resp *CategoryRes, err error)

	// 获取商品可选规格
	GetGoodsSize(ctx context.Context, req *GoodsDetailReq) (resp *GoodsSizeList, err error)

	// 获取商品轮播图
	GetGoodsGallery(ctx context.Context, req *GoodsDetailReq) (resp *GoodsGalleryList, err error)

	// 获取商品特点
	GetGoodsAttribute(ctx context.Context, req *GoodsDetailReq) (resp *AttributeList, err error)

	// 获取商品常见疑问
	GetGoodsIssueList(ctx context.Context, req *GoodsDetailReq) (resp *GoodsIssueList, err error)

	// 获取商品产品列表
	GetGoodsProductList(ctx context.Context, req *GoodsDetailReq) (resp *GoodsProductList, err error)
}

var GoodsSvc GoodsBMServer

func goodsGetGoodsDetail(c *bm.Context) {
	p := new(GoodsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetGoodsDetail(c, p)
	c.JSON(resp, err)
}

func goodsGetGoodsList(c *bm.Context) {
	p := new(GoodsReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetGoodsList(c, p)
	c.JSON(resp, err)
}

func goodsGetBrandDetail(c *bm.Context) {
	p := new(BrandsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetBrandDetail(c, p)
	c.JSON(resp, err)
}

func goodsGetBrandList(c *bm.Context) {
	p := new(BrandsListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetBrandList(c, p)
	c.JSON(resp, err)
}

func goodsGetCategoryList(c *bm.Context) {
	p := new(CategoryChildReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetCategoryList(c, p)
	c.JSON(resp, err)
}

func goodsGetCategory(c *bm.Context) {
	p := new(CategoryReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetCategory(c, p)
	c.JSON(resp, err)
}

func goodsGetGoodsSize(c *bm.Context) {
	p := new(GoodsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetGoodsSize(c, p)
	c.JSON(resp, err)
}

func goodsGetGoodsGallery(c *bm.Context) {
	p := new(GoodsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetGoodsGallery(c, p)
	c.JSON(resp, err)
}

func goodsGetGoodsAttribute(c *bm.Context) {
	p := new(GoodsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetGoodsAttribute(c, p)
	c.JSON(resp, err)
}

func goodsGetGoodsIssueList(c *bm.Context) {
	p := new(GoodsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetGoodsIssueList(c, p)
	c.JSON(resp, err)
}

func goodsGetGoodsProductList(c *bm.Context) {
	p := new(GoodsDetailReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := GoodsSvc.GetGoodsProductList(c, p)
	c.JSON(resp, err)
}

// RegisterGoodsBMServer Register the blademaster route
func RegisterGoodsBMServer(e *bm.Engine, server GoodsBMServer) {
	GoodsSvc = server
	e.GET("/goods", goodsGetGoodsDetail)
	e.GET("/goods/list", goodsGetGoodsList)
	e.GET("/brand", goodsGetBrandDetail)
	e.GET("/brand/list", goodsGetBrandList)
	e.GET("/category/list", goodsGetCategoryList)
	e.GET("/category", goodsGetCategory)
	e.GET("/goods/sizeList", goodsGetGoodsSize)
	e.GET("/goods/gallery", goodsGetGoodsGallery)
	e.GET("/goods/attribute", goodsGetGoodsAttribute)
	e.GET("/goods/issue", goodsGetGoodsIssueList)
	e.GET("/goods/product", goodsGetGoodsProductList)
}
