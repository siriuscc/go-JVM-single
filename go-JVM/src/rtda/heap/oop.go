package heap

type OopDesc struct {
	oopType  *Klass 	 // 实例自身属于什么类
	metaData *Klass
	fields   Slots
	vtable   interface{} // 数组的 item 是一个数组，如 []int
	ext      interface{} // 扩展数据位
	hashCode int32       // 不等于 0
}

type JavaMirror OopDesc

func (recv *OopDesc) SetExt(data interface{}) {
	recv.ext = data
}

func (recv *OopDesc) GetExt() interface{} {
	return recv.ext
}

func (recv *OopDesc) GetFields() Slots {

	return recv.fields
}

func (recv *OopDesc) GetKlass() *Klass {
	return recv.oopType
}

func (recv *OopDesc) IsInstanceOf(klass *Klass) bool {

	return recv.oopType.IsAssignableTo(klass)
}

func (recv *OopDesc) SetRefVar(name string, descriptor string, ref *OopDesc) {

	// 得到对应的属性
	field := recv.oopType.GetInstanceField(name, descriptor)
	recv.fields.SetRef(field.slotId, ref)
}

func (recv *OopDesc) GetRefVar(name string, descriptor string) *OopDesc {

	// 得到对应的属性
	field := recv.oopType.GetInstanceField(name, descriptor)
	return recv.fields.GetRef(field.slotId)
}

func (recv *OopDesc) GetMetaData() *Klass {
	return recv.metaData
}

func (recv *OopDesc) GetByteTable() []int8 {
	return recv.vtable.([]int8)
}

func (recv *OopDesc) GetShortTable() []int16 {
	return recv.vtable.([]int16)
}

func (recv *OopDesc) GetCharTable() []uint16 {
	return recv.vtable.([]uint16)
}

func (recv *OopDesc) GetIntTable() []int32 {
	return recv.vtable.([]int32)
}
func (recv *OopDesc) GetFloatTable() []float32 {
	return recv.vtable.([]float32)
}
func (recv *OopDesc) GetDoubleTable() []float64 {
	return recv.vtable.([]float64)
}
func (recv *OopDesc) GetLongTable() []int64 {
	return recv.vtable.([]int64)
}
func (recv *OopDesc) GetRefTable() []*OopDesc {
	return recv.vtable.([]*OopDesc)
}

// copy
func (recv *OopDesc) Clone() *OopDesc {
	return &OopDesc{
		oopType:  recv.oopType,
		metaData: recv.metaData,
		vtable:   recv.CloneVtable(),
		fields:   recv.CloneFields(),
	}

}

func (recv *OopDesc) GetHashCode() int32 {
	if recv == nil {
		return 0
	}
	return recv.hashCode
}

func (recv *OopDesc) CloneVtable() interface{} {

	switch recv.vtable.(type) {
	case []int8:
		table := recv.GetByteTable()
		newtTable := make([]int8, len(table))
		copy(newtTable, table)
		return newtTable
	case []int16:
		table := recv.GetShortTable()
		newtTable := make([]int16, len(table))
		copy(newtTable, table)
		return newtTable
	case []uint16:
		table := recv.GetCharTable()
		newtTable := make([]uint16, len(table))
		copy(newtTable, table)
		return newtTable
	case []int32:
		table := recv.GetIntTable()
		newtTable := make([]int32, len(table))
		copy(newtTable, table)
		return newtTable
	case []int64:
		table := recv.GetLongTable()
		newtTable := make([]int64, len(table))
		copy(newtTable, table)
		return newtTable
	case []float32:
		table := recv.GetFloatTable()
		newtTable := make([]float32, len(table))
		copy(newtTable, table)
		return newtTable
	case []*OopDesc:
		table := recv.GetRefTable()
		newtTable := make([]*OopDesc, len(table))
		copy(newtTable, table)
		return newtTable
	case nil:
		return nil
	default:
		panic("error, vtable type don't know")

	}
}

func (recv *OopDesc) CloneFields() Slots {

	newSlots := createSlots(uint(len(recv.fields)))
	copy(newSlots, recv.fields)

	return newSlots
}

// 获取实例属性的 int value
func (recv *OopDesc) GetIntField(name, desc string) int32 {

	field := recv.GetKlass().GetInstanceField(name, desc)

	return recv.fields.GetInt(field.slotId)
}

func (recv *OopDesc) SetVTable(vtable interface{}) {
	recv.vtable = vtable
}

func (recv *OopDesc) GetObjectField(name, desc string) *OopDesc {

	field := recv.GetKlass().GetInstanceField(name, desc)
	return recv.fields.GetRef(field.slotId)
}

func (recv *OopDesc) SetIntVar(name string, desc string, value int) {

	field := recv.GetKlass().GetInstanceField(name, desc)
	recv.fields.SetInt(field.slotId, int32(value))
}

func ArrayCopy(src *OopDesc, srcPos int32, dest *OopDesc, destPos int32, length int32) {

	switch src.vtable.(type) {

	case []int8:
		tmp := src.vtable.([]int8)[srcPos : srcPos+length]
		dst := dest.vtable.([]int8)[destPos : destPos+length]
		copy(dst, tmp)
	case []int16:
		tmp := src.vtable.([]int16)[srcPos : srcPos+length]
		dst := dest.vtable.([]int16)[destPos : destPos+length]
		copy(dst, tmp)
	case []uint16:
		tmp := src.vtable.([]uint16)[srcPos : srcPos+length]
		dst := dest.vtable.([]uint16)[destPos : destPos+length]
		copy(dst, tmp)
	case []int32:
		tmp := src.vtable.([]int32)[srcPos : srcPos+length]
		dst := dest.vtable.([]int32)[destPos : destPos+length]
		copy(dst, tmp)
	case []int64:
		tmp := src.vtable.([]int64)[srcPos : srcPos+length]
		dst := dest.vtable.([]int64)[destPos : destPos+length]
		copy(dst, tmp)
	case []float32:
		tmp := src.vtable.([]float32)[srcPos : srcPos+length]
		dst := dest.vtable.([]float32)[destPos : destPos+length]
		copy(dst, tmp)
	case []float64:
		tmp := src.vtable.([]float64)[srcPos : srcPos+length]
		dst := dest.vtable.([]float64)[destPos : destPos+length]
		copy(dst, tmp)
	case []*OopDesc:
		tmp := src.vtable.([]*OopDesc)[srcPos : srcPos+length]
		dst := dest.vtable.([]*OopDesc)[destPos : destPos+length]
		copy(dst, tmp)
	}
}
