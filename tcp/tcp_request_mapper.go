package tcp

import (
	"encoding/binary"

	. "github.com/eldpcn/iggy-go/contracts"
)

func GetMessages(request MessageFetchRequest) []byte {
	bytes := make([]byte, 31)
	bytes[0] = 0

	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.ConsumerId))
	binary.LittleEndian.PutUint32(bytes[5:9], uint32(request.StreamId))
	binary.LittleEndian.PutUint32(bytes[9:13], uint32(request.TopicId))
	binary.LittleEndian.PutUint32(bytes[13:17], uint32(request.PartitionId))

	switch request.PollingStrategy {
	case Offset:
		bytes[17] = 0
	case Timestamp:
		bytes[17] = 1
	case First:
		bytes[17] = 2
	case Last:
		bytes[17] = 3
	case Next:
		bytes[17] = 4
	}

	binary.LittleEndian.PutUint64(bytes[18:26], uint64(request.Value))
	binary.LittleEndian.PutUint32(bytes[26:30], uint32(request.Count))

	if request.AutoCommit {
		bytes[30] = 1
	} else {
		bytes[30] = 0
	}

	return bytes
}

func CreateMessage(streamId, topicId int, request MessageSendRequest) []byte {
	messageBytesCount := 0
	for _, message := range request.Messages {
		messageBytesCount += 16 + 1 + 4 + len(message.Payload)
	}

	bytes := make([]byte, 17+messageBytesCount)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))

	switch request.Key.KeyKind {
	case PartitionId:
		bytes[8] = 0
	case EntityId:
		bytes[8] = 1
	}

	bytes[9] = 4 // default message length
	binary.LittleEndian.PutUint32(bytes[10:14], uint32(request.Key.Value))
	binary.LittleEndian.PutUint32(bytes[14:18], uint32(len(request.Messages)))

	position := 18
	for _, message := range request.Messages {
		copy(bytes[position:position+16], message.Id[:])
		binary.LittleEndian.PutUint32(bytes[position+16:position+20], uint32(len(message.Payload)))
		copy(bytes[position+20:position+20+len(message.Payload)], message.Payload)
		position += 20 + len(message.Payload)
	}

	return bytes
}

func CreateStream(request StreamRequest) []byte {
	bytes := make([]byte, 4+len(request.Name))
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(request.StreamId))
	copy(bytes[4:], []byte(request.Name))
	return bytes
}

func CreateGroup(streamId, topicId int, request CreateConsumerGroupRequest) []byte {
	bytes := make([]byte, 12)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))
	binary.LittleEndian.PutUint32(bytes[8:12], uint32(request.ConsumerGroupId))
	return bytes
}

func JoinGroup(request JoinConsumerGroupRequest) []byte {
	bytes := make([]byte, 12)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(request.StreamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(request.TopicId))
	binary.LittleEndian.PutUint32(bytes[8:12], uint32(request.ConsumerGroupId))
	return bytes
}

func LeaveGroup(request LeaveConsumerGroupRequest) []byte {
	bytes := make([]byte, 12)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(request.StreamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(request.TopicId))
	binary.LittleEndian.PutUint32(bytes[8:12], uint32(request.ConsumerGroupId))
	return bytes
}

func DeleteGroup(streamId, topicId, groupId int) []byte {
	bytes := make([]byte, 12)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))
	binary.LittleEndian.PutUint32(bytes[8:12], uint32(groupId))
	return bytes
}

func GetGroups(streamId, topicId int) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))
	return bytes
}

func GetGroup(streamId, topicId, groupId int) []byte {
	bytes := make([]byte, 12)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))
	binary.LittleEndian.PutUint32(bytes[8:12], uint32(groupId))
	return bytes
}

func CreateTopic(streamId int, request TopicRequest) []byte {
	bytes := make([]byte, 12+len(request.Name))
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(request.TopicId))
	binary.LittleEndian.PutUint32(bytes[8:12], uint32(request.PartitionsCount))
	copy(bytes[12:], []byte(request.Name))
	return bytes
}

func GetTopicByIdMessage(streamId, topicId int) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))
	return bytes
}

func DeleteTopic(streamId, topicId int) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))
	return bytes
}

func UpdateOffset(streamId, topicId int, contract OffsetContract) []byte {
	bytes := make([]byte, 17)
	bytes[0] = 0
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[5:9], uint32(topicId))
	binary.LittleEndian.PutUint32(bytes[9:13], uint32(contract.ConsumerId))
	binary.LittleEndian.PutUint32(bytes[13:17], uint32(contract.PartitionId))
	binary.LittleEndian.PutUint64(bytes[17:25], uint64(contract.Offset))
	return bytes
}

func GetOffset(request OffsetRequest) []byte {
	bytes := make([]byte, 17)
	bytes[0] = 0
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.StreamId))
	binary.LittleEndian.PutUint32(bytes[5:9], uint32(request.TopicId))
	binary.LittleEndian.PutUint32(bytes[9:13], uint32(request.ConsumerId))
	binary.LittleEndian.PutUint32(bytes[13:17], uint32(request.PartitionId))
	return bytes
}
