package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("GET TOPIC BY ID:", func() {
	prefix := "GetTopicById"
	When("User is logged in", func() {
		Context("and tries to get existing topic", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			topicId, name := successfullyCreateTopic(streamId, client)
			topic, err := client.GetTopicById(iggcon.NewIdentifier(streamId), iggcon.NewIdentifier(topicId))

			itShouldNotReturnError(err)
			itShouldReturnSpecificTopic(topicId, name, *topic)
		})

		Context("and tries to get non-existing topic", func() {
			client := createAuthorizedStream()
			streamId := int(createRandomUInt32())

			_, err := client.GetTopicById(iggcon.NewIdentifier(streamId), iggcon.NewIdentifier(int(createRandomUInt32())))

			itShouldReturnSpecificError(err, "stream_id_not_found")
		})

		Context("and tries to get non-existing topic", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)

			_, err := client.GetTopicById(iggcon.NewIdentifier(streamId), iggcon.NewIdentifier(int(createRandomUInt32())))

			itShouldReturnSpecificError(err, "topic_id_not_found")
		})
	})

	When("User is not logged in", func() {
		Context("and tries to get topic by id", func() {
			client := createMessageStream()
			_, err := client.GetTopicById(iggcon.NewIdentifier(int(createRandomUInt32())), iggcon.NewIdentifier(int(createRandomUInt32())))

			itShouldReturnUnauthenticatedError(err)
		})
	})
})
