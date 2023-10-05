package binaryserialization

import (
	"encoding/binary"

	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func CreateGroup(request CreateConsumerGroupRequest) []byte {
	customIdOffset := 4 + request.StreamId.Length + request.TopicId.Length + 1 + len(request.Name)
	bytes := make([]byte, customIdOffset+4)
	copy(bytes[0:customIdOffset], SerializeIdentifiers(request.StreamId, request.TopicId))
	binary.LittleEndian.PutUint32(bytes[customIdOffset:customIdOffset+4], uint32(request.ConsumerGroupId))
	bytes[customIdOffset+4] = byte(len(request.Name))
	copy(bytes[customIdOffset+5:], []byte(request.Name))
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
