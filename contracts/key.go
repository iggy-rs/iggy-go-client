package iggcon

type Key struct {
	KeyKind Keykind `json:"keyKind"`
	Value   int     `json:"value"`
}

type Keykind int

const (
	PartitionId Keykind = iota
	EntityId
)
