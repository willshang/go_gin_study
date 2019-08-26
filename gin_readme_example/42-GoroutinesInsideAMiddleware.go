package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	// http://127.0.0.1:8080/long_async
	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		// 创建在 goroutine 中使用的副本
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	// http://127.0.0.1:8080/long_sync
	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		// 用 time.Sleep() 模拟一个长任务。
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
