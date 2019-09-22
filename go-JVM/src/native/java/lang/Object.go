package lang

import (
	"native"
	"rtda"
)

func init() {
	const class = "java/lang/Object"
	native.Register(class, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(class, "clone", "()Ljava/lang/Object;", clone)
	native.Register(class, "hashCode", "()I", java_lang_Object_hashCode)
	native.Register(class, "notifyAll", "()V", notifyAll)
}

func notifyAll(frame *rtda.Frame) {

}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	oop := this.GetKlass().GetJavaMirror()
	frame.OpStack().PushRef(oop)
}

func clone(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()

	cloneable := this.GetKlass().GetClassLoader().LoadClass("java/lang/Cloneable")

	if !this.GetKlass().IsImplementsOf(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OpStack().PushRef(this.Clone())
}

//  public native int hashCode();
func java_lang_Object_hashCode(frame *rtda.Frame) {

	this := frame.LocalVars().GetThis()
	frame.OpStack().PushInt(this.GetHashCode())
}
