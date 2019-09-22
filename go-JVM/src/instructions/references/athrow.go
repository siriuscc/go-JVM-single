package references

import (
	"fmt"
	"instructions/base"
	"logger"
	"native/java/lang"
	"rtda"
	"rtda/heap"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (recv *ATHROW) Execute(frame *rtda.Frame) {

	e := frame.OpStack().PopRef()

	if e == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.GetThread()

	if !findAndGotoExceptionHandler(thread, e) {
		handleUncaughtException(thread, e)
	}

}

// 没有处理代码
func handleUncaughtException(thread *rtda.Thread, ex *heap.OopDesc) {

	thread.ClearStack()
	// 获取属性
	detailMessage := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	msg := heap.GoString(detailMessage)

	println(fmt.Sprintf("Exception in thread \"%s\" %s: %s", thread.GetName(), ex.GetKlass().GetJavaName(), msg))


	traces := ex.GetExt().(lang.StackTraces)

	for i := 0; i < len(traces); i++ {

		if logger.DEBUG || !traces[i].IsJVMTrace() {
			println("\t at", traces[i].String())
		}
	}
}

func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.OopDesc) bool {

	for {
		currentFrame := thread.GetCurrentFrame()
		pc := currentFrame.GetNextPC() - 1
		handlerPC := currentFrame.GetMethod().FindExceptionHandler(ex.GetKlass(), pc)
		if handlerPC > 0 {
			opStack := currentFrame.OpStack()
			opStack.Clear()

			opStack.PushRef(ex)
			currentFrame.SetNextPC(uint16(handlerPC))
			return true
		}
		thread.PopFrame()
		if thread.HasNoFrame() {
			break
		}
	}
	return false
}
