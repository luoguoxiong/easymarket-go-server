package dao

import (
	"context"
	goods "easymarket-go-server/common/goods/api"
)

// GetGoodsList 获取商品列表
func (d *Dao) GetGoodsList(ctx context.Context, req *goods.GoodsReq) (res *goods.GoodsListRes, err error) {
	res, err = d.goodsGrpc.GetGoodsList(ctx, req)
	return
}

// GetGoodsDetail 获取商品详情
func (d *Dao) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailReq) (res *goods.GoodsRes, err error) {
	res, err = d.goodsGrpc.GetGoodsDetail(ctx, req)
	return
}

// GetBrandList 获取制造商列表
func (d *Dao) GetBrandList(ctx context.Context, req *goods.BrandsListReq) (res *goods.BrandListRes, err error) {
	res, err = d.goodsGrpc.GetBrandList(ctx, req)
	return
}

// GetBrandDetail 获取制造商详情
func (d *Dao) GetBrandDetail(ctx context.Context, req *goods.BrandsDetailReq) (res *goods.BrandsDetailRes, err error) {
	res, err = d.goodsGrpc.GetBrandDetail(ctx, req)
	return
}
