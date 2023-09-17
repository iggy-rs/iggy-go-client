package tcp

import (
	"encoding/binary"
	tcpserialization "github.com/iggy-rs/iggy-go-client/tcp/serialization"
	"net"

	. "github.com/iggy-rs/iggy-go-client/contracts"
	ierror "github.com/iggy-rs/iggy-go-client/errors"
)

type TcpMessageStream struct {
	client *net.TCPConn
}

const (
	InitialBytesLength   = 4
	ExpectedResponseSize = 8
)

func NewTcpMessageStream(url string) (*TcpMessageStream, error) {
	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &TcpMessageStream{client: conn}, nil
}

func (tms *TcpMessageStream) GetStats() (*Stats, error) {
	buffer, err := tms.sendAndFetchResponse([]byte{}, GetStatsCode)
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

	stats := &tcpserialization.TcpStats{}
	err = stats.Deserialize(responseBuffer)

	return &stats.Stats, err
}

func (tms *TcpMessageStream) GetStreams() ([]StreamResponse, error) {
	buffer, err := tms.sendAndFetchResponse([]byte{}, GetStreamsCode)
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

	return MapStreams(responseBuffer), nil
}

func (tms *TcpMessageStream) GetStreamById(request GetStreamRequest) (*StreamResponse, error) {
	message := tcpserialization.SerializeIdentifier(request.StreamID)
	buffer, err := tms.sendAndFetchResponse(message, GetStreamCode)
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

	return MapStream(responseBuffer), nil
}

func (tms *TcpMessageStream) DeleteStream(id Identifier) error {
	message := tcpserialization.SerializeIdentifier(id)
	_, err := tms.sendAndFetchResponse(message, DeleteStreamCode)
	return err
}

func (tms *TcpMessageStream) GetTopicById(streamId Identifier, topicId Identifier) (*TopicResponse, error) {
	message := GetTopicByIdMessage(streamId, topicId)
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

	return MapTopic(responseBuffer)
}

func (tms *TcpMessageStream) GetTopics(streamId Identifier) ([]TopicResponse, error) {
	message := tcpserialization.SerializeIdentifier(streamId)
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

	return MapTopics(responseBuffer)
}

func (tms *TcpMessageStream) CreateTopic(request CreateTopicRequest) error {
	message := CreateTopic(request)
	_, err := tms.sendAndFetchResponse(message, CreateTopicCode)
	return err
}

func (tms *TcpMessageStream) UpdateTopic(request UpdateTopicRequest) error {
	serializedRequest := tcpserialization.TcpUpdateTopicRequest{UpdateTopicRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), CreateTopicCode)
	return err
}

func (tms *TcpMessageStream) DeleteTopic(streamId, topicId Identifier) error {
	message := DeleteTopic(streamId, topicId)
	_, err := tms.sendAndFetchResponse(message, DeleteTopicCode)
	return err
}

func (tms *TcpMessageStream) CreateStream(request CreateStreamRequest) error {
	serializedRequest := tcpserialization.TcpCreateStreamRequest{CreateStreamRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), CreateStreamCode)
	return err
}

func (tms *TcpMessageStream) UpdateStream(request UpdateStreamRequest) error {
	serializedRequest := tcpserialization.TcpUpdateStreamRequest{UpdateStreamRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), UpdateStreamCode)
	return err
}

func (tms *TcpMessageStream) SendMessages(request SendMessagesRequest) error {
	message := CreateMessage(request)
	_, err := tms.sendAndFetchResponse(message, SendMessagesCode)
	return err
}

func (tms *TcpMessageStream) PollMessages(request FetchMessagesRequest) (*FetchMessagesResponse, error) {
	serializedRequest := tcpserialization.TcpFetchMessagesRequest{FetchMessagesRequest: request}
	buffer, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), PollMessagesCode)
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

	return MapMessages(responseBuffer)
}

func (tms *TcpMessageStream) CreateConsumerGroup(request CreateConsumerGroupRequest) error {
	message := CreateGroup(request)
	_, err := tms.sendAndFetchResponse(message, CreateGroupCode)
	return err
}

func (tms *TcpMessageStream) DeleteConsumerGroup(request DeleteConsumerGroupRequest) error {
	message := DeleteGroup(request)
	_, err := tms.sendAndFetchResponse(message, DeleteGroupCode)
	return err
}

func (tms *TcpMessageStream) GetConsumerGroupById(streamId Identifier, topicId Identifier, groupId int) (*ConsumerGroupResponse, error) {
	message := GetGroup(streamId, topicId, groupId)
	buffer, err := tms.sendAndFetchResponse(message, GetGroupCode)
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

	return MapConsumerGroup(responseBuffer)
}

func (tms *TcpMessageStream) GetConsumerGroups(streamId Identifier, topicId Identifier) ([]ConsumerGroupResponse, error) {
	message := GetGroups(streamId, topicId)
	buffer, err := tms.sendAndFetchResponse(message, GetGroupsCode)
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

	return MapConsumerGroups(responseBuffer), err
}

func (tms *TcpMessageStream) GetOffset(request GetOffsetRequest) (*OffsetResponse, error) {
	message := GetOffset(request)
	buffer, err := tms.sendAndFetchResponse(message, GetOffsetCode)
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

	return MapOffsets(responseBuffer), nil
}

func (tms *TcpMessageStream) JoinConsumerGroup(request JoinConsumerGroupRequest) error {
	message := JoinGroup(request)
	_, err := tms.sendAndFetchResponse(message, JoinGroupCode)
	return err
}

func (tms *TcpMessageStream) LeaveConsumerGroup(request LeaveConsumerGroupRequest) error {
	message := LeaveGroup(request)
	_, err := tms.sendAndFetchResponse(message, LeaveGroupCode)
	return err
}

func (tms *TcpMessageStream) StoreOffset(request StoreOffsetRequest) error {
	message := UpdateOffset(request)
	_, err := tms.sendAndFetchResponse(message, StoreOffsetCode)
	return err
}

func (tms *TcpMessageStream) sendAndFetchResponse(message []byte, command CommandCode) ([]byte, error) {
	payload := createPayload(message, command)

	if _, err := tms.client.Write(payload); err != nil {
		return nil, err
	}

	buffer := make([]byte, ExpectedResponseSize)
	if _, err := tms.client.Read(buffer); err != nil {
		return nil, err
	}

	if responseCode := getResponseCode(buffer); responseCode != 0 {
		return nil, ierror.MapFromCode(responseCode)
	}

	return buffer, nil
}

func createPayload(message []byte, command CommandCode) []byte {
	messageLength := len(message) + 4
	messageBytes := make([]byte, InitialBytesLength+messageLength)
	binary.LittleEndian.PutUint32(messageBytes[:4], uint32(messageLength))
	binary.LittleEndian.PutUint32(messageBytes[4:8], uint32(command))
	copy(messageBytes[8:], message)
	return messageBytes
}

func getResponseCode(buffer []byte) int {
	return int(binary.LittleEndian.Uint32(buffer[:4]))
}

func getResponseLength(buffer []byte) (int, error) {
	length := int(binary.LittleEndian.Uint32(buffer[4:]))
	if length <= 1 {
		return 0, &ierror.IggyError{
			Code:    0,
			Message: "Received empty response.",
		}
	}
	return length, nil
}
