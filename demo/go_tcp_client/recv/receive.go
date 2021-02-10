package recv

import (
	"tcp_sticking_demo/demo/go_tcp_client/header"
	"encoding/binary"
	"fmt"
	"net"
)

var savedData []byte

func init(){
	savedData = make([]byte, 0)
	savedData = []byte{0,6}
}

func LoopRead(connection net.Conn, pkgHeaderOption *header.PkgHeaderOption)  {
	readBufferSize := 8
	for{
		readBuffer := make([]byte, readBufferSize)
		readLen, err := connection.Read(readBuffer)

		if ( err != nil) {
			fmt.Println("read error ", err)
			return
		}

		buffer := readBuffer[:readLen]
		fmt.Println(buffer)

		totalData := append(savedData, buffer...)
		DealWithData(totalData, pkgHeaderOption)
	}
}

func DealWithData(totalData []byte, pkgHeaderOption *header.PkgHeaderOption)  {
	for{
		totalDataLen := len(totalData)
		if ( totalDataLen <= pkgHeaderOption.HeaderSize) {
			savedData = totalData[:len(totalData):len(totalData)]
			break
		};

		if ( totalDataLen > pkgHeaderOption.HeaderSize ){
			dataLen := GetDataLenthFromHeader(totalData, pkgHeaderOption)

			if ( totalDataLen < pkgHeaderOption.HeaderSize + dataLen) {
				savedData = totalData[:len(totalData):len(totalData)]
				break
			}
			if ( totalDataLen >= pkgHeaderOption.HeaderSize + dataLen) {

				frameLen := pkgHeaderOption.HeaderSize + dataLen

				pkg := totalData[pkgHeaderOption.HeaderSize : frameLen]
				fmt.Println(pkg)

				lastPosition := len(totalData)
				capacity := len(totalData)

				// 去掉一个数据帧的数据
				totalData = totalData[frameLen : lastPosition : capacity]
				savedData = totalData[:]
			}
		}
	}
}

func GetDataLenthFromHeader(buffer []byte, pkgHeaderOption *header.PkgHeaderOption) int{
	dataLenth := 0

	if pkgHeaderOption.HeaderSize == 2{
		dataLenth = int(binary.BigEndian.Uint16(buffer[0:]))
	}
	if pkgHeaderOption.HeaderSize == 4{
		dataLenth = int(binary.BigEndian.Uint32(buffer[0:]))
	}

	return dataLenth
}