package modle

import (
	"easymarket/inits/mysqlinit"
)

// Topic 专题Model
type Topic struct {
	ID          int    `json:"id" db:"id"`                       // id
	Title       string `json:"title" db:"title"`                 // 标题
	Subtitle    string `json:"subtitle" db:"subtitle"`           // 副标题
	PriceInfo   string `json:"price_info" db:"price_info"`       // 价格
	ScenePicURL string `json:"scene_pic_url" db:"scene_pic_url"` // 图片
	Content     string `json:"content" db:"content"`             // 富文本
}

// GetTopic 获取专题列表
func (t *Topic) GetTopic(page, size int) (ts []Topic, err error) {
	sqlStr := "select id, title, subtitle, price_info,scene_pic_url from nideshop_topic limit ?,?"
	err = mysqlinit.DB.Select(&ts, sqlStr, (page-1)*size, size)
	return
}

// GetTopicDetail 获取专题详情
func (t *Topic) GetTopicDetail(id int) (ts Topic, err error) {
	sqlStr := "SELECT id, title, subtitle, price_info,scene_pic_url,content FROM nideshop_topic where id = ?"
	err = mysqlinit.DB.QueryRowx(sqlStr, id).StructScan(&ts)
	return
}

// GetTopicCount 获取专题列表数量
func (t *Topic) GetTopicCount() (count int, err error) {
	sqlStr := "SELECT COUNT(*) FROM nideshop_topic"
	err = mysqlinit.DB.QueryRow(sqlStr).Scan(&count)
	return
}
