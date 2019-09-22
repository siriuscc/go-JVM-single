package heap

import "classfile"

type ExceptionHandler struct {
	startPC   uint16
	endPC     uint16
	handler   uint16
	catchType *ClassRef
}
type ExceptionTable []*ExceptionHandler

func (recv ExceptionTable) GetCatchTypes() []*Klass {

	klasses := make([]*Klass, len(recv))

	for i, handler := range recv {

		klasses[i] = handler.catchType.ResolvedClass()
	}
	return klasses
}

// cp: 运行时常量池
func CreateExceptionTable(entries []*classfile.ExceptionInfo, cp *ConstantPool) ExceptionTable {

	handlers := make([]*ExceptionHandler, len(entries))

	for i, info := range entries {
		handlers[i] = &ExceptionHandler{
			startPC:   info.GetStartPC(),
			endPC:     info.GetEndPC(),
			handler:   info.GetHandlerPC(),
			catchType: getCatchType(info.GetCatchType(), cp),
		}
	}
	return handlers
}

func getCatchType(index uint16, cp *ConstantPool) *ClassRef {

	if index == 0 {
		return nil
	} else {
		return cp.GetConstant(uint(index)).(*ClassRef)
	}
}

// 找到了就是非空
// exClass 异常类
// pc		异常发生对应的执行位置
func (recv ExceptionTable) findExceptionHandler(exClass *Klass, pc int) *ExceptionHandler {

	for _, handler := range recv {

		if pc >= int(handler.startPC) && pc < int(handler.endPC) {
			if handler.catchType == nil { // catch-all
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || exClass.IsSubClassOf(catchClass) {
				return handler
			}
		}
	}
	return nil
}
