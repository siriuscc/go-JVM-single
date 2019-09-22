package classfile

import (
	"fmt"
	"math"
	"unicode/utf16"
)

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (recv *ConstantStringInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.stringIndex = reader.readUint16()
}

func (recv *ConstantStringInfo) GetString() string {

	return recv.cp.GetUtf8(recv.stringIndex)
}

func (recv *ConstantStringInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t#%d", "String", recv.stringIndex)

}

/////////////////////////////////////////////////////////////////////

type ConstantIntegerInfo struct {
	val int32
}

func (recv *ConstantIntegerInfo) read(reader *ClassReader) {
	recv.val = int32(reader.readUint32())
}

func (recv *ConstantIntegerInfo) GetValue() int32 {
	return recv.val
}

func (recv *ConstantIntegerInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t%d", "Integer", recv.val)
}

/////////////////////////////////////////////////////////////////////

type ConstantFloatInfo struct {
	val float32
}

func (recv *ConstantFloatInfo) read(reader *ClassReader) {
	recv.val = math.Float32frombits(reader.readUint32())
}

func (recv *ConstantFloatInfo) GetValue() float32 {
	return recv.val
}

func (recv *ConstantFloatInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t%f", "Float", recv.val)
}

/////////////////////////////////////////////////////////////////////

type ConstantLongInfo struct {
	val int64
}

func (recv *ConstantLongInfo) read(reader *ClassReader) {
	recv.val = int64(reader.readUint64())
}

func (recv *ConstantLongInfo) GetValue() int64 {
	return recv.val
}

func (recv *ConstantLongInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t%s", "Long", "string")
}

type ConstantDoubleInfo struct {
	val float64
}

func (recv *ConstantDoubleInfo) read(reader *ClassReader) {
	recv.val = math.Float64frombits(reader.readUint64())
}

func (recv *ConstantDoubleInfo) GetValue() float64 {
	return recv.val
}

func (recv *ConstantDoubleInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t%f", "Double", recv.val)
}

type ConstantUtf8Info struct {
	val string
}

func (recv *ConstantUtf8Info) read(reader *ClassReader) {
	length := reader.readUint16()
	bytes := reader.readBytes(uint32(length))
	recv.val = convertMUTF8(bytes)
}

func convertMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}

func (recv *ConstantUtf8Info) ToString() string {
	return fmt.Sprintf("\t%s:\t\t%s", "Utf8", recv.val)
}
