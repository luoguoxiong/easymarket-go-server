package modle

import (
	"easymarket/inits/mysqlinit"
)

// Goods 商品model
type Goods struct {
	ID          int     `json:"id" db:"id"`                     // id
	Name        string  `json:"name" db:"name"`                 // 名称
	ListPicURL  string  `json:"list_pic_url" db:"list_pic_url"` // 图片
	RetailPrice float64 `json:"retail_price" db:"retail_price"` // 价格
	GoodsBrief  string  `json:"goods_brief" db:"goods_brief"`   // 商品描述
}

// Search 查找商品
func (g *Goods) Search() (good Goods, err error) {
	sqlStr := "select id, name from nideshop_goods where id=?"
	err = mysqlinit.DB.Get(&good, sqlStr, 1006002)
	return
}

// NewGoods 新品首发
func (g *Goods) NewGoods(page int, size int) (goods []Goods, err error) {
	sqlStr := "select id, name, list_pic_url, retail_price from nideshop_goods where is_new=1  limit ?,?"
	err = mysqlinit.DB.Select(&goods, sqlStr, (page-1)*size, size)
	return
}

// HotGoods 热门商品
func (g *Goods) HotGoods(page int, size int) (goods []Goods, err error) {
	sqlStr := "select id, name, list_pic_url, retail_price,goods_brief from nideshop_goods where is_hot=1  limit ?,?"
	err = mysqlinit.DB.Select(&goods, sqlStr, (page-1)*size, size)
	return
}
