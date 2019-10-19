package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `binding:"required,email"`
}

func UserStructLevelValidation(v *validator.Validate, structLevel *validator.StructLevel) {
	user := structLevel.CurrentStruct.Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		structLevel.ReportError(reflect.ValueOf(user.FirstName), "FirstName", "first_name", "first or last name")
		structLevel.ReportError(reflect.ValueOf(user.LastName), "LastName", "last_name", "first or last name")
	}
}

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, User{})
	}

	router.POST("/user", validateUser)
	router.Run()
}

// curl -s -X POST http://localhost:8080/user -H 'content-type: application/json' -d '{}'
// curl -s -X POST http://localhost:8080/user -H 'content-type: application/json' -d '{"email": "george@vandaley.com"}'
// curl -X POST http://localhost:8080/user -H 'content-type: application/json' -d '{"fname": "George", "email": "george@vandaley.com"}'
// curl -X POST http://localhost:8080/user -H 'content-type: application/json' -d '{"lname": "Contanza", "email": "george@vandaley.com"}'
// curl -X POST http://localhost:8080/user -H 'content-type: application/json' -d '{"fname": "George", "lname": "Costanza", "email": "george@vandaley.com"}'

func validateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "User validation successful.",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User validation failed",
			"error":   err.Error(),
		})
	}
}
