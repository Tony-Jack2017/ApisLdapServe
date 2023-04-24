package base

import (
	"ApisLdapServe/common/middleware"
	routerGroup "ApisLdapServe/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	RegisterRouter(router)
	router.Run()
}

func RegisterRouter(router *gin.Engine) {
	router.Use(middleware.CORSMiddleware())
	back := router.Group("/api/back")
	user := back.Group("/user")
	admin := back.Group("/admin")
	account := back.Group("/account")
	routerGroup.RegisterUserRoutes(user)
	routerGroup.RegisterAdminRoutes(admin)
	routerGroup.RegisterAccountRoutes(account)
}
