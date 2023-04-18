package main

import "ApisLdapServe/common/base"

func main() {
	base.InitViper("dev")
	base.ReadConfig()
}
