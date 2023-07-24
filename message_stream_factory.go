package iggy

import "errors"

type IMessageStreamFactory interface {
	CreateStream(config MessageStreamConfiguration) (IMessageStream, error)
}

type MessageStreamFactory struct{}

func (msf *MessageStreamFactory) CreateMessageStream(config MessageStreamConfiguration) (IMessageStream, error) {
	if config.Protocol == Tcp {
		tcpMessageStream, err := NewTcpMessageStream(config.BaseAddress)
		if err != nil {
			return nil, err
		}
		return tcpMessageStream, nil
	}

	return nil, errors.New("unsupported protocol")
}
