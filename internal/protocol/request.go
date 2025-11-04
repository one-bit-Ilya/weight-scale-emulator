package protocol

type RequestHeader struct {
	FirstByte  byte
	SecondByte byte
	ThirdByte  byte
}
type RequestPayload struct {
	PayloadLength uint16
	CmdComandCode CmdCode
}

type Crc struct {
	CrcSum uint16
}
type Request struct {
	Header  RequestHeader
	Payload RequestPayload
	Crc     Crc
}
