package tcp

import (
	. "github.com/iggy-rs/iggy-go-client/contracts"
)

func (tms *IggyTcpClient) GetMe() (*MeResponse, error) {
	//TODO implement me
	panic(GetMeCode)
}

func (tms *IggyTcpClient) GetClients() ([]ClientResponse, error) {
	//TODO implement me
	panic(GetClientsCode)
}

func (tms *IggyTcpClient) GetClientById(clientId int) (*ClientResponse, error) {
	//TODO implement me
	panic(GetClientCode)
}
