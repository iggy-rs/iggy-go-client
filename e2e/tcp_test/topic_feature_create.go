package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("CREATE TOPIC:", func() {
	When("User is logged in", func() {
		Context("and tries to create topic unique name and id", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(client)

			request := iggcon.CreateTopicRequest{
				TopicId:         1,
				StreamId:        iggcon.NewIdentifier(streamId),
				Name:            createRandomString(32),
				MessageExpiry:   1000,
				PartitionsCount: 2,
			}
			err := client.CreateTopic(request)

			itShouldNotReturnError(err)
			itShouldSuccessfullyCreateTopic(streamId, request.TopicId, request.Name, client)
		})

		Context("and tries to create topic with duplicate topic name", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(client)
			_, name := successfullyCreateTopic(streamId, client)

			request := iggcon.CreateTopicRequest{
				TopicId:         int(createRandomUInt32()),
				StreamId:        iggcon.NewIdentifier(streamId),
				Name:            name,
				MessageExpiry:   0,
				PartitionsCount: 2,
			}
			err := client.CreateTopic(request)
			itShouldReturnSpecificError(err, "topic_name_already_exists")
		})

		Context("and tries to create topic with duplicate topic id", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(client)
			topicId, _ := successfullyCreateTopic(streamId, client)

			request := iggcon.CreateTopicRequest{
				TopicId:         topicId,
				StreamId:        iggcon.NewIdentifier(streamId),
				Name:            createRandomString(32),
				MessageExpiry:   0,
				PartitionsCount: 2,
			}
			err := client.CreateTopic(request)
			itShouldReturnSpecificError(err, "topic_id_already_exists")
		})

		Context("and tries to create topic with name that's over 255 characters", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(client)

			request := iggcon.CreateTopicRequest{
				TopicId:         int(createRandomUInt32()),
				StreamId:        iggcon.NewIdentifier(streamId),
				Name:            createRandomString(256),
				MessageExpiry:   0,
				PartitionsCount: 2,
			}
			err := client.CreateTopic(request)

			itShouldReturnError(err)
		})
	})

	When("User is not logged in", func() {
		Context("and tries to create topic", func() {
			client := createMessageStream()
			err := client.CreateTopic(iggcon.CreateTopicRequest{
				TopicId:         1,
				StreamId:        iggcon.NewIdentifier(10),
				Name:            "name",
				MessageExpiry:   0,
				PartitionsCount: 2,
			})

			itShouldReturnUnauthenticatedError(err)
		})
	})
})
