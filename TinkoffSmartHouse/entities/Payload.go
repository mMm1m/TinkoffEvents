package entities

import "TinkoffSmartHouse/constants"

type Payload struct {
	Src     int                `json:"src"`
	Dst     int                `json:"dst"`
	Serial  int                `json:"serial"`
	DevType constants.DEV_TYPE `json:"dev_type"`
	Cmd     constants.CMD      `json:"cmd"`
	CmdBody ByteInterface      `json:"cmd_body,omitempty"`
}
