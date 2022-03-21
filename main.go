package main

import (
	"db-diff/pkg/env"
	"db-diff/pkg/logging"
	"db-diff/route"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// diff 功能
func main() {

	var err error
	fmt.Println("---\n\n\n")
	port := ":8001"

	// 生产环境使用随机端口
	if env.IsDebug() == false {
		pt, err := env.GetFreePort()
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
