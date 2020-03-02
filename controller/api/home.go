package api

import (
	"easymarket/controller/response"
	"easymarket/modle"
	"sync"

	"github.com/gin-gonic/gin"
)

type categoryList struct {
	ID        int           `json:"id"`        // 分类ID
	Name      string        `json:"name"`      // 分类名称
	GoodsList []modle.Goods `json:"goodsList"` // 分类商品
}

type homeRes struct {
	NewGoodsList []modle.Goods   `json:"newGoodsList"` // 新品
	Banner       []modle.Bannel  `json:"banner"`       // 轮播图
	Channel      []modle.Channel `json:"channel"`      // 渠道
	HotGoodsList []modle.Goods   `json:"hotGoodsList"` // 热销商品
	BrandList    []modle.Brand   `json:"brandList"`    // 热门制造商
	TopicList    []modle.Topic   `json:"topicList"`    // 专题
	CategoryList []categoryList  `json:"categoryList"` // 分类商品
}

// GetHome 获取首页数据
func GetHome(c *gin.Context) {
	c1 := make(chan []modle.Goods, 1)
	c2 := make(chan []modle.Bannel, 1)
	c3 := make(chan []modle.Channel, 1)
	c4 := make(chan []modle.Goods, 1)
	c5 := make(chan []modle.Brand, 1)
	c6 := make(chan []modle.Topic, 1)
	c7 := make(chan []categoryList, 7)
	var g modle.Goods
	var b modle.Bannel
	var br modle.Brand
	var ch modle.Channel
	var t modle.Topic
	var ca modle.Category
	go getNewGoods(g, c1)
	go getBanners(b, c2)
	go getChannel(ch, c3)
	go getHotGoods(g, c4)
	go getBrands(br, c5)
	go getTopics(t, c6)
	go getCategoryGoods(ca, c7)
	res := homeRes{
		NewGoodsList: <-c1,
		Banner:       <-c2,
		Channel:      <-c3,
		HotGoodsList: <-c4,
		BrandList:    <-c5,
		TopicList:    <-c6,
		CategoryList: <-c7,
	}
	response.ReSuccess(c, res, "")
}

// 首页新品
func getNewGoods(g modle.Goods, c chan<- []modle.Goods) {
	v, _ := g.NewGoods(1, 5)
	c <- v
}

// 获取首页轮播图
func getBanners(b modle.Bannel, c chan<- []modle.Bannel) {
	v, _ := b.GetBannel(3)
	c <- v
}

// 获取首页渠道
func getChannel(ch modle.Channel, c chan<- []modle.Channel) {
	v, _ := ch.GetChannel(5)
	c <- v
}

// 获取热销商品
func getHotGoods(g modle.Goods, c chan<- []modle.Goods) {
	v, _ := g.HotGoods(1, 5)
	c <- v
}

// 获取制造商
func getBrands(b modle.Brand, c chan<- []modle.Brand) {
	v, _ := b.GetBrand(5)
	c <- v
}

// 获取专题
func getTopics(t modle.Topic, c chan<- []modle.Topic) {
	v, _ := t.GetTopic(1, 5)
	c <- v
}

// 获取分类推荐商品
func getCategoryGoods(ca modle.Category, ch chan<- []categoryList) {
	c, _ := ca.GetAllCategory(0)
	categoryLists := make([]categoryList, len(c))
	var wg sync.WaitGroup
	wg.Add(len(c))
	for key, v := range c {
		go func(index int, v modle.Category) {
			c2, _ := ca.GetAllCategory(v.ID)
			ids := []interface{}{}
			for _, v1 := range c2 {
				ids = append(ids, v1.ID)
			}
			var g modle.Goods
			gs, _ := g.GetTypeGoods(ids, 1, 7)
			categoryLists[index] = categoryList{v.ID, v.Name, gs}
			wg.Done()
		}(key, v)
	}
	wg.Wait()
	ch <- categoryLists
}
