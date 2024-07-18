package entities

type Name struct {
	DevName string `json:"dev_name"`
}

func (name Name) toBytes() []byte {
	return []byte(name.DevName)
}
