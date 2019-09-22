package base

type ByteCodeReader struct {
	code []byte
	pc   uint16
}

func (recv *ByteCodeReader) GetPC() uint16 {
	return recv.pc
}

func (recv *ByteCodeReader) Reset(code []byte, pc uint16) {
	recv.code = code
	recv.pc = pc
}

func (recv *ByteCodeReader) ReadUint8() uint8 {

	i := recv.code[recv.pc]
	recv.pc++
	return i
}

func (recv *ByteCodeReader) ReadInt8() int8 {

	i := recv.code[recv.pc]
	recv.pc++
	return int8(i)
}

func (recv *ByteCodeReader) ReadUint16() uint16 {
	// 先高后低
	high := uint16(recv.ReadUint8())
	low := uint16(recv.ReadUint8())
	return high<<8 | low
}

func (recv *ByteCodeReader) ReadInt16() int16 {
	return int16(recv.ReadUint16())
}

func (recv *ByteCodeReader) ReadUint32() uint32 {
	// 先高后低
	high := uint32(recv.ReadUint16())
	low := uint32(recv.ReadUint16())
	return high<<16 | low
}

func (recv *ByteCodeReader) ReadInt32() int32 {

	return int32(recv.ReadUint32())
}

func (recv *ByteCodeReader) ReadInt32s(count uint) []int32 {

	data := make([]int32, count)
	for i := range data {
		data[i] = recv.ReadInt32()
	}
	return data
}

// 为了保证defaultOffset的地址是4的倍数
func (recv *ByteCodeReader) SkipPadding() {

	for recv.pc%4 != 0 {
		recv.ReadUint8()
	}
}
