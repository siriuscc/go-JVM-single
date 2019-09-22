package classpath

import (
	"github.com/pkg/errors"
	"strings"
)

/**
 * 多个实体
 */
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {

	compositeEntry := []Entry{}

	for _, filepath := range strings.Split(pathList, pathListSep) {

		entry := newEntry(filepath)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {

	for _, entry := range self {

		bytes, entry, err := entry.readClass(className)
		if err == nil {
			return bytes, entry, nil
		}
	}

	return nil, nil, errors.New("class not found")
}

// return path1;path2;path3...
func (self CompositeEntry) getPath() string {

	paths := []string{}

	for _, entry := range self {

		path := entry.getPath()

		paths = append(paths, path)
	}
	return strings.Join(paths, pathListSep)
}
