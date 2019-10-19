package main

import (
	"go_gin_study/lesson/Gin入门实战/common/lib"
	"go_gin_study/lesson/Gin入门实战/demo/public"
	"go_gin_study/lesson/Gin入门实战/demo/router"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	lib.InitModule("./conf/dev/", []string{
		"base",
		"mysql",
		"redis",
	})
	defer lib.Destroy()

	public.InitMysql()
	public.InitValidate()
	router.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
