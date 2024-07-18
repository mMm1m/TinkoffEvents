package util

import "TinkoffSmartHouse/entities"

func encodeULEB128(int) []byte {
	return make([]byte, 0)
}

// ans/ bytes_parsed
func decodeULEB128([]byte) (int, int) {
	return 0, 0
}

func CRC8Simple([]byte) byte {
	return 0
}

func getConnection(url string) string {
	return ""
}

func removeSpaces(string) string {
	return ""
}

func findTime(packets *entities.Packets) int {
	return 0
}
