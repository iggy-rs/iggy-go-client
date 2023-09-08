package iggcon

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
	Topics        []TopicResponse `json:"topics,omitempty"`
}
