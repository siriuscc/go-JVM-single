package io

import (
	"native"
	"rtda"
)

func init() {

	const class = "java/io/FileDescriptor"

	native.Register(class, "initIDs", "()V", native_FileDescriptor_initIDs)
	native.Register(class, "set", "(I)J", native_FileDescriptor_set)

}

// 初始化 field.slotID
// This routine initializes JNI field offsets for the class
func native_FileDescriptor_initIDs(frame *rtda.Frame) {

	//inStreamClass := frame.GetMethod().GetOwner()
	//field := inStreamClass.GetInstanceField("fd", "Ljava/io/FileDescriptor;")
	//fis_id= int(field.GetSlotId())
}

// (I)J
//private static native long set(int d);
func native_FileDescriptor_set(frame *rtda.Frame) {

	frame.OpStack().PushLong(0)
}
