package account

import (
	"ApisLdapServe/common/constant"
	"ApisLdapServe/controller"
	"ApisLdapServe/model"
	"ApisLdapServe/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SignIn(context *gin.Context) {
	var AcRep model.AccountReqSignIn
	if err := context.ShouldBind(&AcRep); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code":  constant.ParamsError,
			"error": util.GetValidMsg(err, &AcRep),
		})
		return
	}
	errCtr, code, data := controller.AccountVerifyController(&AcRep)
	if errCtr != nil {
		if strings.HasPrefix(code, "AC_ERR") {
			context.JSON(http.StatusBadRequest, gin.H{
				"code":  code,
				"error": errCtr.Error(),
			})
			return
		}
		if strings.HasPrefix(code, "AS_ERR") {
			context.JSON(http.StatusInternalServerError, gin.H{
				"code":  code,
				"error": errCtr.Error(),
			})
			return
		}
	}
	context.JSON(http.StatusOK, model.AccountRepSignIn{
		Code:    code,
		Message: "Login Success !!!",
		Data:    data,
	})
}

func ForgetPassword(context *gin.Context) {
}

func ModifyPassword(context *gin.Context) {
}
