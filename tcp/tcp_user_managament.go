package tcp

import (
	binaryserialization "github.com/iggy-rs/iggy-go-client/binary_serialization"
	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func (tms *IggyTcpClient) GetUser(identifier Identifier) (*UserResponse, error) {
	//TODO implement me
	panic(GetUserCode)
}

func (tms *IggyTcpClient) GetUsers() ([]UserResponse, error) {
	//TODO implement me
	panic(GetUsersCode)
}

func (tms *IggyTcpClient) CreateUser(request CreateUserRequest) error {
	//TODO implement me
	panic(CreateUserCode)
}

func (tms *IggyTcpClient) UpdateUser(request UpdateUserRequest) error {
	//TODO implement me
	panic(UpdateUserCode)
}

func (tms *IggyTcpClient) DeleteUser(identifier Identifier) error {
	message := binaryserialization.SerializeIdentifier(identifier)
	_, err := tms.sendAndFetchResponse(message, DeleteUserCode)
	return err
}

func (tms *IggyTcpClient) UpdateUserPermissions(request UpdateUserPermissions) error {
	//TODO implement me
	panic(UpdatePermissionsCode)
}

func (tms *IggyTcpClient) ChangePassword(request ChangePasswordRequest) error {
	//TODO implement me
	panic(ChangePasswordCode)
}
