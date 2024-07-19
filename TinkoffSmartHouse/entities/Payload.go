package entities

import (
	"TinkoffSmartHouse/constants"
	"TinkoffSmartHouse/util"
)

type Payload struct {
	Src     int                `json:"src"`
	Dst     int                `json:"dst"`
	Serial  int                `json:"serial"`
	DevType constants.DEV_TYPE `json:"dev_type"`
	Cmd     constants.CMD      `json:"cmd"`
	CmdBody ByteInterface      `json:"cmd_body,omitempty"`
}

func parseCmdBody(device constants.DEV_TYPE, cmd constants.CMD, cmdBodyBytes []byte) ByteInterface {
	return &Packet{}
}

func (pl Payload) toBytes() []byte {
	byteArr := make([]byte, 0)
	byteArr = append(byteArr, util.EncodeULEB128(pl.Src)...)
	byteArr = append(byteArr, util.EncodeULEB128(pl.Dst)...)
	byteArr = append(byteArr, util.EncodeULEB128(pl.Serial)...)
	byteArr = append(byteArr, []byte{byte(pl.DevType), byte(pl.Cmd)}...)
	if pl.CmdBody != nil {
		byteArr = append(byteArr, (pl.CmdBody).toBytes()...)
	}
	return byteArr
}

func payloadFromBytes(data []byte) *Payload {
	src, skip1 := util.DecodeULEB128(data)
	dst, skip2 := util.DecodeULEB128(data[skip1:])
	serial, skip3 := util.DecodeULEB128(data[(skip1 + skip2):])

	skip := skip1 + skip2 + skip3
	pl := Payload{
		Src:     src,
		Dst:     dst,
		Serial:  serial,
		DevType: constants.DEV_TYPE(data[skip]),
		Cmd:     constants.CMD(data[skip+1]),
		CmdBody: nil,
	}
	cmdBodyBytes := data[(skip + 2):]
	pl.CmdBody = parseCmdBody(pl.DevType, pl.Cmd, cmdBodyBytes)
	return &pl
}
