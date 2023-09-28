package binaryserialization

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type TcpLogInRequest struct {
	iggcon.LogInRequest
}

func (request *TcpLogInRequest) Serialize() []byte {
	serialized := make([]byte, 2+len(request.Password)+len(request.Username))

	serialized[0] = byte(len(request.Username))
	copy(serialized[1:1+len(request.Username)], []byte(request.Username))
	serialized[1+len(request.Username)] = byte(len(request.Password))
	copy(serialized[2+len(request.Username):], []byte(request.Username))

	return serialized
}
