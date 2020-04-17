package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetAttruibute 获取商品特点
func (d *Dao) GetAttruibute(id int32) (res *pb.AttributeList, err error) {
	var attribute []*pb.Attribute

	err = d.db.Table("easymarket_attribute").Where("goods_id=?", id).Find(&attribute).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	res = &pb.AttributeList{
		AttributeList: attribute,
	}

	return
}
