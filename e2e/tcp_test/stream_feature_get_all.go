package tcp_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("GET ALL STREAMS:", func() {
	When("User is logged in", func() {
		Context("and tries to get all streams", func() {
			client := createAuthorizedStream()
			streamId, name := successfullyCreateStream(client)
			streams, err := client.GetStreams()

			itShouldNotReturnError(err)
			itShouldContainSpecificStream(streamId, name, streams)
		})
	})

	When("User is not logged in", func() {
		Context("and tries to delete stream", func() {
			client := createMessageStream()
			_, err := client.GetStreams()

			itShouldReturnUnauthenticatedError(err)
		})
	})
})
