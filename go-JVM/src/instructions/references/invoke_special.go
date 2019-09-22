package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 用于调用一些需要特殊处理的实例方法，包括实例初始化方法，私有方法和父类方法
//Invoke instance method; special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

//hack
func (recv *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {

	currentClass := frame.GetMethod().GetOwner()
	cp := currentClass.GetConstantPool()
	methodRef := cp.GetConstant(uint(recv.Index)).(*heap.MethodRef)
	invokeMethod := methodRef.ResolveMethod()
	invokeClass := invokeMethod.GetOwner()

	count := invokeMethod.ArgSlotCount()
	ref := frame.OpStack().GetTopRef(count - 1)

	//校验
	if invokeMethod.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	// 不能调用私有方法
	if invokeMethod.GetAccessFlags().IsPrivate() && invokeClass != currentClass {
		panic("java.lang.NoSuchMethodError")
	}

	if invokeMethod.GetAccessFlags().IsProtected() &&
		invokeClass.GetPackageName() != currentClass.GetPackageName() {
		if invokeClass != currentClass && !currentClass.IsSubClassOf(invokeClass) {

			panic("java.lang.NoSuchMethodError")
		}

	}

	//// 如果是初始化方法，只要不是私有的
	//if invokeMethod.IsInitMethod() && invokeClass != currentClass {
	//
	//	if !currentClass.IsSubClassOf(invokeClass) {
	//		// 其他类不能调用本类的init方法
	//		panic("java.lang.NoSuchMethodError")
	//	}
	//
	//}

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

	realInvokeMethod := invokeMethod

	// 如果调用的是 超类中的函数，但不是构造函数，且当前的ACC_SUPER 标志被设置，需要一个额外的过程查找最终要调用的方法
	if currentClass.IsSuper() &&
		currentClass.IsSubClassOf(invokeClass) &&
		!invokeMethod.IsInitMethod() {

		realInvokeMethod = heap.LookupMethodInExtendsTree(currentClass.GetSuper(), methodRef.GetName(), methodRef.GetDescriptor())
	}

	if realInvokeMethod == nil || realInvokeMethod.GetAccessFlags().IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeAMethod(frame, realInvokeMethod)
}
