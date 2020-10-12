package service

import (
	"context"
	pb "easymarket-go-server/common/goods/api"
)

// GetGoodsSize 获取商品可选规格
func (s *Service) GetGoodsSize(ctx context.Context, req *pb.GoodsDetailReq) (res *pb.GoodsSizeList, err error) {
	// 所有商品可选属性
	goodsSellSize, daoErr := s.dao.GetGoodsSellSize(req.Id)

	// 商品属性分类
	goodsSizeMap := make(map[int32]int32)
	for _, val := range goodsSellSize {
		if _, ok := goodsSizeMap[val.SpecificationId]; !ok {
			goodsSizeMap[val.SpecificationId] = val.SpecificationId
		}
	}

	// 商品规格
	goodsSizeList := make([]*pb.GoodsSize, 0)

	for _, size := range goodsSizeMap {
		var valueList = make([]*pb.GoodsSpec, 0)
		var id int32
		var name string
		for _, val := range goodsSellSize {
			if val.SpecificationId == size {
				valueList = append(valueList, val)
				id = val.SpecificationId
				name = val.Name
			}
		}
		goodsSizeList = append(goodsSizeList, &pb.GoodsSize{
			Id:        id,
			Name:      name,
			ValueList: valueList,
		})
	}

	res = &pb.GoodsSizeList{
		GoodsSizeList: goodsSizeList,
	}

	err = daoErr

	return
}
