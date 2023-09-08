package iggy

import . "github.com/iggy-rs/iggy-go-client/contracts"

type IMessageStream interface {
	CreateStream(request StreamRequest) error
	GetStreamById(request GetStreamRequest) (*StreamResponse, error)
	GetStreams() ([]StreamResponse, error)
	DeleteStream(id int) error

	CreateTopic(streamId int, request TopicRequest) error
	GetTopicById(streamId, topicId int) (*TopicResponse, error)
	GetTopics(streamId int) ([]TopicResponse, error)
	DeleteTopic(streamId int, topicId int) error

	SendMessages(streamId int, topicId int, request MessageSendRequest) error
	PollMessages(request MessageFetchRequest) ([]MessageResponse, error)

	StoreOffset(streamId int, topicId int, offset OffsetContract) error
	GetOffset(request OffsetRequest) (*OffsetResponse, error)

	GetConsumerGroups(streamId int, topicId int) ([]ConsumerGroupResponse, error)
	GetConsumerGroupById(streamId int, topicId int, groupId int) (*ConsumerGroupResponse, error)
	CreateConsumerGroup(streamId int, topicId int, request CreateConsumerGroupRequest) error
	DeleteConsumerGroup(streamId int, topicId int, groupId int) error
	JoinConsumerGroup(request JoinConsumerGroupRequest) error
	LeaveConsumerGroup(request LeaveConsumerGroupRequest) error

	GetStats() (*Stats, error)
}
