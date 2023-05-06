package admin

import (
	"ApisLdapServe/common/constant"
	model2 "ApisLdapServe/common/model"
	"ApisLdapServe/controller"
	"ApisLdapServe/model"
	"ApisLdapServe/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAdmin(context *gin.Context) {
	var AdRep model.AdminReqCreate
	if err := context.ShouldBind(&AdRep); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  constant.ParamsError,
			"error": util.GetValidMsg(err, &AdRep),
		})
		return
	}
	errCtr, code, data := controller.AdminCreateController(&AdRep)
	if errCtr != nil {
		return
	}
	context.JSON(http.StatusOK, model2.Response{
		Code: code,
		Data: data,
	})
}

func GetAdminInfo(context *gin.Context) {

}
