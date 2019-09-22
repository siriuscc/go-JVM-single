package io

import (
	"native"
	"os"
	"rtda"
	"unsafe"
)

func init() {
	class := "java/io/FileOutputStream"
	native.Register(class, "initIDs", "()V", native_FileOutputStream_initIDs)
	native.Register(class, "writeBytes", "([BIIZ)V", native_FileOutputStream_writeBytes)
}

func native_FileOutputStream_initIDs(frame *rtda.Frame) {

	//inStreamClass := frame.GetMethod().GetOwner()
	//field := inStreamClass.GetInstanceField("fd", "Ljava/io/FileDescriptor;")
	//fis_id= int(field.GetSlotId())
}

//private native void writeBytes(byte b[], int off, int len, boolean append)  throws IOException;
func native_FileOutputStream_writeBytes(frame *rtda.Frame) {

	locals := frame.LocalVars()

	b := locals.GetRef(1)
	off := locals.GetInt(2)
	len := locals.GetInt(3)
	//append:=locals.GetBoolean(4)

	bytes := b.GetByteTable()
	goBytes := convertInt8ToUint8s(bytes)

	goBytes = goBytes[off : off+len]
	os.Stdout.Write(goBytes)
}

func convertInt8ToUint8s(int8s []int8) []byte {

	ptr := unsafe.Pointer(&int8s)
	goBytes := *(*[]byte)(ptr)
	return goBytes
}
