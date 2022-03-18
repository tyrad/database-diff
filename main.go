package main

import (
	"db-diff/model"
	"db-diff/pkg/db"
	"db-diff/pkg/env"
	"db-diff/pkg/logging"
	"db-diff/route"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

func DebugDB() {
	dbConfig := model.DbConfig{
		Host:     "10.67.78.133",
		User:     "phpg_nbphs_read",
		Password: "A4hPa8ECuUD7rSmFGNGFdmrU",
		DbName:   "chisapp",
		Port:     5432,
	}
	enable, err := db.CheckConnectEnable(dbConfig)
	fmt.Println("数据库连接是否可用:", enable, err)

	d2, err := db.Connect(dbConfig)
	row := db.QueryRow(d2, "select * from phpg_nbphs.sys_config where config_name = $1", "enableMentalIllnessFaceReg")
	print(row)
}

// diff 功能
func main() {

	var err error

	fmt.Println("---\n\n\n")

	DebugDB()
	fmt.Println("---\n\n\n")
	port := ":8001"

	// 生产环境使用随机端口
	if env.IsDebug() == false {
		pt, err := GetFreePort()
		if err == nil {
			port = fmt.Sprintf(":%v", pt)
		}
	}

	logging.SetUp()
	router := route.InitRouter()

	go func() {
		if env.IsDebug() == false {
			timeAfterTrigger := time.After(time.Second * 1)
			curTime, _ := <-timeAfterTrigger
			fmt.Println(curTime.Format("2006-01-02 15:04:05"))
			err = Open("http://localhost" + port + "/manager")
			fmt.Println("请访问: http://localhost" + port + "/manager")
			if err != nil {
				fmt.Println("open browser err:", err)
			}
		}
	}()

	err = router.Run(port)
	if err != nil {
		logging.Info("Sever err: %v", err)
	}
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	if runtime.GOOS == "windows" {
		cmd := exec.Command(`cmd`, `/c`, `start`, uri)
		return cmd.Start()
	} else {
		cmd := exec.Command(run, uri)
		return cmd.Start()
	}
}

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
