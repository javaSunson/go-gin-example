package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/javaSunson/go-gin-example/pkg/logging"
	"go-gin-example/models"
	"go-gin-example/pkg/gredis"
	"go-gin-example/routers"
	"log"
	"syscall"

	"github.com/javaSunson/go-gin-example/pkg/setting"
)

func main() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
