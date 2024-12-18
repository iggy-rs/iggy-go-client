package iggcon

import (
	"context"
)

type IggyMessageCompression string

const (
	MESSAGE_COMPRESSION_NONE IggyMessageCompression = "none"
	MESSAGE_COMPRESSION_S2   IggyMessageCompression = "s2"
	// MESSAGE_COMPRESSION_ZSTD IggyMessageCompression = "zstd"
)

type IggyConfiguration struct {
	context.Context
	BaseAddress        string                 `json:"baseAddress"`
	Protocol           Protocol               `json:"protocol"`
	MessageCompression IggyMessageCompression `json:"compression"`
}

type Protocol string

const (
	Http Protocol = "Http"
	Tcp  Protocol = "Tcp"
	Quic Protocol = "Quic"
)
