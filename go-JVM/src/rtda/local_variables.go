package rtda

import (
	"math"
	"rtda/heap"
)

type LocalVariables []Slot

func createLocalVariables(maxLocals uint16) LocalVariables {

	if maxLocals < 1 {
		return nil
	}
	return make([]Slot, maxLocals)
}

func (recv LocalVariables) GetBoolean(index uint16) bool {

	return recv[index].num != 0
}

func (recv LocalVariables) SetInt(index uint16, val int32) {
	recv[index].num = val
}

func (recv LocalVariables) GetInt(index uint16) int32 {
	return recv[index].num
}

func (recv LocalVariables) SetFloat(index uint16, val float32) {

	recv[index].num = int32(math.Float32bits(val))
}

func (recv LocalVariables) GetFloat(index uint16) float32 {

	return math.Float32frombits(uint32(recv[index].num))
}

func (recv LocalVariables) SetLong(index uint16, val int64) {

	recv[index].num = int32(val)
	recv[index+1].num = int32(val >> 32)

}

func (recv LocalVariables) GetLong(index uint16) int64 {

	low := int64(uint32(recv[index].num)) // 一定要先转回去uint32,否则会出现 符号位变了的情况
	high := int64(uint32(recv[index+1].num)) << 32
	return high | low
}

func (recv LocalVariables) SetRef(index uint16, ref *heap.OopDesc) {

	recv[index].ref = ref
}

func (recv LocalVariables) GetRef(index uint16) *heap.OopDesc {
	return recv[index].ref
}

func (recv LocalVariables) SetDouble(index uint16, val float64) {

	longVal := math.Float64bits(val)
	recv.SetLong(index, int64(longVal))
}

func (recv LocalVariables) GetDouble(index uint16) float64 {

	longVal := recv.GetLong(index)
	return math.Float64frombits(uint64(longVal))
}

func (recv LocalVariables) SetSlot(index uint16, slot Slot) {

	recv[index] = slot
}

func (recv LocalVariables) GetThis() *heap.OopDesc {

	return recv.GetRef(0)
}

func (recv LocalVariables) ToString() string {

	msg := ""
	for _, s := range recv {
		if s.ref == nil {
			msg += string(s.num) + ","
		} else {
			msg += s.ToString() + ","
		}
	}
	return msg
}
