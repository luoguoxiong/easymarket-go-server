package service

import (
	pb "easymarket-go-server/common/topic/api"
	"easymarket-go-server/common/topic/internal/dao"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, wire.Bind(new(pb.TopicServer), new(*Service)))

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
