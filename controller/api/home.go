package api

import (
	"easymarket/controller/response"

	"github.com/gin-gonic/gin"
)

// Student test
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Goods []Good `json:"goods"`
}

// Good 商品信息
type Good struct {
	Name  string `json:"name"` // 年龄
	Price int    `json:"price"`
}

// GetHome example
// @Summary Add a new pet to the store
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param   some_id query   int     true  "Some ID"
// @Failure 200 {object} Good "Can not find ID"
// @Router /api/home [get]
func GetHome(c *gin.Context) {
	goods := []Good{Good{Name: "商品名称", Price: 1}}
	stu := Student{Name: "w2d", Age: 22, Goods: goods}
	response.ReSuccess(c, "", stu)
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
func PostHome(c *gin.Context) {
	goods := []Good{Good{Name: "商品名称", Price: 2}}
	stu := Student{Name: "lgx", Age: 22, Goods: goods}
	response.ReSuccess(c, "22", stu)
}
