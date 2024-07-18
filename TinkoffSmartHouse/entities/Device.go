package entities

import "TinkoffSmartHouse/constants"

type Device struct {
	Address      int                `json:"address"`
	DevName      string             `json:"dev_name"`
	DevType      constants.DEV_TYPE `json:"dev_type"`
	IsOn         bool               `json:"status"`
	IsPresent    bool               `json:"is_present"`
	ConnDevs     []string           `json:"conn_devs"`
	SensorValues []int              `json:"sensor_values"`
	Sensors      byte               `json:"sensors"`
	Triggers     []Trigger          `json:"triggers"`
	AnswerTime   int                `json:"answer_time"`
}
