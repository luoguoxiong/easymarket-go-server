package modle

import (
	"easymarket/inits/mysqlinit"
)

// Goods 商品model
type Goods struct {
	ID         int    `json:"id" db:"id"`                  // id
	Name       string `json:"name" db:"name"`              // 名称
	GoodsBrief string `json:"goodsBrief" db:"goods_brief"` // 描述
}

// Search 查找商品
func (g *Goods) Search() (good Goods, err error) {
	sqlStr := "select id, name from nideshop_goods where id=?"
	err = mysqlinit.DB.Get(&good, sqlStr, 1006002)
	return
}

// NewGoods 新品首发
func (g *Goods) NewGoods(page int, size int) (goods []Goods, err error) {
	sqlStr := "select id, name, goods_brief from nideshop_goods where is_new=1  limit ?,?"
	err = mysqlinit.DB.Select(&goods, sqlStr, (page-1)*size, size)
	return
}
