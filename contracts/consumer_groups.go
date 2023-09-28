package iggcon

type ConsumerGroupResponse struct {
	Id              int `json:"id"`
	MembersCount    int `json:"membersCount"`
	PartitionsCount int `json:"partitionsCount"`
}

type CreateConsumerGroupRequest struct {
	StreamId        Identifier `json:"streamId"`
	TopicId         Identifier `json:"topicId"`
	ConsumerGroupId int        `json:"consumerGroupId"`
}

type DeleteConsumerGroupRequest struct {
	StreamId        Identifier `json:"streamId"`
	TopicId         Identifier `json:"topicId"`
	ConsumerGroupId int        `json:"consumerGroupId"`
}

type JoinConsumerGroupRequest struct {
	StreamId        Identifier `json:"streamId"`
	TopicId         Identifier `json:"topicId"`
	ConsumerGroupId int        `json:"consumerGroupId"`
}

type LeaveConsumerGroupRequest struct {
	StreamId        Identifier `json:"streamId"`
	TopicId         Identifier `json:"topicId"`
	ConsumerGroupId int        `json:"consumerGroupId"`
}

type ConsumerGroupInfo struct {
	StreamId        int `json:"streamId"`
	TopicId         int `json:"topicId"`
	ConsumerGroupId int `json:"consumerGroupId"`
}
