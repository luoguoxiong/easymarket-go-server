package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetGoodsList 获取商品列表
func (d *Dao) GetGoodsList(isHot, isNew, categoryID, page, size int32) (listRes *pb.GoodsListRes, err error) {
	goodsList := make([]*pb.GoodsRes, 0)
	count := int32(0)

	query := d.db.Table("easymarket_goods")

	if isHot > 0 {
		query = query.Where("is_hot=?", isHot)
	}
	if isNew > 0 {
		query = query.Where("is_new=?", isNew)
	}

	if categoryID > 0 {
		query = query.Where("category_id=?", categoryID)
	}

	if page != 0 && size != 0 {
		err = query.Count(&count).Limit(size).Offset((page - 1) * size).Find(&goodsList).Error
	} else {
		err = query.Count(&count).Find(&goodsList).Error
	}

	listRes = &pb.GoodsListRes{
		GoodsList: goodsList,
		Total:     count,
	}
	return
}

// GetGoodsDetail 获取商品详情
func (d *Dao) GetGoodsDetail(id int32) (goods *pb.GoodsRes, err error) {
	goods = &pb.GoodsRes{}

	err = d.db.Table("easymarket_goods").Where("id=?", id).Find(goods).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return
}
