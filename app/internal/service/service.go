package service

import (
	"context"

	pb "easymarketgoserve/app/api"
	"easymarketgoserve/app/internal/dao"

	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/google/wire"

	goods_service_v1 "easymarketgoserve/common/goods/api"
)

// Provider ...
var Provider = wire.NewSet(New, wire.Bind(new(pb.AppServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao *dao.Dao
}

// New new a service and return.
func New(d *dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = func() {}
	err = paladin.Watch("application.toml", s.ac)
	return
}

//GetGoodsList ...
func (s *Service) GetGoodsList(ctx context.Context, req *goods_service_v1.GoodsReq) (res *goods_service_v1.GoodsListRes, err error) {
	res, err = s.dao.GetGoodsDetail(ctx, req)
	return
}
