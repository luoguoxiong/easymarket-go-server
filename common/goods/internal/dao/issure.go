package dao

import (
	pb "easymarket-go-server/common/goods/api"

	"github.com/jinzhu/gorm"
)

// GetIssure 获取商品常见疑问
func (d *Dao) GetIssure(id int32) (res *pb.GoodsIssueList, err error) {
	var goodsIssue []*pb.GoodsIssue

	err = d.db.Table("easymarket_goods_issue").Where("goods_id=?", id).Find(&goodsIssue).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	res = &pb.GoodsIssueList{
		GoodsIssue: goodsIssue,
	}

	return
}
