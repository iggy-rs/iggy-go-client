package tcp_test

import (
	"github.com/iggy-rs/iggy-go-client"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

//operations

func successfullyCreateStream(prefix string, client iggy.MessageStream) (int, string) {
	streamId := int(createRandomUInt32())
	name := createRandomStringWithPrefix(prefix, 128)

	err := client.CreateStream(iggcon.CreateStreamRequest{
		StreamId: streamId,
		Name:     name,
	})

	itShouldNotReturnError(err)
	itShouldSuccessfullyCreateStream(streamId, name, client)
	return streamId, name
}

//assertions

func itShouldReturnSpecificStream(id int, name string, stream iggcon.StreamResponse) {
	It("should fetch stream with id "+string(rune(id)), func() {
		Expect(stream.Id).To(Equal(id))
	})

	It("should fetch stream with name "+name, func() {
		Expect(stream.Name).To(Equal(name))
	})
}

func itShouldContainSpecificStream(id int, name string, streams []iggcon.StreamResponse) {
	It("should fetch at least one stream", func() {
		Expect(len(streams)).NotTo(Equal(0))
	})

	var stream iggcon.StreamResponse
	found := false

	for _, s := range streams {
		if s.Id == id && s.Name == name {
			stream = s
			found = true
			break
		}
	}

	It("should fetch stream with id "+strconv.Itoa(id), func() {
		Expect(found).To(BeTrue(), "Stream with id %d and name %s not found", id, name)
		Expect(stream.Id).To(Equal(id))
	})

	It("should fetch stream with name "+name, func() {
		Expect(found).To(BeTrue(), "Stream with id %d and name %s not found", id, name)
		Expect(stream.Name).To(Equal(name))
	})
}

func itShouldSuccessfullyCreateStream(id int, expectedName string, client iggy.MessageStream) {
	stream, err := client.GetStreamById(iggcon.GetStreamRequest{StreamID: iggcon.NewIdentifier(id)})

	itShouldNotReturnError(err)
	It("should create stream with id "+string(rune(id)), func() {
		Expect(stream.Id).To(Equal(id))
	})

	It("should create stream with name "+expectedName, func() {
		Expect(stream.Name).To(Equal(expectedName))
	})
}

func itShouldSuccessfullyUpdateStream(id int, expectedName string, client iggy.MessageStream) {
	stream, err := client.GetStreamById(iggcon.GetStreamRequest{StreamID: iggcon.NewIdentifier(id)})

	itShouldNotReturnError(err)
	It("should update stream with id "+string(rune(id)), func() {
		Expect(stream.Id).To(Equal(id))
	})

	It("should update stream with name "+expectedName, func() {
		Expect(stream.Name).To(Equal(expectedName))
	})
}

func itShouldSuccessfullyDeletedStream(id int, client iggy.MessageStream) {
	stream, err := client.GetStreamById(iggcon.GetStreamRequest{StreamID: iggcon.NewIdentifier(id)})

	itShouldReturnSpecificError(err, "stream_id_not_found")
	It("should not return stream", func() {
		Expect(stream).To(BeNil())
	})
}

func deleteStreamAfterTests(streamId int, client iggy.MessageStream) {
	_ = client.DeleteStream(iggcon.NewIdentifier(streamId))
}
