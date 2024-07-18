package entities

type Packet struct {
	Length  byte    `json:"length"`
	Payload Payload `json:"payload"`
	Crc8    byte    `json:"crc8"`
}

type Packets []Packet

func (pcs Packets) toBytes() []byte {
	return make([]byte, len(pcs)*2)
}

func packetsFromBytes(data []byte) *Packets {
	return &Packets{}
}

func (pcs Packet) toBytes() []byte {
	return pcs.toBytes()
}

func packetFromBytes(data []byte) *Packet {
	return &Packet{}
}
