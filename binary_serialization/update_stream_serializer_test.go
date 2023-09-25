package binaryserialization

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	"testing"
)

func TestSerialize_UpdateStream(t *testing.T) {
	request := TcpUpdateStreamRequest{
		iggcon.UpdateStreamRequest{
			StreamId: iggcon.NewIdentifier("stream"),
			Name:     "update_stream",
		},
	}

	serialized1 := request.Serialize()

	expected := []byte{
		0x02,                               // StreamId Kind (StringId)
		0x06,                               // StreamId Length (2)
		0x73, 0x74, 0x72, 0x65, 0x61, 0x6D, // StreamId Value ("stream")
		0x0D,                                                                         // Name Length (13)
		0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5F, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6D, // Name ("update_stream")
	}

	if !areBytesEqual(serialized1, expected) {
		t.Errorf("Test case 1 failed. \nExpected:\t%v\nGot:\t\t%v", expected, serialized1)
	}
}
