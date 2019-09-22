package rtda

import (
	"fmt"
	"rtda/heap"
)

type Slot struct {
	num int32
	ref *heap.OopDesc
}

func (recv *Slot) ToString() string {

	if recv.ref != nil {

		return fmt.Sprintf("(%s,%s)", recv.ref.GetKlass().GetName(), recv.ref.GetMetaData().GetName())
	}
	return string(recv.num)
}
