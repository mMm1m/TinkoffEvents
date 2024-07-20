package entities

import (
	"TinkoffSmartHouse/errors"
	"TinkoffSmartHouse/util"
)

type Packet struct {
	Length  byte    `json:"length"`
	Payload Payload `json:"payload"`
	Crc8    byte    `json:"crc8"`
}

type Packets []Packet

func (pcs Packets) ToBytes() []byte {
	bytesArr := make([]byte, 0)
	for _, packet := range pcs {
		bytesArr = append(bytesArr, packet.toBytes()...)
	}
	return bytesArr
}

func PacketsFromBytes(data []byte) *Packets {
	size, offset := len(data), 0
	var pcts Packets
	for offset < size {
		curr, sz, _ := PacketFromBytes(data[offset:])
		offset += sz
		pcts = append(pcts, *curr)
	}
	return &pcts
}

func (pcs Packet) toBytes() []byte {
	bytesArr := make([]byte, 0)
	bytesArr = append(bytesArr, pcs.Length)
	bytesArr = append(bytesArr, pcs.Payload.ToBytes()...)
	bytesArr = append(bytesArr, pcs.Crc8)
	return bytesArr
}

// указатель на пакет / количество байт в пакете
func PacketFromBytes(data []byte) (*Packet, int, error) {
	if len(data) == 0 {
		return nil, 0, errors.EmptyArrayOfBytes()
	}
	if len(data) != (int(data[0]) + 2) {
		return nil, 0, errors.IncorrectSize()
	}
	length := data[0]
	payload := data[1 : length+1]
	crc8 := data[length+1]
	correctCrc8 := util.CRC8Simple(payload)
	if crc8 != correctCrc8 {
		return nil, 0, errors.IncorrectCRC8Code()
	}
	pld := payloadFromBytes(payload)
	pct := Packet{Length: length, Payload: *pld, Crc8: crc8}
	return &pct, int(length + 2), nil
}
