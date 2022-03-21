package env

import (
	"net"
	"os"
)

// IsDebug 判断当前环境是否是debug环境， 当debug是环境变量添加 DEBUG=1
func IsDebug() bool {
	return os.Getenv("DEBUG") != ""
}

// GetFreePort 获取未被占用的端口
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
