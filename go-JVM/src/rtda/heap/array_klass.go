package heap

func (recv *Klass) CreateBaseArrayClass(count int32) *OopDesc {

	if !recv.IsArray() {
		panic("Not array class:" + recv.name)
	}

	switch recv.name {
	case "[Z":
		return &OopDesc{oopType: recv, vtable: make([]int8, count)}
	case "[B":
		return &OopDesc{oopType: recv, vtable: make([]int8, count)}
	case "[C":
		return &OopDesc{oopType: recv, vtable: make([]uint16, count)}
	case "[S":
		return &OopDesc{oopType: recv, vtable: make([]int16, count)}
	case "[I":
		return &OopDesc{oopType: recv, vtable: make([]int32, count)}
	case "[J":
		return &OopDesc{oopType: recv, vtable: make([]int64, count)}
	case "[F":
		return &OopDesc{oopType: recv, vtable: make([]float32, count)}
	case "[D":
		return &OopDesc{oopType: recv, vtable: make([]float64, count)}

	default:
		return &OopDesc{oopType: recv, vtable: make([]*OopDesc, count)}
	}
}

func (recv *Klass) IsArray() bool {

	return recv.name[0] == '['
}

func (recv *Klass) CreateArrayObject(count int32) *OopDesc {

	// [I
	arrName := getArrayClassName(recv.name)
	arrClass := recv.classLoader.LoadClass(arrName)

	oop := arrClass.CreateBaseArrayClass(count)
	return oop
}

// 创建多维数组，
func (recv *Klass) CreateMultiArrayObject(counts []int32) *OopDesc {

	arrObj := recv.CreateArrayObject(counts[0])

	if len(counts) > 1 {

		refs := arrObj.Refs()
		for i := range refs {
			refs[i] = recv.ComponentClass().CreateMultiArrayObject(counts[1:])
		}
	}
	return arrObj
}
