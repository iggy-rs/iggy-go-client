package tcp

import (
	"encoding/binary"

	. "github.com/iggy-rs/iggy-go-client/contracts"
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
	nameLength := len(request.Name)
	bytes := make([]byte, nameLength+5)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(request.StreamId))
	bytes[4] = byte(nameLength)
	copy(bytes[5:], []byte(request.Name))
	return bytes
}

func CreateGroup(request CreateConsumerGroupRequest) []byte {
	return baseGroupMapping(request.StreamId, request.TopicId, request.ConsumerGroupId)
}

func JoinGroup(request JoinConsumerGroupRequest) []byte {
	return baseGroupMapping(request.StreamId, request.TopicId, request.ConsumerGroupId)
}

func LeaveGroup(request LeaveConsumerGroupRequest) []byte {
	return baseGroupMapping(request.StreamId, request.TopicId, request.ConsumerGroupId)
}

func DeleteGroup(request DeleteConsumerGroupRequest) []byte {
	return baseGroupMapping(request.StreamId, request.TopicId, request.ConsumerGroupId)
}

func GetGroup(streamId Identifier, topicId Identifier, groupId int) []byte {
	return baseGroupMapping(streamId, topicId, groupId)
}

// this is extracted for later refactoring
func baseGroupMapping(streamId Identifier, topicId Identifier, groupId int) []byte {
	customIdOffset := 4 + streamId.Length + topicId.Length
	bytes := make([]byte, customIdOffset+4)
	copy(bytes[0:2+streamId.Length], GetBytesFromIdentifier(streamId))
	copy(bytes[2+streamId.Length:4+streamId.Length+topicId.Length], GetBytesFromIdentifier(topicId))
	binary.LittleEndian.PutUint32(bytes[customIdOffset:customIdOffset+4], uint32(groupId))
	return bytes
}

func GetGroups(streamId, topicId Identifier) []byte {
	bytes := make([]byte, 4+streamId.Length+topicId.Length)
	copy(bytes[0:2+streamId.Length], GetBytesFromIdentifier(streamId))
	copy(bytes[2+streamId.Length:], GetBytesFromIdentifier(topicId))
	return bytes
}

func CreateTopic(request CreateTopicRequest) []byte {
	totalIdSize := 2 + request.StreamId.Length
	totalNameSize := len(request.Name)

	bytes := make([]byte, 15+totalIdSize+totalNameSize)

	copy(bytes[0:totalIdSize], GetBytesFromIdentifier(request.StreamId))

	position := totalIdSize
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.TopicId))

	position += 4
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionsCount))

	position += 4
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.MessageExpiry))

	position += 4
	bytes[position] = byte(totalNameSize)

	position++
	copy(bytes[position:], []byte(request.Name))

	return bytes
}

func GetTopicByIdMessage(streamId, topicId Identifier) []byte {
	bytes := make([]byte, 4+streamId.Length+topicId.Length)
	copy(bytes[0:2+streamId.Length], GetBytesFromIdentifier(streamId))
	copy(bytes[2+topicId.Length:], GetBytesFromIdentifier(topicId))
	return bytes
}

func DeleteTopic(streamId, topicId Identifier) []byte {
	bytes := make([]byte, 4+streamId.Length+topicId.Length)
	copy(bytes[0:2+streamId.Length], GetBytesFromIdentifier(streamId))
	copy(bytes[2+topicId.Length:], GetBytesFromIdentifier(topicId))
	return bytes
}

func UpdateOffset(request StoreOffsetRequest) []byte {
	bytes := make([]byte, 4+request.StreamId.Length+request.TopicId.Length+17)
	bytes[0] = byte(request.Consumer.Kind)
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.Consumer.ID))
	copy(bytes[5:7+request.StreamId.Length], GetBytesFromIdentifier(request.StreamId))
	copy(bytes[7+request.StreamId.Length:9+request.StreamId.Length+request.TopicId.Length], GetBytesFromIdentifier(request.TopicId))
	position := 9 + request.StreamId.Length + request.TopicId.Length
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	binary.LittleEndian.PutUint64(bytes[position+4:position+12], uint64(request.Offset))
	return bytes
}

func GetOffset(request GetOffsetRequest) []byte {
	bytes := make([]byte, 4+request.StreamId.Length+request.TopicId.Length+9)
	bytes[0] = byte(request.Consumer.Kind)
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.Consumer.ID))
	copy(bytes[5:7+request.StreamId.Length], GetBytesFromIdentifier(request.StreamId))
	copy(bytes[7+request.StreamId.Length:9+request.StreamId.Length+request.TopicId.Length], GetBytesFromIdentifier(request.TopicId))
	position := 9 + request.StreamId.Length + request.TopicId.Length
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	return bytes
}

func GetBytesFromIdentifier(identifier Identifier) []byte {
	bytes := make([]byte, int(identifier.Length)+2)
	bytes[0] = byte(identifier.Kind)
	bytes[1] = byte(identifier.Length)

	if identifier.Kind == StringId {
		valAsString := identifier.Value.(string)
		for i := 0; i < int(identifier.Length); i++ {
			bytes[i+2] = valAsString[i]
		}
	} else if identifier.Kind == NumericId {
		valAsInt := identifier.Value.(int)
		binary.LittleEndian.PutUint32(bytes[2:6], uint32(valAsInt))
	}
	return bytes
}
