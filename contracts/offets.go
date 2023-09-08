package iggcon

type OffsetContract struct {
	ConsumerId  int    `json:"consumerId"`
	PartitionId int    `json:"partitionId"`
	Offset      uint64 `json:"offset"`
}

type OffsetRequest struct {
	StreamId    int `json:"streamId"`
	TopicId     int `json:"topicId"`
	ConsumerId  int `json:"consumerId"`
	PartitionId int `json:"partitionId"`
}

type OffsetResponse struct {
	ConsumerId int `json:"consumerId"`
	Offset     int `json:"offset"`
}
