package modle

import "easymarket/inits/mysqlinit"

// Brand 制造商model
type Brand struct {
	ID        int    `json:"id" db:"id"`                   // id
	Name      string `json:"name" db:"name"`               // 名称
	NewPicURL string `json:"new_pic_url" db:"new_pic_url"` // 图片地址
}

// GetBrand 获取新制造商
func (b *Brand) GetBrand(size int) (bs []Brand, err error) {
	sqlStr := "SELECT id, name, new_pic_url FROM `nideshop_brand` WHERE is_new =1  ORDER BY sort_order ASC LIMIT ?"
	err = mysqlinit.DB.Select(&bs, sqlStr, size)
	return
}
