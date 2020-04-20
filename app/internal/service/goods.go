package service

import (
	"context"
	"sync"

	pb "easymarket-go-server/app/api"
	goods "easymarket-go-server/common/goods/api"
)

//GetGoodsList 获取商品列表
func (s *Service) GetGoodsList(ctx context.Context, req *goods.GoodsReq) (res *goods.GoodsListRes, err error) {
	res, err = s.dao.GetGoodsList(ctx, req)
	return
}

//GetGoodsDetail 商品详情
func (s *Service) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailReq) (res *goods.GoodsRes, err error) {
	res, err = s.dao.GetGoodsDetail(ctx, req)
	return
}

// GetBrandList 获取制造商列表
func (s *Service) GetBrandList(ctx context.Context, req *goods.BrandsListReq) (res *goods.BrandListRes, err error) {
	res, err = s.dao.GetBrandList(ctx, req)
	return
}

// GetBrandDetail 获取制造商详情
func (s *Service) GetBrandDetail(ctx context.Context, req *goods.BrandsDetailReq) (res *goods.BrandsDetailRes, err error) {
	res, err = s.dao.GetBrandDetail(ctx, req)
	return
}

// GetCategoryList 获取子商品分类列表
func (s *Service) GetCategoryList(ctx context.Context, req *goods.CategoryChildReq) (res *goods.CateGoryListRes, err error) {
	res, err = s.dao.GetCategoryList(ctx, req)
	return
}

// GetCategory 获取商品分类详情
func (s *Service) GetCategory(ctx context.Context, req *goods.CategoryReq) (res *goods.CategoryRes, err error) {
	res, err = s.dao.GetCategory(ctx, req)
	return
}

// GetGoodsSell 获取商品分类详情
func (s *Service) GetGoodsSell(ctx context.Context, req *goods.GoodsDetailReq) (res *pb.GoodsSellRes, err error) {
	var wg sync.WaitGroup
	var sizeListChan = make(chan []*goods.GoodsSize, 1)
	var galleryChan = make(chan []*goods.GoodsGallery, 1)
	var attributeChan = make(chan []*goods.Attribute, 1)
	var productChan = make(chan []*goods.Product, 1)
	var issueChan = make(chan []*goods.GoodsIssue, 1)
	wg.Add(5)

	go func() {
		sizeList, _ := s.dao.GetGoodsSize(ctx, req)
		sizeListChan <- sizeList.GoodsSizeList
		wg.Done()
	}()
	go func() {
		gallery, _ := s.dao.GetGoodsGallery(ctx, req)
		galleryChan <- gallery.GoodsGallery
		wg.Done()
	}()

	go func() {
		attribute, _ := s.dao.GetGoodsAttribute(ctx, req)
		attributeChan <- attribute.AttributeList
		wg.Done()
	}()

	go func() {
		product, _ := s.dao.GetGoodsProductList(ctx, req)
		productChan <- product.ProductList
		wg.Done()
	}()

	go func() {
		issue, _ := s.dao.GetGoodsIssueList(ctx, req)
		issueChan <- issue.GoodsIssue
		wg.Done()
	}()
	wg.Wait()
	res = &pb.GoodsSellRes{
		GoodsSize:    <-sizeListChan,
		GoodsGallery: <-galleryChan,
		Attribute:    <-attributeChan,
		ProductList:  <-productChan,
		Issue:        <-issueChan,
	}

	return res, err
}
