package heap

import "math"

// 存放变量

type Slot struct {
	num int32
	ref *OopDesc
}

type Slots []Slot

func createSlots(size uint) Slots {

	if size < 1 {
		return nil
	}
	return make([]Slot, size)
}

func (recv Slots) SetInt(index uint, val int32) {
	recv[index].num = val
}

func (recv Slots) GetInt(index uint) int32 {
	return recv[index].num
}

func (recv Slots) SetFloat(index uint, val float32) {

	recv[index].num = int32(math.Float32bits(val))
}

func (recv Slots) GetFloat(index uint) float32 {

	return math.Float32frombits(uint32(recv[index].num))
}

func (recv Slots) SetLong(index uint, val int64) {

	recv[index].num = int32(val)
	recv[index+1].num = int32(val >> 32)

}

func (recv Slots) GetLong(index uint) int64 {

	low := int64(uint32(recv[index].num)) // 一定要先转回去uint32,否则会出现 符号位变了的情况
	high := int64(uint32(recv[index+1].num)) << 32
	return high | low
}

func (recv Slots) SetRef(index uint, ref *OopDesc) {

	recv[index].ref = ref
}

func (recv Slots) GetRef(index uint) *OopDesc {
	return recv[index].ref
}

func (recv Slots) SetDouble(index uint, val float64) {

	longVal := math.Float64bits(val)
	recv.SetLong(index, int64(longVal))
}

func (recv Slots) GetDouble(index uint) float64 {

	longVal := recv.GetLong(index)
	return math.Float64frombits(uint64(longVal))
}
