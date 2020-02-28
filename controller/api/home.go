package api

import (
	"easymarket/controller/response"
	"easymarket/modle"

	"github.com/gin-gonic/gin"
)

type homeRes struct {
	NewGoodsList []modle.Goods   `json:"newGoodsList"` // 新品
	Banner       []modle.Bannel  `json:"banner"`       // 轮播图
	Channel      []modle.Channel `json:"channel"`      // 渠道
	HotGoodsList []modle.Goods   `json:"hotGoodsList"` // 热销商品
}

// GetHome 获取首页数据
func GetHome(c *gin.Context) {
	c1 := make(chan []modle.Goods, 1)
	c2 := make(chan []modle.Bannel, 1)
	c3 := make(chan []modle.Channel, 1)
	c4 := make(chan []modle.Goods, 1)
	var g modle.Goods
	var b modle.Bannel
	var ch modle.Channel
	go func() {
		v, _ := g.NewGoods(1, 5)
		c1 <- v
	}()
	go func() {
		v, _ := b.GetBannel(3)
		c2 <- v
	}()
	go func() {
		v, _ := ch.GetChannel(5)
		c3 <- v
	}()
	go func() {
		v, _ := g.HotGoods(1, 3)
		c4 <- v
	}()

	res := &homeRes{
		NewGoodsList: <-c1,
		Banner:       <-c2,
		Channel:      <-c3,
		HotGoodsList: <-c4,
	}
	response.ReSuccess(c, res, "")
}
