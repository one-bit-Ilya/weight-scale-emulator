package protocol

const (
	FirstHeaderByte  byte = 0xF8
	SecondHeaderByte byte = 0x55
	ThirdHeaderByte  byte = 0xCE
)

type CmdCode byte

const (
	CmdGetWeightRequest  CmdCode = 0xA0
	CmdGetWeightResponse CmdCode = 0x10
)
