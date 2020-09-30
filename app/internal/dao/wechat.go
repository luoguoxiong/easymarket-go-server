package dao

import (
	"context"
	wechat "easymarket-go-server/common/wechat/api"
)

//WeChatLogin 微信小程序登录
func (d *Dao) WeChatLogin(ctx context.Context, req *wechat.LoginReq) (resp *wechat.LoginRes, err error) {
	return d.wechatClient.Login(ctx, req)
}

// WeChatGetOpenID 获取openId
func (d *Dao) WeChatGetOpenID(ctx context.Context, req *wechat.CodeReq) (resp *wechat.OpenIdRes, err error) {
	return d.wechatClient.GetWeChatOpenID(ctx, req)
}
