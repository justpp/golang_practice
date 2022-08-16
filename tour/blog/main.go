package main

import (
	"giao/tour/blog/global"
	"giao/tour/blog/internal/model"
	"giao/tour/blog/internal/routers"
	"giao/tour/blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Printf("init.setupSetting err:%v", err)
	}

	err = setupDBEngine()
	if err != nil {
		return
	}
}

func main() {
	log.Printf("database %v", global.DatabaseSetting)
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	r := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := r.ListenAndServe()
	if err != nil {
		return
	}
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	// 这里不能写成 :=
	// := 会声明新变量 那么全局global.DBEngine != global.DBEngine
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return err
}
