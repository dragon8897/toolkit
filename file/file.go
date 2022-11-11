package file

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Exist(f string) (exist bool) {
	_, err := os.Stat(f)
	if err != nil {
		return
	}
	exist = true
	return
}

func MD5(f string) (m string, err error) {
	if !Exist(f) {
		err = fmt.Errorf("%s is not exist!", f)
		return
	}
	file, err := os.Open(f)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	m = hex.EncodeToString(hashInBytes)
	return
}

func NameNoExt(f string) (n string) {
	ext := filepath.Ext(f)
	return f[:len(f)-len(ext)]
}

func BaseNoExt(f string) (n string) {
	ext := filepath.Ext(f)
	return filepath.Base(f[:len(f)-len(ext)])
}
