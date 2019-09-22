package classfile

// class
const (
	ACC_PUBLIC     = 0x0001
	ACC_FINAL      = 0x0010
	ACC_SUPER      = 0x0020 //Treat superclass methods specially when invoked by the invokespecial instruction.
	ACC_INTERFACE  = 0x0200 //Is an interface, not a class.
	ACC_ABSTRACT   = 0x0400 //Declared abstract; must not be instantiated.
	ACC_SYNTHETIC  = 0x1000 //Declared synthetic; not present in the source code.
	ACC_ANNOTATION = 0x2000 //Declared as an annotation type.
	ACC_ENUM       = 0x4000 // Declared as an enum type.
)

// method
const (
	//ACC_PUBLIC       = 0x0001
	ACC_PRIVATE   = 0x0002
	ACC_PROTECTED = 0x0004
	ACC_STATIC    = 0x0008
	//ACC_FINAL        = 0x0010
	ACC_SYNCHRONIZED = 0x0020 // 复用
	ACC_BRIDGE       = 0x0040 //
	ACC_VARARGS      = 0x0080
	ACC_NATIVE       = 0x0100
	//ACC_ABSTRACT     = 0x0400
	ACC_STRICT = 0x0800
	//ACC_SYNTHETIC    = 0x1000
)

//field
const (
	//ACC_PUBLIC 		=0x0001
	//ACC_PRIVATE 	=0x0002
	//ACC_PROTECTED 	=0x0004
	//ACC_STATIC 		=0x0008
	//ACC_FINAL 		=0x0010
	ACC_VOLATILE  = 0x0040 //复用
	ACC_TRANSIENT = 0x0080 //复用
	//ACC_SYNTHETIC 	=0x1000
	//ACC_ENUM 		=0x4000
)

//type AccessFlags interface {
//	validate() bool
//	logger.Debuggable
//	Readable
//}

const (
	ACCESS_CLASS       = 1
	ACCESS_METHOD      = 2
	ACCESS_FIELD       = 3
	ACCESS_INNER_CLASS = 4
)

type AccessFlags struct {
	accessType uint
	flags      uint16 // 0x0000
}

func CreateAccessFlag(accessType uint, flags ...uint16) *AccessFlags {

	access := &AccessFlags{accessType: accessType}

	for _, flag := range flags {
		access.flags |= flag
	}
	return access
}

func (recv *AccessFlags) read(reader *ClassReader) {
	recv.flags = reader.readUint16()
	recv.validate()
}

func (recv *AccessFlags) validate() bool {

	switch recv.accessType {

	case ACCESS_CLASS:
		recv.validateClass()
	case ACCESS_METHOD:
	case ACCESS_FIELD:
	case ACCESS_INNER_CLASS:

	}

	return true
}
func (recv *AccessFlags) specialToString() string {

	switch recv.accessType {

	case ACCESS_CLASS:
		return recv.classToString()
	case ACCESS_METHOD:
		return recv.methodToString()
	case ACCESS_FIELD:
		return recv.fieldToString()
	case ACCESS_INNER_CLASS:
		return recv.innerClassToString()
	}

	return ""
}

// 默认输出二进制
func (recv *AccessFlags) ToString() string {

	s := "AccessFlags: "
	if recv.flags&ACC_PUBLIC > 0 {
		s += "ACC_PUBLIC "
	}
	if recv.flags&ACC_PRIVATE > 0 {
		s += "ACC_PRIVATE "
	}
	if recv.flags&ACC_PROTECTED > 0 {
		s += "ACC_PROTECTED "
	}
	if recv.flags&ACC_STATIC > 0 {
		s += "ACC_STATIC "
	}
	if recv.flags&ACC_FINAL > 0 {
		s += "ACC_FINAL "
	}
	if recv.flags&ACC_NATIVE > 0 {
		s += "ACC_NATIVE "
	}
	if recv.flags&ACC_INTERFACE > 0 {
		s += "ACC_INTERFACE "
	}
	if recv.flags&ACC_ABSTRACT > 0 {
		s += "ACC_ABSTRACT "
	}
	if recv.flags&ACC_STRICT > 0 {
		s += "ACC_STRICT "
	}
	if recv.flags&ACC_SYNTHETIC > 0 {
		s += "ACC_SYNTHETIC "
	}
	if recv.flags&ACC_ANNOTATION > 0 {
		s += "ACC_ANNOTATION "
	}
	if recv.flags&ACC_ENUM > 0 {
		s += "ACC_ENUM "
	}
	s += recv.specialToString()

	return s
}

