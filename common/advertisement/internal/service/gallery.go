package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetGoodsGallery 获取商品轮播图
func (s *Service) GetGoodsGallery(ctx context.Context, req *pb.GoodsDetailReq) (res *pb.GoodsGalleryList, err error) {
	return s.dao.GetGallery(req.Id)
}
