package rtda

import (
	"logger"
	"rtda/heap"
)

type Thread struct {
	name      string
	pc        uint16 // 指令执行前指向要执行的代码，执行后下一回合开始前，指向已执行的代码
	stack     *ThreadStack
	threadOop *heap.OopDesc
}

func (recv *Thread) GetPC() uint16 {
	return recv.pc
}

func (recv *Thread) SetPC(pc uint16) {
	recv.pc = pc
}

func (recv *Thread) PushFrame(frame *Frame) {
	recv.stack.push(frame)
}

func (recv *Thread) PopFrame() *Frame {
	return recv.stack.pop()
}

func (recv *Thread) GetCurrentFrame() *Frame {
	return recv.stack.getTop()
}

// 没有frame可用
func (recv *Thread) HasNoFrame() bool {
	return recv.stack.isEmpty()
}

func (recv *Thread) GetThreadFrameCount() uint {

	return recv.stack.getSize()
}

func CreateThread() *Thread {
	return &Thread{
		pc:    0,
		name:  "main",
		stack: CreateThreadStack(1024),
	}
}

func (recv *Thread) CreateFrame(method *heap.Method) *Frame {

	frame := &Frame{
		lower:          nil,
		method:         method,
		localVariables: make([]Slot, method.GetMaxLocals()),
		operandStack:   CreateOperandStack(uint16(method.GetMaxStack())),
		nextPC:         0,
		currentThread:  recv}

	recv.PushFrame(frame)
	return frame
}

func (recv *Thread) LogFrames() {

	for !recv.HasNoFrame() {

		frame := recv.PopFrame()
		method := frame.method

		className := method.GetOwner().GetName()
		methodName := method.GetName()

		logger.Printf(">> pc:%4d %v.%v%v \n", frame.nextPC, className, methodName, method.GetDescriptor())
	}
}

func (recv *Thread) ClearStack() {

	recv.stack.clear()
}

func (recv *Thread) GetFrames() []*Frame {

	return recv.stack.getFrames()
}

func (recv *Thread) GetJavaMirror() *heap.OopDesc {

	return recv.threadOop
}

func (recv *Thread) GetName() string {

	return recv.name
}
