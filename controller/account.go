package controller

import (
	"ApisLdapServe/common/constant"
	"ApisLdapServe/model"
	"errors"
	"gorm.io/gorm"
)

func AccountVerifyController(accountReqSignIn *model.AccountReqSignIn) (error, string, interface{}) {
	var err error
	var account *model.Account
	if accountReqSignIn.LoginType == "email" {
		err, account = model.AccountCheckSQL(0, accountReqSignIn.Account, "")
	}
	if accountReqSignIn.LoginType == "phone" {
		err, account = model.AccountCheckSQL(0, "", accountReqSignIn.Account)
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user isn't existed"), constant.ExistedError, nil
		} else {
			return err, constant.SQLError, nil
		}
	}
	if accountReqSignIn.Password != account.Password {
		return errors.New("the account or password is incorrect"), constant.AccountOrPasswordError, nil
	}
	return nil, constant.Success, account
}

func AccountModifyController() error {
	return nil
}
