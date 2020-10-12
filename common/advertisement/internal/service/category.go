package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetCategoryList  获取子商品分类列表
func (s *Service) GetCategoryList(ctx context.Context, req *pb.CategoryChildReq) (res *pb.CateGoryListRes, err error) {
	res, err = s.dao.GetCategoryList(req.ParentId)
	return
}

// GetCategory 获取商品分类详情
func (s *Service) GetCategory(ctx context.Context, req *pb.CategoryReq) (res *pb.CategoryRes, err error) {
	res, err = s.dao.GetCategory(req.Id)
	return
}
