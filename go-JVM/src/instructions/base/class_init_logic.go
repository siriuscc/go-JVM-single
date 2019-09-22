package base

import (
	"rtda"
	"rtda/heap"
)

// 类的初始化逻辑

func InitClass(thread *rtda.Thread, klass *heap.Klass) {

	klass.Init()
	scheduleClinit(thread, klass)
	initSuperClass(thread, klass)
}

// 准备执行类的初始化
func scheduleClinit(thread *rtda.Thread, klass *heap.Klass) {

	clinit := klass.GetClinitMethod()

	if clinit != nil {
		thread.CreateFrame(clinit)
	}
}

// JVM 保证，如果一个类已经初始化了，他的所有父类都是已初始化的。
func initSuperClass(thread *rtda.Thread, klass *heap.Klass) {

	// 父类没有初始化
	if !klass.IsInterface() && klass.GetSuper() != nil && !klass.GetSuper().IsInitialized() {
		InitClass(thread, klass.GetSuper())
	}
}
