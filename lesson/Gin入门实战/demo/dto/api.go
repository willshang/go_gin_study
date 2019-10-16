package dto

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"go_gin_study/lesson/Gin入门实战/demo/public"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type LoginInput struct {
	Username string `form:"username" json:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}

func (l *LoginInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(l); err != nil {
		return err
	}

	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(l)
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

type ListPageInput struct {
	Page string `form:"page" json:"page" validate:"required"`
	Name string `form:"name" json:"name" validate:""`
}

func (l *ListPageInput) BindingVaildParams(c *gin.Context) error {
	if err := c.ShouldBind(l); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(l)
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

type AddUserInput struct {
	Name  string `form:"name" json:"name" validate:"required"`
	Sex   int    `form:"sex" json:"sex" validate:""`
	Age   int    `form:"age" json:"age" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" validate:"required"`
	Addr  string `form:"addr" json:"addr" validate:"required"`
}

func (a *AddUserInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(a); err != nil {
		return err
	}

	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(a)
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

type EditUserInput struct {
	ID    int    `form:"id" json:"id" validate:"required"`
	Name  string `form:"name" json:"name" validate:"required"`
	Sex   int    `form:"sex" json:"sex" validate:""`
	Age   int    `form:"age" json:"age" validate:"required,gt=10"`
	Birth string `form:"birth" json:"birth" validate:"required"`
	Addr  string `form:"addr" json:"addr" validate:"required"`
}

func (e *EditUserInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(e); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(e)
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

type RemoveUserInput struct {
	IDS string `form:"ids" json:"ids" validate:"required"`
}

func (r *RemoveUserInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(r); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(r)
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
