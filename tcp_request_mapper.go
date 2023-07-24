package iggy

import (
	"encoding/binary"
	"unsafe"
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
		messageBytesCount += 16 + 4 + len(message.Payload)
	}

	bytes := make([]byte, 17+messageBytesCount)
	binary.LittleEndian.PutUint32(bytes[0:4], uint32(streamId))
	binary.LittleEndian.PutUint32(bytes[4:8], uint32(topicId))

	switch request.KeyKind {
	case PartitionId:
		bytes[8] = 0
	case EntityId:
		bytes[8] = 1
	}

	binary.LittleEndian.PutUint32(bytes[9:13], uint32(request.KeyValue))
	binary.LittleEndian.PutUint32(bytes[13:17], uint32(len(request.Messages)))

	position := 17
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

func MapStats(payload []byte) *Stats {
	var stats *Stats

	stats.ProcessId = int(binary.LittleEndian.Uint32(payload[0:4]))
	stats.CpuUsage = *(*float32)(unsafe.Pointer(&payload[4]))
	stats.MemoryUsage = binary.LittleEndian.Uint64(payload[8:16])
	stats.TotalMemory = binary.LittleEndian.Uint64(payload[16:24])
	stats.AvailableMemory = binary.LittleEndian.Uint64(payload[24:32])
	stats.RunTime = binary.LittleEndian.Uint64(payload[32:40])
	stats.StartTime = binary.LittleEndian.Uint64(payload[40:48])
	stats.ReadBytes = binary.LittleEndian.Uint64(payload[48:56])
	stats.WrittenBytes = binary.LittleEndian.Uint64(payload[56:64])
	stats.MessagesSizeBytes = binary.LittleEndian.Uint64(payload[64:72])
	stats.StreamsCount = int(binary.LittleEndian.Uint32(payload[72:76]))
	stats.TopicsCount = int(binary.LittleEndian.Uint32(payload[76:80]))
	stats.PartitionsCount = int(binary.LittleEndian.Uint32(payload[80:84]))
	stats.SegmentsCount = int(binary.LittleEndian.Uint32(payload[84:88]))
	stats.MessagesCount = binary.LittleEndian.Uint64(payload[88:96])
	stats.ClientsCount = int(binary.LittleEndian.Uint32(payload[96:100]))
	stats.ConsumerGroupsCount = int(binary.LittleEndian.Uint32(payload[100:104]))

	position := 104
	hostnameLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.Hostname = string(payload[position+4 : position+4+hostnameLength])
	position += 4 + hostnameLength

	osNameLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.OsName = string(payload[position+4 : position+4+osNameLength])
	position += 4 + osNameLength

	osVersionLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.OsVersion = string(payload[position+4 : position+4+osVersionLength])
	position += 4 + osVersionLength

	kernelVersionLength := int(binary.LittleEndian.Uint32(payload[position : position+4]))
	stats.KernelVersion = string(payload[position+4 : position+4+kernelVersionLength])

	return stats
}
