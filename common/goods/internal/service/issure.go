package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetGoodsIssueList 获取商品常见疑问
func (s *Service) GetGoodsIssueList(ctx context.Context, req *pb.GoodsDetailReq) (res *pb.GoodsIssueList, err error) {
	res, err = s.dao.GetIssure(req.Id)
	return
}
