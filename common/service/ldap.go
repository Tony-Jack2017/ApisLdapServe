package service

import (
	"ApisLdapServe/common/base"
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

var LD *ldap.Conn

func InitLdap() {

}

func ConnectLdap() {
	url := fmt.Sprintf("%s:%d", base.Conf.Ldap.Host, base.Conf.Ldap.Port)
	//创建与ldap服务器的链接
	conn, err := ldap.DialTLS("tcp", url, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		panic("connect ldap server failed !!!")
		return
	} else {
		LD = conn
	}
}
