package iggcon

type TopicRequest struct {
	TopicId         int    `json:"topicId"`
	Name            string `json:"name"`
	PartitionsCount int    `json:"partitionsCount"`
}

type TopicResponse struct {
	Id              int                 `json:"id"`
	Name            string              `json:"name"`
	SizeBytes       uint64              `json:"sizeBytes"`
	MessagesCount   uint64              `json:"messagesCount"`
	PartitionsCount int                 `json:"partitionsCount"`
	Partitions      []PartitionContract `json:"partitions,omitempty"`
}
