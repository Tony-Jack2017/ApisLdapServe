package base

import (
	routerGroup "ApisLdapServe/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	RegisterRouter(router)
	router.Run()
}

func RegisterRouter(router *gin.Engine) {
	user := router.Group("/user")
	admin := router.Group("/admin")
	account := router.Group("/account")
	routerGroup.RegisterUserRoutes(user)
	routerGroup.RegisterAdminRoutes(admin)
	routerGroup.RegisterAccountRoutes(account)
}
