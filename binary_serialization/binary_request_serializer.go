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

//USERS

func SerializeCreateUserRequest(request CreateUserRequest) []byte {
	capacity := 4 + len(request.Username) + len(request.Password)
	if request.Permissions != nil {
		capacity += 1 + 4 + CalculatePermissionsSize(request.Permissions)
	}

	bytes := make([]byte, capacity)
	position := 0

	bytes[position] = byte(len(request.Username))
	position += 1
	copy(bytes[position:position+len(request.Username)], []byte(request.Username))
	position += len(request.Username)

	bytes[position] = byte(len(request.Password))
	position += 1
	copy(bytes[position:position+len(request.Password)], []byte(request.Password))
	position += len(request.Password)

	statusByte := byte(0)
	if request.Status == Active {
		statusByte = byte(1)
	} else if request.Status == Inactive {
		statusByte = byte(2)
	}
	bytes[position] = statusByte
	position += 1

	if request.Permissions != nil {
		bytes[position] = byte(1)
		position += 1
		permissionsBytes := GetBytesFromPermissions(request.Permissions)
		binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(len(permissionsBytes)))
		position += 4
		copy(bytes[position:position+len(permissionsBytes)], permissionsBytes)
		position += len(permissionsBytes)
	} else {
		bytes[position] = byte(0)
		position += 1
	}

	return bytes
}

func GetBytesFromPermissions(data *Permissions) []byte {
	size := CalculatePermissionsSize(data)
	bytes := make([]byte, size)

	bytes[0] = boolToByte(data.Global.ManageServers)
	bytes[1] = boolToByte(data.Global.ReadServers)
	bytes[2] = boolToByte(data.Global.ManageUsers)
	bytes[3] = boolToByte(data.Global.ReadUsers)
	bytes[4] = boolToByte(data.Global.ManageStreams)
	bytes[5] = boolToByte(data.Global.ReadStreams)
	bytes[6] = boolToByte(data.Global.ManageTopics)
	bytes[7] = boolToByte(data.Global.ReadTopics)
	bytes[8] = boolToByte(data.Global.PollMessages)
	bytes[9] = boolToByte(data.Global.SendMessages)

	position := 10

	if data.Streams != nil {
		bytes[position] = byte(1)
		position += 1

		for streamID, stream := range data.Streams {
			binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(streamID))
			position += 4

			bytes[position] = boolToByte(stream.ManageStream)
			bytes[position+1] = boolToByte(stream.ReadStream)
			bytes[position+2] = boolToByte(stream.ManageTopics)
			bytes[position+3] = boolToByte(stream.ReadTopics)
			bytes[position+4] = boolToByte(stream.PollMessages)
			bytes[position+5] = boolToByte(stream.SendMessages)
			position += 6

			if stream.Topics != nil {
				bytes[position] = byte(1)
				position += 1

				for topicID, topic := range stream.Topics {
					binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(topicID))
					position += 4

					bytes[position] = boolToByte(topic.ManageTopic)
					bytes[position+1] = boolToByte(topic.ReadTopic)
					bytes[position+2] = boolToByte(topic.PollMessages)
					bytes[position+3] = boolToByte(topic.SendMessages)
					position += 4

					bytes[position] = byte(0)
					position += 1
				}
			} else {
				bytes[position] = byte(0)
				position += 1
			}
		}
	} else {
		bytes[0] = byte(0)
	}

	return bytes
}

func CalculatePermissionsSize(data *Permissions) int {
	size := 10

	if data.Streams != nil {
		size += 1

		for _, stream := range data.Streams {
			size += 4
			size += 6
			size += 1

			if stream.Topics != nil {
				size += 1
				size += len(stream.Topics) * 9
			} else {
				size += 1
			}
		}
	} else {
		size += 1
	}

	return size
}

func boolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func SerializeUpdateUser(request UpdateUserRequest) []byte {
	length := request.UserID.Length + 2

	if len(request.Username) != 0 {
		length += 2 + len(request.Username)
	}

	if request.Status != nil {
		length += 2
	}

	bytes := make([]byte, length+1)
	position := 0

	copy(bytes[position:position+request.UserID.Length+2], SerializeIdentifier(request.UserID))
	position += position + request.UserID.Length + 2

	if len(request.Username) != 0 {
		bytes[position] = 1
		position++
		bytes[position] = byte(len(request.Username))
		position++
		copy(bytes[position:position+len(request.Username)], []byte(request.Username))
		position += len(request.Username)
	} else {
		bytes[position] = 0
		position++
	}

	if request.Status != nil {
		bytes[position] = 1
		position++
		statusByte := byte(0)
		if *request.Status == Active {
			statusByte = 1
		} else if *request.Status == Inactive {
			statusByte = 2
		}
		bytes[position] = statusByte
	} else {
		bytes[position] = 0
	}

	return bytes
}

func SerializeChangePasswordRequest(request ChangePasswordRequest) []byte {
	length := request.UserID.Length + 2 + len(request.CurrentPassword) + len(request.NewPassword) + 2
	bytes := make([]byte, length)
	position := 0

	copy(bytes[position:position+request.UserID.Length+2], SerializeIdentifier(request.UserID))
	position += request.UserID.Length + 2

	bytes[position] = byte(len(request.CurrentPassword))
	position++
	copy(bytes[position:position+len(request.CurrentPassword)], []byte(request.CurrentPassword))
	position += len(request.CurrentPassword)

	bytes[position] = byte(len(request.NewPassword))
	position++
	copy(bytes[position:position+len(request.NewPassword)], []byte(request.NewPassword))
	position += len(request.NewPassword)

	return bytes
}

func SerializeUpdateUserPermissionsRequest(request UpdateUserPermissionsRequest) []byte {
	length := request.UserID.Length + 2

	if request.Permissions != nil {
		length += 1 + 4 + CalculatePermissionsSize(request.Permissions)
	}

	bytes := make([]byte, length)
	position := 0

	copy(bytes[position:position+request.UserID.Length+2], SerializeIdentifier(request.UserID))
	position += request.UserID.Length + 2

	if request.Permissions != nil {
		bytes[position] = 1
		position++
		permissionsBytes := GetBytesFromPermissions(request.Permissions)
		binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(len(permissionsBytes)))
		position += 4
		copy(bytes[position:position+len(permissionsBytes)], permissionsBytes)
		position += len(permissionsBytes)
	} else {
		bytes[position] = 0
		position++
	}

	return bytes
}

func SerializeInt(value int) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(value))
	return bytes
}
