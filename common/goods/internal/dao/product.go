package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetProduct 获取商品产品
func (d *Dao) GetProduct(id int32) (res *pb.GoodsProductList, err error) {
	var product []*pb.Product

	err = d.db.Table("easymarket_product").Where("goods_id=?", id).Find(&product).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	res = &pb.GoodsProductList{
		ProductList: product,
	}

	return
}
