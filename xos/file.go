package xos

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	MODE_DIR  = 0
	MODE_FILE = 1
)

func ListSubFiles(path string, mode int) ([]string, error) {
	var r []string
	d, err := ioutil.ReadDir(path)
	if err != nil {
		return r, err
	}
	for _, d := range d {
		if mode == MODE_DIR {
			if d.IsDir() {
				r = append(r, d.Name())
			}
		}
		if mode == MODE_FILE {
			if !d.IsDir() {
				r = append(r, d.Name())
			}
		}
	}
	return r, nil
}

func ListSubFilesRecur(path string) (files []string, err error) {
	files = make([]string, 0, 30)
	err = filepath.Walk(path, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() { // 忽略目录
			return nil
		}
		files = append(files, filename)
		return nil
	})
	return files, err
}
