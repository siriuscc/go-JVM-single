package native

import "rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDesc string, method NativeMethod) {
	key := className + "~" + methodName + "-" + methodDesc
	registry[key] = method
}

func FindNativeMethod(className string, methodName string, methodDesc string) NativeMethod {
	key := className + "~" + methodName + "-" + methodDesc

	if method, ok := registry[key]; ok {
		return method
	}

	if methodDesc == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}

	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}
