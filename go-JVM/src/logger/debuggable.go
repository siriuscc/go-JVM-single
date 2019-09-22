package logger

import (
	"fmt"
)

// 可调试的
type Debuggable interface {
	ToString() string
}

var DEBUG = bool(false)

func Debug(debuggable Debuggable) {

	if DEBUG == true {
		fmt.Println(debuggable.ToString())
	}
}

func Debugs(tag string, a ...Debuggable) {

	if DEBUG == true {
		fmt.Print(tag)
		for _, i := range a {
			fmt.Print(i.ToString())
		}
		fmt.Println("")
	}

}

func Println(a ...interface{}) (n int, err error) {
	if DEBUG == true {
		return fmt.Println(a...)
	}
	return len(a), nil
}

func Printf(format string, a ...interface{}) {
	if DEBUG == true {
		fmt.Printf(format, a...)
	}
}
