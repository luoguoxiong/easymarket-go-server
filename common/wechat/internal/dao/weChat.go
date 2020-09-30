package dao

import (
	pb "easymarket-go-server/common/wechat/api"
	"easymarket-go-server/common/wechat/constvariable"
	"easymarket-go-server/libary"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func genHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 3 * time.Second,
	}
}

// GetWeChatOpenID 获取微信OpenId
func (d *Dao) GetWeChatOpenID(code string) (sessionData *pb.OpenIdRes, err error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	url = fmt.Sprintf(url, d.wechat.Appid, d.wechat.Secret, code)

	resp, err := genHTTPClient().Get(url)

	if err != nil {
		log.Error("code2Session err:%s, appId:%v, code:%v", err, d.wechat.Appid, code)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("code2Session ioutil err:%s, appId:%v, code:%v", err, d.wechat.Appid, code)
		return
	}

	sessionData = &pb.OpenIdRes{}

	err = json.Unmarshal(body, &sessionData)

	if err != nil {
		log.Error("Code2Session Unmarshal err body:%s, errInfo:%s, code:%v", string(body), err, code)
		return
	}

	return
}

// UserLogin 用户登录
func (d *Dao) UserLogin(login *pb.LoginReq) (user *pb.LoginRes, err error) {
	user = &pb.LoginRes{}
	err = d.db.Table("easymarket_user").Where("openId=?", login.OpenID).Find(user).Error
	// 如果没有则是新用户
	if err == gorm.ErrRecordNotFound {
		err = d.db.Table("easymarket_user").Select("nickname", "openId").Create(&pb.LoginReq{
			NickName: login.NickName,
			OpenID:   login.OpenID,
		}).Error
		err = d.db.Table("easymarket_user").Where("openId=?", login.OpenID).Find(user).Error
	} else {
		err = d.db.Table("easymarket_user").Select("nickname").Updates(login).Error
		user.NickName = login.NickName
	}
	var token string
	token, err = libary.GenerateToken(int(user.ID), constvariable.TOKENTIMEOUT)

	user.Token = token

	key := fmt.Sprintf("TOKEN_%s", strconv.Itoa(int(user.ID)))

	err = d.redis.Set(key, token, constvariable.TOKENTIMEOUT).Err()
	if err != nil {
		return nil, err
	}
	return
}
