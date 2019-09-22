package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 调用静态方法
type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (recv *INVOKE_STATIC) Execute(frame *rtda.Frame) {

	// 运行时常量池
	cp := frame.GetMethod().GetOwner().GetConstantPool()
	methodRef := cp.GetConstant(uint(recv.Index)).(*heap.MethodRef)

	resolvedMethod := methodRef.ResolveMethod()

	if !resolvedMethod.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	klass := resolvedMethod.GetOwner()
	// 调用之前，如果 还没有初始化，因为指令已经读入，但是参数没有，
	// 这时，thread.pc指向指令， 把frame pc回退到thread.pc，然后push clinit frame
	if !klass.IsInitialized() {
		frame.RevertNextPC()
		base.InitClass(frame.GetThread(), klass)
		return
	}

	base.InvokeAMethod(frame, resolvedMethod)
}
