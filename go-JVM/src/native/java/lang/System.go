package lang

import (
	"instructions/base"
	"native"
	"rtda"
	"rtda/heap"
	"runtime"
	"time"
)

func init() {

	const class = "java/lang/System"
	native.Register(class, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register(class, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.Register(class, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(class, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	native.Register(class, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)

	native.Register(class, "currentTimeMillis", "()J", currentTimeMillis)

	native.Register(class, "mapLibraryName", "(Ljava/lang/String;)Ljava/lang/String;", mapLibraryName)
	native.Register(class, "identityHashCode", "(Ljava/lang/Object;)I", java_lang_Object_hashCode)

}

//(Ljava/lang/String;)Ljava/lang/String;
//public static native String mapLibraryName(String libname);
func mapLibraryName(frame *rtda.Frame) {

	locals := frame.LocalVars()
	libName := locals.GetRef(0)

	frame.OpStack().PushRef(libName)
}

func currentTimeMillis(frame *rtda.Frame) {

	millis := time.Now().UnixNano() / int64(time.Millisecond)
	stack := frame.OpStack()
	stack.PushLong(millis)
}

func setIn0(frame *rtda.Frame) {

	inObj := frame.LocalVars().GetRef(0)

	sysClass := frame.GetMethod().GetOwner()
	sysClass.SetStaticFieldRef("in", "Ljava/io/InputStream;", inObj)
}

//private static native void setOut0(PrintStream out);
func setOut0(frame *rtda.Frame) {
	out := frame.LocalVars().GetRef(0)
	sysClass := frame.GetMethod().GetOwner()
	sysClass.SetStaticFieldRef("out", "Ljava/io/PrintStream;", out)
}

//private static native void setErr0(PrintStream err);
func setErr0(frame *rtda.Frame) {
	errObj := frame.LocalVars().GetRef(0)
	sysClass := frame.GetMethod().GetOwner()
	sysClass.SetStaticFieldRef("err", "Ljava/io/PrintStream;", errObj)
}

// src[srcPos,srcPos+length-1] to dest[destPos,destPos+length-1]
//   当src和dest是同一个数组时，需要先将src[srcPos,srcPos+length-1] copy 到缓存数组
//    if dest或者src 是 null, 抛出NullPointerException
//    符合下列条件抛出ArrayStoreException，不改变内容
//      src 不是数组，或者dest 不是数组
//      src 的组件类型和dest 的组件类型 是不同的基本类型
// 	    一个是基本类型，一个是ref类型
//    符合下列条件之一，IndexOutOfBoundsException，不改变内容
//      srcPos <0 or descPos <0
//      length <0
//		srcPos+length > len(src) or destPos+length >len(dest)
//    如果对于每个s（item 类型），s 无法转为d，那么抛出 ArrayStoreException，比如在k位置发生错误，那么前面的都已经复制过去了
//public static native void arraycopy(Object src,  int  srcPos, Object dest, int destPos, int length);

func arraycopy(frame *rtda.Frame) {

	locals := frame.LocalVars()
	src := locals.GetRef(0)
	srcPos := locals.GetInt(1)
	dest := locals.GetRef(2)
	destPos := locals.GetInt(3)
	length := locals.GetInt(4)

	// 检查
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	if !src.GetKlass().IsArray() || !dest.GetKlass().IsArray() {
		panic("java.lang.ArrayStoreException")
	}

	s := src.GetKlass().ComponentClass()
	d := dest.GetKlass().ComponentClass()

	if s.IsPrimitive() && !d.IsPrimitive() || !s.IsPrimitive() && d.IsPrimitive() {
		panic("java.lang.ArrayStoreException")
	} else if !s.IsPrimitive() && !d.IsPrimitive() && s != d && !s.IsSubClassOf(d) {
		panic("java.lang.ArrayStoreException")
	}

	if srcPos < 0 || srcPos+length > src.ArrayLength() ||
		destPos < 0 || destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, srcPos, dest, destPos, length)
}

var sysProps = map[string]string{
	"java.version":         "1.8.0",
	"java.vendor":          "jvm.go",
	"java.vendor.url":      "https://github.com/zxh0/jvm.go",
	"java.home":            "todo",
	"java.class.version":   "52.0",
	"java.class.path":      "todo",
	"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
	"os.name":              runtime.GOOS,   // todo
	"os.arch":              runtime.GOARCH, // todo
	"os.version":           "",             // todo
	"file.separator":       "/",            // todo os.PathSeparator
	"path.separator":       ":",            // todo os.PathListSeparator
	"line.separator":       "\n",           // todo
	"user.name":            "",             // todo
	"user.home":            "",             // todo
	"user.dir":             ".",            // todo
	"user.country":         "CN",           // todo
	"file.encoding":        "UTF-8",
	"sun.stdout.encoding":  "UTF-8",
	"sun.stderr.encoding":  "UTF-8",
}

//  private static native Properties initProperties(Properties props);

// 1. 先得到private static Properties props;
// 2. 调用 props的setProperty 方法：   public synchronized Object setProperty(String key, String value) {
//
func initProperties(frame *rtda.Frame) {

	sysClass := frame.GetMethod().GetOwner()

	propsRef := sysClass.GetStaticFieldRef("props", "Ljava/util/Properties;")
	propsClass := sysClass.GetClassLoader().LoadClass("java/util/Properties")
	//public synchronized Object setProperty(String key, String value)
	setPropMethod := propsClass.GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")

	thread := frame.GetThread()

	for k, v := range sysProps {

		key := heap.JString(sysClass.GetClassLoader(), k)
		val := heap.JString(sysClass.GetClassLoader(), v)

		driverFrame := thread.CreateDriverFrame(3)

		driverFrame.OpStack().PushRef(propsRef) // this, key,value
		driverFrame.OpStack().PushRef(key)
		driverFrame.OpStack().PushRef(val)

		base.InvokeAMethod(driverFrame, setPropMethod)
	}

	frame.OpStack().PushRef(propsRef)

}
