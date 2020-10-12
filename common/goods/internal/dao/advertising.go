package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)
// GetAdsSwiper 获取广告轮播图
func (d *Dao) GetAdsSwiper() (res *pb.AdsSwiper, err error) {
	var banner []*pb.SwiperItem

	err = d.db.Table("easymarket_ad").Find(&banner).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	res = &pb.AdsSwiper{
		Banner: banner,
	}

	return
}