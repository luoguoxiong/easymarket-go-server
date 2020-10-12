package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetGoodsAttribute 获取商品常见疑问
func (s *Service) GetGoodsAttribute(ctx context.Context, req *pb.GoodsDetailReq) (res *pb.AttributeList, err error) {
	return s.dao.GetAttruibute(req.Id)
}
