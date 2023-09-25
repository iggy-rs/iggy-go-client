package binaryserialization

import (
	"encoding/binary"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type TcpCreateTopicRequest struct {
	iggcon.CreateTopicRequest
}

func (request *TcpCreateTopicRequest) Serialize() []byte {
	totalIdSize := 2 + request.StreamId.Length
	totalNameSize := len(request.Name)

	bytes := make([]byte, 15+totalIdSize+totalNameSize)

	copy(bytes[0:totalIdSize], SerializeIdentifier(request.StreamId))

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
