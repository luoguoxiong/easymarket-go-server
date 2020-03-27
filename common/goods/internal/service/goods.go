package service

import (
	"context"
	pb "easymarketgoserve/common/goods/api"
)

// GetGoodsDetail 获取商品列表
func (s *Service) GetGoodsDetail(ctx context.Context, req *pb.GoodsReq) (res *pb.GoodsRes, err error) {
	return res, nil
}

// GetGoodsList 获取商品列表
func (s *Service) GetGoodsList(ctx context.Context, req *pb.GoodsReq) (res *pb.GoodsListRes, err error) {
	goods, count := s.dao.GetGoodsList(req.IsHot, req.Page, req.Size_)
	res = &pb.GoodsListRes{
		Total:     count,
		GoodsList: goods,
	}
	return res, nil
}
