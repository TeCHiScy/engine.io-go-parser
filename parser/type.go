package parser

import (
	"github.com/zishang520/engine.io-go-parser/packet"
	"github.com/zishang520/engine.io-go-parser/types"
)

type Parser interface {
	Protocol() int
	EncodePacket(*packet.Packet, bool, ...bool) (types.BufferInterface, error)
	DecodePacket(types.BufferInterface, ...bool) (*packet.Packet, error)
	EncodePayload([]*packet.Packet, ...bool) (types.BufferInterface, error)
	DecodePayload(types.BufferInterface) ([]*packet.Packet, error)
}

const SEPARATOR byte = 0x1E

// Packet types.
var (
	PacketTypes map[packet.Type]byte = map[packet.Type]byte{
		packet.Open:    '0',
		packet.Close:   '1',
		packet.Ping:    '2',
		packet.Pong:    '3',
		packet.Message: '4',
		packet.Upgrade: '5',
		packet.Noop:    '6',
	}

	PacketTypesReverse map[byte]packet.Type = map[byte]packet.Type{
		'0': packet.Open,
		'1': packet.Close,
		'2': packet.Ping,
		'3': packet.Pong,
		'4': packet.Message,
		'5': packet.Upgrade,
		'6': packet.Noop,
	}

	// Premade error packet.
	ErrorPacket = &packet.Packet{Type: packet.Error, Data: types.NewStringBufferString(`parser error`)}
)
