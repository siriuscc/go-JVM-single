package rtda

import "rtda/heap"

// 驱动方法，就是空方法

func (recv *Thread) CreateDriverFrame(maxStack int) *Frame {

	frame := &Frame{
		lower:         nil,
		method:        heap.DriverMethod,
		operandStack:  CreateOperandStack(uint16(maxStack)),
		nextPC:        0,
		currentThread: recv}
	recv.PushFrame(frame)
	return frame
}

func (recv *Thread) CreateThrowFrame(maxStackSize uint16) *Frame {
	frame := &Frame{
		lower:         nil,
		method:        heap.AThrowMethod,
		operandStack:  CreateOperandStack(maxStackSize),
		nextPC:        0,
		currentThread: recv}
	recv.PushFrame(frame)
	return frame
}
