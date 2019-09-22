package heap

import "unicode/utf16"

// 字符池
var internStrings = map[string]*OopDesc{}

func JString(loader *ClassLoader, goStr string) *OopDesc {

	if strObj, ok := internStrings[goStr]; ok {
		return strObj
	}

	utf8Chars := stringToUTF16(goStr)

	charArr := &OopDesc{oopType: loader.LoadClass("[C"), vtable: utf8Chars}
	strObj := loader.LoadClass("java/lang/String").CreateObject()
	// 这里没有初始化类，应该是错误的

	strObj.SetRefVar("value", "[C", charArr)

	internStrings[goStr] = strObj
	return strObj
}

func stringToUTF16(str string) []uint16 {

	runes := []rune(str)
	return utf16.Encode(runes)
}

func GoString(strObj *OopDesc) string {

	charsObj := strObj.GetRefVar("value", "[C")

	return utf16ToString(charsObj.Chars())
}

func utf16ToString(chars []uint16) string {

	return string(utf16.Decode(chars))
}

func InternString(this *OopDesc) *OopDesc {

	s := GoString(this)
	if strObj, ok := internStrings[s]; ok {
		return strObj
	}
	internStrings[s] = this

	return this
}
