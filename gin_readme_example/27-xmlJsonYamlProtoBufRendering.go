package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	// gin.H 是 map[string]interface{} 简化方法
	// http://127.0.0.1:8080/someJSON
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
		return
	})

	// http://127.0.0.1:8080/moreJSON
	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	// http://127.0.0.1:8080/someXML
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
		return
	})

	// http://127.0.0.1:8080/someYAML
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
		return
	})

	// http://127.0.0.1:8080/someProtoBuf
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
