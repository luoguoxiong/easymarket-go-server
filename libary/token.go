package libary

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const secret = "easymarket"

type jwtCustomClaims struct {
	jwt.StandardClaims
	UserID int `json:"UserID"`
}

// ParsToken token解析
type ParsToken struct {
	Exp    int64 // 过期时间
	UserID int   // 用户Id
}

// GenerateToken 生成token
func GenerateToken(userID int, exp time.Duration) (string, error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(exp).Unix()),
		},
		userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken 解析Topken
func ParseToken(tokenSrt string) (parsToken *ParsToken, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	fmt.Println(token)
	if err != nil {
		return
	}
	parsToken = &ParsToken{
		int64(token.Claims.(jwt.MapClaims)["exp"].(float64)),
		int(token.Claims.(jwt.MapClaims)["UserID"].(float64)),
	}
	return
}
