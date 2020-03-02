package modle

import (
	"easymarket/inits/mysqlinit"
)

// Topic 专题Model
type Topic struct {
	ID        int    `json:"id" db:"id"`                 // id
	Title     string `json:"title" db:"title"`           // 标题
	Subtitle  string `json:"subtitle" db:"subtitle"`     // 副标题
	PriceInfo string `json:"price_info" db:"price_info"` // 价格
}

// GetTopic 获取专题列表
func (t *Topic) GetTopic(page, size int) (ts []Topic, err error) {
	sqlStr := "select id, title, subtitle, price_info from nideshop_topic limit ?,?"
	err = mysqlinit.DB.Select(&ts, sqlStr, (page-1)*size, size)
	return
}
