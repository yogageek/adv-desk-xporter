package main

import (
	"fmt"
	"net/http"
	logic "porter/pkg/logic/gqlclient"
	"porter/routers"
)

func init() {
	logic.RefreshToken()
	logic.InitGqlClientAndToken()
}

func main() {
	// logic.Export()
	// logic.Import()
	startServer()
}

func startServer() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
		// ReadTimeout:    ReadTimeout,
		// WriteTimeout:   WriteTimeout,
		// MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

//接下來掛server 3支api

//先測匯出是否可以導出一個檔案

//在測匯入是否可以選擇一個檔案

//最後做取完成狀態率( long polling API，可以讓所有使用者能夠即時知道現在是否有匯入匯出的工作正在做。)
