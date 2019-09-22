package io

import (
	"native"
	"rtda"
)

func init() {

	class := "java/io/WinNTFileSystem"
	native.Register(class, "initIDs", "()V", native_WinNTFileSystem_initIDs)
}

func native_WinNTFileSystem_initIDs(frame *rtda.Frame) {

	//inStreamClass := frame.GetMethod().GetOwner()
	//field := inStreamClass.GetInstanceField("fd", "Ljava/io/FileDescriptor;")
	//fis_id= int(field.GetSlotId())
}
