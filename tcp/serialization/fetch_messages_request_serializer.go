package tcpserialization

import (
	"encoding/binary"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

const (
	headerSize            = 5
	streamTopicIdHeader   = 2
	partitionStrategySize = 5
	offsetSize            = 12
	commitFlagSize        = 1
)

type TcpFetchMessagesRequest struct {
	iggcon.FetchMessagesRequest
}

func (request *TcpFetchMessagesRequest) Serialize() []byte {
	streamTopicIdLength := streamTopicIdHeader + request.StreamId.Length + streamTopicIdHeader + request.TopicId.Length
	messageSize := headerSize + offsetSize + streamTopicIdLength + partitionStrategySize + offsetSize + commitFlagSize
	bytes := make([]byte, messageSize)

	bytes[0] = byte(request.Consumer.Kind)
	binary.LittleEndian.PutUint32(bytes[1:5], uint32(request.Consumer.Id))

	position := headerSize

	copy(bytes[position:position+streamTopicIdLength], append(append([]byte{}, SerializeIdentifier(request.StreamId)...), SerializeIdentifier(request.TopicId)...))

	position += streamTopicIdLength
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.PartitionId))
	bytes[position+4] = byte(request.PollingStrategy.Kind)

	position += partitionStrategySize
	binary.LittleEndian.PutUint64(bytes[position:position+8], uint64(request.PollingStrategy.Value))
	binary.LittleEndian.PutUint32(bytes[position+8:position+12], uint32(request.Count))

	position += offsetSize

	if request.AutoCommit {
		bytes[position] = 1
	} else {
		bytes[position] = 0
	}

	return bytes
}
