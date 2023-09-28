package iggcon

type ClientResponse struct {
	ID                  uint                `json:"id"`
	Address             string              `json:"address"`
	UserID              uint                `json:"userId"`
	Transport           string              `json:"transport"`
	ConsumerGroupsCount int                 `json:"consumerGroupsCount"`
	ConsumerGroups      []ConsumerGroupInfo `json:"consumerGroups"`
}
