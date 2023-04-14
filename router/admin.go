package router

import (
	"ApisLdapServe/api/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(adminGroup *gin.RouterGroup) {
	adminGroup.POST("/create", admin.AddAdmin)
}
