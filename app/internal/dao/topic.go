package dao

import (
	"context"
	topic "easymarket-go-server/common/topic/api"
)

// GetTopicList 获取专题列表
func (d *Dao) GetTopicList(ctx context.Context, req *topic.TopicListReq) (res *topic.TopicListRes, err error) {
	res, err = d.topicClient.GetTopicList(ctx, req)
	return
}

// GetTopic 获取专题详情
func (d *Dao) GetTopic(ctx context.Context, req *topic.TopicReq) (res *topic.TopicRes, err error) {
	res, err = d.topicClient.GetTopic(ctx, req)
	return
}

// GetTopicRelated 获取相似专题列表
func (d *Dao) GetTopicRelated(ctx context.Context, req *topic.TopicRelatedReq) (res *topic.TopicListRes, err error) {
	res, err = d.topicClient.GetTopicRelated(ctx, req)
	return
}
