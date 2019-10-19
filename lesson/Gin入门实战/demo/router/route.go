package router

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_gin_study/lesson/Gin入门实战/demo/controller"
	"go_gin_study/lesson/Gin入门实战/demo/middleware"
)

func InitRouter(middleWares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middleWares...)

	v1 := router.Group("/demo")
	v1.Use(
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.IPAuthMiddleware(),
		middleware.TranslationMiddleware(),
	)
	{
		controller.DemoRegister(v1)
	}

	store := sessions.NewCookieStore([]byte("secret"))
	apiNormalGroup := router.Group("/api")
	apiController := &controller.Api{}
	apiNormalGroup.Use(
		sessions.Sessions("mySession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware(),
	)
	apiNormalGroup.POST("/login", apiController.Login)
	apiNormalGroup.GET("/loginOut", apiController.LoginOut)

	apiAuthGroup := router.Group("/api")
	apiAuthGroup.Use(
		sessions.Sessions("mySession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware(),
	)
	apiAuthGroup.GET("/user/listPage", apiController.ListPage)
	apiAuthGroup.GET("/user/add", apiController.AddUser)
	apiAuthGroup.GET("/user/edit", apiController.EditUser)
	apiAuthGroup.GET("/user/remove", apiController.RemoveUser)
	apiAuthGroup.GET("/user/batchRemove", apiController.RemoveUser)
	return router
}
