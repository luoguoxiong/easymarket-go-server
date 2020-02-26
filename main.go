package main

import (
	"easymarket/inits"
	"easymarket/utils"
	"fmt"
	"net/http"
	"time"
)

func main() {
	Router := inits.InitRouter()
	utils.InitLog()
	utils.QMLog.Info("服务器开启")
	s := &http.Server{
		Addr:           ":8888",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`
	欢迎使用 easymarket-go-server
	作者：peroluo
	github: https://github.com/Peroluo/easymarket-go-server
	默认自动化文档地址: http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址: http://127.0.0.1:8888
`, s.Addr)
	_ = s.ListenAndServe()
}
