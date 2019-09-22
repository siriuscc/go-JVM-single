package reserved

import (
	"fmt"
	"instructions/base"
	"native"
	_ "native/java/io"
	_ "native/java/lang"
	_ "native/java/security"
	_ "native/java/util/concurrent/atomic"
	_ "native/sun/io"
	_ "native/sun/misc"
	_ "native/sun/reflect"
	"rtda"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (recv *INVOKE_NATIVE) Execute(frame *rtda.Frame) {

	method := frame.GetMethod()
	klass := method.GetOwner()

	nativeMethod := native.FindNativeMethod(klass.GetName(), method.GetName(), method.GetDescriptor())

	if nativeMethod == nil {
		//省略错误处理部分代码
		panic(fmt.Sprintf("java.lang.UnsatisfiedLinkError:%s.%s:%s"+klass.GetName(), method.GetName(), method.GetDescriptor()))
	}

	nativeMethod(frame)
}
