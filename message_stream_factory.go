package iggy

import (
	"errors"

	iggcon "github.com/eldpcn/iggy-go/contracts"
	tcp "github.com/eldpcn/iggy-go/tcp"
)

type IMessageStreamFactory interface {
	CreateStream(config iggcon.MessageStreamConfiguration) (IMessageStream, error)
}

type MessageStreamFactory struct{}

func (msf *MessageStreamFactory) CreateMessageStream(config iggcon.MessageStreamConfiguration) (IMessageStream, error) {
	if config.Protocol == iggcon.Tcp {
		tcpMessageStream, err := tcp.NewTcpMessageStream(config.BaseAddress)
		if err != nil {
			return nil, err
		}
		return tcpMessageStream, nil
	}

	return nil, errors.New("unsupported protocol")
}
