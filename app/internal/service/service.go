package service

import (
	"context"
	"fmt"

	pb "easymarketgoserve/app/api"
	"easymarketgoserve/app/internal/dao"

	"github.com/bilibili/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

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
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)
	res, _ := s.dao.GetGoodsDetail(ctx)
	fmt.Printf("hello %s", req.Name)
	fmt.Println(res)
	return
}

// SayHelloURL bm demo func.
func (s *Service) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {

	res, _ := s.dao.GetGoodsDetail(ctx)
	fmt.Println(res)
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

// Close close the resource.
func (s *Service) Close() {
}
