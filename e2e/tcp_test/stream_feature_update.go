package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("UPDATE STREAM:", func() {
	When("User is logged in", func() {
		Context("and tries to update existing stream with a valid name", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(client)
			newName := createRandomString(128)

			err := client.UpdateStream(iggcon.UpdateStreamRequest{
				StreamId: iggcon.NewIdentifier(streamId),
				Name:     newName,
			})
			itShouldNotReturnError(err)
			itShouldSuccessfullyUpdateStream(streamId, newName, client)
		})

		Context("and tries to update stream with duplicate stream name", func() {
			client := createAuthorizedStream()
			_, stream1Name := successfullyCreateStream(client)
			stream2Id, _ := successfullyCreateStream(client)

			err := client.UpdateStream(iggcon.UpdateStreamRequest{
				StreamId: iggcon.NewIdentifier(stream2Id),
				Name:     stream1Name,
			})

			itShouldReturnSpecificError(err, "stream_name_already_exists")
		})

		Context("and tries to update non-existing stream", func() {
			client := createAuthorizedStream()
			err := client.UpdateStream(iggcon.UpdateStreamRequest{
				StreamId: iggcon.NewIdentifier(int(createRandomUInt32())),
				Name:     createRandomString(128),
			})

			itShouldReturnSpecificError(err, "stream_id_not_found")
		})

		Context("and tries to update existing stream with a name that's over 255 characters", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(client)

			err := client.UpdateStream(iggcon.UpdateStreamRequest{
				StreamId: iggcon.NewIdentifier(streamId),
				Name:     createRandomString(256),
			})
			itShouldReturnError(err)
		})
	})

	When("User is not logged in", func() {
		Context("and tries to update stream", func() {
			client := createMessageStream()
			err := client.UpdateStream(iggcon.UpdateStreamRequest{
				StreamId: iggcon.NewIdentifier(int(createRandomUInt32())),
				Name:     createRandomString(128),
			})

			itShouldReturnUnauthenticatedError(err)
		})
	})
})
