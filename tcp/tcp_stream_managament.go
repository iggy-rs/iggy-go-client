package tcp

import (
	"github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
	ierror "github.com/iggy-rs/iggy-go-client/errors"
)

func (tms *IggyTcpClient) GetStreams() ([]StreamResponse, error) {
	buffer, err := tms.sendAndFetchResponse([]byte{}, GetStreamsCode)
	if err != nil {
		return nil, err
	}

	responseLength, err := getResponseLength(buffer)
	if err != nil {
		return nil, err
	}

	responseBuffer := make([]byte, responseLength)
	_, err = tms.client.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return binaryserialization.DeserializeStreams(responseBuffer), nil
}

func (tms *IggyTcpClient) GetStreamById(request GetStreamRequest) (*StreamResponse, error) {
	message := binaryserialization.SerializeIdentifier(request.StreamID)
	buffer, err := tms.sendAndFetchResponse(message, GetStreamCode)
	if err != nil {
		return nil, err
	}

	responseLength, err := getResponseLength(buffer)
	if err != nil {
		return nil, err
	}

	responseBuffer := make([]byte, responseLength)
	_, err = tms.client.Read(responseBuffer)
	if err != nil {
		return nil, err
	}

	return binaryserialization.DeserializerStream(responseBuffer), nil
}

func (tms *IggyTcpClient) CreateStream(request CreateStreamRequest) error {
	if MaxStringLength < len(request.Name) {
		return ierror.TextTooLong("stream_name")
	}
	serializedRequest := binaryserialization.TcpCreateStreamRequest{CreateStreamRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), CreateStreamCode)
	return err
}

func (tms *IggyTcpClient) UpdateStream(request UpdateStreamRequest) error {
	if MaxStringLength <= len(request.Name) {
		return ierror.TextTooLong("stream_name")
	}
	serializedRequest := binaryserialization.TcpUpdateStreamRequest{UpdateStreamRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), UpdateStreamCode)
	return err
}

func (tms *IggyTcpClient) DeleteStream(id Identifier) error {
	message := binaryserialization.SerializeIdentifier(id)
	_, err := tms.sendAndFetchResponse(message, DeleteStreamCode)
	return err
}
