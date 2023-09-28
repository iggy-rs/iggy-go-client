package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/iggy-rs/iggy-go-client"
	. "github.com/iggy-rs/iggy-go-client/contracts"
	sharedDemoContracts "github.com/iggy-rs/iggy-go-client/samples/shared"
)

// config
const (
	DefaultStreamId = 1
	TopicId         = 1
	Partition       = 1
	Interval        = 1000
	ConsumerId      = 1
)

func main() {
	factory := &IggyClientFactory{}
	config := IggyConfiguration{
		BaseAddress: "127.0.0.1:8090",
		Protocol:    Tcp,
	}

	messageStream, err := factory.CreateMessageStream(config)
	if err != nil {
		panic(err)
	}

	_, err = messageStream.LogIn(LogInRequest{
		Username: "iggy",
		Password: "iggy",
	})
	if err != nil {
		panic("COULD NOT LOG IN")
	}

	if err = EnsureInsfrastructureIsInitialized(messageStream); err != nil {
		panic(err)
	}

	if err := ConsumeMessages(messageStream); err != nil {
		panic(err)
	}
}

func EnsureInsfrastructureIsInitialized(messageStream MessageStream) error {
	if _, streamErr := messageStream.GetStreamById(GetStreamRequest{
		StreamID: NewIdentifier(DefaultStreamId),
	}); streamErr != nil {
		streamErr = messageStream.CreateStream(CreateStreamRequest{
			StreamId: DefaultStreamId,
			Name:     "Test Producer Stream",
		})

		if streamErr != nil {
			panic(streamErr)
		}

		fmt.Printf("Created stream with ID: %d.\n", DefaultStreamId)
	}

	fmt.Printf("Stream with ID: %d exists.\n", DefaultStreamId)

	if _, topicErr := messageStream.GetTopicById(NewIdentifier(DefaultStreamId), NewIdentifier(TopicId)); topicErr != nil {
		topicErr = messageStream.CreateTopic(CreateTopicRequest{
			TopicId:         TopicId,
			Name:            "Test Topic From Producer Sample",
			PartitionsCount: 12,
			StreamId:        NewIdentifier(DefaultStreamId),
		})

		if topicErr != nil {
			panic(topicErr)
		}

		fmt.Printf("Created topic with ID: %d.\n", TopicId)
	}

	fmt.Printf("Topic with ID: %d exists.\n", TopicId)

	return nil
}

func ConsumeMessages(messageStream MessageStream) error {
	fmt.Printf("Messages will be polled from stream '%d', topic '%d', partition '%d' with interval %d ms.\n", DefaultStreamId, TopicId, Partition, Interval)

	for {
		messagesWrapper, err := messageStream.PollMessages(FetchMessagesRequest{
			Count:           1,
			StreamId:        NewIdentifier(DefaultStreamId),
			TopicId:         NewIdentifier(TopicId),
			Consumer:        Consumer{Kind: ConsumerSingle, Id: ConsumerId},
			PartitionId:     Partition,
			PollingStrategy: NextPollingStrategy(),
			AutoCommit:      true,
		})
		if err != nil {
			return err
		}
		if messagesWrapper == nil {
			panic("Something went wrong - this edge case handling will be refactored")
		}
		messages := messagesWrapper.Messages
		if len(messages) != 0 {
			for _, message := range messages {
				if err := HandleMessage(message); err != nil {
					fmt.Printf("Error when consuming message: %s\n", err.Error())
				}
			}
		} else {
			fmt.Println("Received 0 messages.")
		}

		time.Sleep(time.Duration(Interval) * time.Millisecond)
	}
}

func HandleMessage(messageResponse MessageResponse) error {
	length := (len(messageResponse.Payload) * 3) / 4
	bytes := make([]byte, length)

	str := string(messageResponse.Payload)
	isBase64 := false

	if _, err := base64.StdEncoding.Decode(bytes, []byte(str)); err == nil {
		isBase64 = true
	}

	var envelope sharedDemoContracts.Envelope

	if isBase64 {
		bytes, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			return err
		}

		jsonStr := string(bytes)
		if err = json.Unmarshal([]byte(jsonStr), &envelope); err != nil {
			return err
		}
	} else {
		if err := json.Unmarshal([]byte(messageResponse.Payload), &envelope); err != nil {
			return err
		}
	}

	fmt.Printf("Handling message type: %s at offset: %d with message Id: %s ", envelope.MessageType, messageResponse.Offset, messageResponse.Id)

	switch envelope.MessageType {
	case "order_created":
		var orderCreated sharedDemoContracts.OrderCreated
		if err := json.Unmarshal([]byte(envelope.Payload), &orderCreated); err != nil {
			return err
		}
		fmt.Printf("%+v\n", orderCreated)
	case "order_confirmed":
		var orderConfirmed sharedDemoContracts.OrderConfirmed
		if err := json.Unmarshal([]byte(envelope.Payload), &orderConfirmed); err != nil {
			return err
		}
		fmt.Printf("%+v\n", orderConfirmed)
	case "order_rejected":
		var orderRejected sharedDemoContracts.OrderRejected
		if err := json.Unmarshal([]byte(envelope.Payload), &orderRejected); err != nil {
			return err
		}
		fmt.Printf("%+v\n", orderRejected)
	}
	return nil
}
