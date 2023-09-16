package tcpserialization

import (
	"encoding/binary"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type TcpCreateStreamRequest struct {
	iggcon.CreateStreamRequest
}

const (
	streamIDOffset   = 0
	nameLengthOffset = 4
	payloadOffset    = 5
)

func (request *TcpCreateStreamRequest) Serialize() []byte {
	nameLength := len(request.Name)
	bytes := make([]byte, nameLength+payloadOffset)
	binary.LittleEndian.PutUint32(bytes[streamIDOffset:streamIDOffset+4], uint32(request.StreamId))
	bytes[nameLengthOffset] = byte(nameLength)
	copy(bytes[payloadOffset:], []byte(request.Name))
	return bytes
}
