package service

import (
	"context"
	wechat "easymarket-go-server/common/wechat/api"
)

//WeChatLogin 微信小程序登录
func (s *Service) WeChatLogin(ctx context.Context, req *wechat.LoginReq) (resp *wechat.LoginRes, err error) {
	return s.dao.WeChatLogin(ctx, req)
}

// WeChatGetOpenID 获取openId
func (s *Service) WeChatGetOpenID(ctx context.Context, req *wechat.CodeReq) (resp *wechat.OpenIdRes, err error) {
	return s.dao.WeChatGetOpenID(ctx, req)
}
