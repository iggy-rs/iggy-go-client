package tcp

import (
	"github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
	ierror "github.com/iggy-rs/iggy-go-client/errors"
)

func (tms *IggyTcpClient) GetTopics(streamId Identifier) ([]TopicResponse, error) {
	message := binaryserialization.SerializeIdentifier(streamId)
	buffer, err := tms.sendAndFetchResponse(message, GetTopicsCode)
	if err != nil {
		return nil, err
	}

	responseLength, err := getResponseLength(buffer)
	if err != nil {
		return nil, err
	}

	responseBuffer := make([]byte, responseLength)
	_, err = tms.client.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return binaryserialization.DeserializeTopics(responseBuffer)
}

func (tms *IggyTcpClient) GetTopicById(streamId Identifier, topicId Identifier) (*TopicResponse, error) {
	message := binaryserialization.SerializeIdentifiers(streamId, topicId)
	buffer, err := tms.sendAndFetchResponse(message, GetTopicCode)
	if err != nil {
		return nil, err
	}

	responseLength, err := getResponseLength(buffer)
	if err != nil {
		return nil, err
	}

	responseBuffer := make([]byte, responseLength)
	if _, err := tms.client.Read(responseBuffer); err != nil {
		return nil, err
	}

	return binaryserialization.DeserializeTopic(responseBuffer)
}

func (tms *IggyTcpClient) CreateTopic(request CreateTopicRequest) error {
	if MaxStringLength < len(request.Name) {
		return ierror.TextTooLong("topic_name")
	}
	serializedRequest := binaryserialization.TcpCreateTopicRequest{CreateTopicRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), CreateTopicCode)
	return err
}

func (tms *IggyTcpClient) UpdateTopic(request UpdateTopicRequest) error {
	if MaxStringLength < len(request.Name) {
		return ierror.TextTooLong("topic_name")
	}
	serializedRequest := binaryserialization.TcpUpdateTopicRequest{UpdateTopicRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), UpdateTopicCode)
	return err
}

func (tms *IggyTcpClient) DeleteTopic(streamId, topicId Identifier) error {
	message := binaryserialization.SerializeIdentifiers(streamId, topicId)
	_, err := tms.sendAndFetchResponse(message, DeleteTopicCode)
	return err
}
