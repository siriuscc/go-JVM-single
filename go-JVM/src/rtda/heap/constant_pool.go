package heap

import "classfile"

// 运行时常量
type Constant interface {
}

type ConstantPool struct {
	klass  *Klass
	consts []Constant
}

func (recv *ConstantPool) GetConstant(index uint) Constant {

	if recv.consts[index] != nil {
		return recv.consts[index]
	}
	return nil
}

func (recv *Klass) loadConstantPool(fileCP classfile.ConstantPool) *ConstantPool {

	poolSize := len(fileCP)
	consts := make([]Constant, poolSize)
	cp := &ConstantPool{klass: recv, consts: consts}

	for i, c := range fileCP {
		switch c.(type) {
		case *classfile.ConstantIntegerInfo:
			info := c.(*classfile.ConstantIntegerInfo)
			consts[i] = info.GetValue()
			break
		case *classfile.ConstantFloatInfo:
			info := c.(*classfile.ConstantFloatInfo)
			consts[i] = info.GetValue()
			break
		case *classfile.ConstantDoubleInfo:
			info := c.(*classfile.ConstantDoubleInfo)
			consts[i] = info.GetValue()
			i++
			break
		case *classfile.ConstantLongInfo:
			info := c.(*classfile.ConstantLongInfo)
			consts[i] = info.GetValue()
			i++
			break
		case *classfile.ConstantStringInfo:
			info := c.(*classfile.ConstantStringInfo)
			consts[i] = info.GetString()
			break

		case *classfile.ConstantClassInfo:
			ref := &ClassRef{}
			consts[i] = ref.init(cp, fileCP, c)
			break

		case *classfile.ConstantFieldRefInfo:
			ref := &FieldRef{}
			consts[i] = ref.init(cp, fileCP, c)
			break
		case *classfile.ConstantMethodRefInfo:
			ref := &MethodRef{}
			consts[i] = ref.init(cp, fileCP, c)
			break
		case *classfile.ConstantInterfaceMethodRefInfo:
			ref := &InterfaceMethodRef{}
			consts[i] = ref.init(cp, fileCP, c)
			break
		}
	}

	recv.constantPool = cp
	return cp
}
