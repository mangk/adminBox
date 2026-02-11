package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetExecPath() (workDir string) {
	// exePath, _ := os.Executable()
	// if strings.HasPrefix(exePath, os.TempDir()) {
	// 	// go run 模式
	// 	workDir, _ = os.Getwd()
	// } else {
	// 	// 二进制模式
	// 	workDir = filepath.Dir(exePath)
	// }

	pwd, _ := os.Getwd()
	args0 := os.Args[0]

	if filepath.Dir(args0) == pwd {
		workDir = pwd
	} else {
		if pwd == "/" {
			workDir = filepath.Dir(args0)
		} else {
			if strings.Contains(args0, os.TempDir()) || (strings.Contains(args0, "Caches") && strings.Contains(args0, "go-build")) {
				workDir = pwd
			} else {
				workDir = filepath.Dir(filepath.Join(pwd, args0))
			}
		}
	}
	return
}
