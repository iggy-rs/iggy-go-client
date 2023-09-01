package tcp

import (
	"encoding/binary"
	"net"

	. "github.com/iggy-rs/iggy-go-client/contracts"
	iggerr "github.com/iggy-rs/iggy-go-client/errors"
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
	var message []byte
	buffer, err := tms.SendAndFetchResponse(message, GetStatsCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)

	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	_, err = tms.client.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return MapStats(responseBuffer), nil
}

func (tms *TcpMessageStream) GetStreams() ([]StreamResponse, error) {
	var message []byte
	buffer, err := tms.SendAndFetchResponse(message, GetStreamsCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)

	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	_, err = tms.client.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return MapStreams(responseBuffer), nil
}

func (tms *TcpMessageStream) GetStreamById(id int) (*StreamResponse, error) {
	message := make([]byte, 4)
	binary.LittleEndian.PutUint32(message, uint32(id))
	buffer, err := tms.SendAndFetchResponse(message, GetStreamCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)

	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	_, err = tms.client.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return MapStream(responseBuffer), nil
}

func (tms *TcpMessageStream) DeleteStream(id int) error {
	message := make([]byte, 4)
	binary.LittleEndian.PutUint32(message, uint32(id))
	_, err := tms.SendAndFetchResponse(message, DeleteStreamCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) GetTopicById(streamId int, topicId int) (*TopicResponse, error) {
	message := GetTopicByIdMessage(streamId, topicId)
	buffer, err := tms.SendAndFetchResponse(message, GetTopicCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)
	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	if _, err := tms.client.Read(responseBuffer); err != nil {
		return nil, err
	}

	return MapTopic(responseBuffer)
}

func (tms *TcpMessageStream) GetTopics(streamId int) ([]TopicResponse, error) {
	message := make([]byte, 4)
	binary.LittleEndian.PutUint32(message, uint32(streamId))
	buffer, err := tms.SendAndFetchResponse(message, GetTopicsCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)

	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	_, err = tms.client.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return MapTopics(responseBuffer)
}

func (tms *TcpMessageStream) CreateTopic(streamId int, request TopicRequest) error {
	message := CreateTopic(streamId, request)
	_, err := tms.SendAndFetchResponse(message, CreateTopicCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) DeleteTopic(streamId int, topicId int) error {
	message := DeleteTopic(streamId, topicId)
	_, err := tms.SendAndFetchResponse(message, DeleteTopicCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) CreateStream(request StreamRequest) error {
	message := CreateStream(request)
	_, err := tms.SendAndFetchResponse(message, CreateStreamCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) SendMessages(streamId int, topicId int, request MessageSendRequest) error {
	message := CreateMessage(streamId, topicId, request)
	_, err := tms.SendAndFetchResponse(message, SendMessagesCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) PollMessages(request MessageFetchRequest) ([]MessageResponse, error) {
	message := GetMessages(request)
	buffer, err := tms.SendAndFetchResponse(message, PollMessagesCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)
	if responseLength <= 1 {
		return []MessageResponse{}, nil
	}

	responseBuffer := make([]byte, responseLength)
	if _, err := tms.client.Read(responseBuffer); err != nil {
		return nil, err
	}

	return MapMessages(responseBuffer)
}

func (tms *TcpMessageStream) CreateConsumerGroup(streamId int, topicId int, request CreateConsumerGroupRequest) error {
	message := CreateGroup(streamId, topicId, request)
	_, err := tms.SendAndFetchResponse(message, CreateGroupCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) DeleteConsumerGroup(streamId int, topicId int, groupId int) error {
	message := DeleteGroup(streamId, topicId, groupId)
	_, err := tms.SendAndFetchResponse(message, DeleteGroupCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) GetConsumerGroupById(streamId int, topicId int, groupId int) (*ConsumerGroupResponse, error) {
	message := GetGroup(streamId, topicId, groupId)
	buffer, err := tms.SendAndFetchResponse(message, GetGroupCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)
	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	if _, err := tms.client.Read(responseBuffer); err != nil {
		return nil, err
	}

	return MapConsumerGroup(responseBuffer)
}

func (tms *TcpMessageStream) GetConsumerGroups(streamId int, topicId int) ([]ConsumerGroupResponse, error) {
	message := GetGroups(streamId, topicId)
	buffer, err := tms.SendAndFetchResponse(message, GetGroupsCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)
	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	if _, err := tms.client.Read(responseBuffer); err != nil {
		return nil, err
	}

	return MapConsumerGroups(responseBuffer), err
}

func (tms *TcpMessageStream) GetOffset(request OffsetRequest) (*OffsetResponse, error) {
	message := GetOffset(request)
	buffer, err := tms.SendAndFetchResponse(message, GetOffsetCode)
	if err != nil {
		return nil, err
	}

	responseLength := GetResponseLength(buffer)
	if responseLength <= 1 {
		return nil, nil
	}

	responseBuffer := make([]byte, responseLength)
	if _, err := tms.client.Read(responseBuffer); err != nil {
		return nil, err
	}

	return MapOffsets(responseBuffer), nil
}

func (tms *TcpMessageStream) JoinConsumerGroup(request JoinConsumerGroupRequest) error {
	message := JoinGroup(request)
	_, err := tms.SendAndFetchResponse(message, JoinGroupCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) LeaveConsumerGroup(request LeaveConsumerGroupRequest) error {
	message := LeaveGroup(request)
	_, err := tms.SendAndFetchResponse(message, LeaveGroupCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) StoreOffset(streamId int, topicId int, offset OffsetContract) error {
	message := UpdateOffset(streamId, topicId, offset)
	_, err := tms.SendAndFetchResponse(message, StoreOffsetCode)
	if err != nil {
		return err
	}

	return nil
}

func (tms *TcpMessageStream) SendAndFetchResponse(message []byte, command int) ([]byte, error) {
	payload := CreatePayload(message, command)

	if _, err := tms.client.Write(payload); err != nil {
		return nil, err
	}

	buffer := make([]byte, ExpectedResponseSize)
	if _, err := tms.client.Read(buffer); err != nil {
		return nil, err
	}

	if responseCode := GetResponseCode(buffer); responseCode != 0 {
		return nil, iggerr.MapFromCode(responseCode)
	}

	return buffer, nil
}

func CreatePayload(message []byte, command int) []byte {
	messageLength := len(message) + 4
	messageBytes := make([]byte, InitialBytesLength+messageLength)
	binary.LittleEndian.PutUint32(messageBytes[:4], uint32(messageLength))
	binary.LittleEndian.PutUint32(messageBytes[4:8], uint32(command))
	copy(messageBytes[8:], message)
	return messageBytes
}

func GetResponseCode(buffer []byte) int {
	return int(binary.LittleEndian.Uint32(buffer[:4]))
}

func GetResponseLength(buffer []byte) int {
	return int(binary.LittleEndian.Uint32(buffer[4:]))
}
