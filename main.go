package main

import (
	"easymarket/inits"
	"easymarket/inits/mysqlinit"
	"easymarket/utils"
	"fmt"
	"net/http"
	"time"
)

func main() {
	Router := inits.InitRouter()  // 初始化路由
	db, err := mysqlinit.InitDb() // 初始化数据库连接
	defer db.Close()
	if err != nil {

	}
	utils.InitLog() // 初始化日志
	utils.QMLog.Info("服务器开启")
	s := &http.Server{
		Addr:           ":8887",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`
欢迎使用 easymarket-go-server
作者:peroluo
github:https://github.com/Peroluo/easymarket-go-server
`)
	_ = s.ListenAndServe()
}
