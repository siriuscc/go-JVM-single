package heap

import "classfile"

type ClassMember struct {
	name  string
	owner *Klass

	annotationData []byte // RuntimeVisibleAnnotations_attribute
	accessFlags    *classfile.AccessFlags
	descriptor     string

	signatures string
}

func (recv *ClassMember) GetSignature() string                   { return recv.signatures }
func (recv *ClassMember) GetAnnotationData() []byte              { return recv.annotationData }
func (recv *ClassMember) GetName() string                        { return recv.name }
func (recv *ClassMember) GetOwner() *Klass                       { return recv.owner }
func (recv *ClassMember) GetAccessFlags() *classfile.AccessFlags { return recv.accessFlags }
func (recv *ClassMember) GetDescriptor() string                  { return recv.descriptor }

func (recv *ClassMember) isAccessibleTo(d *Klass) bool {

	// 如果字段是public，任何 类都可以访问
	if recv.accessFlags.IsPublic() {
		return true
	}
	c := recv.owner

	if recv.accessFlags.IsProtected() {
		// 子类和同一个包下的类可以访问
		return d == c || d.IsSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}

	if !recv.accessFlags.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}

	return c == d
}
