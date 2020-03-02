package api

import (
	"easymarket/controller/response"
	"easymarket/modle"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type topicRes struct {
	Count int           `json:"count"`
	Data  []modle.Topic `json:"data"`
}

// GetTopicList 获取专题列表
func GetTopicList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	var tModle modle.Topic
	var t topicRes
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		v, _ := tModle.GetTopic(page, size)
		t.Data = v
		wg.Done()
	}()
	go func() {
		count, _ := tModle.GetTopicCount()
		t.Count = count
	}()
	wg.Wait()
	response.ReSuccess(c, t, "")
}
