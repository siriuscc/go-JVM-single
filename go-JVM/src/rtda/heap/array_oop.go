package heap

func (recv *OopDesc) Bytes() []int8      { return recv.vtable.([]int8) }
func (recv *OopDesc) Shorts() []int16    { return recv.vtable.([]int16) }
func (recv *OopDesc) Ints() []int32      { return recv.vtable.([]int32) }
func (recv *OopDesc) Chars() []uint16    { return recv.vtable.([]uint16) }
func (recv *OopDesc) Floats() []float32  { return recv.vtable.([]float32) }
func (recv *OopDesc) Doubles() []float64 { return recv.vtable.([]float64) }
func (recv *OopDesc) Longs() []int64     { return recv.vtable.([]int64) }
func (recv *OopDesc) Refs() []*OopDesc   { return recv.vtable.([]*OopDesc) }

func (recv *OopDesc) ArrayLength() int32 {

	switch recv.vtable.(type) {
	case []int8:
		return int32(len(recv.Bytes()))
	case []int16:
		return int32(len(recv.Shorts()))
	case []int32:
		return int32(len(recv.Ints()))
	case []uint16:
		return int32(len(recv.Chars()))
	case []float32:
		return int32(len(recv.Floats()))
	case []float64:
		return int32(len(recv.Doubles()))
	case []int64:
		return int32(len(recv.Longs()))
	case []*OopDesc:
		return int32(len(recv.Refs()))
	default:
		panic("Not array")
	}
}
