package _1_gin_sentry

import (
	"errors"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func Recovery(client *raven.Client, onlyCrashes bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			flags := map[string]string{
				"endpoint": c.Request.RequestURI,
			}

			if r := recover(); r != nil {
				debug.PrintStack()
				rvalue := fmt.Sprint(r)

				packet := raven.NewPacket(rvalue,
					raven.NewException(errors.New(rvalue),
						raven.NewStacktrace(2, 3, nil)))
				client.Capture(packet, flags)
				c.Writer.WriteHeader(http.StatusInternalServerError)
			}

			if !onlyCrashes {
				for _, item := range c.Errors {
					packet := raven.NewPacket(item.Error(), &raven.Message{
						Message: item.Error(),
						Params: []interface{}{
							item.Meta,
						},
					})
					client.Capture(packet, flags)
				}
			}
		}()
		c.Next()
	}
}
