package api

import (
	"easymarket/controller/response"
	"easymarket/modle"

	"github.com/gin-gonic/gin"
)

// GetHome example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   offset     query    int     true        "Offset"
// @Success 200 {object} modle.Goods "Can not find ID"
// @Router /api/home [get]
func GetHome(c *gin.Context) {
	var c1 = make(chan modle.Goods, 1)
	var c2 = make(chan []modle.Goods, 1)
	res := make(map[string]interface{}, 2)
	var g modle.Goods
	go func() {
		goods, _ := g.NewGoods(1, 5)
		c2 <- goods
	}()
	go func() {
		good, _ := g.Search()
		c1 <- good
	}()
	res["good"] = <-c1
	res["goods"] = <-c2
	response.ReSuccess(c, "", res)
}

// PostHome example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param account body Good true "Add account"
// @Success  200 {object} Good "Can not find ID"
// @Router /api/home [post]
