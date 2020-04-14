package service

import (
	"context"
	pb "easymarket-go-server/common/topic/api"
)

// GetTopic 获取专题详情
func (s *Service) GetTopic(ctx context.Context, req *pb.TopicReq) (res *pb.TopicRes, err error) {
	res, err = s.dao.GetTopic(req.Id)
	return
}

// GetTopicList 获取专题列表
func (s *Service) GetTopicList(ctx context.Context, req *pb.TopicListReq) (res *pb.TopicListRes, err error) {
	res, err = s.dao.GetTopicList(req.Page, req.Size_)
	return
}

// GetTopicRelated 获取相似专题列表
func (s *Service) GetTopicRelated(ctx context.Context, req *pb.TopicRelatedReq) (res *pb.TopicListRes, err error) {
	res, err = s.dao.GetTopicRelated(req.Id, req.Page, req.Size_)
	return
}
