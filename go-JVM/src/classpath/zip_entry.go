package classpath

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

func (receiver *ZipEntry) getPath() string {
	return receiver.absPath
}

func (receiver *ZipEntry) readClass(classname string) ([]byte, Entry, error) {

	r, err := zip.OpenReader(receiver.absPath)

	if err != nil {
		return nil, nil, err
	}
	// 退出前关闭资源
	defer r.Close()

	for _, f := range r.File {
		if f.Name == classname {
			rc, err := f.Open()

			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)

			if err != nil {
				return nil, nil, err
			}
			return data, receiver, nil
		}
	}

	return nil, nil, errors.New("class no found:" + classname)
}
