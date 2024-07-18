package entities

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
	return []byte{0}
}

func (sr Sensor) toBytes() []byte {
	return []byte{0}
}
