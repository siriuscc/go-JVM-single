package heap

import (
	"classfile"
	"strings"
)

type Method struct {
	ClassMember
	constValueIndex uint
	maxStack        uint
	maxLocal        uint
	code            []byte
	argSlotCount    uint

	exceptionTable ExceptionTable
	lineNumbers    LineNumberTable

	returnDesc string
	paramsDesc []string

	parameterAnnotationData []byte // RuntimeVisibleParameterAnnotations_attribute
	annotationDefaultData   []byte // AnnotationDefault_attribute
}

func (recv *Method) GetCode() []byte                  { return recv.code }
func (recv *Method) GetMaxStack() uint                { return recv.maxStack }
func (recv *Method) GetMaxLocals() uint               { return recv.maxLocal }
func (recv *Method) IsInitMethod() bool               { return recv.name == "<init>" }
func (recv *Method) IsClinitMethod() bool             { return recv.name == "<clinit>" }
func (recv *Method) ArgSlotCount() uint               { return recv.argSlotCount }
func (recv *Method) GetReturnTypeDesc() string        { return recv.returnDesc }
func (recv *Method) GetCheckExceptions() []*Klass     { return recv.exceptionTable.GetCatchTypes() }
func (recv *Method) GetParameterAnnotations() []byte  { return recv.parameterAnnotationData }
func (recv *Method) GetAnnotationDefaultData() []byte { return recv.annotationDefaultData }
func (recv *Method) GetLineNumber(pc uint16) int      { return recv.lineNumbers.GetLineNumber(uint(pc)) }

func (recv *Method) GetParameterTypesDesc() []string {

	recv.initMethodDescriptor()
	return recv.paramsDesc
}

func (recv *Method) calcArgSlotCount() {

	recv.initMethodDescriptor()
	count := len(recv.paramsDesc)

	for _, param := range recv.paramsDesc {
		if param == "J" || param == "D" {
			count++
		}
	}

	if !recv.accessFlags.IsStatic() {
		count++
	}
	recv.argSlotCount = uint(count)
}

// ()V
// ([Ljava.lang.String;])V
// ()
func (recv *Method) initMethodDescriptor() {

	// 先找出左右括号，截断

	if recv.returnDesc != "" {
		return
	}

	recv.paramsDesc = make([]string, 0)

	left := strings.LastIndex(recv.descriptor, "(") + 1
	right := strings.LastIndex(recv.descriptor, ")")

	recv.returnDesc = recv.descriptor[right+1:]

	if left == right {
		return
	} else if left > right {
		panic("error index")
	}
	params_str := recv.descriptor[left:right]

	var item string

	for len(params_str) > 0 {
		item, params_str = splitOneParam(params_str)

		if len(item) > 0 {
			recv.paramsDesc = append(recv.paramsDesc, item)
		}
	}

}

// 给native方法注入code
func (recv *Method) injectCodeAttribute() {

	recv.maxStack = 4
	recv.maxLocal = recv.argSlotCount

	switch recv.returnDesc[0] {
	case 'V':
		recv.code = []byte{0xfe, 0xb1} // return
	case 'D':
		recv.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		recv.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		recv.code = []byte{0xfe, 0xad} // lreturn
	case 'L', '[':
		recv.code = []byte{0xfe, 0xb0} // areturn
	default:
		recv.code = []byte{0xfe, 0xac} // ireturn
	}
}

func (recv *Method) FindExceptionHandler(ex *Klass, pc uint16) int {

	handler := recv.exceptionTable.findExceptionHandler(ex, int(pc))

	if handler != nil {
		return int(handler.handler)
	}
	return -1
}

func (recv *Klass) loadMethods(methodInfos []classfile.MethodInfo) {

	//logger.Println("loading :", recv.name)

	methods := make([]*Method, len(methodInfos))
	for i, info := range methodInfos {
		methods[i] = &Method{}

		methods[i].accessFlags = info.GetAccessFlags()
		methods[i].descriptor = info.GetDescriptor()
		methods[i].name = info.GetMethodName()

		//logger.Println("%%%%%",methods[i].name)

		methods[i].owner = recv
		methods[i].calcArgSlotCount()

		if codeAttr := info.GetCodeAttribute(); codeAttr != nil {
			methods[i].code = codeAttr.GetCode()
			methods[i].maxStack = uint(codeAttr.GetMaxStack())
			methods[i].maxLocal = uint(codeAttr.GetMaxLocals())
		}

		if methods[i].accessFlags.IsNative() {
			methods[i].injectCodeAttribute()
		}

		//logger.Printf("LineNumberTableAttr: %v \n",info.GetLineNumberTableAttr())

		methods[i].loadLineNumbers(info.GetLineNumberTableAttr())
		methods[i].exceptionTable = CreateExceptionTable(info.GetExceptionTableAttr(), recv.constantPool)
		methods[i].signatures = info.GetSignaturesAttr().GetSignature()

		methods[i].annotationData = info.RuntimeVisibleAnnotationsAttributeData()
		methods[i].parameterAnnotationData = info.RuntimeVisibleParameterAnnotationsAttributeData()
		methods[i].annotationDefaultData = info.AnnotationDefaultAttributeData()
	}

	recv.methods = methods
}
