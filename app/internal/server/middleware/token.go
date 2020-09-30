package middleware

import (
	"easymarket-go-server/app/api"
	"easymarket-go-server/app/internal/service"
	"easymarket-go-server/libary"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"time"
)

// TokenValidate token校验
func TokenValidate(s *service.Service) func(ctx *bm.Context) {
	return func(ctx *bm.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "" {

			tokenData, err := libary.ParseToken(token)
			if err != nil {
				ctx.JSON(nil, api.TokenIsError)
				ctx.Abort()
				return
			}

			now := time.Now().Unix()
			if now >= tokenData.Exp {
				ctx.JSON(nil, api.TokenTimeOut)
				ctx.Abort()
				return
			}

			oldToken := s.GetTokenByUserID(tokenData.UserID)
			if token != oldToken {
				ctx.JSON(nil, api.TokenIsNew)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
