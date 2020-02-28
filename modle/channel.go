package modle

import "easymarket/inits/mysqlinit"

// Channel 渠道
type Channel struct {
	ID      int    `json:"id" db:"id"`             // 渠道ID
	Name    string `json:"name" db:"name"`         // 渠道名称
	IconURL string `json:"icon_url" db:"icon_url"` // 渠道图标
}

// GetChannel 获取渠道
func (c *Channel) GetChannel(size int) (cs []Channel, err error) {
	sqlStr := "SELECT id, name, icon_url FROM nideshop_channel LIMIT ?"
	err = mysqlinit.DB.Select(&cs, sqlStr, size)
	return
}
