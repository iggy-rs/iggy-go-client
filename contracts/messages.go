package iggcon

import (
	"errors"

	"github.com/google/uuid"
)

type MessageFetchRequest struct {
	ConsumerKind    ConsumerKind   `json:"consumerKind"`
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
	Headers map[HeaderKey]HeaderValue
}

type MessageSendRequest struct {
	StreamId     Identifier   `json:"streamId"`
	TopicId      Identifier   `json:"topicId"`
	Partitioning Partitioning `json:"partitioning"`
	Messages     []Message    `json:"messages"`
}

type MessagePolling int

const (
	Offset MessagePolling = iota
	Timestamp
	First
	Last
	Next
)

type HeaderValue struct {
	Kind  HeaderKind
	Value []byte
}

type HeaderKey struct {
	Value string
}

func NewHeaderKey(val string) (HeaderKey, error) {
	if len(val) == 0 || len(val) > 255 {
		return HeaderKey{}, errors.New("Value has incorrect size, must be between 1 and 255")
	}
	return HeaderKey{Value: val}, nil
}

type Guid struct {
	Value string
}

type HeaderKind int

const (
	Raw     HeaderKind = 1
	String  HeaderKind = 2
	Bool    HeaderKind = 3
	Int32   HeaderKind = 6
	Int64   HeaderKind = 7
	Int128  HeaderKind = 8
	Uint32  HeaderKind = 11
	Uint64  HeaderKind = 12
	Uint128 HeaderKind = 13
	Float   HeaderKind = 14
	Double  HeaderKind = 15
)
