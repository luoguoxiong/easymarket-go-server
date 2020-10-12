package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetGoodsSellSize 获取商品售卖规格
func (d *Dao) GetGoodsSellSize(id int32) (goods []*pb.GoodsSpec, err error) {
	goods = make([]*pb.GoodsSpec, 0)

	err = d.db.Raw("SELECT s.id,s.goods_id,s.pic_url,s.`value`,s.specification_id,z.`name` FROM easymarket_goods_spec s LEFT JOIN easymarket_size z  ON z.id=s.specification_id WHERE (s.goods_id=?)", id).Scan(&goods).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return
}
