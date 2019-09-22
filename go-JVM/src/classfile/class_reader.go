package classfile

import (
	"encoding/binary"
	"fmt"
)

// 可读接口
type Readable interface {
	read(reader *ClassReader)
}

// 只负责 字节层面的解析
type ClassReader struct {
	data []byte
	cp   ConstantPool
}

func (recv *ClassReader) readUint8() uint8 {

	data := recv.data[0]
	recv.data = recv.data[1:]
	return data
}

func (recv *ClassReader) readUint16() uint16 {

	data := binary.BigEndian.Uint16(recv.data)
	recv.data = recv.data[2:]
	return data
}

func (recv *ClassReader) readUint32() uint32 {
	data := binary.BigEndian.Uint32(recv.data)
	recv.data = recv.data[4:]
	return data
}

func (recv *ClassReader) readUint64() uint64 {
	data := binary.BigEndian.Uint64(recv.data)
	recv.data = recv.data[8:]
	return data
}

// 数组的长度为开头的u16指出，每一个都是u16
func (recv *ClassReader) readUint16s() []uint16 {
	n := recv.readUint16()

	data := make([]uint16, n)

	for i := range data {
		data[i] = recv.readUint16()
	}

	return data
}

// 读取n个字节
func (recv *ClassReader) readBytes(n uint32) []byte {

	data := recv.data[:n]
	recv.data = recv.data[n:]
	return data
}

func (recv *ClassReader) ToString() string {

	return fmt.Sprintf("ClassReader:data len:%d, cp size:%d", len(recv.data), len(recv.cp))
}
