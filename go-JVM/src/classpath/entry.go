package classpath

import (
	"os"
	"strings"
)

const pathListSep = string(os.PathListSeparator)

type Entry interface {
	/**
	 * 读取文件为二进制流
	 */
	readClass(className string) ([]byte, Entry, error)
	getPath() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSep) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") ||
		strings.HasSuffix(path, ".ZIP") {

		return newZipEntry(path)
	}

	return newDirEntry(path)
}
