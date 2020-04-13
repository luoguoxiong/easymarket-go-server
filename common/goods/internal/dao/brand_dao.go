package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetBrandList 获取制造商列表
func (d *Dao) GetBrandList(page, size int32) (listRes *pb.BrandListRes, err error) {
	var brandList []*pb.BrandsDetailRes
	var count int32

	err = d.db.Table("nideshop_brand").Count(&count).Limit(size).Offset((page - 1) * size).Find(&brandList).Error

	listRes = &pb.BrandListRes{
		BrandList: brandList,
		Total:     count,
	}

	return
}

// GetBrandDetail 获取制造商详情
func (d *Dao) GetBrandDetail(id int32) (brand *pb.BrandsDetailRes, err error) {
	brand = &pb.BrandsDetailRes{}

	err = d.db.Table("nideshop_brand").Where("id=?", id).Find(brand).Error
	if err == gorm.ErrRecordNotFound {
		return brand, nil
	}

	return
}
