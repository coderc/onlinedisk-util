package model

type UserModel struct {
	Id         int64  `gorm:"id"`
	UUID       int64  `gorm:"uuid"`
	Username   string `gorm:"username"`
	Password   string `gorm:"password"`
	CreateTime string `gorm:"create_time"`
	UpdateTime string `gorm:"update_time"`
}
