package entities

type Timestamp struct {
	Timestamp int `json:"timestamp"`
}

func (ts Timestamp) toBytes() []byte {
	return make([]byte, 0)
}

func findTime(packets *Packets) int {
	return 0
}
