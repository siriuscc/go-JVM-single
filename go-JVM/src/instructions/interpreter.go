package instructions

import (
	"instructions/base"
	"instructions/comparisons"
	"instructions/constants"
	"instructions/control"
	"instructions/conversions"
	"instructions/extended"
	"instructions/loads"
	"instructions/math"
	"instructions/references"
	"instructions/reserved"
	"instructions/stack"
	"instructions/stores"
	"logger"
	"reflect"
	"rtda"
	"strings"
)

// 解释器，
type Interpreter struct {
	verbose        bool                       // 是否打印错误信息
	instructionMap map[uint8]base.Instruction // 指令集
}

// 解析器主流程, 解释器只负责解释，应该把参数的注入移出去
func (recv *Interpreter) Interpret(thread *rtda.Thread) {

	defer recv.catchError(thread)

	recv.loop(thread)
}

// 异常处理
func (recv *Interpreter) catchError(thread *rtda.Thread) {

	// catch 所有的panic
	if r := recover(); r != nil {
		thread.LogFrames()

		panic(r)
	}
}

// 命令执行
func (recv *Interpreter) loop(currentThread *rtda.Thread) {

	// 构建一个reader
	reader := &base.ByteCodeReader{}
	for {

		// 准备阶段，pc指向将要执行的指令
		frame := currentThread.GetCurrentFrame()
		nextPC := frame.GetNextPC()
		currentThread.SetPC(nextPC)
		reader.Reset(frame.GetMethod().GetCode(), nextPC)

		// 获取指令，获取指令对应的数据
		opCode := reader.ReadUint8()
		inst := recv.instructionMap[opCode]
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.GetPC()) // frame执行完有更新
		if recv.verbose {
			recv.logInstruction(frame, inst)
		}

		// 执行阶段
		inst.Execute(frame)
		if currentThread.HasNoFrame() {
			break
		}
	}
}

