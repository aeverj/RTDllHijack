package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CopyFile(dstName, srcName string) (writeen int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer src.Close()
	err = os.MkdirAll(filepath.Dir(dstName), os.ModePerm)
	if err != nil {
		return 0, err
	}
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Mlogger.Error(err.Error())
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func Contains(arr []string, str string) bool {
	for _, v := range arr {
		if strings.ToLower(str) == strings.ToLower(v) {
			return true
		}
	}
	return false
}
