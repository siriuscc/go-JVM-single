package io

import (
	"native"
	"rtda"
)

func init() {
	native.Register("java/io/FileInputStream",
		"initIDs", "()V", initIDs)
}

//private final FileDescriptor fd;
var fis_id int

// 获取 fd对象在 FileInputStream 中的id，类似于实例槽。
//private static native void initIDs();
func initIDs(frame *rtda.Frame) {

	inStreamClass := frame.GetMethod().GetOwner()
	field := inStreamClass.GetInstanceField("fd", "Ljava/io/FileDescriptor;")
	fis_id = int(field.GetSlotId())
}
