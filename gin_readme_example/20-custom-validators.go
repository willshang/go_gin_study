package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

// Booking contains binded and validated data.
// Booking包含绑定和验证的数据
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"` // 大于CheckIn字段
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	// http://localhost:8085/bookable?check_in=2020-04-16&check_out=2020-04-17 // 大于当前时间,验证通过
	// http://localhost:8085/bookable?check_in=2018-03-08&check_out=2018-03-09 // 验证失败
	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Booking dates are valid!",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "getBookable",
		})
	}
	return
}

func bookableDate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		// 今天大于指定日期，验证失败
		if today.After(date) {
			return false
		}
	}
	return true
}
