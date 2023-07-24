package iggy

type IMessageStream interface {
	CreateStream(request StreamRequest) error
	GetStreamById(id int) (*StreamResponse, error)
	GetStreams() ([]StreamResponse, error)
	DeleteStream(id int) error

	CreateTopic(streamId int, request TopicRequest) error
	GetTopicById(streamId, topicId int) (*TopicResponse, error)
	GetTopics(streamId int) ([]TopicResponse, error)
	DeleteTopic(streamId int, topicId int) error

	SendMessages(streamId int, topicId int, request MessageSendRequest) error
	PollMessages(request MessageFetchRequest) ([]MessageResponse, error)

	GetStats() (*Stats, error)
}
