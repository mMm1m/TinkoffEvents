package entities

type Name struct {
	DevName string `json:"dev_name"`
}

func (name Name) toBytes() []byte {
	byteArr := []byte{byte(len(name.DevName))}
	return append(byteArr, []byte(name.DevName)...)
}
