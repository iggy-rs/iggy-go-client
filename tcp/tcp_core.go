package tcp

import (
	"context"
	"encoding/binary"
	"net"
	"sync"

	. "github.com/iggy-rs/iggy-go-client/contracts"
	ierror "github.com/iggy-rs/iggy-go-client/errors"
)

type IggyTcpClient struct {
	client *net.TCPConn
	mtx    sync.Mutex
}

const (
	InitialBytesLength   = 4
	ExpectedResponseSize = 8
	MaxStringLength      = 255
)

func NewTcpMessageStream(ctx context.Context, url string) (*IggyTcpClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return nil, err
	}

	var d = net.Dialer{
		KeepAlive: -1,
	}
	conn, err := d.DialContext(ctx, "tcp", addr.String())
	if err != nil {
		return nil, err
	}

	return &IggyTcpClient{client: conn.(*net.TCPConn)}, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (tms *IggyTcpClient) read(expectedSize int) (int, []byte, error) {
	var totalRead int
	buffer := make([]byte, expectedSize)

	for totalRead < expectedSize {
		readSize := expectedSize - totalRead
		n, err := tms.client.Read(buffer[totalRead : totalRead+readSize])
		if err != nil {
			return totalRead, buffer[:totalRead], err
		}
		totalRead += n
	}

	return totalRead, buffer, nil
}

func (tms *IggyTcpClient) write(payload []byte) (int, error) {
	n, err := tms.client.Write(payload)
	if err != nil {
		return n, err
	}

	return n, nil
}

func (tms *IggyTcpClient) sendAndFetchResponse(message []byte, command CommandCode) ([]byte, error) {
	// ! TODO: aditional locks may be required for multiple tcp conns
	// tms.mtx.Lock()
	// defer tms.mtx.Unlock()

	payload := createPayload(message, command)
	if _, err := tms.write(payload); err != nil {
		return nil, err
	}

	_, buffer, err := tms.read(ExpectedResponseSize)
	if err != nil {
		return nil, err
	}

	length := int(binary.LittleEndian.Uint32(buffer[4:]))

	if responseCode := getResponseCode(buffer); responseCode != 0 {
		// TEMP: See https://github.com/iggy-rs/iggy/pull/604 for context.
		// from: https://github.com/iggy-rs/iggy/blob/master/sdk/src/tcp/client.rs#L217
		if responseCode == 2012 ||
			responseCode == 2013 ||
			responseCode == 1011 ||
			responseCode == 1012 ||
			responseCode == 46 ||
			responseCode == 51 ||
			responseCode == 5001 ||
			responseCode == 5004 {
		} else {
			return nil, ierror.MapFromCode(responseCode)
		}

		// ! TODO: Should handle full support for decoding these messages
		// for now still need to read bytes to stop comply with spec
		_, _, err := tms.read(length)
		if err != nil {
			return nil, err
		}

		return buffer, ierror.MapFromCode(responseCode)
	}

	// Added to support messages that do not send back bytes
	// from: https://github.com/iggy-rs/iggy/blob/214f0ca9368a74164caa4aa5cc55320dfa49ee6a/sdk/src/tcp/client.rs#L363
	if length <= 1 {
		return []byte{}, nil
	}

	_, buffer, err = tms.read(length)
	if err != nil {
		return nil, err
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
