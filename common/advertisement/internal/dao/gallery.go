package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetGallery 获取商品轮播图
func (d *Dao) GetGallery(id int32) (res *pb.GoodsGalleryList, err error) {
	var gallery []*pb.GoodsGallery

	err = d.db.Table("easymarket_gallery").Where("goods_id=?", id).Find(&gallery).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	res = &pb.GoodsGalleryList{
		GoodsGallery: gallery,
	}

	return
}
