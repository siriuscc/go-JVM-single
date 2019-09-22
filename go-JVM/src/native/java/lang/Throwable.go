package lang

import (
	"fmt"
	"instructions/base"
	"logger"
	"native"
	"rtda"
	"rtda/heap"
)

func init() {

	class := "java/lang/Throwable"
	native.Register(class, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
	native.Register(class, "getStackTraceDepth", "()I", getStackTraceDepth)
	native.Register(class, "getStackTraceElement", "(I)Ljava/lang/StackTraceElement;", getStackTraceElement)

}

//native StackTraceElement getStackTraceElement(int index);
// (I)Ljava/lang/StackTraceElement;
func getStackTraceElement(frame *rtda.Frame) {

	locals := frame.LocalVars()
	this := locals.GetThis()
	index := locals.GetInt(1)

	const initDesc = "(Ljava/lang/String;Ljava/lang/String;Ljava/lang/String;I)V"

	loader := frame.GetMethod().GetOwner().GetClassLoader()

	stackTraceElementKlass := loader.LoadClass("java/lang/StackTraceElement")
	element := stackTraceElementKlass.CreateObject()
	frame.OpStack().PushRef(element)

	initMethod := stackTraceElementKlass.GetInitMethod(initDesc)
	traces := this.GetExt().(StackTraces)

	//  执行初始化函数
	driverFrame := frame.GetThread().CreateDriverFrame(5)
	opStack := driverFrame.OpStack()

	opStack.PushRef(element)
	opStack.PushRef(heap.JString(loader, traces[index].className))
	opStack.PushRef(heap.JString(loader, traces[index].methodName))
	opStack.PushRef(heap.JString(loader, traces[index].fileName))
	opStack.PushInt(int32(traces[index].lineNumber))

	base.InvokeAMethod(driverFrame, initMethod)

}

//    native int getStackTraceDepth();
func getStackTraceDepth(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()
	traces := this.GetExt().(StackTraces)
	frame.OpStack().PushInt(int32(traces.GetDepth()))
}

// 栈信息
type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

type StackTraces []*StackTraceElement

func (recv StackTraces) GetDepth() int {

	return len(recv)
}

func (recv *StackTraceElement) IsJVMTrace() bool {
	if recv.className[0] == '<' {
		return true
	}
	return false
}

func (recv *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)", recv.className, recv.methodName, recv.fileName, recv.lineNumber)
}

//     private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()
	frame.OpStack().PushRef(this)
	stes := createStackTraceElements(this, frame.GetThread())
	this.SetExt(stes)
}

func createStackTraceElements(this *heap.OopDesc, thread *rtda.Thread) StackTraces {

	// 需要跳过两帧，因为 栈顶正在执行  fillInStackTrace(),fillInStackTrace(int)
	// 下面还有 n帧，在执行异常类的构造函数，所以要计算一下 继承树的层级
	skip := this.GetKlass().GetExtendsLayerCount() + 2

	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, 0)

	for _, frame := range frames {
		// 跳过 内部的帧
		if logger.DEBUG || frame.GetMethod().GetOwner().GetName()[0] != '<' {

			stes = append(stes, createStackTraceElement(frame))
		}
	}
	return stes
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {

	method := frame.GetMethod()
	class := method.GetOwner()

	return &StackTraceElement{
		fileName:   class.GetSourceFileName(),
		className:  class.GetName(),
		methodName: method.GetName(),
		lineNumber: method.GetLineNumber(frame.GetNextPC() - 1),
	}
}
