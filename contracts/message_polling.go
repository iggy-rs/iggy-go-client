package iggcon

type PollingStrategy struct {
	Kind  MessagePolling
	Value uint64
}

type MessagePolling int

const (
	Offset    MessagePolling = 1
	Timestamp MessagePolling = 2
	First     MessagePolling = 3
	Last      MessagePolling = 4
	Next      MessagePolling = 5
)

func NewPollingStrategy(kind MessagePolling, value uint64) PollingStrategy {
	return PollingStrategy{
		Kind:  kind,
		Value: value,
	}
}

func OffsetPollingStrategy(value uint64) PollingStrategy {
	return NewPollingStrategy(Offset, value)
}

func TimestampPollingStrategy(value uint64) PollingStrategy {
	return NewPollingStrategy(Timestamp, value)
}

func FirstPollingStrategy() PollingStrategy {
	return NewPollingStrategy(First, 0)
}

func LastPollingStrategy() PollingStrategy {
	return NewPollingStrategy(Last, 0)
}

func NextPollingStrategy() PollingStrategy {
	return NewPollingStrategy(Next, 0)
}
