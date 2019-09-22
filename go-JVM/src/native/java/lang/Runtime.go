package lang

import (
	"native"
	"rtda"
	"runtime"
)

func init() {
	const jlRuntime = "java/lang/Runtime"
	native.Register(jlRuntime, "availableProcessors", "()I", availableProcessors)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *rtda.Frame) {
	numCPU := runtime.NumCPU()

	stack := frame.OpStack()
	stack.PushInt(int32(numCPU))
}
