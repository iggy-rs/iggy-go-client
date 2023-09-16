package iggy

import . "github.com/iggy-rs/iggy-go-client/contracts"

type IMessageStream interface {
	CreateStream(request StreamRequest) error
	GetStreamById(request GetStreamRequest) (*StreamResponse, error)
	GetStreams() ([]StreamResponse, error)
	DeleteStream(id Identifier) error

	CreateTopic(request CreateTopicRequest) error
	GetTopicById(streamId, topicId Identifier) (*TopicResponse, error)
	GetTopics(streamId Identifier) ([]TopicResponse, error)
	DeleteTopic(streamId, topicId Identifier) error

	//TODO update contracts
	SendMessages(streamId int, topicId int, request MessageSendRequest) error
	PollMessages(request MessageFetchRequest) ([]MessageResponse, error)

	//TODO update contracts
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
