package model

type FileModel struct {
	Id         int64  `gorm:"id" json:"id"`
	UUID       int64  `gorm:"uuid" json:"uuid,string"`
	SHA1       string `gorm:"sha1" json:"sha1"`
	UserId     int64  `gorm:"user_id" json:"user_id,string"`
	Name       string `gorm:"name" json:"name"`
	Path       string `gorm:"path" json:"path"`
	Size       int64  `gorm:"size" json:"size"`
	Status     int    `gorm:"status" json:"status"`
	CreateTime string `gorm:"create_time" json:"create_time"`
	UpdateTime string `gorm:"update_time" json:"update_time"`
	Ex1        int    `gorm:"ex1" json:"ex1"`
	Ex2        string `gorm:"ex2" json:"ex2"`
}
