package references

import (
	"fmt"
	"instructions/base"
	"rtda"
	"rtda/heap"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

// Invoke instance method; dispatch based on class
// 用于调用对象的实例方法，根据对象的实际类型进行分派。
func (recv *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {

	currentClass := frame.GetMethod().GetOwner()
	cp := currentClass.GetConstantPool()
	methodRef := cp.GetConstant(uint(recv.Index)).(*heap.MethodRef)
	invokeMethod := methodRef.ResolveMethod()
	invokeClass := invokeMethod.GetOwner()

	count := invokeMethod.ArgSlotCount()

	//校验
	//if invokeMethod.IsInitMethod() && invokeMethod.GetOwner() != currentClass {
	//	// 其他类不能调用本类的init方法
	//	panic("java.lang.NoSuchMethodError")
	//}
	if invokeMethod.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	this := frame.OpStack().GetTopRef(count - 1)

	//// debug hack
	//if methodRef.GetName() == "append" {
	//	fmt.Println(this)
	//}

	if this == nil {
		//if methodRef.GetName() == "println" {
		//
		//	hackPrintln(frame, methodRef)
		//	return
		//}
		panic("java.lang.NullPointerException")
	}

	// protected method(this,...) 方法只能被声明该方法的类或子类调用。
	//	 protected 可以是子类访问
	if invokeMethod.GetAccessFlags().IsProtected() {
		// protected 方法是包作用域，必须包名相同
		if currentClass.GetPackageName() != invokeClass.GetPackageName() {
			// 如果包名不同，必须是子类访问
			if currentClass != invokeClass && !currentClass.IsSubClassOf(invokeClass) {
				panic("java.lang.IllegalAccessError")
			}
		} //else{ //packagename 相同，可以访问	}
	}

	realInvokeMethod := heap.LookupMethodInExtendsTree(this.GetKlass(), methodRef.GetName(), methodRef.GetDescriptor())

	if realInvokeMethod == nil || realInvokeMethod.GetAccessFlags().IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeAMethod(frame, realInvokeMethod)

}

func hackPrintln(frame *rtda.Frame, methodRef *heap.MethodRef) {

	stack := frame.OpStack()
	switch methodRef.GetDescriptor() {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(B)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(I)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		jStr := stack.PopRef()
		goStr := heap.GoString(jStr)
		fmt.Println(goStr)
	default:
		panic("println:" + methodRef.GetDescriptor())
	}
	stack.PopRef()

}
