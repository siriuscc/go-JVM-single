package base

import (
	"logger"
	"rtda"
	"rtda/heap"
	"strings"
)

func InvokeAMethod(callerFrame *rtda.Frame, invokeMethod *heap.Method) {

	invokeFrame := callerFrame.GetThread().CreateFrame(invokeMethod)

	// copy params
	locals := invokeFrame.LocalVars()
	opStack := callerFrame.OpStack()
	slotCount := int(invokeMethod.ArgSlotCount())
	for i := slotCount - 1; i >= 0; i-- {
		locals.SetSlot(uint16(i), opStack.PopSlot())
	}

	logger.Printf("%s + invoke:%s: %s,%s:%s\n", strings.Repeat("    ", int(callerFrame.GetThread().GetThreadFrameCount())), invokeMethod.GetOwner().GetName(), invokeMethod.GetName(), invokeMethod.GetDescriptor(), locals)


	if invokeMethod.GetAccessFlags().IsNative() {
		if invokeMethod.GetName() == "registerNatives" {
			callerFrame.GetThread().PopFrame()
		}
	}

}
