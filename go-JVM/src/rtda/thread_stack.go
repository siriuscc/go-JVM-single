package rtda

// 单向链表,头插法入栈，尾巴出栈
type ThreadStack struct {
	maxSize uint
	size    uint
	top     *Frame
}

func CreateThreadStack(maxStackSize uint) *ThreadStack {

	stack := &ThreadStack{maxSize: maxStackSize, size: 0, top: nil}

	return stack
}

func (recv *ThreadStack) push(frame *Frame) bool {

	if frame == nil {
		return false
	}

	frame.lower = recv.top
	recv.top = frame

	recv.size++

	return true
}

func (recv *ThreadStack) pop() *Frame {

	if recv.top == nil {
		return nil
	}

	tmp := recv.top
	recv.top = tmp.lower
	recv.size--

	return tmp
}

func (recv *ThreadStack) getTop() *Frame {

	return recv.top
}

func (recv *ThreadStack) getSize() uint {
	return recv.size
}

func (recv *ThreadStack) isEmpty() bool {
	return recv.top == nil
}

func (recv *ThreadStack) clear() {

	recv.top = nil
	recv.size = 0

	//for !recv.isEmpty(){
	//	recv.pop()
	//}
}

func (recv *ThreadStack) getFrames() []*Frame {

	frames := make([]*Frame, 0)

	for f := recv.top; f != nil; f = f.lower {
		frames = append(frames, f)
	}
	return frames
}