func (recv *Interpreter) Init(verbose bool) {
	recv.verbose = verbose

	recv.instructionMap = make(map[uint8]base.Instruction)

	// 常量部分
	recv.instructionMap[0x00] = &control.NOP{}
	recv.instructionMap[0x01] = &control.ACONST_NULL{}
	recv.instructionMap[0x02] = &control.ICONST_M1{}
	recv.instructionMap[0x03] = &control.ICONST_0{}
	recv.instructionMap[0x04] = &control.ICONST_1{}
	recv.instructionMap[0x05] = &control.ICONST_2{}
	recv.instructionMap[0x06] = &control.ICONST_3{}
	recv.instructionMap[0x07] = &control.ICONST_4{}
	recv.instructionMap[0x08] = &control.ICONST_5{}
	recv.instructionMap[0x09] = &control.LCONST_0{}
	recv.instructionMap[0x0A] = &control.LCONST_1{}
	recv.instructionMap[0x0B] = &control.FCONST_0{}
	recv.instructionMap[0x0C] = &control.FCONST_1{}
	recv.instructionMap[0x0D] = &control.FCONST_2{}
	recv.instructionMap[0x0E] = &control.DCONST_0{}
	recv.instructionMap[0x0F] = &control.DCONST_1{}

	recv.instructionMap[0x10] = &constants.BIPUSH{}
	recv.instructionMap[0x11] = &constants.SIPUSH{}
	recv.instructionMap[0x12] = &constants.LDC{}
	recv.instructionMap[0x13] = &constants.LDC_W{}
	recv.instructionMap[0x14] = &constants.LDC2_W{}

	// loads
	recv.instructionMap[0x15] = &loads.ILOAD{}
	recv.instructionMap[0x16] = &loads.LLOAD{}
	recv.instructionMap[0x17] = &loads.FLOAD{}
	recv.instructionMap[0x18] = &loads.DLOAD{}
	recv.instructionMap[0x19] = &loads.ALOAD{}
	recv.instructionMap[0x1A] = &loads.ILOAD_0{}
	recv.instructionMap[0x1B] = &loads.ILOAD_1{}
	recv.instructionMap[0x1C] = &loads.ILOAD_2{}
	recv.instructionMap[0x1D] = &loads.ILOAD_3{}
	recv.instructionMap[0x1E] = &loads.LLOAD_0{}
	recv.instructionMap[0x1F] = &loads.LLOAD_1{}
	recv.instructionMap[0x20] = &loads.LLOAD_2{}
	recv.instructionMap[0x21] = &loads.LLOAD_3{}
	recv.instructionMap[0x22] = &loads.FLOAD_0{}
	recv.instructionMap[0x23] = &loads.FLOAD_1{}
	recv.instructionMap[0x24] = &loads.FLOAD_2{}
	recv.instructionMap[0x25] = &loads.FLOAD_3{}
	recv.instructionMap[0x26] = &loads.DLOAD_0{}
	recv.instructionMap[0x27] = &loads.DLOAD_1{}
	recv.instructionMap[0x28] = &loads.DLOAD_2{}
	recv.instructionMap[0x29] = &loads.DLOAD_3{}
	recv.instructionMap[0x2A] = &loads.ALOAD_0{}
	recv.instructionMap[0x2B] = &loads.ALOAD_1{}
	recv.instructionMap[0x2C] = &loads.ALOAD_2{}
	recv.instructionMap[0x2D] = &loads.ALOAD_3{}

	recv.instructionMap[0x2E] = &loads.IALOAD{}
	recv.instructionMap[0x2F] = &loads.LALOAD{}
	recv.instructionMap[0x30] = &loads.FALOAD{}
	recv.instructionMap[0x31] = &loads.DALOAD{}
	recv.instructionMap[0x32] = &loads.AALOAD{}
	recv.instructionMap[0x33] = &loads.BALOAD{}
	recv.instructionMap[0x34] = &loads.CALOAD{}
	recv.instructionMap[0x35] = &loads.SALOAD{}

	// stores
	recv.instructionMap[0x36] = &stores.ISTORE{}
	recv.instructionMap[0x37] = &stores.LSTORE{}
	recv.instructionMap[0x38] = &stores.FSTORE{}
	recv.instructionMap[0x39] = &stores.DSTORE{}
	recv.instructionMap[0x3A] = &stores.ASTORE{}
	recv.instructionMap[0x3B] = &stores.ISTORE_0{}
	recv.instructionMap[0x3C] = &stores.ISTORE_1{}
	recv.instructionMap[0x3D] = &stores.ISTORE_2{}
	recv.instructionMap[0x3E] = &stores.ISTORE_3{}
	recv.instructionMap[0x3F] = &stores.LSTORE_0{}
	recv.instructionMap[0x40] = &stores.LSTORE_1{}
	recv.instructionMap[0x41] = &stores.LSTORE_2{}
	recv.instructionMap[0x42] = &stores.LSTORE_3{}
	recv.instructionMap[0x43] = &stores.FSTORE_0{}
	recv.instructionMap[0x44] = &stores.FSTORE_1{}
	recv.instructionMap[0x45] = &stores.FSTORE_2{}
	recv.instructionMap[0x46] = &stores.FSTORE_3{}
	recv.instructionMap[0x47] = &stores.DSTORE_0{}
	recv.instructionMap[0x48] = &stores.DSTORE_1{}
	recv.instructionMap[0x49] = &stores.DSTORE_2{}
	recv.instructionMap[0x4A] = &stores.DSTORE_3{}
	recv.instructionMap[0x4B] = &stores.ASTORE_0{}
	recv.instructionMap[0x4C] = &stores.ASTORE_1{}
	recv.instructionMap[0x4D] = &stores.ASTORE_2{}
	recv.instructionMap[0x4E] = &stores.ASTORE_3{}
	recv.instructionMap[0x4F] = &stores.IASTORE{}
	recv.instructionMap[0x50] = &stores.LASTORE{}
	recv.instructionMap[0x51] = &stores.FASTORE{}
	recv.instructionMap[0x52] = &stores.DASTORE{}
	recv.instructionMap[0x53] = &stores.AASTORE{}
	recv.instructionMap[0x54] = &stores.BASTORE{}
	recv.instructionMap[0x55] = &stores.CASTORE{}
	recv.instructionMap[0x56] = &stores.SASTORE{}

	// Stack
	recv.instructionMap[0x57] = &stack.POP{}
	recv.instructionMap[0x58] = &stack.POP2{}
	recv.instructionMap[0x59] = &stack.DUP{}
	recv.instructionMap[0x5A] = &stack.DUP_X1{}
	recv.instructionMap[0x5B] = &stack.DUP_X2{}
	recv.instructionMap[0x5C] = &stack.DUP2{}
	recv.instructionMap[0x5D] = &stack.DUP2_X1{}
	recv.instructionMap[0x5E] = &stack.DUP2_X2{}
	recv.instructionMap[0x5F] = &stack.SWAP{}

	//math
	recv.instructionMap[0x60] = &math.IADD{}
	recv.instructionMap[0x61] = &math.LADD{}
	recv.instructionMap[0x62] = &math.FADD{}
	recv.instructionMap[0x63] = &math.DADD{}
	recv.instructionMap[0x64] = &math.ISUB{}
	recv.instructionMap[0x65] = &math.LSUB{}
	recv.instructionMap[0x66] = &math.FSUB{}
	recv.instructionMap[0x67] = &math.DSUB{}
	recv.instructionMap[0x68] = &math.IMUL{}
	recv.instructionMap[0x69] = &math.LMUL{}
	recv.instructionMap[0x6A] = &math.FMUL{}
	recv.instructionMap[0x6B] = &math.DMUL{}
	recv.instructionMap[0x6C] = &math.IDIV{}
	recv.instructionMap[0x6D] = &math.LDIV{}
	recv.instructionMap[0x6E] = &math.FDIV{}
	recv.instructionMap[0x6F] = &math.DDIV{}
	recv.instructionMap[0x70] = &math.IREM{}
	recv.instructionMap[0x71] = &math.LREM{}
	recv.instructionMap[0x72] = &math.FREM{}
	recv.instructionMap[0x73] = &math.DREM{}
	recv.instructionMap[0x74] = &math.INEG{}
	recv.instructionMap[0x75] = &math.LNEG{}
	recv.instructionMap[0x76] = &math.FNEG{}
	recv.instructionMap[0x77] = &math.DNEG{}
	recv.instructionMap[0x78] = &math.ISHL{}
	recv.instructionMap[0x79] = &math.LSHL{}
	recv.instructionMap[0x7A] = &math.ISHR{}
	recv.instructionMap[0x7B] = &math.LSHR{}
	recv.instructionMap[0x7C] = &math.IUSHR{}
	recv.instructionMap[0x7D] = &math.LUSHR{}
	recv.instructionMap[0x7E] = &math.IAND{}
	recv.instructionMap[0x7F] = &math.LAND{}
	recv.instructionMap[0x80] = &math.IOR{}
	recv.instructionMap[0x81] = &math.LOR{}
	recv.instructionMap[0x82] = &math.IXOR{}
	recv.instructionMap[0x83] = &math.LXOR{}
	recv.instructionMap[0x84] = &math.IINC{}

	// conversions
	recv.instructionMap[0x85] = &conversions.I2L{}
	recv.instructionMap[0x86] = &conversions.I2F{}
	recv.instructionMap[0x87] = &conversions.I2D{}
	recv.instructionMap[0x88] = &conversions.L2I{}
	recv.instructionMap[0x89] = &conversions.L2F{}
	recv.instructionMap[0x8A] = &conversions.L2D{}
	recv.instructionMap[0x8B] = &conversions.F2I{}
	recv.instructionMap[0x8C] = &conversions.F2L{}
	recv.instructionMap[0x8D] = &conversions.F2D{}
	recv.instructionMap[0x8E] = &conversions.D2I{}
	recv.instructionMap[0x8F] = &conversions.D2L{}
	recv.instructionMap[0x90] = &conversions.D2F{}
	recv.instructionMap[0x91] = &conversions.I2B{}
	recv.instructionMap[0x92] = &conversions.I2C{}
	recv.instructionMap[0x93] = &conversions.I2S{}

	// Comparisons
	recv.instructionMap[0x94] = &comparisons.LCMP{}
	recv.instructionMap[0x95] = &comparisons.FCMPL{}
	recv.instructionMap[0x96] = &comparisons.FCMPG{}
	recv.instructionMap[0x97] = &comparisons.DCMPL{}
	recv.instructionMap[0x98] = &comparisons.DCMPG{}
	recv.instructionMap[0x99] = &comparisons.IFEQ{}
	recv.instructionMap[0x9A] = &comparisons.IFNE{}
	recv.instructionMap[0x9B] = &comparisons.IFLT{}
	recv.instructionMap[0x9C] = &comparisons.IFGE{}
	recv.instructionMap[0x9D] = &comparisons.IFGT{}
	recv.instructionMap[0x9E] = &comparisons.IFLE{}
	recv.instructionMap[0x9F] = &comparisons.IF_ICMPEQ{}
	recv.instructionMap[0xA0] = &comparisons.IF_ICMPNE{}
	recv.instructionMap[0xA1] = &comparisons.IF_ICMPLT{}
	recv.instructionMap[0xA2] = &comparisons.IF_ICMPGE{}
	recv.instructionMap[0xA3] = &comparisons.IF_ICMPGT{}
	recv.instructionMap[0xA4] = &comparisons.IF_ICMPLE{}
	recv.instructionMap[0xA5] = &comparisons.IF_ACMPEQ{}
	recv.instructionMap[0xA6] = &comparisons.IF_ACMPNE{}

	// Control
	recv.instructionMap[0xA7] = &control.GOTO{}
	recv.instructionMap[0xA8] = &control.JSR{}
	//recv.instructionMap[0xA9] = &control.RET{}
	recv.instructionMap[0xAA] = &control.TABLESWITCH{}
	recv.instructionMap[0xAB] = &control.LOOPUPSWITCH{}
	recv.instructionMap[0xAC] = &control.IRETURN{}
	recv.instructionMap[0xAD] = &control.LRETURN{}
	recv.instructionMap[0xAE] = &control.FRETURN{}
	recv.instructionMap[0xAF] = &control.DRETURN{}
	recv.instructionMap[0xB0] = &control.ARETURN{}
	recv.instructionMap[0xB1] = &control.RETURN{}

	//References
	recv.instructionMap[0xB2] = &references.GETSTATIC{}
	recv.instructionMap[0xB3] = &references.PUTSTATIC{}
	recv.instructionMap[0xB4] = &references.GETFIELD{}
	recv.instructionMap[0xB5] = &references.PUTFIELD{}
	recv.instructionMap[0xB6] = &references.INVOKE_VIRTUAL{}
	recv.instructionMap[0xB7] = &references.INVOKE_SPECIAL{}
	recv.instructionMap[0xB8] = &references.INVOKE_STATIC{}
	recv.instructionMap[0xB9] = &references.INVOKE_INTERFACE{}
	//recv.instructionMap[0xBA] = &references.INVOKE_DYNAMIC{}
	recv.instructionMap[0xBB] = &references.NEW{}
	recv.instructionMap[0xBC] = &references.NEW_ARRAY{}
	recv.instructionMap[0xBD] = &references.ANEW_ARRAY{}
	recv.instructionMap[0xBE] = &references.ARRAY_LENGTH{}
	recv.instructionMap[0xBF] = &references.ATHROW{}
	recv.instructionMap[0xC0] = &references.CHECKCAST{}
	recv.instructionMap[0xC1] = &references.INSTANCE_OF{}
	recv.instructionMap[0xC2] = &references.MONITOR_ENTER{}
	recv.instructionMap[0xC3] = &references.MONITOR_EXIT{}
	// Extended
	recv.instructionMap[0xC4] = &extended.WIDE{}
	recv.instructionMap[0xC5] = &references.MULTI_ANEW_ARRAY{}
	recv.instructionMap[0xC6] = &extended.IFNULL{}
	recv.instructionMap[0xC7] = &extended.IFNONNULL{}
	recv.instructionMap[0xC8] = &extended.GOTO_W{}
	recv.instructionMap[0xC9] = &extended.JSR_W{}

	// Reserved
	recv.instructionMap[0xCA] = &reserved.BREAKPOINT{}

	recv.instructionMap[0xFE] = &reserved.INVOKE_NATIVE{}
	//recv.instructionMap[0xFE] = &reserved.IMPDEP1{}
	recv.instructionMap[0xFF] = &reserved.IMPDEP2{}
}

func (recv *Interpreter) GetInstruction(opCode uint8) base.Instruction {

	return recv.instructionMap[opCode]
}

func (recv *Interpreter) GetWideInstruction(opCode uint8) base.WideInstruction {

	inst, ok := recv.instructionMap[opCode].(base.WideInstruction)
	if ok {
		return inst
	}
	return nil
}

func (recv *Interpreter) logInstruction(frame *rtda.Frame, inst base.Instruction) {

	method := frame.GetMethod()
	className := method.GetOwner().GetName()
	methodName := method.GetName()
	pc := frame.GetNextPC()

	count := frame.GetThread().GetThreadFrameCount()

	logger.Printf("%s |- %v.%v() #%2d %s %v\n", strings.Repeat("    ", int(count)), className, methodName, pc, recv.GetInstructionName(inst), inst)
}

func (recv *Interpreter) GetInstructionName(inst base.Instruction) string {

	split := strings.Split(reflect.TypeOf(inst).String(), ".")
	instName := split[len(split)-1]
	instName = strings.ToLower(instName)

	return instName
}
