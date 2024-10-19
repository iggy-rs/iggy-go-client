package iggcon

type CreateTopicRequest struct {
	StreamId             Identifier `json:"streamId"`
	TopicId              int        `json:"topicId"`
	PartitionsCount      int        `json:"partitionsCount"`
	CompressionAlgorithm uint8      `json:"compressionAlgorithm"`
	Name                 string     `json:"name"`
	MessageExpiry        int        `json:"messageExpiry"`
	MaxTopicSize         uint64     `json:"maxTopicSize"`
	ReplicationFactor    uint8      `json:"replicationFactor"`
}

type UpdateTopicRequest struct {
	TopicId       Identifier `json:"topicId"`
	StreamId      Identifier `json:"streamId"`
	Name          string     `json:"name"`
	MessageExpiry int        `json:"messageExpiry"`
}

type TopicResponse struct {
	Id              int                 `json:"id"`
	CreatedAt       int                 `json:"createdAt"`
	Name            string              `json:"name"`
	SizeBytes       uint64              `json:"sizeBytes"`
	MessageExpiry   int                 `json:"messageExpiry"`
	MessagesCount   uint64              `json:"messagesCount"`
	PartitionsCount int                 `json:"partitionsCount"`
	Partitions      []PartitionContract `json:"partitions,omitempty"`
}
