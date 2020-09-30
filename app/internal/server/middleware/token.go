package middleware

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"time"
)

const secret = "secret"

type jwtCustomClaims struct {
	jwt.StandardClaims
	UserID int32 `json:"userId"`
}

// TokenValidate 用户token校验
func TokenValidate(ctx *bm.Context) {
	Authorization := ctx.Request.Header.Get("Authorization")
	str, _ := GenerateToken()
	fmt.Println(str)
	// fmt.Println(Authorization)
	claims, _ := ParseToken(Authorization)

	startTime := int64(claims.(jwt.MapClaims)["exp"].(float64))

	now := time.Now().Unix()

	if now >= startTime {
		fmt.Println("超时")
	}

	fmt.Println(claims)

	ctx.Next()
}

// GenerateToken 生产token
func GenerateToken() (string, error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Minute * 1).Unix()),
		},
		173,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken /
func ParseToken(tokenSrt string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return secret, nil
	})
	claims = token.Claims
	return
}
