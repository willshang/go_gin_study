package main

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"net/http"
)

type Person struct {
	Agt     int    `form:"age" validate:"required,gt=10"`
	Name    string `form:"name" validate:"require"`
	Address string `form:"address" validate:"required"`
}

var (
	Uni      *ut.UniversalTrasnlator
	Validate *validator.Validate
)

func main() {
	Validate := validator.New()

	zh := zh2.New()
	en := en2.New()
	Uni := ut.New(zh, en)

	r := gin.Default()

	r.GET("/testing", func(c *gin.Context) {
		locate := c.DefaultQuery("locale", "zh")
		trans, _ := Uni.GetTranslator(locate)
		switch locate {
		case "zh":
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		case "en":
			en_translations.RegisterDefaultTranslations(Validate, trans)
		default:
			zh_translations.RegisterDefaultTranslations(Validate, trans)
		}

		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			c.Abort()
			return
		}
		if err := Validate.Struct(person); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := make([]string, 0)
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(500, "%v", sliceErrs)
			c.Abort()
			return
		}
		c.String(http.StatusOK, "")
	})
}
