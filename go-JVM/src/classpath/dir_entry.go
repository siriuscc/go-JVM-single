package classpath

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//go语言不需要显式实现接口，只要方法匹配即可。没有专门的构造函数，统一使用newClassName函数来做

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {

	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}

}

func (receiver *DirEntry) readClass(className string) ([]byte, Entry, error) {

	filepath := filepath.Join(receiver.absDir, className)
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		panic(fmt.Sprintf("Error: can't find main class %s", className))
	}
	return data, receiver, err
}

func (receiver *DirEntry) getPath() string {
	return receiver.absDir
}
