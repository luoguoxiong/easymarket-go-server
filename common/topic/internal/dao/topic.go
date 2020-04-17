package dao

import (
	pb "easymarket-go-server/common/topic/api"

	"github.com/jinzhu/gorm"
)

// GetTopicList 获取专题列表
func (d *Dao) GetTopicList(page, size int32) (listRes *pb.TopicListRes, err error) {
	var topicList []*pb.TopicRes
	var count int32

	err = d.db.Table("easymarket_topic").Count(&count).Limit(size).Offset((page - 1) * size).Find(&topicList).Error

	listRes = &pb.TopicListRes{
		TopicList: topicList,
		Total:     count,
	}

	return
}

// GetTopic 获取专题详情
func (d *Dao) GetTopic(id int32) (topic *pb.TopicRes, err error) {
	topic = &pb.TopicRes{}

	err = d.db.Table("easymarket_topic").Where("id=?", id).Find(topic).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return
}

// GetTopicRelated 获取专题相似列表
func (d *Dao) GetTopicRelated(id, page, size int32) (listRes *pb.TopicListRes, err error) {
	var topicList []*pb.TopicRes
	var count int32

	err = d.db.Table("easymarket_topic").Where("id!=?", id).Count(&count).Limit(size).Offset((page - 1) * size).Find(&topicList).Error

	listRes = &pb.TopicListRes{
		TopicList: topicList,
		Total:     count,
	}

	return
}
