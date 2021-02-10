package header

type PkgHeaderOption struct {
	HeaderSize 			int
	MaxDataSize 		int
	HeaderFrameLenth 	int
}

func GetPkgOptionWithHeaderSize(headerSize int) *PkgHeaderOption {
	pkgHeaderOption := PkgHeaderOption{}

	if headerSize == 0 {
		pkgHeaderOption.HeaderSize = 2
	}
	if headerSize !=2 && headerSize !=4 { return nil }
	pkgHeaderOption.HeaderSize = headerSize

	switch pkgHeaderOption.HeaderSize {
	case 2:
		pkgHeaderOption.MaxDataSize = 0xFFFF
		pkgHeaderOption.HeaderFrameLenth = 2
	case 4:
		pkgHeaderOption.MaxDataSize = 0x7FFFFFFF
		pkgHeaderOption.HeaderFrameLenth = 4
	}

	return &pkgHeaderOption
}