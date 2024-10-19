package tcp

import (
	binaryserialization "github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func (tms *IggyTcpClient) GetStats() (*Stats, error) {
	buffer, err := tms.sendAndFetchResponse([]byte{}, GetStatsCode)
	if err != nil {
		return nil, err
	}

	responseLength, err := getResponseLength(buffer)
	if err != nil {
		return nil, err
	}

	_, responseBuffer, err := tms.read(responseLength)
	if err != nil {
		return nil, err
	}

	stats := &binaryserialization.TcpStats{}
	err = stats.Deserialize(responseBuffer)

	return &stats.Stats, err
}

func (tms *IggyTcpClient) Ping() error {
	_, err := tms.sendAndFetchResponse([]byte{}, PingCode)
	return err
}
