package protocol

type ResponseHeader struct {
	FirstByte  byte
	SecondByte byte
	ThirdByte  byte
}

type WeightResponsePayload struct {
	PayloadLength uint16
	CmdComandCode CmdCode
	Weight        [4]byte
	Division      byte
	Stable        byte
}

type WeightResponse struct {
	Header  ResponseHeader
	Payload WeightResponsePayload
	Crc     Crc
}
