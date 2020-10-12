package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetBrandList 获取制造商列表
func (s *Service) GetBrandList(ctx context.Context, req *pb.BrandsListReq) (res *pb.BrandListRes, err error) {
	res, err = s.dao.GetBrandList(req.Page, req.Size_)
	return
}

// GetBrandDetail 获取制造商详情
func (s *Service) GetBrandDetail(ctx context.Context, req *pb.BrandsDetailReq) (res *pb.BrandsDetailRes, err error) {
	res, err = s.dao.GetBrandDetail(req.Id)
	return
}
