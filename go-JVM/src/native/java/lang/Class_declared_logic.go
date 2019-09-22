package lang

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

//private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// 得到所有的构造器
func getDeclaredConstructors0(frame *rtda.Frame) {

	const _constructorDesc = "" +
		"(Ljava/lang/Class;" +
		"[Ljava/lang/Class;" +
		"[Ljava/lang/Class;" +
		"II" +
		"Ljava/lang/String;" +
		"[B[B)V"

	locals := frame.LocalVars()
	oop := locals.GetThis()
	publicOnly := locals.GetBoolean(1)

	classLoader := frame.GetMethod().GetOwner().GetClassLoader()

	constructorKlass := classLoader.LoadClass("java/lang/reflect/Constructor")
	constructorArrKlass := constructorKlass.ArrayClass()

	constructorMethods := oop.GetMetaData().GetConstructors(publicOnly)
	constructorCount := uint(len(constructorMethods))

	arrayOop := constructorArrKlass.CreateArrayObject(int32(constructorCount))
	frame.OpStack().PushRef(arrayOop)

	if constructorCount > 0 {
		thread := frame.GetThread()
		refTable := make([]*heap.OopDesc, constructorCount)
		constructorInitMethod := constructorKlass.GetInitMethod(_constructorDesc)

		for i, constructorMethod := range constructorMethods {
			constructorObj := constructorKlass.CreateObject()
			refTable[i] = constructorObj

			driverFrame := thread.CreateDriverFrame(9)

			ops := driverFrame.OpStack()
			ops.PushRef(constructorObj)                                                      // this
			ops.PushRef(oop)                                                                 // declaringClass
			ops.PushRef(toClassArr(classLoader, constructorMethod.GetParameterTypesDesc()))  // parameterTypes
			ops.PushRef(toClassArrObj(classLoader, constructorMethod.GetCheckExceptions()))  // checkedExceptions
			ops.PushInt(int32(constructorMethod.GetAccessFlags().GetFlags()))                // modifiers
			ops.PushInt(int32(0))                                                            // todo slot
			ops.PushRef(getSignatureStr(classLoader, constructorMethod.GetSignature()))      // signature
			ops.PushRef(toByteArr(classLoader, constructorMethod.GetAnnotationData()))       // annotations
			ops.PushRef(toByteArr(classLoader, constructorMethod.GetParameterAnnotations())) // parameterAnnotations

			// init constructorObj
			base.InvokeAMethod(driverFrame, constructorInitMethod)
		}
		arrayOop.SetVTable(refTable)
	}

}

// [idea] 对于每个method
//    创建一个调用者帧，callerFrame
//    调用其初始化函数
//
// (B)[Ljava/lang/reflect/Method;
// private native Method[] getDeclaredMethods0(boolean publicOnly);
func getDeclaredMethods0(frame *rtda.Frame) {

	const _methodInitDesc = "" +
		"(Ljava/lang/Class;" +
		"Ljava/lang/String;" +
		"[Ljava/lang/Class;" +
		"Ljava/lang/Class;" +
		"[Ljava/lang/Class;" +
		"II" +
		"Ljava/lang/String;" +
		"[B[B[B)V"

	thread := frame.GetThread()
	locals := frame.LocalVars()

	oop := locals.GetThis()
	publicOnly := locals.GetBoolean(1)

	methods := oop.GetMetaData().GetMethods(publicOnly)

	classLoader := frame.GetMethod().GetOwner().GetClassLoader()
	methodKlass := classLoader.LoadClass("java/lang/reflect/Method")
	methodArrKlass := methodKlass.ArrayClass()

	methodCount := uint(len(methods))

	arrayOop := methodArrKlass.CreateArrayObject(int32(methodCount))
	refTable := make([]*heap.OopDesc, len(methods))
	methodConstructor := methodKlass.GetInitMethod(_methodInitDesc)

	// 准备参数
	for i, method := range methods {

		methodObj := methodKlass.CreateObject()
		refTable[i] = methodObj

		callerFrame := thread.CreateDriverFrame(12)

		checkedExceptions := method.GetCheckExceptions()
		signature := method.GetSignature()
		annotations := method.GetAnnotationData()
		parameterAnnotations := method.GetParameterAnnotations()
		annotationDefault := method.GetAnnotationDefaultData()

		ops := callerFrame.OpStack()
		ops.PushRef(methodObj) //todo
		ops.PushRef(methodKlass.GetJavaMirror())
		ops.PushRef(heap.JString(classLoader, method.GetName()))                       // name
		ops.PushRef(toClassArr(classLoader, method.GetParameterTypesDesc()))           // parameterTypes
		ops.PushRef(classLoader.LoadClass(method.GetReturnTypeDesc()).GetJavaMirror()) // returnType
		ops.PushRef(toClassArrObj(classLoader, checkedExceptions))                     // checkedExceptions
		ops.PushInt(int32(method.GetAccessFlags().GetFlags()))                         // modifiers
		ops.PushInt(int32(0))                                                          // todo: slot
		ops.PushRef(getSignatureStr(classLoader, signature))                           // signature
		ops.PushRef(toByteArr(classLoader, annotations))                               // annotations
		ops.PushRef(toByteArr(classLoader, parameterAnnotations))                      // parameterAnnotations
		ops.PushRef(toByteArr(classLoader, annotationDefault))                         // annotationDefault

		// 调用method 的初始化方法
		base.InvokeAMethod(callerFrame, methodConstructor)
	}

	arrayOop.SetVTable(refTable)
}

// 获取this 的所有field
// private native Field[] getDeclaredFields0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Field;
func getDeclaredFields0(frame *rtda.Frame) {
	const _fieldConstructorDescriptor = "" +
		"(Ljava/lang/Class;" +
		"Ljava/lang/String;" +
		"Ljava/lang/Class;" +
		"II" +
		"Ljava/lang/String;" +
		"[B)V"

	thread := frame.GetThread()
	locals := frame.LocalVars()
	classLoader := frame.GetMethod().GetOwner().GetClassLoader()

	oop := locals.GetThis()
	publicOnly := locals.GetBoolean(1)

	metaData := oop.GetMetaData()
	fields := metaData.GetFields(publicOnly)
	fieldCount := uint(len(fields))

	fieldKlass := classLoader.LoadClass("java/lang/reflect/Field")
	fieldArr := fieldKlass.CreateArrayObject(int32(fieldCount)) // []Field

	opStack := frame.OpStack()
	opStack.PushRef(fieldArr)

	if fieldCount > 0 {
		fieldObjs := fieldArr.Refs()

		fieldInitMethod := fieldKlass.GetInitMethod(_fieldConstructorDescriptor)

		for i, field := range fields {

			callerFrame := thread.CreateDriverFrame(8)

			fieldObj := fieldKlass.CreateObject()
			fieldObjs[i] = fieldObj

			ops := callerFrame.OpStack()
			ops.PushRef(fieldObj)                                           // this
			ops.PushRef(oop)                                                // declaringClass
			ops.PushRef(heap.JString(classLoader, field.GetName()))         // name
			ops.PushRef(field.GetType())                                    // type
			ops.PushInt(int32(field.GetAccessFlags().GetFlags()))           // modifiers
			ops.PushInt(int32(field.GetSlotId()))                           // slot
			ops.PushRef(getSignatureStr(classLoader, field.GetSignature())) // signature
			ops.PushRef(toByteArr(classLoader, field.GetAnnotationData()))  // annotations

			// init fieldObj
			base.InvokeAMethod(callerFrame, fieldInitMethod)
		}
	}
}
