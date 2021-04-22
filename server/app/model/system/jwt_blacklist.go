package model

type JwtBlacklist struct {
	ID  uint   `orm:"id" json:"ID" gorm:"primarykey"`
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (j *JwtBlacklist) TableName() string {
	return "jwt_blacklists"
}
