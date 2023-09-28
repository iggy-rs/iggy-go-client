package tcp_test

import (
	"github.com/iggy-rs/iggy-go-client"
	iggcon "github.com/iggy-rs/iggy-go-client/contracts"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
)

//operations

func successfullyCreateTopic(streamId int, client iggy.MessageStream) (int, string) {
	request := iggcon.CreateTopicRequest{
		TopicId:         int(createRandomUInt32()),
		StreamId:        iggcon.NewIdentifier(streamId),
		Name:            createRandomString(128),
		MessageExpiry:   0,
		PartitionsCount: 2,
	}
	err := client.CreateTopic(request)

	itShouldSuccessfullyCreateTopic(streamId, request.TopicId, request.Name, client)
	itShouldNotReturnError(err)
	return request.TopicId, request.Name
}

//assertions

func itShouldReturnSpecificTopic(id int, name string, topic iggcon.TopicResponse) {
	It("should fetch topic with id "+string(rune(id)), func() {
		Expect(topic.Id).To(Equal(id))
	})

	It("should fetch topic with name "+name, func() {
		Expect(topic.Name).To(Equal(name))
	})
}

func itShouldContainSpecificTopic(id int, name string, topics []iggcon.TopicResponse) {
	It("should fetch at least one topic", func() {
		Expect(len(topics)).NotTo(Equal(0))
	})

	var topic iggcon.TopicResponse
	found := false

	for _, s := range topics {
		if s.Id == id && s.Name == name {
			topic = s
			found = true
			break
		}
	}

	It("should fetch topic with id "+strconv.Itoa(id), func() {
		Expect(found).To(BeTrue(), "Topic with id %d and name %s not found", id, name)
		Expect(topic.Id).To(Equal(id))
	})

	It("should fetch topic with name "+name, func() {
		Expect(found).To(BeTrue(), "Topic with id %d and name %s not found", id, name)
		Expect(topic.Name).To(Equal(name))
	})
}

func itShouldSuccessfullyCreateTopic(streamId int, topicId int, expectedName string, client iggy.MessageStream) {
	topic, err := client.GetTopicById(iggcon.NewIdentifier(streamId), iggcon.NewIdentifier(topicId))

	It("should create topic with id "+string(rune(topicId)), func() {
		Expect(topic).NotTo(BeNil())
		Expect(topic.Id).To(Equal(topicId))
	})

	It("should create topic with name "+expectedName, func() {
		Expect(topic).NotTo(BeNil())
		Expect(topic.Name).To(Equal(expectedName))
	})
	itShouldNotReturnError(err)
}

func itShouldSuccessfullyUpdateTopic(streamId int, topicId int, expectedName string, client iggy.MessageStream) {
	topic, err := client.GetTopicById(iggcon.NewIdentifier(streamId), iggcon.NewIdentifier(topicId))

	It("should update topic with id "+string(rune(topicId)), func() {
		Expect(topic).NotTo(BeNil())
		Expect(topic.Id).To(Equal(topicId))
	})

	It("should update topic with name "+expectedName, func() {
		Expect(topic).NotTo(BeNil())
		Expect(topic.Name).To(Equal(expectedName))
	})
	itShouldNotReturnError(err)
}

func itShouldSuccessfullyDeletedTopic(streamId int, topicId int, client iggy.MessageStream) {
	topic, err := client.GetTopicById(iggcon.NewIdentifier(streamId), iggcon.NewIdentifier(topicId))

	itShouldReturnSpecificError(err, "topic_id_not_found")
	It("should not return topic", func() {
		Expect(topic).To(BeNil())
	})
}
