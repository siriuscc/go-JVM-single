package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	appClasspath  Entry
}

func ParseClasspath(jreOption string, cpOption string) *Classpath {

	cp := &Classpath{}

	cp.parseJREClassPath(jreOption)
	cp.parseAppClasspath(cpOption)

	return cp
}

func (receiver *Classpath) parseJREClassPath(jreOption string) {
	jrePath := getJREPath(jreOption)
	receiver.bootClasspath = newEntry(filepath.Join(jrePath, "lib", "*"))
	receiver.extClasspath = newEntry(filepath.Join(jrePath, "lib", "ext", "*"))
}

// 获取jre的真实路径
func getJREPath(jreOption string) string {

	if jreOption != "" && dirExist(jreOption) {
		return jreOption
	} else if dirExist(".jre") {
		return "./jre"
	} else {
		javaHome := os.Getenv("JAVA_HOME")

		if javaHome == "" {
			panic("Can't find jre path")
		}
		return filepath.Join(javaHome, "jre")
	}
}

// 路径是否存在
func dirExist(path string) bool {

	if _, e := os.Stat(path); e != nil {
		if os.IsNotExist(e) {
			return false
		}
		panic("dirExist:: error")
	}
	return true
}

func (receiver *Classpath) parseAppClasspath(cpOption string) {

	if cpOption == "" {
		cpOption = "."
	}

	receiver.appClasspath = newEntry(cpOption)
}

func (receiver *Classpath) ReadClass(className string) ([]byte, Entry, error) {

	className = className + ".class"
	//logger.Println("className:" + className)

	// 双亲委派的实现
	if bytes, entry, e := receiver.bootClasspath.readClass(className); e == nil {

		//fmt.Println("debug: boot load:" + className)
		return bytes, entry, e
	}
	if bytes, entry, e := receiver.extClasspath.readClass(className); e == nil {
		//fmt.Println("debug: ext load:" + className)

		return bytes, entry, e
	}

	return receiver.appClasspath.readClass(className)
}

func (receiver *Classpath) GetPath() string {
	return receiver.appClasspath.getPath()
}
