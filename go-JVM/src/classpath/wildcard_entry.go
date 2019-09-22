package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 支持 lib/*
func newWildcardEntry(path string) CompositeEntry {

	baseDir := path[:len(path)-1] //remove *
	compositeEntries := CompositeEntry{}

	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {

		//fmt.Println(path)
		if err != nil {
			return err
		} else if info.IsDir() && path != baseDir { // 跳过
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntries = append(compositeEntries, jarEntry)
		}

		return nil
	})

	return compositeEntries
}
