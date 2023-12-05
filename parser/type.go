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
		packet.OPEN:    '0',
		packet.CLOSE:   '1',
		packet.PING:    '2',
		packet.PONG:    '3',
		packet.MESSAGE: '4',
		packet.UPGRADE: '5',
		packet.NOOP:    '6',
	}

	PacketTypesReverse map[byte]packet.Type = map[byte]packet.Type{
		'0': packet.OPEN,
		'1': packet.CLOSE,
		'2': packet.PING,
		'3': packet.PONG,
		'4': packet.MESSAGE,
		'5': packet.UPGRADE,
		'6': packet.NOOP,
	}

	// Premade error packet.
	ErrorPacket = &packet.Packet{Type: packet.ERROR, Data: types.NewStringBufferString(`parser error`)}
)
