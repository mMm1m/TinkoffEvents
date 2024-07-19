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
	devNameLen := byte(len(sw.DevName))
	stringBytes := []byte(sw.DevName)
	var byteArr []byte = []byte{devNameLen}
	byteArr = append(byteArr, stringBytes...)
	byteArr = append(byteArr, byte(len(sw.DevProps.DevNames)))
	for _, i := range sw.DevProps.DevNames {
		tmp := []byte{byte(len(i))}
		tmp = append(tmp, []byte(i)...)
		byteArr = append(byteArr, tmp...)
	}
	return byteArr
}

func (v Value) toBytes() []byte {
	return []byte{v.Value}
}
