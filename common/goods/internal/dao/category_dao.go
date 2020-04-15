package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetCategoryList 获取子商品分类列表
func (d *Dao) GetCategoryList(parentID int32) (listRes *pb.CateGoryListRes, err error) {
	var categoryList []*pb.CategoryRes

	err = d.db.Table("easymarket_category").Where("parent_id=?", parentID).Find(&categoryList).Error

	listRes = &pb.CateGoryListRes{
		CategoryList: categoryList,
	}

	return
}

// GetCategory 获取商品分类详情
func (d *Dao) GetCategory(id int32) (res *pb.CategoryRes, err error) {
	res = &pb.CategoryRes{}
	err = d.db.Table("easymarket_category").Where("id=?", id).Find(&res).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}
