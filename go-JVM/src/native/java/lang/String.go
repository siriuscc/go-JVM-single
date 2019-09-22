package lang

import (
	"native"
	"rtda"
	"rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern",
		"()Ljava/lang/String;", intern)
}

//public native String intern();
func intern(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()

	interned := heap.InternString(this)

	frame.OpStack().PushRef(interned)
}
