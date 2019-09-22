package heap

type SymRef struct {
	cp        *ConstantPool
	className string // 类的全限定名称： java/lang/Object
	klass     *Klass
}

func (recv *SymRef) ResolvedClass() *Klass {
	if recv.klass == nil {
		recv.resolveClassRef()
	}

	return recv.klass
}

func (recv *SymRef) resolveClassRef() {

	d := recv.cp.klass // 触发的类
	c := d.classLoader.LoadClass(recv.className)

	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	recv.klass = c

}

func (recv *SymRef) GetClassName() string {
	return recv.className
}

func (recv *SymRef) GetKlass() *Klass {
	return recv.klass
}
