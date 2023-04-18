package model

type Code struct {
	CID         string `gorm:"primaryKey;column:c_id"`
	Type        string `gorm:"column:type"`
	CodeType    string `gorm:"column:code_type"`
	Status      string `gorm:"column:status"`
	ExpiredTime int8   `gorm:"column:expired_time"`
}

type Token struct {
	TID         string `gorm:"primaryKey;column:t_id"`
	TokenString string `gorm:"column:token_string"`
}

func (code *Code) TableName() string {
	return "ldap_codes"
}

func (token *Token) TableName() string {
	return "ldap_tokens"
}
