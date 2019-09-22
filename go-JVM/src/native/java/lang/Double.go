package lang

import (
	"math"
	"native"
	"rtda"
)

func init() {

	native.Register("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)

}

//public static native long doubleToRawLongBits(double value);
func doubleToRawLongBits(frame *rtda.Frame) {

	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OpStack().PushLong(int64(bits))
}

//public static native double longBitsToDouble(long bits);
func longBitsToDouble(frame *rtda.Frame) {

	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits))
	frame.OpStack().PushDouble(value)
}
