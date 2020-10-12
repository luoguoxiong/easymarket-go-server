package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetAttruibute 获取商品特点
func (d *Dao) GetAttruibute(id int32) (res *pb.AttributeList, err error) {
	var attribute = make([]*pb.Attribute, 0)
	err = d.db.Table("easymarket_attribute_goods").Where("goods_id=?", id).Find(&attribute).Error

	err = d.db.Raw("SELECT g.value, a.name FROM `easymarket_attribute_goods` g LEFT JOIN easymarket_attribute a ON g.attribute_id = a.id WHERE g.goods_id=?", id).Scan(&attribute).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	res = &pb.AttributeList{
		AttributeList: attribute,
	}

	return
}
