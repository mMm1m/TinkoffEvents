package server

import (
	"TinkoffSmartHouse/entities"
)

func requestServer(commUrl, reqString string) ([]byte, int, error) {
	return make([]byte, 0), 0, nil
}

func setState(pcts *entities.Packets, database map[int]*entities.Device, devs []string,
	state byte, src int, serial *int) {

}

func pingSwitches(pcts *entities.Packets, database map[int]*entities.Device, src int, serial *int) {

}

func handleResponse(database map[int]*entities.Device, requestTimes map[int][]int, pcts,
	tasks *entities.Packets, src int, serial *int) {

}

func SimulateServer() {

}
