package controller

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gin_study/lesson/Gin入门实战/demo/dto"
	"go_gin_study/lesson/Gin入门实战/demo/middleware"
)

type Api struct{}

func (a *Api) Login(c *gin.Context) {
	api := &dto.LoginInput{}
	if err := api.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	if api.Username == "admin" && api.Password == "123456" {
		session := sessions.Default(c)
		session.Set("user", api.Username)
		session.Save()
		middleware.ResponseSuccess(c, api)
	} else {
		middleware.ResponseError(c, 2002, errors.New("账户或者密码错误"))
	}
	return
}

func (a *Api) LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	session.Save()
	return
}

func (a *Api) ListPage(c *gin.Context) {
	listInput := &dto.ListPageInput{}

}
