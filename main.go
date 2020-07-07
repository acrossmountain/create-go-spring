package main

import (
	"time"

	"github.com/go-spring/go-spring-parent/spring-logger"
)

func main() {
	SpringLogger.Info("create-go-spring => " + time.Now().Format("2006-01-02 15:04:05"))
	// 1. 获取命令行参数
	// 1.1 设置提示用法
	// 1.2 判断是否有参数传入
	// 1.3 显示提示帮助
	// 2. 获取所有参数
	// 2.1 获取详细命令
	
}