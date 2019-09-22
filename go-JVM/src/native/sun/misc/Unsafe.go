package misc

import (
	"native"
	"rtda"
	"rtda/heap"
	"unsafe"
)

func init() {

	const class = "sun/misc/Unsafe"
	native.Register(class, "arrayBaseOffset", "(Ljava/lang/Class;)I", unsafe_arrayBaseOffset)
	native.Register(class, "arrayIndexScale", "(Ljava/lang/Class;)I", unsafe_arrayIndexScale)
	native.Register(class, "addressSize", "()I", unsafe_addressSize)
	native.Register(class, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", unsafe_objectFieldOffset)
	native.Register(class, "getIntVolatile", "(Ljava/lang/Object;J)I", unsafe_getIntVolatile)
	native.Register(class, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;", unsafe_getObject)

	native.Register(class, "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", unsafe_compareAndSwapObject)
	native.Register(class, "compareAndSwapInt", "(Ljava/lang/Object;JII)Z", unsafe_compareAndSwapInt)
	native.Register(class, "compareAndSwapLong", "(Ljava/lang/Object;JJJ)Z", unsafe_compareAndSwapLong)

	native.Register(class, "allocateMemory", "(J)J", allocateMemory)
	native.Register(class, "reallocateMemory", "(JJ)J", reallocateMemory)
	native.Register(class, "freeMemory", "(J)V", freeMemory)
	native.Register(class, "getByte", "(J)B", getByte)
	native.Register(class, "putLong", "(JJ)V", putLong)

}

// public native int arrayBaseOffset(Class<?> klass);
// 获取数组第一个元素的偏移地址
func unsafe_arrayBaseOffset(frame *rtda.Frame) {

	//locals := frame.LocalVars()

	//classInstance:=locals.GetRef(1)
	//classInstance.GetArrayBaseOffset()

	// 这里 本来应该是  this.vtable-this 得到偏移量的，go不支持指针运算。姑且返回0
	frame.OpStack().PushInt(0)
}

// 获取步长
// 这样做不过是权宜之计
// public native int unsafe_arrayIndexScale(Class<?> var1);
func unsafe_arrayIndexScale(frame *rtda.Frame) {
	stack := frame.OpStack()
	stack.PushInt(1)
}

// 指针的size, sizeof(p),应该是8
// public native int unsafe_addressSize();
func unsafe_addressSize(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()
	sizeof := unsafe.Sizeof(this)
	frame.OpStack().PushInt(int32(sizeof))
}

// 获取field 的 slot属性
// 对于实例，field是以slot占位置的，long和double占两个slot，可以理解为slot是一个int32
// public native long unsafe_objectFieldOffset(Field var1);
func unsafe_objectFieldOffset(frame *rtda.Frame) {

	locals := frame.LocalVars()
	javaFieldOop := locals.GetRef(1)

	offset := javaFieldOop.GetIntField("slot", "I")

	stack := frame.OpStack()
	stack.PushLong(int64(offset))
}

// 著名的CAS操作
// public final native boolean compareAndSwapObject(Object oop, long offset, Object expected, Object updateValue);
func unsafe_compareAndSwapObject(frame *rtda.Frame) {

	locals := frame.LocalVars()

	obj := locals.GetRef(1)
	offset := locals.GetLong(2)
	expected := locals.GetRef(4)
	updateValue := locals.GetRef(5)

	//obj 在offset 位置的对象 等不等于 expected,等于就 赋值：updateValue

	// oop 可能是对象
	// 可能是数组

	fields := obj.GetFields()

	if !obj.GetKlass().IsArray() {
		// object,
		swapped := _casObj(obj, fields, offset, expected, updateValue)
		frame.OpStack().PushBool(swapped)
	} else if !obj.GetKlass().ComponentClass().IsPrimitive() { // 组件类型是 Object

		swapped := _casArr(obj, offset, expected, updateValue)
		frame.OpStack().PushBool(swapped)
	} else {
		panic("compareAndSwapObject: component is not ref")
	}
}

// public final native boolean compareAndSwapInt(Object obj, long offset, int expected, int updateValue);
func unsafe_compareAndSwapInt(frame *rtda.Frame) {

	locals := frame.LocalVars()
	opStack := frame.OpStack()

	obj := locals.GetRef(1)
	offset := locals.GetLong(2) // 两个槽
	expected := locals.GetInt(4)
	updateValue := locals.GetInt(5)

	if obj.GetKlass().IsArray() {
		// [I   arr[i]=
		ints := obj.GetIntTable()
		if ints[offset] == expected {
			ints[offset] = updateValue
			opStack.PushBool(true)
		} else {
			opStack.PushBool(false)
		}
	} else {
		// object.int
		fields := obj.GetFields()
		v := fields.GetInt(uint(offset))
		if v == expected {
			fields.SetInt(uint(offset), updateValue)
			opStack.PushBool(true)
		} else {
			opStack.PushBool(false)
		}
	}
}

//public final native boolean compareAndSwapLong(Object obj, long offset, long expected, long updateValue);

func unsafe_compareAndSwapLong(frame *rtda.Frame) {

	locals := frame.LocalVars()
	opStack := frame.OpStack()

	obj := locals.GetRef(1)
	offset := locals.GetLong(2)
	expected := locals.GetLong(4)
	updateValue := locals.GetLong(6)

	if obj.GetKlass().IsArray() {
		// [I   arr[i]=
		longs := obj.GetLongTable()
		if longs[offset] == expected {
			longs[offset] = updateValue
			opStack.PushBool(true)
		} else {
			opStack.PushBool(false)
		}
	} else {
		// object.int
		fields := obj.GetFields()
		v := fields.GetLong(uint(offset))
		if v == expected {
			fields.SetLong(uint(offset), updateValue)
			opStack.PushBool(true)
		} else {
			opStack.PushBool(false)
		}
	}
}

//public native Object getObject(Object obj, long offset);
func unsafe_getObject(frame *rtda.Frame) {

	locals := frame.LocalVars()
	opStack := frame.OpStack()
	obj := locals.GetRef(1)
	offset := locals.GetLong(2)

	if obj.GetKlass().IsArray() {
		// [I   arr[i]=
		refs := obj.GetRefTable()
		opStack.PushRef(refs[offset])

	} else {
		// object.int
		fields := obj.GetFields()
		ref := fields.GetRef(uint(offset))
		opStack.PushRef(ref)
	}
}

//public native int getIntVolatile(Object obj, long offset);
func unsafe_getIntVolatile(frame *rtda.Frame) {

	locals := frame.LocalVars()
	opStack := frame.OpStack()

	obj := locals.GetRef(1)
	offset := locals.GetInt(2)

	if obj.GetKlass().IsArray() {
		table := obj.GetIntTable()
		opStack.PushInt(table[offset])
	} else {
		// object,
		v := obj.GetFields().GetInt(uint(offset))
		opStack.PushInt(v)
	}
}

//public final native boolean compareAndSwapLong(Object var1, long var2, long var4, long var6);

func _casArr(arr *heap.OopDesc, index int64, expected *heap.OopDesc, updateValue *heap.OopDesc) bool {

	refTable := arr.GetRefTable()

	if refTable[index] == expected {
		refTable[index] = updateValue
		return true
	}
	return false
}

func _casObj(obj *heap.OopDesc, fields heap.Slots, offset int64, expected *heap.OopDesc, updateValue *heap.OopDesc) bool {

	ref := fields.GetRef(uint(offset))

	if ref == expected {
		fields.SetRef(uint(offset), updateValue)
		return true
	} else {
		return false
	}
}
