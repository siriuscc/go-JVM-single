package main

import (
	"classpath"
	"fmt"
	"instructions"
	"instructions/base"
	"rtda"
	"rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
	interpreter *instructions.Interpreter
}

func (recv *JVM) start() {

	recv.init()
	recv.exeMain()
}

func (recv *JVM) init() {

	// 注入 解释器
	recv.interpreter = &instructions.Interpreter{}
	recv.interpreter.Init(recv.cmd.verbose)

	// 加载 VM 类， 会生成多个 clinit 的frame, 执行类的初始化
	vmClass := recv.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(recv.mainThread, vmClass)
	recv.interpreter.Interpret(recv.mainThread)

}

func (recv *JVM) exeMain() {

	className := strings.Replace(recv.cmd.class, ".", "/", -1)
	mainClass := recv.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class [%s] \n", recv.cmd.class)
		return
	}

	// 创建 main 函数对应的 frame 并注入 参数
	jArgs := createArgsArray(recv.classLoader, recv.cmd.args)

	frame := recv.mainThread.CreateFrame(mainMethod)
	frame.LocalVars().SetRef(0, jArgs)

	recv.interpreter.Interpret(recv.mainThread)

}

func createJVM(cmd *Cmd) *JVM {
	classpath := classpath.ParseClasspath(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.CreateClassLoader(classpath, cmd.verbose)

	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  rtda.CreateThread(),
	}

}

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.OopDesc {

	stringClass := loader.LoadClass("java/lang/String")
	stringArrObj := stringClass.ArrayClass().CreateArrayObject(int32(len(args)))
	refs := stringArrObj.Refs()

	for i, arg := range args {
		refs[i] = heap.JString(loader, arg)
	}

	return stringArrObj
}
