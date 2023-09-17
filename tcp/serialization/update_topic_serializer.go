package tcpserialization

import (
	"encoding/binary"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type TcpUpdateTopicRequest struct {
	iggcon.UpdateTopicRequest
}

func (request *TcpUpdateTopicRequest) Serialize() []byte {
	bytes := make([]byte, len(request.Name)+request.StreamId.Length+request.TopicId.Length+9)
	copy(bytes[0:2+request.StreamId.Length], SerializeIdentifier(request.StreamId))
	copy(bytes[2+request.StreamId.Length:4+request.StreamId.Length+request.TopicId.Length], SerializeIdentifier(request.TopicId))
	position := 4 + request.StreamId.Length + request.TopicId.Length
	binary.LittleEndian.PutUint32(bytes[position:position+4], uint32(request.MessageExpiry))
	bytes[position+4] = byte(len(request.Name))
	copy(bytes[position+5:], []byte(request.Name))
	return bytes
}
