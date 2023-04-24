package controller

import (
	"ApisLdapServe/common/constant"
	"ApisLdapServe/model"
)

func AccountVerifyController(accountReqSignIn *model.AccountReqSignIn) (error, string, interface{}) {
	return nil, constant.Success, nil
}

func AccountModifyController() error {
	return nil
}
