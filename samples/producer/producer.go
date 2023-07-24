package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	iggy "github.com/eldpcn/iggy-go"
	sharedDemoContracts "github.com/eldpcn/iggy-go/samples/shared"
)

const (
	StreamId          = 1
	TopicId           = 1
	MessageBatchCount = 1
	PartitionId       = 1
	Interval          = 1000
)

func main() {
	factory := &iggy.MessageStreamFactory{}
	config := iggy.MessageStreamConfiguration{
		BaseAddress: "127.0.0.1:8090",
		Protocol:    iggy.Tcp,
	}

	messageStream, err := factory.CreateMessageStream(config)
	if err != nil {
		panic(err)
	}

	if err = EnsureInsfrastructureIsInitialized(messageStream); err != nil {
		panic(err)
	}

	if err = PublishMessages(messageStream); err != nil {
		panic(err)
	}
}

func EnsureInsfrastructureIsInitialized(messageStream iggy.IMessageStream) error {
	if _, streamErr := messageStream.GetStreamById(StreamId); streamErr != nil {
		streamErr = messageStream.CreateStream(iggy.StreamRequest{
			StreamId: StreamId,
			Name:     "Test Producer Stream",
		})

		if streamErr != nil {
			panic(streamErr)
		}

		fmt.Printf("Created stream with ID: %d.\n", StreamId)
	}

	fmt.Printf("Stream with ID: %d exists.\n", StreamId)

	if _, topicErr := messageStream.GetTopicById(StreamId, TopicId); topicErr != nil {
		topicErr = messageStream.CreateTopic(StreamId, iggy.TopicRequest{
			TopicId:         TopicId,
			Name:            "Test Topic From Producer Sample",
			PartitionsCount: 12,
		})

		if topicErr != nil {
			panic(topicErr)
		}

		fmt.Printf("Created topic with ID: %d.\n", TopicId)
	}

	fmt.Printf("Topic with ID: %d exists.\n", TopicId)

	return nil
}

func PublishMessages(messageStream iggy.IMessageStream) error {
	fmt.Printf("Messages will be sent to stream '%d', topic '%d', partition '%d' with interval %d ms.\n", StreamId, TopicId, PartitionId, Interval)
	messageGenerator := NewMessageGenerator()

	for {
		var debugMessages []sharedDemoContracts.ISerializableMessage
		var messages []iggy.Message

		for i := 0; i < MessageBatchCount; i++ {
			message := messageGenerator.GenerateMessage()
			json := message.ToBytes()

			debugMessages = append(debugMessages, message)
			messages = append(messages, iggy.Message{
				Id:      uuid.New(),
				Payload: json,
			})
		}

		err := messageStream.SendMessages(StreamId, TopicId, iggy.MessageSendRequest{
			Messages: messages,
			KeyKind:  iggy.PartitionId,
			KeyValue: PartitionId,
		})
		if err != nil {
			return nil
		}

		for _, m := range debugMessages {
			fmt.Println("Sent messages:", m.ToJson())
		}

		time.Sleep(time.Millisecond * time.Duration(Interval))
	}
}
