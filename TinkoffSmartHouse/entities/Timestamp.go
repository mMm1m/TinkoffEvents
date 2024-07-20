package entities

import (
	"TinkoffSmartHouse/constants"
	"TinkoffSmartHouse/util"
)

type Timestamp struct {
	Timestamp int `json:"timestamp"`
}

func (ts Timestamp) toBytes() []byte {
	return util.EncodeULEB128(ts.Timestamp)
}

func FindTime(packets *Packets) int {
	for _, packet := range *packets {
		if packet.Payload.DevType == constants.CLOCK && packet.Payload.Cmd == constants.TICK {
			return packet.Payload.CmdBody.(Timestamp).Timestamp
		}
	}
	return -1
}
