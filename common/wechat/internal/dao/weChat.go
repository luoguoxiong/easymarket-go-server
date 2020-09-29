package dao

import (
	pb "easymarket-go-server/common/wechat/api"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/pkg/log"
	"io/ioutil"
	"net/http"
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
func (d *Dao) UserLogin(login *pb.LoginReq)(user *pb.LoginRes,err error){
	user = &pb.LoginRes{}
	err = d.db.Table("easymarket_user").Where("openId=?", login.OpenID).Find(user).Error
	fmt.Println(user.ID)
 return
}
