package tcpserialization

import (
	"encoding/binary"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	"reflect"
	"testing"
)

func TestSerialize_TcpCreateStreamRequest(t *testing.T) {
	// Create a sample TcpCreateStreamRequest
	request := TcpCreateStreamRequest{
		CreateStreamRequest: iggcon.CreateStreamRequest{
			StreamId: 123,
			Name:     "test_stream",
		},
	}

	// Serialize the request
	serialized := request.Serialize()

	// Expected serialized bytes
	expectedStreamID := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedStreamID, 123)
	expectedNameLength := byte(11) // Length of "test_stream"
	expectedPayload := []byte("test_stream")

	// Check the length of the serialized bytes
	if len(serialized) != int(payloadOffset+len(request.Name)) {
		t.Errorf("Serialized length is incorrect. Expected: %d, Got: %d", payloadOffset+len(request.Name), len(serialized))
	}

	// Check the StreamID field
	if !reflect.DeepEqual(serialized[streamIDOffset:streamIDOffset+4], expectedStreamID) {
		t.Errorf("StreamID is incorrect. Expected: %v, Got: %v", expectedStreamID, serialized[streamIDOffset:streamIDOffset+4])
	}

	// Check the NameLength field
	if serialized[nameLengthOffset] != expectedNameLength {
		t.Errorf("NameLength is incorrect. Expected: %d, Got: %d", expectedNameLength, serialized[nameLengthOffset])
	}

	// Check the Payload field
	if !reflect.DeepEqual(serialized[payloadOffset:], expectedPayload) {
		t.Errorf("Payload is incorrect. Expected: %v, Got: %v", expectedPayload, serialized[payloadOffset:])
	}
}
