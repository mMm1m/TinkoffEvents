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
	switch {
	case (device == constants.HUB || device == constants.SOCKET || device == constants.LAMP || device == constants.CLOCK) && (cmd == constants.WHOISHERE || cmd == constants.IAMHERE):
		nameLen := cmdBodyBytes[0]
		name := cmdBodyBytes[1 : nameLen+1]
		return Name{string(name)}
	case (device == constants.SENSOR) && (cmd == constants.WHOISHERE || cmd == constants.IAMHERE):
		nameLength := cmdBodyBytes[0]
		name := cmdBodyBytes[1 : nameLength+1]
		sensors := cmdBodyBytes[nameLength+1]
		triggerLength := cmdBodyBytes[nameLength+2]
		triggers := make([]Trigger, triggerLength)
		curSkip := int(nameLength) + 3
		for i := 0; i < int(triggerLength); i++ {
			curOp := cmdBodyBytes[curSkip]
			curSkip++
			curVal, skipULEB := util.DecodeULEB128(cmdBodyBytes[curSkip:])
			curNameLen := cmdBodyBytes[curSkip+skipULEB]
			curName := string(cmdBodyBytes[curSkip+skipULEB+1 : curSkip+skipULEB+1+int(curNameLen)])
			curSkip += skipULEB + int(curNameLen) + 1
			curTrigger := Trigger{
				Op:    curOp,
				Value: curVal,
				Name:  curName,
			}
			triggers[i] = curTrigger
		}
		return Sensors{
			DevName: string(name),
			DevProps: EnvSensorsProperties{
				Sensors:  sensors,
				Triggers: triggers,
			},
		}
	case cmd == constants.GETSTATUS && (device == constants.SOCKET || device == constants.LAMP || device == constants.SENSOR || device == constants.SWITCH):
		return nil
	case cmd == constants.STATUS && device == constants.SENSOR:
		valuesSize := int(cmdBodyBytes[0])
		values := make([]int, valuesSize)
		curSkip := 1
		for i := 0; i < valuesSize; i++ {
			curVal, skipULEB := util.DecodeULEB128(cmdBodyBytes[curSkip:])
			values[i] = curVal
			curSkip += skipULEB
		}
		return Sensor{Values: values}
	case device == constants.SWITCH && (cmd == constants.WHOISHERE || cmd == constants.IAMHERE):
		nameLength := cmdBodyBytes[0]
		name := cmdBodyBytes[1 : nameLength+1]
		devNamesLen := int(cmdBodyBytes[nameLength+1])
		devNames := make([]string, devNamesLen)
		curSkip := nameLength + 2
		for i := 0; i < devNamesLen; i++ {
			curNameLen := cmdBodyBytes[curSkip]
			curName := cmdBodyBytes[curSkip+1 : curSkip+1+curNameLen]
			devNames[i] = string(curName)
			curSkip += curNameLen + 1
		}
		return Switch{
			DevName:  string(name),
			DevProps: DevProps{DevNames: devNames},
		}
	case device == constants.CLOCK && cmd == constants.TICK:
		timeTick, _ := util.DecodeULEB128(cmdBodyBytes)
		return Timestamp{Timestamp: timeTick}
	case ((cmd == constants.STATUS) && (device == constants.LAMP || device == constants.SOCKET || device == constants.SWITCH)) || ((cmd == constants.SETSTATUS) && (device == constants.LAMP || device == constants.SOCKET)):
		v := cmdBodyBytes[0]
		return Value{Value: v}
	}
	return nil
}

func (pl Payload) ToBytes() []byte {
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
