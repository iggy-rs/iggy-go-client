package tcp

import (
	"encoding/binary"

	. "github.com/iggy-rs/iggy-go-client/contracts"
	tcpserialization "github.com/iggy-rs/iggy-go-client/tcp/serialization"
)

func CreateMessage(request SendMessagesRequest) []byte {
	streamTopicIdLength := 2 + request.StreamId.Length + 2 + request.TopicId.Length
	messageBytesCount := calculateMessageBytesCount(request.Messages)
	totalSize := streamTopicIdLength + messageBytesCount + request.Partitioning.Length + 2
	bytes := make([]byte, totalSize)
	position := 0
	//ids
	copy(bytes[position:2+request.StreamId.Length], tcpserialization.SerializeIdentifier(request.StreamId))
	copy(bytes[position+2+request.StreamId.Length:streamTopicIdLength], tcpserialization.SerializeIdentifier(request.TopicId))
	position = streamTopicIdLength

	//partitioning
	bytes[streamTopicIdLength] = byte(request.Partitioning.Kind)
	bytes[streamTopicIdLength+1] = byte(request.Partitioning.Length)
	copy(bytes[streamTopicIdLength+2:streamTopicIdLength+2+request.Partitioning.Length], []byte(request.Partitioning.Value))
	position = streamTopicIdLength + 2 + request.Partitioning.Length

	emptyHeaders := make([]byte, 4)

	for _, message := range request.Messages {

		copy(bytes[position:position+16], message.Id[:])
		if message.Headers != nil {
			headersBytes := GetHeadersBytes(message.Headers)
			binary.LittleEndian.PutUint32(bytes[position+16:position+20], uint32(len(headersBytes)))
			copy(bytes[position+20:position+20+len(headersBytes)], headersBytes)
			position += len(headersBytes) + 20
		} else {
			copy(bytes[position+16:position+16+len(emptyHeaders)], emptyHeaders)
			position += 20
		}

		binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(len(message.Payload)))
		copy(bytes[position+4:position+4+len(message.Payload)], message.Payload)
		position += len(message.Payload) + 4
	}

	return bytes
}

func calculateMessageBytesCount(messages []Message) int {
	count := 0
	for _, msg := range messages {
		count += 16 + 4 + len(msg.Payload) + 4
		for key, header := range msg.Headers {
			count += 4 + len(key.Value) + 1 + 4 + len(header.Value)
		}
	}
	return count
}

func GetHeadersBytes(headers map[HeaderKey]HeaderValue) []byte {
	headersLength := 0
	for key, header := range headers {
		headersLength += 4 + len(key.Value) + 1 + 4 + len(header.Value)
	}
	headersBytes := make([]byte, headersLength)
	position := 0
	for key, value := range headers {
		headerBytes := GetBytesFromHeader(key, value)
		copy(headersBytes[position:position+len(headerBytes)], headerBytes)
		position += len(headerBytes)
	}
	return headersBytes
}

func GetBytesFromHeader(key HeaderKey, value HeaderValue) []byte {
	headerBytesLength := 4 + len(key.Value) + 1 + 4 + len(value.Value)
	headerBytes := make([]byte, headerBytesLength)

	binary.LittleEndian.PutUint32(headerBytes[:4], uint32(len(key.Value)))
	copy(headerBytes[4:4+len(key.Value)], key.Value)

	headerBytes[4+len(key.Value)] = byte(value.Kind)

	binary.LittleEndian.PutUint32(headerBytes[4+len(key.Value)+1:4+len(key.Value)+1+4], uint32(len(value.Value)))
	copy(headerBytes[4+len(key.Value)+1+4:], value.Value)

	return headerBytes
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
	copy(bytes[0:2+streamId.Length], tcpserialization.SerializeIdentifier(streamId))
	copy(bytes[2+streamId.Length:4+streamId.Length+topicId.Length], tcpserialization.SerializeIdentifier(topicId))
	binary.LittleEndian.PutUint32(bytes[customIdOffset:customIdOffset+4], uint32(groupId))
	return bytes
}

func GetGroups(streamId, topicId Identifier) []byte {
	bytes := make([]byte, 4+streamId.Length+topicId.Length)
	copy(bytes[0:2+streamId.Length], tcpserialization.SerializeIdentifier(streamId))
	copy(bytes[2+streamId.Length:], tcpserialization.SerializeIdentifier(topicId))
	return bytes
}

func CreateTopic(request CreateTopicRequest) []byte {
	totalIdSize := 2 + request.StreamId.Length
	totalNameSize := len(request.Name)

	bytes := make([]byte, 15+totalIdSize+totalNameSize)

	copy(bytes[0:totalIdSize], tcpserialization.SerializeIdentifier(request.StreamId))

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
	copy(bytes[0:2+streamId.Length], tcpserialization.SerializeIdentifier(streamId))
	copy(bytes[2+topicId.Length:], tcpserialization.SerializeIdentifier(topicId))
	return bytes
}

func DeleteTopic(streamId, topicId Identifier) []byte {
	bytes := make([]byte, 4+streamId.Length+topicId.Length)
	copy(bytes[0:2+streamId.Length], tcpserialization.SerializeIdentifier(streamId))
	copy(bytes[2+topicId.Length:], tcpserialization.SerializeIdentifier(topicId))
	return bytes
}

func UpdateOffset(request StoreOffsetRequest) []byte {
	bytes := make([]byte, 4+request.StreamId.Length+request.TopicId.Length+17)
	bytes[0] = byte(request.Consumer.Kind)
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.Consumer.Id))
	copy(bytes[5:7+request.StreamId.Length], tcpserialization.SerializeIdentifier(request.StreamId))
	copy(bytes[7+request.StreamId.Length:9+request.StreamId.Length+request.TopicId.Length], tcpserialization.SerializeIdentifier(request.TopicId))
	position := 9 + request.StreamId.Length + request.TopicId.Length
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	binary.LittleEndian.PutUint64(bytes[position+4:position+12], uint64(request.Offset))
	return bytes
}

func GetOffset(request GetOffsetRequest) []byte {
	bytes := make([]byte, 4+request.StreamId.Length+request.TopicId.Length+9)
	bytes[0] = byte(request.Consumer.Kind)
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.Consumer.Id))
	copy(bytes[5:7+request.StreamId.Length], tcpserialization.SerializeIdentifier(request.StreamId))
	copy(bytes[7+request.StreamId.Length:9+request.StreamId.Length+request.TopicId.Length], tcpserialization.SerializeIdentifier(request.TopicId))
	position := 9 + request.StreamId.Length + request.TopicId.Length
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	return bytes
}
