package dao

import (
	pb "easymarketgoserve/common/goods/api"
)

// GetGoodsList 获取商品列表
func (d *Dao) GetGoodsList(isHot, page, size int32) (goodsList []*pb.GoodsRes, count int32) {
	d.db.Table("nideshop_goods").Where("is_hot=?", isHot).Count(&count).Limit(size).Offset((page - 1) * size).Find(&goodsList)

	return goodsList, count
}
