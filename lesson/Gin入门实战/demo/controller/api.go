package controller

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gin_study/lesson/Gin入门实战/demo/dao"
	"go_gin_study/lesson/Gin入门实战/demo/dto"
	"go_gin_study/lesson/Gin入门实战/demo/middleware"
	"strconv"
	"strings"
	"time"
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
	middleware.ResponseSuccess(c, "success")
	return
}

func (a *Api) ListPage(c *gin.Context) {
	listInput := &dto.ListPageInput{}
	if err := listInput.BindingVaildParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	user := &dao.User{}
	pageInt, err := strconv.ParseInt(listInput.Page, 10, 64)
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	if userList, total, err := user.PageList(listInput.Name, int(pageInt), 20); err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	} else {
		m := map[string]interface{}{
			"list":  userList,
			"total": total,
		}
		middleware.ResponseSuccess(c, m)
	}
	return
}

func (a *Api) AddUser(c *gin.Context) {
	addInput := &dto.AddUserInput{}
	if err := addInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	}
	user := &dao.User{}
	user.Name = addInput.Name
	user.Sex = addInput.Sex
	user.Age = addInput.Age
	user.Birth = addInput.Birth
	user.Addr = addInput.Addr
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	if err := user.Save(); err != nil {
		middleware.ResponseError(c, 2007, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (a *Api) EditUser(c *gin.Context) {
	editInput := &dto.EditUserInput{}
	if err := editInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	}

	user := &dao.User{}
	if userDb, err := user.Find(int64(editInput.ID)); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	} else {
		user = userDb
	}

	user.Name = editInput.Name
	user.Sex = editInput.Sex
	user.Age = editInput.Age
	user.Birth = editInput.Birth
	user.Addr = editInput.Addr
	user.UpdateAt = time.Now()
	if err := user.Save(); err != nil {
		middleware.ResponseError(c, 2007, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}

func (a *Api) RemoveUser(c *gin.Context) {
	removeInput := &dto.RemoveUserInput{}
	if err := removeInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2006, err)
		return
	}

	user := &dao.User{}
	if err := user.Del(strings.Split(removeInput.IDS, ",")); err != nil {
		middleware.ResponseError(c, 2007, err)
		return
	}
	middleware.ResponseSuccess(c, "")
	return
}
