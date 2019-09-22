package atomic

import (
	"native"
	"rtda"
)

func init() {
	const class = "java/util/concurrent/atomic/AtomicLong"
	native.Register(class, "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *rtda.Frame) {
	frame.OpStack().PushBool(false)
}
