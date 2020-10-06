package dao

import (
	"fmt"
	"strconv"
)

// GetTokenByUserID 获取用户Token
func (d *Dao) GetTokenByUserID(userID int) (token string) {
	key := fmt.Sprintf("token:user:%s", strconv.Itoa(userID))
	token, err := d.redis.Get(key).Result()
	if err != nil {
		return ""
	}
	return
}
