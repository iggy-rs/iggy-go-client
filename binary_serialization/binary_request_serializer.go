package binaryserialization

import (
	"encoding/binary"

	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func CreateGroup(request CreateConsumerGroupRequest) []byte {
	customIdOffset := 4 + request.StreamId.Length + request.TopicId.Length
	bytes := make([]byte, 4+request.StreamId.Length+request.TopicId.Length+1+4+len(request.Name))
	copy(bytes[0:customIdOffset], SerializeIdentifiers(request.StreamId, request.TopicId))
	binary.LittleEndian.PutUint32(bytes[customIdOffset:customIdOffset+4], uint32(request.ConsumerGroupId))
	bytes[customIdOffset+4] = byte(len(request.Name))
	copy(bytes[customIdOffset+5:], []byte(request.Name))
	return bytes
}

func UpdateOffset(request StoreOffsetRequest) []byte {
	bytes := make([]byte, 6+request.StreamId.Length+request.TopicId.Length+request.Consumer.Id.Length+13)
	bytes[0] = byte(request.Consumer.Kind)
	position := 7 + request.StreamId.Length + request.TopicId.Length + request.Consumer.Id.Length
	copy(bytes[1:position], SerializeIdentifiers(request.Consumer.Id, request.StreamId, request.TopicId))

	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	binary.LittleEndian.PutUint64(bytes[position+4:position+12], uint64(request.Offset))
	return bytes
}

func GetOffset(request GetOffsetRequest) []byte {
	bytes := make([]byte, 6+request.StreamId.Length+request.TopicId.Length+request.Consumer.Id.Length+5)
	bytes[0] = byte(request.Consumer.Kind)
	position := 7 + request.StreamId.Length + request.TopicId.Length + request.Consumer.Id.Length
	copy(bytes[1:position], SerializeIdentifiers(request.Consumer.Id, request.StreamId, request.TopicId))
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	return bytes
}

func CreatePartitions(request CreatePartitionsRequest) []byte {
	bytes := make([]byte, 8+request.StreamId.Length+request.TopicId.Length)
	position := 4 + request.StreamId.Length + request.TopicId.Length
	copy(bytes[0:position], SerializeIdentifiers(request.StreamId, request.TopicId))
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionsCount))

	return bytes
}

func DeletePartitions(request DeletePartitionRequest) []byte {
	bytes := make([]byte, 8+request.StreamId.Length+request.TopicId.Length)
	position := 4 + request.StreamId.Length + request.TopicId.Length
	copy(bytes[0:position], SerializeIdentifiers(request.StreamId, request.TopicId))
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionsCount))

	return bytes
}
