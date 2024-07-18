package entities

type Switch struct {
	DevName  string   `json:"dev_name"`
	DevProps DevProps `json:"dev_prop"`
}

type DevProps struct {
	DevNames []string `json:"dev_names"`
}

type Value struct {
	Value byte `json:"value"`
}

func (sw Switch) toBytes() []byte {
	return make([]byte, 0)
}

func (sw Value) toBytes() []byte {
	return make([]byte, 0)
}
