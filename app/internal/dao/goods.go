package dao

import (
	"context"
	goods "easymarket-go-server/common/goods/api"
)

// GetGoodsList 获取商品列表
func (d *Dao) GetGoodsList(ctx context.Context, req *goods.GoodsReq) (res *goods.GoodsListRes, err error) {
	res, err = d.goodsClient.GetGoodsList(ctx, req)
	// 不是很喜欢这样处理slice零值问题...
	goodsList := make([]*goods.GoodsRes, 0)
	if len(res.GoodsList) > 0 {
		goodsList = res.GoodsList
	}
	res = &goods.GoodsListRes{
		GoodsList: goodsList,
		Total:     res.Total,
	}
	return
}

// GetGoodsDetail 获取商品详情
func (d *Dao) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailReq) (res *goods.GoodsRes, err error) {
	res, err = d.goodsClient.GetGoodsDetail(ctx, req)
	return
}

// GetBrandList 获取制造商列表
func (d *Dao) GetBrandList(ctx context.Context, req *goods.BrandsListReq) (res *goods.BrandListRes, err error) {
	res, err = d.goodsClient.GetBrandList(ctx, req)
	return
}

// GetBrandDetail 获取制造商详情
func (d *Dao) GetBrandDetail(ctx context.Context, req *goods.BrandsDetailReq) (res *goods.BrandsDetailRes, err error) {
	res, err = d.goodsClient.GetBrandDetail(ctx, req)
	return
}

// GetCategoryList 获取子商品分类列表
func (d *Dao) GetCategoryList(ctx context.Context, req *goods.CategoryChildReq) (resp *goods.CateGoryListRes, err error) {
	resp, err = d.goodsClient.GetCategoryList(ctx, req)
	return
}

// GetCategory 获取商品分类详情
func (d *Dao) GetCategory(ctx context.Context, req *goods.CategoryReq) (resp *goods.CategoryRes, err error) {
	resp, err = d.goodsClient.GetCategory(ctx, req)
	return
}
