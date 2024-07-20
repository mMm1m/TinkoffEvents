package server

import (
	"TinkoffSmartHouse/constants"
	"TinkoffSmartHouse/entities"
	"TinkoffSmartHouse/util"
	"encoding/base64"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func requestServer(commUrl, reqString string) ([]byte, int, error) {
	client := &http.Client{}
	req := new(http.Request)
	var err error
	if reqString == "" {
		req, err = http.NewRequest(http.MethodPost, util.GetConnection(commUrl), nil)
	} else {
		req, err = http.NewRequest(http.MethodPost, util.GetConnection(commUrl), strings.NewReader(reqString))
	}
	if err != nil {
		return []byte{}, http.StatusBadRequest, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, http.StatusBadRequest, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	code := resp.StatusCode
	if err != nil {
		return []byte{}, http.StatusBadRequest, err
	}
	return body, code, nil
}

func setState(pcts *entities.Packets, database map[int]*entities.Device, devs []string,
	state byte, src int, serial *int) {
	for _, item := range database {
		name := item.DevName
		for _, dev := range devs {
			if name == dev {
				var cmdBody entities.ByteInterface = entities.Value{Value: state}
				newPacket := entities.Packet{
					Length: 0,
					Payload: entities.Payload{
						Src:     src,
						Dst:     item.Address,
						Serial:  *serial,
						DevType: item.DevType,
						Cmd:     constants.SETSTATUS,
						CmdBody: cmdBody,
					},
					Crc8: 0,
				}
				*serial++
				newPacket.Crc8 = util.CRC8Simple(newPacket.Payload.ToBytes())
				newPacket.Length = byte(len(newPacket.Payload.ToBytes()))
				*pcts = append(*pcts, newPacket)
				break
			}
		}
	}
}

func pingSwitches(pcts *entities.Packets, database map[int]*entities.Device, src int, serial *int) {
	for _, dev := range database {
		if dev.DevType == constants.SWITCH && dev.IsPresent {
			newPacket := entities.Packet{
				Length: 0,
				Payload: entities.Payload{
					Src:     src,
					Dst:     dev.Address,
					Serial:  *serial,
					DevType: constants.HUB,
					Cmd:     constants.GETSTATUS,
					CmdBody: nil,
				},
				Crc8: 0,
			}
			*serial++
			newPacket.Crc8 = util.CRC8Simple(newPacket.Payload.ToBytes())
			newPacket.Length = byte(len(newPacket.Payload.ToBytes()))
			*pcts = append(*pcts, newPacket)
		}
	}
}

func handleResponse(database map[int]*entities.Device, requestTimes map[int][]int, pcts,
	tasks *entities.Packets, src int, serial *int) {
	answerTime := entities.FindTime(pcts)
	for _, pct := range *pcts {
		val, ok := database[pct.Payload.Src]
		if ok && !val.IsPresent && pct.Payload.Cmd != constants.WHOISHERE {
			continue
		}
		switch pct.Payload.Cmd {
		case constants.IAMHERE:
			dvt := pct.Payload.DevType
			adr := pct.Payload.Src
			isAlive := answerTime-requestTimes[constants.TO_ALL][0] <= 300
			switch {
			case dvt == constants.SWITCH:
				body := pct.Payload.CmdBody.(entities.Switch)
				database[adr] = &entities.Device{
					Address:      adr,
					DevName:      body.DevName,
					DevType:      pct.Payload.DevType,
					IsOn:         false,
					IsPresent:    isAlive,
					ConnDevs:     body.DevProps.DevNames,
					SensorValues: nil,
					Sensors:      0,
					Triggers:     nil,
					AnswerTime:   answerTime,
				}
			case dvt == constants.SENSOR:
				body := pct.Payload.CmdBody.(entities.Sensors)
				database[adr] = &entities.Device{
					Address:      adr,
					DevName:      body.DevName,
					DevType:      pct.Payload.DevType,
					IsOn:         false,
					IsPresent:    isAlive,
					ConnDevs:     nil,
					SensorValues: nil,
					Sensors:      body.DevProps.Sensors,
					Triggers:     body.DevProps.Triggers,
					AnswerTime:   answerTime,
				}
			default:
				body := pct.Payload.CmdBody.(entities.Name)
				database[adr] = &entities.Device{
					Address:      adr,
					DevName:      body.DevName,
					DevType:      pct.Payload.DevType,
					IsOn:         false,
					IsPresent:    isAlive,
					ConnDevs:     nil,
					SensorValues: nil,
					Sensors:      0,
					Triggers:     nil,
					AnswerTime:   answerTime,
				}
			}
		case constants.WHOISHERE:
			var cmdBody entities.ByteInterface = entities.Name{DevName: constants.HUB_NAME}
			newPacket := entities.Packet{
				Length: 0,
				Payload: entities.Payload{
					Src:     src,
					Dst:     constants.TO_ALL,
					Serial:  *serial,
					DevType: constants.HUB,
					Cmd:     constants.IAMHERE,
					CmdBody: cmdBody,
				},
				Crc8: 0,
			}
			*serial++
			newPacket.Crc8 = util.CRC8Simple(newPacket.Payload.ToBytes())
			newPacket.Length = byte(len(newPacket.Payload.ToBytes()))
			*tasks = append(*tasks, newPacket)

			dvt := pct.Payload.DevType
			adr := pct.Payload.Src
			switch {
			case dvt == constants.SWITCH:
				body := pct.Payload.CmdBody.(entities.Switch)
				database[adr] = &entities.Device{
					Address:      adr,
					DevName:      body.DevName,
					DevType:      pct.Payload.DevType,
					IsOn:         false,
					IsPresent:    true,
					ConnDevs:     body.DevProps.DevNames,
					SensorValues: nil,
					Sensors:      0,
					Triggers:     nil,
					AnswerTime:   answerTime,
				}
			case dvt == constants.SENSOR:
				body := pct.Payload.CmdBody.(entities.Sensors)
				database[adr] = &entities.Device{
					Address:      adr,
					DevName:      body.DevName,
					DevType:      pct.Payload.DevType,
					IsOn:         false,
					IsPresent:    true,
					ConnDevs:     nil,
					SensorValues: nil,
					Sensors:      body.DevProps.Sensors,
					Triggers:     body.DevProps.Triggers,
					AnswerTime:   answerTime,
				}
			default:
				body := pct.Payload.CmdBody.(entities.Name)
				database[adr] = &entities.Device{
					Address:      adr,
					DevName:      body.DevName,
					DevType:      pct.Payload.DevType,
					IsOn:         false,
					IsPresent:    true,
					ConnDevs:     nil,
					SensorValues: nil,
					Sensors:      0,
					Triggers:     nil,
					AnswerTime:   answerTime,
				}
			}
		case constants.STATUS:
			if pct.Payload.Src != constants.TO_ALL {
				if len(requestTimes[pct.Payload.Src]) >= 2 {
					requestTimes[pct.Payload.Src] = requestTimes[pct.Payload.Src][1:]
				} else {
					delete(requestTimes, pct.Payload.Src)
				}
			}
			switch pct.Payload.DevType {
			case constants.LAMP:
				cbv := pct.Payload.CmdBody.(entities.Value)
				if cbv.Value == 1 {
					database[pct.Payload.Src].IsOn = true
				} else {
					database[pct.Payload.Src].IsOn = false
				}
			case constants.SOCKET:
				cbv := pct.Payload.CmdBody.(entities.Value)
				if cbv.Value == 1 {
					database[pct.Payload.Src].IsOn = true
				} else {
					database[pct.Payload.Src].IsOn = false
				}
			case constants.SWITCH:
				cbv := pct.Payload.CmdBody.(entities.Value)
				switch cbv.Value {
				case 1:
					database[pct.Payload.Src].IsOn = true
					devNames2TurnOn := database[pct.Payload.Src].ConnDevs
					setState(tasks, database, devNames2TurnOn, 1, src, serial)
				default:
					database[pct.Payload.Src].IsOn = false
					devNames2TurnOff := database[pct.Payload.Src].ConnDevs
					setState(tasks, database, devNames2TurnOff, 0, src, serial)
				}
			case constants.SENSOR:
				values := pct.Payload.CmdBody.(entities.Sensor).Values
				database[pct.Payload.Src].SensorValues = values
				valuesAll := [4]int{-1, -1, -1, -1}
				envSensor := database[pct.Payload.Src]
				sensorTypeMask := envSensor.Sensors
				idx := 0
				for i := 0; i < 4; i++ {
					if sensorTypeMask&1 == 1 {
						valuesAll[i] = values[idx]
						idx++
					}
					sensorTypeMask = sensorTypeMask >> 1
				}
				triggers := envSensor.Triggers
				for _, trigger := range triggers {
					thresh := trigger.Value
					device := trigger.Name
					opBits := trigger.Op

					state := opBits & 1
					opBits = opBits >> 1
					greaterThen := opBits & 1
					opBits = opBits >> 1
					sensorType := opBits

					if greaterThen == 1 {
						if valuesAll[sensorType] > thresh {
							setState(tasks, database, []string{device}, state, src, serial)
						}
					} else {
						if valuesAll[sensorType] < thresh && valuesAll[sensorType] != -1 {
							setState(tasks, database, []string{device}, state, src, serial)
						}
					}
				}
			}
		}
	}
}

func SimulateServer() {
	args := os.Args[1:]
	if len(args) < 2 {
		os.Exit(99)
	}
	commUrl := args[0]
	hubAddress, err := strconv.ParseInt(args[1], 16, 64)
	if err != nil {
		os.Exit(99)
	}

	database := make(map[int]*entities.Device)
	requestTimes := make(map[int][]int)

	serialCounter := 1
	var statusCode int
	var reqStr string
	var respRawBytes, respRawBytesTrimmed, respBytes []byte

	pendingTasks := entities.Packets{}
	var hubTime int

	for {
		var cbn entities.ByteInterface = entities.Name{DevName: constants.HUB_NAME}
		pcts := entities.Packets{
			entities.Packet{
				Length: 0,
				Payload: entities.Payload{
					Src:     int(hubAddress),
					Dst:     constants.TO_ALL,
					Serial:  serialCounter,
					DevType: constants.HUB,
					Cmd:     constants.WHOISHERE,
					CmdBody: cbn,
				},
				Crc8: 0,
			},
		}
		serialCounter++
		pcts[0].Length = byte(len(pcts[0].Payload.ToBytes()))
		pcts[0].Crc8 = util.CRC8Simple(pcts[0].Payload.ToBytes())
		reqStr = base64.RawURLEncoding.EncodeToString(pcts.ToBytes())
		respRawBytes, statusCode, err = requestServer(commUrl, reqStr)
		if err != nil {
			os.Exit(99)
		}

		if statusCode == http.StatusOK {
			respRawBytesTrimmed = []byte(util.RemoveSpaces(string(respRawBytes)))
			respBytes, err = base64.RawURLEncoding.DecodeString(string(respRawBytesTrimmed))
			if err != nil {
				continue
			}
			respPcts := entities.PacketsFromBytes(respBytes)
			hubTime = entities.FindTime(respPcts)
			requestTimes[constants.TO_ALL] = []int{hubTime}
			handleResponse(database, requestTimes, respPcts, &pendingTasks, int(hubAddress), &serialCounter)
			for _, dev := range database {
				dev.IsPresent = true
			}
			break
		} else if statusCode == http.StatusNoContent {
			os.Exit(0)
		} else {
			os.Exit(99)
		}
	}

	for _, dev := range database {
		if dev.DevType == constants.SENSOR {
			getStatusReq := entities.Packet{
				Length: 0,
				Payload: entities.Payload{
					Src:     int(hubAddress),
					Dst:     dev.Address,
					Serial:  serialCounter,
					DevType: constants.HUB,
					Cmd:     constants.GETSTATUS,
					CmdBody: nil,
				},
				Crc8: 0,
			}
			serialCounter++
			getStatusReq.Length = byte(len(getStatusReq.Payload.ToBytes()))
			getStatusReq.Crc8 = util.CRC8Simple(getStatusReq.Payload.ToBytes())
			pendingTasks = append(pendingTasks, getStatusReq)
		}
	}

	for statusCode == http.StatusOK {
		pingSwitches(&pendingTasks, database, int(hubAddress), &serialCounter)
		for _, pct := range pendingTasks {
			curCmd := pct.Payload.Cmd
			curDst := pct.Payload.Dst
			if curCmd == constants.GETSTATUS || curCmd == constants.SETSTATUS {
				requestTimes[curDst] = append(requestTimes[curDst], hubTime)
			}
		}
		reqStr = base64.RawURLEncoding.EncodeToString(pendingTasks.ToBytes())
		pendingTasks = entities.Packets{}

		respRawBytes, statusCode, err = requestServer(commUrl, reqStr)

		if err != nil {
			os.Exit(99)
		}
		respRawBytesTrimmed = []byte(util.RemoveSpaces(string(respRawBytes)))
		respBytes, err = base64.RawURLEncoding.DecodeString(string(respRawBytesTrimmed))
		if err != nil {
			continue
		}

		respPcts := entities.PacketsFromBytes(respBytes)
		hubTime = entities.FindTime(respPcts)

		for address, timeQueue := range requestTimes {
			if hubTime-timeQueue[0] > 300 {
				if _, ok := database[address]; ok {
					database[address].IsPresent = false
				}
				delete(requestTimes, address)
			}
		}
		handleResponse(database, requestTimes, respPcts, &pendingTasks, int(hubAddress), &serialCounter)
	}

	if statusCode == http.StatusNoContent {
		os.Exit(0)
	}
	os.Exit(99)
}
