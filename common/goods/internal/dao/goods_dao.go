package dao

import (
	pb "easymarketgoserve/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetGoodsList 获取商品列表
func (d *Dao) GetGoodsList(isHot, page, size int32) (listRes *pb.GoodsListRes, err error) {
	var goodsList []*pb.GoodsRes
	var count int32

	err = d.db.Table("nideshop_goods").Where("is_hot=?", isHot).Count(&count).Limit(size).Offset((page - 1) * size).Find(&goodsList).Error

	listRes = &pb.GoodsListRes{
		GoodsList: goodsList,
		Total:     count,
	}

	return
}

// GetGoodsDetail 获取商品详情
func (d *Dao) GetGoodsDetail(id int32) (goods *pb.GoodsRes, err error) {
	goods = &pb.GoodsRes{}

	err = d.db.Table("nideshop_goods").Where("id=?", id).Find(goods).Error
	if err == gorm.ErrRecordNotFound {
		return goods, nil
	}

	return
}
