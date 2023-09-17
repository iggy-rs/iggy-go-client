package iggy

import . "github.com/iggy-rs/iggy-go-client/contracts"

type IMessageStream interface {
	GetStreamById(request GetStreamRequest) (*StreamResponse, error)
	GetStreams() ([]StreamResponse, error)
	CreateStream(request CreateStreamRequest) error
	UpdateStream(request UpdateStreamRequest) error
	DeleteStream(id Identifier) error

	CreateTopic(request CreateTopicRequest) error
	GetTopicById(streamId, topicId Identifier) (*TopicResponse, error)
	GetTopics(streamId Identifier) ([]TopicResponse, error)
	DeleteTopic(streamId, topicId Identifier) error

	SendMessages(request SendMessagesRequest) error
	PollMessages(request FetchMessagesRequest) (*FetchMessagesResponse, error)

	StoreOffset(request StoreOffsetRequest) error
	GetOffset(request GetOffsetRequest) (*OffsetResponse, error)

	GetConsumerGroups(streamId Identifier, topicId Identifier) ([]ConsumerGroupResponse, error)
	GetConsumerGroupById(streamId Identifier, topicId Identifier, groupId int) (*ConsumerGroupResponse, error)
	CreateConsumerGroup(request CreateConsumerGroupRequest) error
	DeleteConsumerGroup(request DeleteConsumerGroupRequest) error
	JoinConsumerGroup(request JoinConsumerGroupRequest) error
	LeaveConsumerGroup(request LeaveConsumerGroupRequest) error

	GetStats() (*Stats, error)
}
