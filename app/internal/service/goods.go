package service

import (
	"context"

	goods_service_v1 "easymarket-go-server/common/goods/api"
)

//GetGoodsList 获取商品列表
func (s *Service) GetGoodsList(ctx context.Context, req *goods_service_v1.GoodsReq) (res *goods_service_v1.GoodsListRes, err error) {
	res, err = s.dao.GetGoodsList(ctx, req)
	return
}

//GetGoodsDetail 商品详情
func (s *Service) GetGoodsDetail(ctx context.Context, req *goods_service_v1.GoodsDetailReq) (res *goods_service_v1.GoodsRes, err error) {
	res, err = s.dao.GetGoodsDetail(ctx, req)
	return
}

// GetBrandList 获取制造商列表
func (s *Service) GetBrandList(ctx context.Context, req *goods_service_v1.BrandsListReq) (res *goods_service_v1.BrandListRes, err error) {
	res, err = s.dao.GetBrandList(ctx, req)
	return
}

// GetBrandDetail 获取制造商详情
func (s *Service) GetBrandDetail(ctx context.Context, req *goods_service_v1.BrandsDetailReq) (res *goods_service_v1.BrandsDetailRes, err error) {
	res, err = s.dao.GetBrandDetail(ctx, req)
	return
}
