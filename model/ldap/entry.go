package ldap

type Entry struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	DN          string `json:"dn"`
	ObjectClass string `json:"object_class"`
}
