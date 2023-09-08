package iggcon

import (
	"github.com/google/uuid"
)

type MessageFetchRequest struct {
	ConsumerType    ConsumerType   `json:"consumerType"`
	StreamId        int            `json:"streamId"`
	TopicId         int            `json:"topicId"`
	ConsumerId      int            `json:"consumerId"`
	PartitionId     int            `json:"partitionId"`
	PollingStrategy MessagePolling `json:"pollingStrategy"`
	Value           uint64         `json:"value"`
	Count           int            `json:"count"`
	AutoCommit      bool           `json:"autoCommit"`
}

type MessageResponse struct {
	Offset    uint64    `json:"offset"`
	Timestamp uint64    `json:"timestamp"`
	Id        uuid.UUID `json:"id"`
	Payload   []byte    `json:"payload"`
}

type Message struct {
	Id      uuid.UUID
	Payload []byte
}

type MessageSendRequest struct {
	Key      Key       `json:"key"`
	Messages []Message `json:"messages"`
}

type MessagePolling int

const (
	Offset MessagePolling = iota
	Timestamp
	First
	Last
	Next
)
