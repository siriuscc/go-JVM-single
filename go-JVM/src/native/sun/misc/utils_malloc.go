package misc

import (
	"encoding/binary"
	"math"
)

// _allocated 维护一个分段 内存映射，并不是所有位置都有值，每一段内存的start位置，在_allocated[index]=mem
// _nextAddress 表示下一个可用位置

var _allocated = map[int64][]byte{} // int64,[]byte
var _nextAddress = int64(64)        // not zero!

// 分配一段size长度的内存
func allocate(size int64) int64 {

	mem := make([]byte, size)

	address := _nextAddress
	_allocated[address] = mem
	_nextAddress += size
	return address
}

// 对 [address,address+size] 重新分配
func reallocate(address, size int64) int64 {
	if size == 0 {
		return 0
	} else if address == 0 {
		return allocate(size)
	} else {
		mem := memoryAt(address)
		if len(mem) >= int(size) {
			return address
		} else {
			delete(_allocated, address)
			newAddress := allocate(size)
			newMem := memoryAt(newAddress)
			copy(newMem, mem)
			return newAddress
		}
	}
}

// 释放 address 对应的内存块
func free(address int64) {
	if _, ok := _allocated[address]; ok {
		delete(_allocated, address)
	} else {
		panic("memory was not allocated!")
	}
}

// 得到  中间的任意一段内存，[address,endAddress]
func memoryAt(address int64) []byte {
	for startAddress, mem := range _allocated {
		endAddress := startAddress + int64(len(mem))
		if address >= startAddress && address < endAddress {
			offset := address - startAddress
			return mem[offset:]
		}
	}
	panic("invalid address!")
	return nil
}

var _bigEndian = binary.BigEndian

func PutInt8(s []byte, val int8) {
	s[0] = uint8(val)
}
func CutInt8(s []byte) int8 {
	return int8(s[0])
}

func PutUint16(s []byte, val uint16) {
	_bigEndian.PutUint16(s, val)
}
func Uint16(s []byte) uint16 {
	return _bigEndian.Uint16(s)
}

func PutInt16(s []byte, val int16) {
	_bigEndian.PutUint16(s, uint16(val))
}
func Int16(s []byte) int16 {
	return int16(_bigEndian.Uint16(s))
}

func PutInt32(s []byte, val int32) {
	_bigEndian.PutUint32(s, uint32(val))
}
func Int32(s []byte) int32 {
	return int32(_bigEndian.Uint32(s))
}

func PutInt64(s []byte, val int64) {
	_bigEndian.PutUint64(s, uint64(val))
}
func Int64(s []byte) int64 {
	return int64(_bigEndian.Uint64(s))
}

func PutFloat32(s []byte, val float32) {
	_bigEndian.PutUint32(s, math.Float32bits(val))
}
func Float32(s []byte) float32 {
	return math.Float32frombits(_bigEndian.Uint32(s))
}

func PutFloat64(s []byte, val float64) {
	_bigEndian.PutUint64(s, math.Float64bits(val))
}
func Float64(s []byte) float64 {
	return math.Float64frombits(_bigEndian.Uint64(s))
}
