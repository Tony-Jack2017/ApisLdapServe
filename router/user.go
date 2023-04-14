package router

import (
	"ApisLdapServe/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(userGroup *gin.RouterGroup) {
	userGroup.POST("/verify", user.VerifyUser)
	userGroup.POST("/create")
}
