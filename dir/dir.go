package dir

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
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

type ListOptions func(name string) bool

func WithExt(exts ...string) func(p string) bool {
	return func(p string) (ret bool) {
		ext := path.Ext(p)
		for _, v := range exts {
			if "."+v == ext {
				ret = true
				return
			}
		}
		return
	}
}

func WithNoExt(exts ...string) func(p string) bool {
	return func(p string) (ret bool) {
		ext := path.Ext(p)
		for _, v := range exts {
			if "."+v == ext {
				ret = false
				return
			}
		}
		ret = true
		return
	}
}

func WithReg(exp string) func(p string) bool {
	return func(p string) (ret bool) {
		re := regexp.MustCompile(exp)
		ret = re.Match([]byte(p))
		return
	}
}

func List(d string, opts ...ListOptions) (files []string) {
	if !IsDir(d) {
		return
	}
	dirs, err := ioutil.ReadDir(d)
	if err != nil {
		panic(err)
	}
	for _, file := range dirs {
		name := file.Name()
		if name == "." || name == ".." || name == "__debug_bin" {
			continue
		}
		isContinue := false
		for _, opt := range opts {
			if !opt(name) {
				isContinue = true
				break
			}
		}
		if isContinue {
			continue
		}
		files = append(files, name)
	}
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

func Exist(d string) (ret bool) {
	s, err := os.Stat(d)
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		return
	}
	ret = s.IsDir()
	return
}

func EnsureDir(d string) (created bool) {
	if Exist(d) {
		return
	}
	err := os.Mkdir(d, os.ModePerm)
	if err != nil {
		return
	}
	created = true
	return
}
