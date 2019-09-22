package rtda

import (
	"fmt"
	"rtda/heap"
)

type Frame struct {
	lower          *Frame
	localVariables LocalVariables
	operandStack   *OperandStack

	method *heap.Method

	currentThread *Thread
	nextPC        uint16 // 指令执行前 和执行后，都指向下一条要执行的指令
}

func (recv *Frame) OpStack() *OperandStack {
	return recv.operandStack
}

func (recv *Frame) LocalVars() LocalVariables {
	return recv.localVariables
}
func (recv *Frame) GetThread() *Thread {
	return recv.currentThread
}

func (recv *Frame) GetNextPC() uint16 {
	return recv.nextPC
}

func (recv *Frame) Test() {

	recv.testLocalVars()
	recv.testOpStack()
}

func (recv *Frame) GetMethod() *heap.Method {
	return recv.method
}

// invokedMethod
// 被调用的方法
func (recv *Frame) InvokeOtherMethod(invokedMethod *heap.Method) {
	currentThread := recv.currentThread
	newFrame := currentThread.CreateFrame(invokedMethod)

	// 数一下形参数目
	argSlotCount := int(invokedMethod.ArgSlotCount())

	// 参数传递
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			slot := recv.OpStack().PopSlot()
			newFrame.localVariables.SetSlot(uint16(i), slot)
		}
	}

}

func (recv *Frame) testLocalVars() {

	recv.localVariables.SetInt(0, 100)
	recv.localVariables.SetDouble(1, -123.789)
	recv.localVariables.SetLong(3, -10110000000)
	recv.localVariables.SetRef(5, nil)
	recv.localVariables.SetFloat(6, -1234.324)
	recv.localVariables.SetInt(7, 1024)

	fmt.Printf("%d\t %d\n", 0, recv.localVariables.GetInt(0))
	fmt.Printf("%d\t %f\n", 1, recv.localVariables.GetDouble(1))
	fmt.Printf("%d\t %d\n", 3, recv.localVariables.GetLong(3))
	fmt.Printf("%d\t %p\n", 5, recv.localVariables.GetRef(5))
	fmt.Printf("%d\t %f\n", 6, recv.localVariables.GetFloat(6))
	fmt.Printf("%d\t %d\n", 7, recv.localVariables.GetInt(7))

}

func (recv *Frame) testOpStack() {

	recv.operandStack.PushInt(123)
	recv.operandStack.PushInt(-1222)
	recv.operandStack.PushDouble(-1235.52323)
	recv.operandStack.PushLong(-77777777)
	recv.operandStack.PushFloat(-12324.5)
	recv.operandStack.PushRef(nil)
	recv.operandStack.PushFloat(2324.15)
	recv.operandStack.PushDouble(1235.52323)

	fmt.Printf(" %f\n", recv.operandStack.PopDouble())
	fmt.Printf(" %f\n", recv.operandStack.PopFloat())
	fmt.Printf(" %p\n", recv.operandStack.PopRef())
	fmt.Printf(" %f\n", recv.operandStack.PopFloat())
	fmt.Printf(" %d\n", recv.operandStack.PopLong())
	fmt.Printf(" %f\n", recv.operandStack.PopDouble())
	fmt.Printf(" %d\n", recv.operandStack.PopInt())
	fmt.Printf(" %d\n", recv.operandStack.PopInt())

}

func (recv *Frame) SetNextPC(pc uint16) {

	recv.nextPC = pc
}

func (recv *Frame) RevertNextPC() {
	recv.nextPC = recv.currentThread.pc
}
