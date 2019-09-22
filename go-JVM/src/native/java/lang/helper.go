package lang

import (
	"rtda/heap"
	"unsafe"
)

// []byte => byte[]
func toByteArr(loader *heap.ClassLoader, goBytes []byte) *heap.OopDesc {
	if goBytes != nil {
		jBytes := castUint8sToInt8s(goBytes)
		return heap.CreateByteArray(loader, jBytes)
	}
	return nil
}

func toClassArrObj(classLoader *heap.ClassLoader, klasses []*heap.Klass) *heap.OopDesc {

	class := classLoader.LoadClass("java/lang/Class")

	count := int32(len(klasses))
	arrayObject := class.CreateArrayObject(count)
	descs := make([]*heap.OopDesc, count)

	for i, klass := range klasses {
		descs[i] = klass.GetJavaMirror()
	}

	arrayObject.SetVTable(descs)
	return arrayObject
}

// []string  转为 oop{ oop[] } 代表 [Ljava/lang/Class;
func toClassArr(classLoader *heap.ClassLoader, parameterTypesDesc []string) *heap.OopDesc {

	class := classLoader.LoadClass("[Ljava/lang/Class;")

	count := int32(len(parameterTypesDesc))
	arrObj := class.CreateArrayObject(count)
	descs := make([]*heap.OopDesc, count)

	for i, desc := range parameterTypesDesc {
		klass := classLoader.LoadClass(desc)
		descs[i] = klass.GetJavaMirror()
	}
	arrObj.SetVTable(descs)

	return arrObj
}

func castUint8sToInt8s(goBytes []byte) (jBytes []int8) {
	ptr := unsafe.Pointer(&goBytes)
	jBytes = *((*[]int8)(ptr))
	return
}

func getSignatureStr(loader *heap.ClassLoader, signature string) *heap.OopDesc {
	if signature != "" {
		return heap.JString(loader, signature)
	}
	return nil
}
