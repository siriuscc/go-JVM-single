package main

import (
	"fmt"
	"syscall"
	"testing"
	"unsafe"
)

func TestLoad(t *testing.T) {

	s := "HelloWorld"
	i := 14

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Println(s)

	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

type Object struct {
	name string
	age  int
}

func TestHashCode(t *testing.T) {

	o := &Object{}

	println(unsafe.Sizeof(o))

}
