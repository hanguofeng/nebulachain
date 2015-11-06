package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type FileWalker struct {
	entrypoint string
}

type FileCallback func(fullpath string, shortpath string, fingerprint string)

func (this *FileWalker) Walk(fileCallback FileCallback) {
	filepath.Walk(this.entrypoint, func(nodeName string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			shortName := strings.TrimPrefix(nodeName, this.entrypoint+"/")
			fingerPrint, _ := this.CalcFileFingerPrint(nodeName)
			fileCallback(nodeName, shortName, fingerPrint)
		}

		return nil
	})
}

func (this *FileWalker) CalcFileFingerPrint(file string) (string, error) {
	f, err := os.Open(file)

	if err == nil {
		dig := md5.New()
		io.Copy(dig, f)
		sum := dig.Sum(nil)
		return fmt.Sprintf("01|%x", sum), nil
	} else {
		return "", err
	}
}
