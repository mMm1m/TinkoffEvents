package util

import (
	"TinkoffSmartHouse/constants"
	"fmt"
	"strings"
	"unicode"
)

func EncodeULEB128(v int) []byte {
	var byteArr []byte
	for {
		byte_ := byte(v & 0x7F)
		v >>= 7
		if v != 0 {
			v |= 0x80
		}
		byteArr = append(byteArr, byte_)
		if v == 0 {
			break
		}
	}
	return byteArr
}

// ans/ bytes_parsed
func DecodeULEB128(bytes []byte) (int, int) {
	res := 0
	shift := 0
	bytesParsed := 0
	for _, bt := range bytes {
		bytesParsed++
		res |= (int(bt) & 0x7f) << shift
		shift += 7
		if bt&0x80 == 0 {
			break
		}
	}
	return res, bytesParsed
}

func CRC8Simple(bytes []byte) byte {
	const mask byte = 0x1D
	crc8 := byte(0)
	for _, curByte := range bytes {
		crc8 ^= curByte
		for i := 0; i < 8; i++ {
			if (crc8 & 0x80) != 0 {
				crc8 = (crc8 << 1) ^ mask
			} else {
				crc8 <<= 1
			}
		}
	}
	return crc8
}

func GetConnection(url string) string {
	if url == "" {
		return fmt.Sprintf("%v://%v:%v", constants.TYPE, constants.HOST, constants.PORT)
	}
	return url
}

func RemoveSpaces(str string) string {
	b := strings.Builder{}
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
