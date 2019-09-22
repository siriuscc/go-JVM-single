package security

import (
	"instructions/base"
	"native"
	"rtda"
)

func init() {

	const class = "java/security/AccessController"
	native.Register(class, "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", doPrivileged)
	native.Register(class, "doPrivileged", "(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;", doPrivileged)
	native.Register(class, "getStackAccessControlContext", "()Ljava/security/AccessControlContext;", getStackAccessControlContext)

}

// (Ljava/security/PrivilegedAction)Ljava/lang/Object
// public static native <T> T doPrivileged(PrivilegedAction<T> action);
func doPrivileged(frame *rtda.Frame) {

	locals := frame.LocalVars()
	action := locals.GetRef(0)

	opStack := frame.OpStack()
	opStack.PushRef(action)
	//T run();
	method := action.GetKlass().GetInstanceMethod("run", "()Ljava/lang/Object;")
	base.InvokeAMethod(frame, method)

}

// private static native AccessControlContext getStackAccessControlContext();
// ()Ljava/security/AccessControlContext;
func getStackAccessControlContext(frame *rtda.Frame) {
	// todo
	frame.OpStack().PushRef(nil)
}
