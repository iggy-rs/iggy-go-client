package binaryserialization

import (
	"encoding/binary"

	. "github.com/iggy-rs/iggy-go-client/contracts"
)

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
	copy(bytes[0:customIdOffset], SerializeIdentifiers(streamId, topicId))
	binary.LittleEndian.PutUint32(bytes[customIdOffset:customIdOffset+4], uint32(groupId))
	return bytes
}

func UpdateOffset(request StoreOffsetRequest) []byte {
	bytes := make([]byte, 4+request.StreamId.Length+request.TopicId.Length+17)
	bytes[0] = byte(request.Consumer.Kind)
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.Consumer.Id))
	copy(bytes[5:7+request.StreamId.Length], SerializeIdentifier(request.StreamId))
	copy(bytes[7+request.StreamId.Length:9+request.StreamId.Length+request.TopicId.Length], SerializeIdentifier(request.TopicId))
	position := 9 + request.StreamId.Length + request.TopicId.Length
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	binary.LittleEndian.PutUint64(bytes[position+4:position+12], uint64(request.Offset))
	return bytes
}

func GetOffset(request GetOffsetRequest) []byte {
	bytes := make([]byte, 4+request.StreamId.Length+request.TopicId.Length+9)
	bytes[0] = byte(request.Consumer.Kind)
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.Consumer.Id))
	copy(bytes[5:7+request.StreamId.Length], SerializeIdentifier(request.StreamId))
	copy(bytes[7+request.StreamId.Length:9+request.StreamId.Length+request.TopicId.Length], SerializeIdentifier(request.TopicId))
	position := 9 + request.StreamId.Length + request.TopicId.Length
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	return bytes
}
