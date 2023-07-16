package model

type UserFileModel struct {
	Id       int64  `gorm:"id" json:"id"`
	UserId   int64  `gorm:"user_id" json:"user_id,string"`
	FileId   int64  `gorm:"file_id" json:"file_id,string"`
	FileName string `gorm:"file_name" json:"file_name"`
}
