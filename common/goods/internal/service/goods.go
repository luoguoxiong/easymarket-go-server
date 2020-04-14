package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetGoodsDetail 获取商品详情
func (s *Service) GetGoodsDetail(ctx context.Context, req *pb.GoodsDetailReq) (res *pb.GoodsRes, err error) {
	res, err = s.dao.GetGoodsDetail(req.Id)
	return
}

// GetGoodsList 获取商品列表
func (s *Service) GetGoodsList(ctx context.Context, req *pb.GoodsReq) (res *pb.GoodsListRes, err error) {
	res, err = s.dao.GetGoodsList(req.IsHot, req.IsNew, req.CategoryId, req.Page, req.Size_)
	return
}
