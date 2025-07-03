package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExecPath() (workDir string) {
	exePath, _ := os.Executable()
	if strings.HasPrefix(exePath, os.TempDir()) {
		// go run 模式
		workDir, _ = os.Getwd()
	} else {
		// 二进制模式
		workDir = filepath.Dir(exePath)
	}
	return
}