func (recv *AccessFlags) validateClass() bool {
	// 接口
	if recv.flags&ACC_INTERFACE != 0 {
		//如果ACC_INTERFACE，则ACC_ABSTRACT 也必须被设置。ACC_FINAL,ACC_SUPER, and ACC_ENUM 不能被设置。
		if recv.flags&ACC_ABSTRACT == 0 {
			panic("java.lang.ClassFormatEror:accessFlags")
		}
		if recv.flags&ACC_ABSTRACT != 0 && recv.flags&ACC_ENUM != 0 {
			panic("java.lang.ClassFormatEror:accessFlags")
		}
	} else { // class
		//if (recv.flags&ACC_ABSTRACT != 0) && (recv.flags&ACC_ABSTRACT != 0) {
		//	panic("java.lang.ClassFormatEror:accessFlags")
		//}
	}
	return true
}

func (recv *AccessFlags) classToString() string {

	s := ""
	if recv.flags&ACC_SUPER > 0 {
		s += "ACC_SUPER "
	}
	return s
}

func (recv *AccessFlags) methodToString() string {

	s := ""
	if recv.flags&ACC_SYNCHRONIZED > 0 {
		s += "ACC_SYNCHRONIZED "
	}
	if recv.flags&ACC_BRIDGE > 0 {
		s += "ACC_BRIDGE "
	}
	if recv.flags&ACC_VARARGS > 0 {
		s += "ACC_VARARGS "
	}
	return s
}

func (recv *AccessFlags) fieldToString() string {

	s := ""
	if recv.flags&ACC_TRANSIENT > 0 {
		s += "ACC_TRANSIENT "
	}
	if recv.flags&ACC_VOLATILE > 0 {
		s += "ACC_VOLATILE "
	}
	return s
}

func (recv *AccessFlags) innerClassToString() string {
	return ""
}

func (recv *AccessFlags) IsPublic() bool {

	return 0 != recv.flags&ACC_PUBLIC
}

func (recv *AccessFlags) IsPrivate() bool {
	return 0 != recv.flags&ACC_PRIVATE
}

func (recv *AccessFlags) IsProtected() bool {
	return 0 != recv.flags&ACC_PROTECTED
}

func (recv *AccessFlags) IsStatic() bool {
	return 0 != recv.flags&ACC_STATIC
}

func (recv *AccessFlags) IsFinal() bool {
	return 0 != recv.flags&ACC_FINAL
}

func (recv *AccessFlags) IsNative() bool {
	return 0 != recv.flags&ACC_NATIVE
}

func (recv *AccessFlags) IsSynthetic() bool {
	return 0 != recv.flags&ACC_SYNTHETIC
}

func (recv *AccessFlags) IsInterface() bool {
	return 0 != recv.flags&ACC_INTERFACE
}
func (recv *AccessFlags) IsAbstract() bool {
	return 0 != recv.flags&ACC_ABSTRACT
}

func (recv *AccessFlags) IsAnnotation() bool {
	return 0 != recv.flags&ACC_ANNOTATION
}

func (recv *AccessFlags) IsEnum() bool {
	return 0 != recv.flags&ACC_ENUM
}

func (recv *AccessFlags) IsSuper() bool {
	return recv.accessType == ACCESS_CLASS && 0 != recv.flags&ACC_SUPER
}

func (recv *AccessFlags) IsSynchronized() bool {
	return recv.accessType == ACCESS_METHOD && 0 != recv.flags&ACC_SYNCHRONIZED
}

func (recv *AccessFlags) IsVarargs() bool {
	return recv.accessType == ACCESS_METHOD && 0 != recv.flags&ACC_VARARGS
}

func (recv *AccessFlags) IsBridge() bool {
	return recv.accessType == ACCESS_METHOD && 0 != recv.flags&ACC_BRIDGE
}

func (recv *AccessFlags) IsTransient() bool {
	return recv.accessType == ACCESS_FIELD && 0 != recv.flags&ACC_TRANSIENT
}
func (recv *AccessFlags) IsVolatile() bool {
	return recv.accessType == ACCESS_FIELD && 0 != recv.flags&ACC_VOLATILE
}

func (recv *AccessFlags) GetFlags() uint16 {
	return recv.flags
}
