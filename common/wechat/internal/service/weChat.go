package service

import (
	"context"
	pb "easymarket-go-server/common/wechat/api"
)

//GetWeChatOpenID 获取OpenId
func (s *Service) GetWeChatOpenID(ctx context.Context, req *pb.CodeReq) (resp *pb.OpenIdRes, err error) {
	resp, err = s.dao.GetWeChatOpenID(req.Code)
	return
}

//Login 微信小程序登录
func (s *Service) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginRes, err error) {
	resp,err = s.dao.UserLogin(req)
	return
}
