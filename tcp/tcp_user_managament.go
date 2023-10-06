package tcp

import (
	binaryserialization "github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func (tms *IggyTcpClient) GetUser(identifier Identifier) (*UserResponse, error) {
	message := binaryserialization.SerializeIdentifier(identifier)
	buffer, err := tms.sendAndFetchResponse(message, GetUserCode)
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

	return binaryserialization.DeserializeUser(responseBuffer)
}

func (tms *IggyTcpClient) GetUsers() ([]*UserResponse, error) {
	buffer, err := tms.sendAndFetchResponse([]byte{}, GetUsersCode)
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

	return binaryserialization.DeserializeUsers(responseBuffer)
}

func (tms *IggyTcpClient) CreateUser(request CreateUserRequest) error {
	var message []byte //TODO ADD SERIALIZATION
	_, err := tms.sendAndFetchResponse(message, CreateUserCode)
	return err
}

func (tms *IggyTcpClient) UpdateUser(request UpdateUserRequest) error {
	var message []byte //TODO ADD SERIALIZATION
	_, err := tms.sendAndFetchResponse(message, UpdateUserCode)
	return err
}

func (tms *IggyTcpClient) DeleteUser(identifier Identifier) error {
	message := binaryserialization.SerializeIdentifier(identifier)
	_, err := tms.sendAndFetchResponse(message, DeleteUserCode)
	return err
}

func (tms *IggyTcpClient) UpdateUserPermissions(request UpdateUserPermissionsRequest) error {
	var message []byte //TODO ADD SERIALIZATION
	_, err := tms.sendAndFetchResponse(message, UpdatePermissionsCode)
	return err
}

func (tms *IggyTcpClient) ChangePassword(request ChangePasswordRequest) error {
	var message []byte //TODO ADD SERIALIZATION
	_, err := tms.sendAndFetchResponse(message, ChangePasswordCode)
	return err
}
