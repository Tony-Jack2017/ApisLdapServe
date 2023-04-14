package router

import (
	"ApisLdapServe/api/account"
	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(accountGroup *gin.RouterGroup) {
	accountGroup.POST("/signIn", account.SignIn)
	accountGroup.POST("/forget", account.ForgetPassword)
	accountGroup.POST("/modify_password", account.ModifyPassword)
}
