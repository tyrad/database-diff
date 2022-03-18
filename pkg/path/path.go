package path

import (
	"db-diff/pkg/env"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GenAppPath(str string) string {
	if env.IsDebug() {
		return str
	}
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index] + "/" + str
}
