package packet

import (
	"io"

	"github.com/zishang520/engine.io-go-parser/types"
)

type (
	Type    string
	Options struct {
		Compress          bool                  `json:"compress" mapstructure:"compress" msgpack:"compress"`
		WsPreEncoded      types.BufferInterface `json:"wsPreEncoded,omitempty" mapstructure:"wsPreEncoded,omitempty" msgpack:"wsPreEncoded,omitempty"`
		WsPreEncodedFrame types.BufferInterface `json:"wsPreEncodedFrame,omitempty" mapstructure:"wsPreEncodedFrame,omitempty" msgpack:"wsPreEncodedFrame,omitempty"`
	}
	Packet struct {
		Type    Type      `json:"type" mapstructure:"type" msgpack:"type"`
		Data    io.Reader `json:"data,omitempty" mapstructure:"data,omitempty" msgpack:"data,omitempty"`
		Options *Options  `json:"options,omitempty" mapstructure:"options,omitempty" msgpack:"options,omitempty"`

		// Deprecated: this method will be removed in the next major release, please use [Options.WsPreEncoded] instead.
		WsPreEncoded types.BufferInterface
	}
)

// Packet types.
const (
	Open    Type = "open"
	Close   Type = "close"
	Ping    Type = "ping"
	Pong    Type = "pong"
	Message Type = "message"
	Upgrade Type = "upgrade"
	Noop    Type = "noop"
	Error   Type = "error"
)
