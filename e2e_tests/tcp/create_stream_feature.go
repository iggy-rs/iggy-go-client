package tcp

import (
	"github.com/iggy-rs/iggy-go-client"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CREATE STREAM:", func() {
	When("User is logged in", func() {
		Context("and tries to create stream unique name and id", func() {
			client := createAuthorizedStream()
			streamId := int(createRandomUInt32())
			name := CreateRandomString(32)

			err := client.CreateStream(iggcon.CreateStreamRequest{
				StreamId: streamId,
				Name:     name,
			})

			itShouldNotReturnError(err)
			itShouldSuccessfullyCreateStream(streamId, name, client)
		})

		Context("and tries to create stream with duplicate stream name", func() {
			client := createAuthorizedStream()
			streamId := int(createRandomUInt32())
			name := CreateRandomString(32)

			err := client.CreateStream(iggcon.CreateStreamRequest{
				StreamId: streamId,
				Name:     name,
			})

			itShouldNotReturnError(err)
			itShouldSuccessfullyCreateStream(streamId, name, client)

			err = client.CreateStream(iggcon.CreateStreamRequest{
				StreamId: int(createRandomUInt32()),
				Name:     name,
			})

			itShouldReturnSpecificError(err, "stream_name_already_exists")
		})

		Context("and tries to create stream with duplicate stream id", func() {
			client := createAuthorizedStream()
			streamId := int(createRandomUInt32())
			name := CreateRandomString(32)

			err := client.CreateStream(iggcon.CreateStreamRequest{
				StreamId: streamId,
				Name:     name,
			})

			itShouldNotReturnError(err)
			itShouldSuccessfullyCreateStream(streamId, name, client)

			err = client.CreateStream(iggcon.CreateStreamRequest{
				StreamId: streamId,
				Name:     CreateRandomString(32),
			})

			itShouldReturnSpecificError(err, "stream_id_already_exists")
		})

		Context("and tries to create stream name that's over 255 characters", func() {
			client := createAuthorizedStream()
			streamId := int(createRandomUInt32())
			name := CreateRandomString(256)

			err := client.CreateStream(iggcon.CreateStreamRequest{
				StreamId: streamId,
				Name:     name,
			})

			itShouldReturnError(err)
		})
	})

	When("User is not logged in", func() {
		Context("and tries to create stream", func() {
			client := createMessageStream()
			err := client.CreateStream(iggcon.CreateStreamRequest{
				StreamId: int(createRandomUInt32()),
				Name:     CreateRandomString(32),
			})

			itShouldReturnUnauthenticatedError(err)
		})
	})
})

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
