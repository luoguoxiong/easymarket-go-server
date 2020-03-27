package service

import (
	pb "easymarketgoserve/common/goods/api"
	"easymarketgoserve/common/goods/internal/dao"

	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, wire.Bind(new(pb.GoodsServer), new(*Service)))

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
