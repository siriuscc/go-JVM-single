package control

import (
	"instructions/base"
	"rtda"
)

// 把 null ref 推入操作数栈顶
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (recv *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OpStack().PushRef(nil)
}

// 把 double 0 推入栈顶
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (recv *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OpStack().PushDouble(0.0000000)
}

// 把 double 1 推入栈顶
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (recv *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OpStack().PushDouble(1.0000000)
}

// 把float 0
type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (recv *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OpStack().PushFloat(0.0000000)
}

// 把float 1
type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (recv *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OpStack().PushFloat(1.0000000)
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (recv *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OpStack().PushFloat(2.0000000)
}

// 把 int -1 入栈
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (recv *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (recv *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (recv *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (recv *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (recv *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (recv *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (recv *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(5)
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (recv *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OpStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (recv *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OpStack().PushLong(1)
}
