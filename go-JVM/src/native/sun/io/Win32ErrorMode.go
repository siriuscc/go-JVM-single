package io

import (
	"native"
	"rtda"
)

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtda.Frame) {
	// todo
	frame.OpStack().PushLong(0)
}
