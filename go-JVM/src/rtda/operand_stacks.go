package rtda

import (
	"math"
	"rtda/heap"
)

type OperandStack struct {
	maxStackSize int32 // max size
	size         int32 // 当前真实大小，use size
	slots        []Slot
}

func CreateOperandStack(maxStackSize uint16) *OperandStack {

	if maxStackSize <= 0 {
		return nil
	}
	return &OperandStack{slots: make([]Slot, maxStackSize), size: 0, maxStackSize: int32(maxStackSize)}
}

func (recv *OperandStack) PushSlot(slot Slot) {

	if recv.size >= recv.maxStackSize {
		panic("java.lang.OperandStack:OutOfSize")
	}
	recv.slots[recv.size] = slot
	recv.size++
}

func (recv *OperandStack) PopSlot() Slot {

	if recv.size < 1 {
		panic("java.lang.OperandStack:Empty")
	}
	recv.size--
	return recv.slots[recv.size]
}

func (recv *OperandStack) PushInt(val int32) {

	if recv.size >= recv.maxStackSize {
		panic("java.lang.OperandStack:OutOfSize")
	}
	recv.slots[recv.size].num = val
	recv.size++
}

func (recv *OperandStack) PopInt() int32 {

	if recv.size < 1 {
		panic("java.lang.OperandStack:Empty")
	}
	recv.size--
	return recv.slots[recv.size].num
}

func (recv *OperandStack) PopInt2() (int32, int32) {

	v2 := recv.PopInt()
	v1 := recv.PopInt()

	return v2, v1
}

func (recv *OperandStack) PushFloat(val float32) {

	if recv.size >= recv.maxStackSize {
		panic("java.lang.OperandStack:OutOfSize")
	}
	recv.slots[recv.size].num = int32(math.Float32bits(val))
	recv.size++
}
func (recv *OperandStack) PopFloat() float32 {
	if recv.size < 1 {
		panic("java.lang.OperandStack:Empty")
	}
	recv.size--
	return math.Float32frombits(uint32(recv.slots[recv.size].num))
}

func (recv *OperandStack) PopFloat2() (float32, float32) {

	v2 := recv.PopFloat()
	v1 := recv.PopFloat()

	return v2, v1
}

func (recv *OperandStack) PushLong(val int64) {

	if recv.size+1 >= recv.maxStackSize {
		panic("java.lang.OperandStack:OutOfSize")
	}

	recv.slots[recv.size].num = int32(uint32(val))         // 低位 截取
	recv.slots[recv.size+1].num = int32(uint64(val) >> 32) // 把符号位也存进去
	recv.size += 2
}

func (recv *OperandStack) PopLong() int64 {
	if recv.size < 2 {
		panic("java.lang.OperandStack:Empty,A Long need two Slots")
	}
	recv.size -= 2
	low := uint64(uint32(recv.slots[recv.size].num))
	height := uint64(uint32(recv.slots[recv.size+1].num)) << 32

	return int64(height | low)
}

func (recv *OperandStack) PopLong2() (int64, int64) {

	v2 := recv.PopLong()
	v1 := recv.PopLong()

	return v2, v1
}

func (recv *OperandStack) PushDouble(val float64) {
	longVal := int64(math.Float64bits(val))
	recv.PushLong(longVal)
}

func (recv *OperandStack) PopDouble() float64 {
	longVal := recv.PopLong()
	return math.Float64frombits(uint64(longVal))
}

func (recv *OperandStack) PopDouble2() (float64, float64) {

	v2 := recv.PopDouble()
	v1 := recv.PopDouble()
	return v2, v1
}

func (recv *OperandStack) PushRef(ref *heap.OopDesc) {

	if recv.size >= recv.maxStackSize {
		panic("java.lang.OperandStack:OutOfSize")
	}

	recv.slots[recv.size].ref = ref
	recv.size++
}

func (recv *OperandStack) PopRef() *heap.OopDesc {

	if int(recv.size) < 1 {
		panic("java.lang.OperandStack:Empty")
	}
	recv.size--
	return recv.slots[recv.size].ref
}

func (recv *OperandStack) TopRef() *heap.OopDesc {

	if int(recv.size) < 1 {
		panic("java.lang.OperandStackEmpty")
	}
	return recv.slots[recv.size-1].ref
}

// top0 等价于 getTop
//
func (recv *OperandStack) GetTopRef(index uint) *heap.OopDesc {

	if recv.size < int32(index)+1 {
		panic("java.lang.OperandStackEmpty")
	}
	return recv.slots[recv.size-int32(index)-1].ref
}

func (recv *OperandStack) PopRef2() (*heap.OopDesc, *heap.OopDesc) {

	v2 := recv.PopRef()
	v1 := recv.PopRef()

	return v2, v1
}

func (recv *OperandStack) Clear() {
	recv.size = 0

	for i := range recv.slots {
		recv.slots[i].ref = nil
		recv.slots[i].num = 0
	}
}

func (recv *OperandStack) PushBool(value bool) {

	if value {
		recv.PushInt(1)
	} else {
		recv.PushInt(0)
	}
}
