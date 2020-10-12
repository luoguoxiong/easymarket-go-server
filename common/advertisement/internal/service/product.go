package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetGoodsProductList 获取商品产品
func (s *Service) GetGoodsProductList(ctx context.Context, req *pb.GoodsDetailReq) (res *pb.GoodsProductList, err error) {
	return s.dao.GetProduct(req.Id)
}
