package tcp_test

import (
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("UPDATE TOPIC:", func() {
	prefix := "UpdateTopic"
	When("User is logged in", func() {
		Context("and tries to update existing topic with a valid data", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			topicId, _ := successfullyCreateTopic(streamId, client)
			newName := createRandomString(128)
			request := iggcon.UpdateTopicRequest{
				TopicId:       iggcon.NewIdentifier(topicId),
				StreamId:      iggcon.NewIdentifier(streamId),
				Name:          newName,
				MessageExpiry: 0,
			}
			err := client.UpdateTopic(request)
			itShouldNotReturnError(err)
			itShouldSuccessfullyUpdateTopic(streamId, topicId, newName, client)
		})

		Context("and tries to create topic with duplicate topic name", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, client)
			_, topic1Name := successfullyCreateTopic(streamId, client)
			topic2Id, _ := successfullyCreateTopic(streamId, client)

			request := iggcon.UpdateTopicRequest{
				TopicId:       iggcon.NewIdentifier(topic2Id),
				StreamId:      iggcon.NewIdentifier(streamId),
				Name:          topic1Name,
				MessageExpiry: 0,
			}
			err := client.UpdateTopic(request)

			itShouldReturnSpecificError(err, "topic_name_already_exists")
		})

		Context("and tries to update non-existing topic", func() {
			client := createAuthorizedStream()
			streamId := int(createRandomUInt32())
			topicId := int(createRandomUInt32())
			request := iggcon.UpdateTopicRequest{
				TopicId:       iggcon.NewIdentifier(topicId),
				StreamId:      iggcon.NewIdentifier(streamId),
				Name:          createRandomString(128),
				MessageExpiry: 0,
			}
			err := client.UpdateTopic(request)

			itShouldReturnSpecificError(err, "stream_id_not_found")
		})

		Context("and tries to update non-existing stream", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, createAuthorizedStream())
			topicId := int(createRandomUInt32())
			request := iggcon.UpdateTopicRequest{
				TopicId:       iggcon.NewIdentifier(topicId),
				StreamId:      iggcon.NewIdentifier(streamId),
				Name:          createRandomString(128),
				MessageExpiry: 0,
			}
			err := client.UpdateTopic(request)

			itShouldReturnSpecificError(err, "topic_id_not_found")
		})

		Context("and tries to update existing topic with a name that's over 255 characters", func() {
			client := createAuthorizedStream()
			streamId, _ := successfullyCreateStream(prefix, client)
			defer deleteStreamAfterTests(streamId, createAuthorizedStream())
			topicId, _ := successfullyCreateTopic(streamId, client)
			request := iggcon.UpdateTopicRequest{
				TopicId:       iggcon.NewIdentifier(topicId),
				StreamId:      iggcon.NewIdentifier(streamId),
				Name:          createRandomString(256),
				MessageExpiry: 0,
			}

			err := client.UpdateTopic(request)

			itShouldReturnSpecificError(err, "topic_name_too_long")
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
