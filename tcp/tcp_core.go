package tcp

import (
	"encoding/binary"
	"net"

	. "github.com/iggy-rs/iggy-go-client/contracts"
	ierror "github.com/iggy-rs/iggy-go-client/errors"
)

type IggyTcpClient struct {
	client *net.TCPConn
}

const (
	InitialBytesLength   = 4
	ExpectedResponseSize = 8
	MaxStringLength      = 255
)

func NewTcpMessageStream(url string) (*IggyTcpClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}

	err = conn.SetKeepAlive(true)
	if err != nil {
		return nil, err
	}

	return &IggyTcpClient{client: conn}, nil
}

func (tms *IggyTcpClient) sendAndFetchResponse(message []byte, command CommandCode) ([]byte, error) {
	payload := createPayload(message, command)

	if _, err := tms.client.Write(payload); err != nil {
		return nil, err
	}

	buffer := make([]byte, ExpectedResponseSize)
	if _, err := tms.client.Read(buffer); err != nil {
		return nil, err
	}

	if responseCode := getResponseCode(buffer); responseCode != 0 {
		return nil, ierror.MapFromCode(responseCode)
	}

	return buffer, nil
}

func createPayload(message []byte, command CommandCode) []byte {
	messageLength := len(message) + 4
	messageBytes := make([]byte, InitialBytesLength+messageLength)
	binary.LittleEndian.PutUint32(messageBytes[:4], uint32(messageLength))
	binary.LittleEndian.PutUint32(messageBytes[4:8], uint32(command))
	copy(messageBytes[8:], message)
	return messageBytes
}

func getResponseCode(buffer []byte) int {
	return int(binary.LittleEndian.Uint32(buffer[:4]))
}

func getResponseLength(buffer []byte) (int, error) {
	length := int(binary.LittleEndian.Uint32(buffer[4:]))
	if length <= 1 {
		return 0, &ierror.IggyError{
			Code:    0,
			Message: "Received empty response.",
		}
	}
	return length, nil
}
