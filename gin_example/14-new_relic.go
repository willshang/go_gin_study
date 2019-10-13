package main

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent"
	"log"
	"net/http"
	"os"
)

const (
	// NewRelicTxnKey是用于从上下文检索NewRelic事务的键
	NewRelicTxnKey = "NewRelicTxnKey"
)

// NewRelicMonitoring是一个中间件，
// 它启动一个newrelic事务，将其存储在上下文中，然后调用下一个处理程序
func NewRelicMonitoring(app newrelic.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path, c.Writer, c.Request)
		defer txn.End()

		c.Set(NewRelicTxnKey, txn)
		c.Next()
	}
}

func main() {
	router := gin.Default()

	cfg := newrelic.NewConfig(os.Getenv("APP_NAME"), os.Getenv("NEW_RELIC_API_KEY"))
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		log.Printf("failed to make new_relice app: %v", err)
	} else {
		router.Use(NewRelicMonitoring(app))
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!\n")
	})
	router.Run()
}
