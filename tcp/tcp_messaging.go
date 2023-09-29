package tcp

import (
	"github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
	ierror "github.com/iggy-rs/iggy-go-client/errors"
)

func (tms *IggyTcpClient) SendMessages(request SendMessagesRequest) error {
	if len(request.Messages) == 0 {
		return ierror.CustomError("messages_count_should_be_greater_than_zero")
	}
	serializedRequest := binaryserialization.TcpSendMessagesRequest{SendMessagesRequest: request}
	_, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), SendMessagesCode)
	return err
}

func (tms *IggyTcpClient) PollMessages(request FetchMessagesRequest) (*FetchMessagesResponse, error) {
	serializedRequest := binaryserialization.TcpFetchMessagesRequest{FetchMessagesRequest: request}
	buffer, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), PollMessagesCode)
	if err != nil {
		return nil, err
	}

	responseLength, err := getResponseLength(buffer)
	if err != nil {
		return nil, err
	}

	responseBuffer := make([]byte, responseLength)
	if _, err := tms.client.Read(responseBuffer); err != nil {
		return nil, err
	}

	return binaryserialization.DeserializeFetchMessagesResponse(responseBuffer)
}
