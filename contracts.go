package iggy

import (
	"github.com/google/uuid"
)

type MessageStreamConfiguration struct {
	BaseAddress string   `json:"baseAddress"`
	Protocol    Protocol `json:"protocol"`
}

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

type MessageFetchRequest struct {
	ConsumerType    ConsumerType   `json:"consumerType"`
	StreamId        int            `json:"streamId"`
	TopicId         int            `json:"topicId"`
	ConsumerId      int            `json:"consumerId"`
	PartitionId     int            `json:"partitionId"`
	PollingStrategy MessagePolling `json:"pollingStrategy"`
	Value           int            `json:"value"`
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
	KeyKind  Keykind   `json:"keyKind"`
	KeyValue int       `json:"keyValue"`
	Messages []Message `json:"messages"`
}

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

type PartitionContract struct {
	Id            int    `json:"id"`
	MessagesCount uint64 `json:"messagesCount"`
	SegmentsCount int    `json:"segmentsCount"`
	CurrentOffset uint64 `json:"currentOffset"`
	SizeBytes     uint64 `json:"sizeBytes"`
}

type Stats struct {
	ProcessId           int     `json:"process_id"`
	CpuUsage            float32 `json:"cpu_usage"`
	MemoryUsage         uint64  `json:"memory_usage"`
	TotalMemory         uint64  `json:"total_memory"`
	AvailableMemory     uint64  `json:"available_memory"`
	RunTime             uint64  `json:"run_time"`
	StartTime           uint64  `json:"start_time"`
	ReadBytes           uint64  `json:"read_bytes"`
	WrittenBytes        uint64  `json:"written_bytes"`
	MessagesSizeBytes   uint64  `json:"messages_size_bytes"`
	StreamsCount        int     `json:"streams_count"`
	TopicsCount         int     `json:"topics_count"`
	PartitionsCount     int     `json:"partitions_count"`
	SegmentsCount       int     `json:"segments_count"`
	MessagesCount       uint64  `json:"messages_count"`
	ClientsCount        int     `json:"clients_count"`
	ConsumerGroupsCount int     `json:"consumer_groups_count"`
	Hostname            string  `json:"hostname"`
	OsName              string  `json:"os_name"`
	OsVersion           string  `json:"os_version"`
	KernelVersion       string  `json:"kernel_version"`
}

type StreamRequest struct {
	StreamId int    `json:"streamId"`
	Name     string `json:"name"`
}

type StreamResponse struct {
	Id            int             `json:"id"`
	Name          string          `json:"name"`
	SizeBytes     uint64          `json:"sizeBytes"`
	MessagesCount uint64          `json:"messagesCount"`
	TopicsCount   int             `json:"topicsCount"`
	Topics        []TopicResponse `json:"topics"`
}

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

type Keykind int

const (
	PartitionId Keykind = iota
	EntityId
)

type MessagePolling int

const (
	Offset MessagePolling = iota
	Timestamp
	First
	Last
	Next
)

type Protocol string

const (
	Http Protocol = "Http"
	Tcp  Protocol = "Tcp"
	Quic Protocol = "Quic"
)

type ConsumerType int

const (
	Consumer ConsumerType = iota
	ConsumerGroup
)

const (
	KillCode         = 0
	PingCode         = 1
	GetStatsCode     = 2
	SendMessagesCode = 10
	PollMessagesCode = 11
	StoreOffsetCode  = 12
	GetOffsetCode    = 13
	GetStreamCode    = 20
	GetStreamsCode   = 21
	CreateStreamCode = 22
	DeleteStreamCode = 23
	GetTopicCode     = 30
	GetTopicsCode    = 31
	CreateTopicCode  = 32
	DeleteTopicCode  = 33
	GetGroupCode     = 40
	GetGroupsCode    = 41
	CreateGroupCode  = 42
	DeleteGroupCode  = 43
	JoinGroupCode    = 44
	LeaveGroupCode   = 45
)
