package service

import (
	"context"

	goods "easymarket-go-server/common/goods/api"
)

//GetGoodsList 获取商品列表
func (s *Service) GetGoodsList(ctx context.Context, req *goods.GoodsReq) (res *goods.GoodsListRes, err error) {
	res, err = s.dao.GetGoodsList(ctx, req)
	return
}

//GetGoodsDetail 商品详情
func (s *Service) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailReq) (res *goods.GoodsRes, err error) {
	res, err = s.dao.GetGoodsDetail(ctx, req)
	return
}

// GetBrandList 获取制造商列表
func (s *Service) GetBrandList(ctx context.Context, req *goods.BrandsListReq) (res *goods.BrandListRes, err error) {
	res, err = s.dao.GetBrandList(ctx, req)
	return
}

// GetBrandDetail 获取制造商详情
func (s *Service) GetBrandDetail(ctx context.Context, req *goods.BrandsDetailReq) (res *goods.BrandsDetailRes, err error) {
	res, err = s.dao.GetBrandDetail(ctx, req)
	return
}

// GetCategoryList 获取子商品分类列表
func (s *Service) GetCategoryList(ctx context.Context, req *goods.CategoryChildReq) (res *goods.CateGoryListRes, err error) {
	res, err = s.dao.GetCategoryList(ctx, req)
	return
}

// GetCategory 获取商品分类详情
func (s *Service) GetCategory(ctx context.Context, req *goods.CategoryReq) (res *goods.CategoryRes, err error) {
	res, err = s.dao.GetCategory(ctx, req)
	return
}
