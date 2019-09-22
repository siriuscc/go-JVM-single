package lang

import (
	"native"
	"rtda"
)

func init() {

	class := "java/lang/Thread"
	native.Register(class, "currentThread", "()Ljava/lang/Thread;", currentThread)
	native.Register(class, "setPriority0", "(I)V", setPriority0)
	native.Register(class, "isAlive", "()Z", isAlive)
	native.Register(class, "start0", "()V", start0)

}

// 线程是否存活
// public final native boolean isAlive();
func isAlive(frame *rtda.Frame) {
	// todo
	frame.OpStack().PushBool(false)
}

// private native void start0();
// ()V
func start0(frame *rtda.Frame) {
	// todo
}

//private native void setPriority0(int newPriority);
func setPriority0(frame *rtda.Frame) {

}

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *rtda.Frame) {

	classLoader := frame.GetMethod().GetOwner().GetClassLoader()

	currentThread := frame.GetThread()

	if t := currentThread.GetJavaMirror(); t != nil {

		frame.OpStack().PushRef(t)
	}

	threadKlass := classLoader.LoadClass("java/lang/Thread")
	threadOop := threadKlass.CreateObject()

	threadGroupClass := classLoader.LoadClass("java/lang/ThreadGroup")
	threadGroupOop := threadGroupClass.CreateObject()

	threadOop.SetRefVar("group", "Ljava/lang/ThreadGroup;", threadGroupOop)
	threadOop.SetIntVar("priority", "I", 1)
	frame.OpStack().PushRef(threadOop)

}
