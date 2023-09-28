package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("DELETE STREAM:", func() {
	When("User is logged in", func() {
		Context("and tries to delete existing stream", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(client)
			err := client.DeleteStream(iggcon.NewIdentifier(streamId))

			itShouldNotReturnError(err)
			itShouldSuccessfullyDeletedStream(streamId, client)
		})

		Context("and tries to delete non-existing stream", func() {
			client := createAuthorizedStream()
			streamId := int(createRandomUInt32())

			err := client.DeleteStream(iggcon.NewIdentifier(streamId))

			itShouldReturnSpecificError(err, "stream_id_not_found")
		})
	})

	When("User is not logged in", func() {
		Context("and tries to delete stream", func() {
			client := createMessageStream()
			err := client.DeleteStream(iggcon.NewIdentifier(int(createRandomUInt32())))

			itShouldReturnUnauthenticatedError(err)
		})
	})
})
