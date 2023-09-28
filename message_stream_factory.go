package iggy

import (
	"errors"
	"github.com/iggy-rs/iggy-go-client/tcp"

	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
)

type IMessageStreamFactory interface {
	CreateStream(config iggcon.IggyConfiguration) (MessageStream, error)
}

type IggyClientFactory struct{}

func (msf *IggyClientFactory) CreateMessageStream(config iggcon.IggyConfiguration) (MessageStream, error) {
	if config.Protocol == iggcon.Tcp {
		tcpMessageStream, err := tcp.NewTcpMessageStream(config.BaseAddress)
		if err != nil {
			return nil, err
		}
		return tcpMessageStream, nil
	}

	return nil, errors.New("unsupported protocol")
}
