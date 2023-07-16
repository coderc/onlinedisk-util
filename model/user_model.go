package model

type UserModel struct {
	Id         int64  `gorm:"id" json:"id"`
	UUID       int64  `gorm:"uuid" json:"uuid,string"`
	Username   string `gorm:"username" json:"username"`
	Password   string `gorm:"password" json:"-"`
	CreateTime string `gorm:"create_time" json:"create_time"`
	UpdateTime string `gorm:"update_time" json:"update_time"`
}
