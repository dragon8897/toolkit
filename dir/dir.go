package dir

import (
	"io/ioutil"
	"os"
	"path"
)

func walkDir(dir string, base string) (files []string) {
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range dirs {
		name := file.Name()
		if name == "." || name == ".." || name == "__debug_bin" {
			continue
		}
		filePath := path.Join(dir, name)
		basePath := path.Join(base, name)
		if file.IsDir() {
			files = append(files, walkDir(filePath, basePath)...)
		} else {
			files = append(files, basePath)
		}
	}
	return files
}

func Walk(d string) (files []string) {
	if !IsDir(d) {
		return
	}
	files = walkDir(d, "")
	return
}

func IsDir(d string) (ret bool) {
	s, err := os.Stat(d)
	if err != nil {
		return
	}
	ret = s.IsDir()
	return
}
