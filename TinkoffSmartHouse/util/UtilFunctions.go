package util

import (
	"TinkoffSmartHouse/constants"
	"fmt"
)

func EncodeULEB128(int) []byte {
	return make([]byte, 0)
}

// ans/ bytes_parsed
func DecodeULEB128([]byte) (int, int) {
	return 0, 0
}

func CRC8Simple([]byte) byte {
	return 0
}

func getConnection(url string) string {
	if url == "" {
		return fmt.Sprintf("%v://%v:%v", constants.TYPE, constants.HOST, constants.PORT)
	}
	return url
}

func removeSpaces(string) string {
	return ""
}
