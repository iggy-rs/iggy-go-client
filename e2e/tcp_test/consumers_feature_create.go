package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("CREATE CONSUMER GROUP:", func() {
	prefix := "CreateConsumerGroup"
	When("User is logged in", func() {
		Context("and tries to create consumer group unique name and id", func() {
			client := createAuthorizedConnection()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			topicId, _ := successfullyCreateTopic(streamId, client)

			request := iggcon.CreateConsumerGroupRequest{
				StreamId:        iggcon.NewIdentifier(streamId),
				TopicId:         iggcon.NewIdentifier(topicId),
				ConsumerGroupId: int(createRandomUInt32()),
				Name:            createRandomString(16),
			}
			err := client.CreateConsumerGroup(request)

			itShouldNotReturnError(err)
			itShouldSuccessfullyCreateConsumer(streamId, topicId, request.ConsumerGroupId, request.Name, client)
		})

		Context("and tries to create consumer group for a non existing stream", func() {
			client := createAuthorizedConnection()
			request := iggcon.CreateConsumerGroupRequest{
				StreamId:        iggcon.NewIdentifier(int(createRandomUInt32())),
				TopicId:         iggcon.NewIdentifier(int(createRandomUInt32())),
				ConsumerGroupId: int(createRandomUInt32()),
				Name:            createRandomString(16),
			}
			err := client.CreateConsumerGroup(request)

			itShouldReturnSpecificError(err, "stream_id_not_found")
		})

		Context("and tries to create consumer group for a non existing topic", func() {
			client := createAuthorizedConnection()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			request := iggcon.CreateConsumerGroupRequest{
				StreamId:        iggcon.NewIdentifier(streamId),
				TopicId:         iggcon.NewIdentifier(int(createRandomUInt32())),
				ConsumerGroupId: int(createRandomUInt32()),
				Name:            createRandomString(16),
			}
			err := client.CreateConsumerGroup(request)

			itShouldReturnSpecificError(err, "stream_id_not_found")
		})

		Context("and tries to create consumer group with duplicate group name", func() {
			client := createAuthorizedConnection()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			topicId, _ := successfullyCreateTopic(streamId, client)
			_, name := successfullyCreateConsumer(streamId, topicId, client)

			request := iggcon.CreateConsumerGroupRequest{
				StreamId:        iggcon.NewIdentifier(streamId),
				TopicId:         iggcon.NewIdentifier(topicId),
				ConsumerGroupId: int(createRandomUInt32()),
				Name:            name,
			}
			err := client.CreateConsumerGroup(request)

			itShouldReturnSpecificError(err, "group_name_already_exists")
		})

		Context("and tries to create consumer group with duplicate group id", func() {
			client := createAuthorizedConnection()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			topicId, _ := successfullyCreateTopic(streamId, client)
			groupId, _ := successfullyCreateConsumer(streamId, topicId, client)

			request := iggcon.CreateConsumerGroupRequest{
				StreamId:        iggcon.NewIdentifier(streamId),
				TopicId:         iggcon.NewIdentifier(topicId),
				ConsumerGroupId: groupId,
				Name:            createRandomString(16),
			}
			err := client.CreateConsumerGroup(request)

			itShouldReturnSpecificError(err, "group_id_already_exists")
		})

		Context("and tries to create group with name that's over 255 characters", func() {
			client := createAuthorizedConnection()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			topicId, _ := successfullyCreateTopic(streamId, client)

			request := iggcon.CreateConsumerGroupRequest{
				StreamId:        iggcon.NewIdentifier(streamId),
				TopicId:         iggcon.NewIdentifier(topicId),
				ConsumerGroupId: int(createRandomUInt32()),
				Name:            createRandomString(256),
			}
			err := client.CreateConsumerGroup(request)

			itShouldNotReturnError(err)

			itShouldReturnSpecificError(err, "group_name_too_long")
		})
	})

	When("User is not logged in", func() {
		Context("and tries to create consumer group", func() {
			client := createConnection()
			request := iggcon.CreateConsumerGroupRequest{
				StreamId:        iggcon.NewIdentifier(int(createRandomUInt32())),
				TopicId:         iggcon.NewIdentifier(int(createRandomUInt32())),
				ConsumerGroupId: int(createRandomUInt32()),
				Name:            createRandomString(16),
			}
			err := client.CreateConsumerGroup(request)

			itShouldReturnUnauthenticatedError(err)
		})
	})
})
