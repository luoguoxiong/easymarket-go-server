package modle

import "easymarket/inits/mysqlinit"

// Category categoryModel
type Category struct {
	ID   int    `json:"id" db:"id"`     // id
	Name string `json:"name" db:"name"` // 名称
}

// GetCategory 获取parent_id=0分类
func (c *Category) GetCategory(page, size int) (ca []Category, err error) {
	sqlStr := "select id, `name` from nideshop_category WHERE parent_id = 0 limit ?,?"
	err = mysqlinit.DB.Select(&ca, sqlStr, (page-1)*size, size)
	return
}

// GetAllCategory 获取某父级ID的所有分类
func (c *Category) GetAllCategory(parentID int) (ca []Category, err error) {
	sqlStr := "select id, `name` from nideshop_category WHERE parent_id = ?"
	err = mysqlinit.DB.Select(&ca, sqlStr, parentID)
	return
}
