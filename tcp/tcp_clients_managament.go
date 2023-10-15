package tcp

import (
	binaryserialization "github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func (tms *IggyTcpClient) GetClients() ([]ClientResponse, error) {
	buffer, err := tms.sendAndFetchResponse([]byte{}, GetClientsCode)
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

	return binaryserialization.DeserializeClients(responseBuffer)
}

func (tms *IggyTcpClient) GetClientById(clientId int) (*ClientResponse, error) {
	message := binaryserialization.SerializeInt(clientId)
	buffer, err := tms.sendAndFetchResponse(message, GetClientCode)
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

	return binaryserialization.DeserializeClient(responseBuffer), nil
}
