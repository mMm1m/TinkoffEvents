package constants

type CMD byte
type DEV_TYPE byte

const (
	CRC8Error     = 555
	ZeroSize      = 550
	IncorrectSize = 556
)

const (
	HUB    DEV_TYPE = 0x01
	SENSOR DEV_TYPE = 0x02
	SWITCH DEV_TYPE = 0x03
	LAMP   DEV_TYPE = 0x04
	SOCKET DEV_TYPE = 0x05
	CLOCK  DEV_TYPE = 0x06
)

const (
	WHOISHERE CMD = 0x01
	IAMHERE   CMD = 0x02
	GETSTATUS CMD = 0x03
	STATUS    CMD = 0x04
	SETSTATUS CMD = 0x05
	TICK      CMD = 0x06
)

const (
	TO_ALL   int    = 0x3FFF
	HUB_NAME string = "HUB01"
	HOST     string = "localhost"
	PORT     string = "9998"
	TYPE     string = "http"
)
