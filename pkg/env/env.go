package env

import "os"

// IsDebug 判断当前环境是否是debug环境， 当debug是环境变量添加 DEBUG=1
func IsDebug() bool {
	return os.Getenv("DEBUG") != ""
}
