package middleware
import (
	"fmt"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

const secret ="secret"

// TokenValidate 用户token校验
func TokenValidate(ctx *bm.Context) {
	// Authorization := ctx.Request.Header.Get("Authorization")
	str,_ := GenerateToken()

	fmt.Println(str)

	claims, err := ParseToken(str)
	if nil != err {
		fmt.Println(" err :", err)
	}
	fmt.Println("userId:", claims.(jwt.MapClaims)["userId"])

	ctx.Next()
}


// GenerateToken 生产token
func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": 172,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
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
