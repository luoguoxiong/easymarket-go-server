package http

import (
	"easymarket-go-server/app/api"
)

var errcode = map[int]string{
	api.TokenTimeOut.Code(): "登录已过期",
	api.TokenIsError.Code(): "token解析失败",
	api.TokenIsNew.Code():   "token已失效",
}
