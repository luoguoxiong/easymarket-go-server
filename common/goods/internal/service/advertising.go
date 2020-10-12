package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// GetAdsSwiper 获取首页广告轮播图
func (s *Service) GetAdsSwiper(ctx context.Context, e *empty.Empty) (res *pb.AdsSwiper, err error) {
	res, err = s.dao.GetAdsSwiper()
	return
}
