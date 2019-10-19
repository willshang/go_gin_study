package dto

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"go_gin_study/lesson/Gin入门实战/demo/public"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type InStruct struct {
	Name     string `form:"name" validate:"required" json:"name" `
	Age      int64  `form:"age" validate:"required" json:"age"`
	Password string `form:"password" validate:"required" json:"password"`
}

func (i *InStruct) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(i); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := make([]string, 0)
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
