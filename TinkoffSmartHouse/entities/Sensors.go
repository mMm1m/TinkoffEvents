package entities

import "TinkoffSmartHouse/util"

type Trigger struct {
	Op    byte   `json:"op"`
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type Sensors struct {
	DevName  string               `json:"dev_name"`
	DevProps EnvSensorsProperties `json:"dev_props"`
}

type EnvSensorsProperties struct {
	Sensors  byte      `json:"sensors"`
	Triggers []Trigger `json:"trigger"`
}

type Sensor struct {
	Values []int `json:"values"`
}

func (srs Sensors) toBytes() []byte {
	devNameLen := byte(len(srs.DevName))
	devName := []byte(srs.DevName)
	sensors := srs.DevProps.Sensors
	triggerLen := byte(len(srs.DevProps.Triggers))

	byteArr := []byte{devNameLen}
	byteArr = append(byteArr, devName...)
	byteArr = append(byteArr, sensors)
	byteArr = append(byteArr, triggerLen)
	for i := 0; i < int(triggerLen); i++ {
		tmp := []byte{srs.DevProps.Triggers[i].Op}
		tmp = append(tmp, util.EncodeULEB128(srs.DevProps.Triggers[i].Value)...)
		tmp = append(tmp, byte(len(srs.DevProps.Triggers[i].Name)))
		tmp = append(tmp, []byte(srs.DevProps.Triggers[i].Name)...)
		byteArr = append(byteArr, tmp...)
	}
	return byteArr
}

func (sr Sensor) toBytes() []byte {
	len := byte(len(sr.Values))
	byteArr := []byte{len}
	for _, val := range sr.Values {
		byteArr = append(byteArr, util.EncodeULEB128(val)...)
	}
	return byteArr
}
