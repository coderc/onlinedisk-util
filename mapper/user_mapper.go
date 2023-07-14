package mapper

import (
	"fmt"

	db "github.com/coderc/onlinedisk-util/db"
	"github.com/coderc/onlinedisk-util/model"
)

const (
	errorUserNotExist = "用户不存在"
)

func InsertUserInfo(uuid int64, username, password string) error {
	_, err := db.GetConn().Exec("insert into table_user (`uuid`, `username`, `password`) values (?, ?, ?)",
		uuid, username, password)
	if err != nil {
		return err
	}

	return nil
}

// GetUser 判断用户是否存在
func GetUser(username, password string) (userModel *model.UserModel, err error) {
	rows, err := db.GetConn().Query("select uuid from table_user where username = ? and password = ? limit 1",
		username, password)
	if err != nil {
		return
	}

	userModel = &model.UserModel{}
	if rows.Next() {
		if err = rows.Scan(&userModel.UUID); err != nil {
			return
		}
		return
	}
	err = fmt.Errorf(errorUserNotExist)
	return
}
