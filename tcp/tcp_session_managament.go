package tcp

import (
	"github.com/iggy-rs/iggy-go-client/binary_serialization"

	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func (tms *IggyTcpClient) LogIn(request LogInRequest) (*LogInResponse, error) {
	serializedRequest := binaryserialization.TcpLogInRequest{LogInRequest: request}
	buffer, err := tms.sendAndFetchResponse(serializedRequest.Serialize(), LoginUserCode)
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

	return binaryserialization.DeserializeLogInResponse(responseBuffer), nil
}

func (tms *IggyTcpClient) LogInWithAccessToken(request LogInAccessTokenRequest) (*LogInResponse, error) {
	message := binaryserialization.SerializeLoginWithPersonalAccessToken(request)
	buffer, err := tms.sendAndFetchResponse(message, LoginWithAccessTokenCode)
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

	return binaryserialization.DeserializeLogInResponse(responseBuffer), nil
}

func (tms *IggyTcpClient) LogOut() error {
	_, err := tms.sendAndFetchResponse([]byte{}, LogoutUserCode)
	return err
}
