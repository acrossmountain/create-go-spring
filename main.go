package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// 1. 获取命令行参数
	// 1.1 设置提示用法
	// 1.2 判断是否有参数传入
	// 1.3 显示提示帮助
	// 2. 获取所有参数
	// 2.1 获取详细命令
	app := &cli.App{
		Name:  "create-go-spring",
		Usage: "create-go-spring -m 'github.com/go-spring/create-go-spring' [ web http-rpc redis gorm ... ]",
	}
	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
