package control

import (
	"instructions/base"
	"rtda"
)

type RETURN struct{ base.NoOperandsInstruction }
type IRETURN struct{ base.NoOperandsInstruction }
type ARETURN struct{ base.NoOperandsInstruction }
type DRETURN struct{ base.NoOperandsInstruction }
type FRETURN struct{ base.NoOperandsInstruction }
type LRETURN struct{ base.NoOperandsInstruction }

func (recv *RETURN) Execute(frame *rtda.Frame) {
	frame.GetThread().PopFrame()
}

func (recv *IRETURN) Execute(frame *rtda.Frame) {

	currentThread := frame.GetThread()
	currentFrame := currentThread.PopFrame()
	invokeFrame := currentThread.GetCurrentFrame()
	retVal := currentFrame.OpStack().PopInt()
	invokeFrame.OpStack().PushInt(retVal)
}

func (recv *LRETURN) Execute(frame *rtda.Frame) {

	currentThread := frame.GetThread()
	currentFrame := currentThread.PopFrame()
	invokeFrame := currentThread.GetCurrentFrame()
	retVal := currentFrame.OpStack().PopLong()
	invokeFrame.OpStack().PushLong(retVal)
}

func (recv *FRETURN) Execute(frame *rtda.Frame) {

	currentThread := frame.GetThread()
	currentFrame := currentThread.PopFrame()
	invokeFrame := currentThread.GetCurrentFrame()
	retVal := currentFrame.OpStack().PopFloat()
	invokeFrame.OpStack().PushFloat(retVal)
}

func (recv *DRETURN) Execute(frame *rtda.Frame) {

	currentThread := frame.GetThread()
	currentFrame := currentThread.PopFrame()
	invokeFrame := currentThread.GetCurrentFrame()
	retVal := currentFrame.OpStack().PopDouble()
	invokeFrame.OpStack().PushDouble(retVal)
}

func (recv *ARETURN) Execute(frame *rtda.Frame) {

	currentThread := frame.GetThread()
	currentFrame := currentThread.PopFrame()
	invokeFrame := currentThread.GetCurrentFrame()
	retVal := currentFrame.OpStack().PopRef()
	invokeFrame.OpStack().PushRef(retVal)

	//if retVal != nil {
	//	logger.Printf(" ~areturn :type:%s,metadata:%s", retVal.GetKlass().GetName(), retVal.GetMetaData().GetName())
	//}
}
