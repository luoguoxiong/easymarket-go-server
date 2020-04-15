package service

import (
	"context"
	topic "easymarket-go-server/common/topic/api"
)

//GetTopic 获取专题详情
func (s *Service) GetTopic(ctx context.Context, req *topic.TopicReq) (res *topic.TopicRes, err error) {
	return
}

//GetTopicList 获取专题列表
func (s *Service) GetTopicList(ctx context.Context, req *topic.TopicListReq) (res *topic.TopicListRes, err error) {
	return
}

// GetTopicRelated 获取相似专题列表
func (s *Service) GetTopicRelated(ctx context.Context, req *topic.TopicRelatedReq) (res *topic.TopicListRes, err error) {
	return
}
