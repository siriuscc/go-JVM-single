package misc

import "rtda"

//public native long reallocateMemory(long var1, long var3);

// (J)J
//public native long allocateMemory(long size);
func allocateMemory(frame *rtda.Frame) {

	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	size := vars.GetLong(1)

	address := allocate(size)
	stack := frame.OpStack()
	stack.PushLong(address)
}

//public native void freeMemory(long address);
func freeMemory(frame *rtda.Frame) {

	locals := frame.LocalVars()
	address := locals.GetLong(1)
	free(address)
}

//public native long reallocateMemory(long address, long size);
func reallocateMemory(frame *rtda.Frame) {

	locals := frame.LocalVars()
	address := locals.GetLong(1)
	size := locals.GetLong(3)

	address = reallocate(address, size)

	frame.OpStack().PushLong(address)
}

//public native byte getByte(Object obj, long address);
func getByte(frame *rtda.Frame) {

	stack, mem := _getStackAndBytes(frame)
	stack.PushInt(int32(CutInt8(mem)))

}

// 往address 写入value
//public native void putLong(long address, long value);
func putLong(frame *rtda.Frame) {

	locals := frame.LocalVars()
	// vars.GetRef(0) // this
	address := locals.GetLong(1)
	value := locals.GetLong(3)

	mem := memoryAt(address)
	PutInt64(mem, value)
}

func _getStackAndBytes(frame *rtda.Frame) (*rtda.OperandStack, []byte) {
	vars := frame.LocalVars()
	// vars.GetRef(0) // this
	address := vars.GetLong(1)

	stack := frame.OpStack()
	mem := memoryAt(address)
	return stack, mem
}
