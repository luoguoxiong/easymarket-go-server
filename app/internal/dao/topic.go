package dao

import (
	"context"
	topic "easymarket-go-server/common/topic/api"
)

// GetTopicList 获取专题列表列表
func (d *Dao) GetTopicList(ctx context.Context, req *topic.TopicListReq) (res *topic.TopicListRes, err error) {
	res, err = d.topicClient.GetTopicList(ctx, req)
	return
}
