package iggcon

type ConsumerGroupResponse struct {
	Id              int `json:"id"`
	MembersCount    int `json:"membersCount"`
	PartitionsCount int `json:"partitionsCount"`
}

type CreateConsumerGroupRequest struct {
	ConsumerGroupId int `json:"consumerGroupId"`
}

type JoinConsumerGroupRequest struct {
	StreamId        int `json:"streamId"`
	TopicId         int `json:"topicId"`
	ConsumerGroupId int `json:"consumerGroupId"`
}

type LeaveConsumerGroupRequest struct {
	StreamId        int `json:"streamId"`
	TopicId         int `json:"topicId"`
	ConsumerGroupId int `json:"consumerGroupId"`
}
