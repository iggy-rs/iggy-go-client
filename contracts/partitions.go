package iggcon

type PartitionContract struct {
	Id            int    `json:"id"`
	MessagesCount uint64 `json:"messagesCount"`
	CreatedAt     uint64 `json:"createdAt"`
	SegmentsCount int    `json:"segmentsCount"`
	CurrentOffset uint64 `json:"currentOffset"`
	SizeBytes     uint64 `json:"sizeBytes"`
}
