package modle

import "easymarket/inits/mysqlinit"

// Bannel 轮播图model
type Bannel struct {
	ID       int    `json:"id" db:"id"`               // id
	ImageURL string `json:"image_url" db:"image_url"` // 图片地址
}

// GetBannel 获取轮播图
func (b *Bannel) GetBannel(size int) (banner []Bannel, err error) {
	sqlStr := "SELECT id, image_url FROM nideshop_ad LIMIT ?"
	err = mysqlinit.DB.Select(&banner, sqlStr, size)
	return
}
