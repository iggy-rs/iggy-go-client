package iggcon

type ConsumerKind int

const (
	ConsumerSingle ConsumerKind = 1
	ConsumerGroup  ConsumerKind = 2
)

type Consumer struct {
	Kind ConsumerKind
	ID   uint32
}

type IdKind int

const (
	NumericId IdKind = 1
	StringId  IdKind = 2
)

type Identifier struct {
	Kind   IdKind
	Length int
	Value  string
}
