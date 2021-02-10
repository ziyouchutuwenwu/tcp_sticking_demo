package send

import (
	"encoding/binary"
	"fmt"
	"tcp_sticking_demo/demo/go_tcp_client/header"
)

func setDataLenthToHeader(buffer []byte, length int, pkgHeaderOption *header.PkgHeaderOption){
	if pkgHeaderOption.HeaderSize == 2{
		binary.BigEndian.PutUint16(buffer[0:], uint16(length))
	}
	if pkgHeaderOption.HeaderSize == 4{
		binary.BigEndian.PutUint32(buffer[0:], uint32(length))
	}
}

func MakeDataToSend(data []byte, pkgHeaderOption *header.PkgHeaderOption) []byte{
	dataLength := len(data)

	if ( dataLength > pkgHeaderOption.MaxDataSize) { return nil}
	buffer := make([]byte, pkgHeaderOption.HeaderFrameLenth + dataLength )
	setDataLenthToHeader(buffer, dataLength, pkgHeaderOption)

	copy(buffer[pkgHeaderOption.HeaderFrameLenth:], data)
	fmt.Println(buffer)

	return buffer
}