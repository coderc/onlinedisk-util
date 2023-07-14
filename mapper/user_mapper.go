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
	sql := "insert into table_user (`uuid`, `username`, `password`) values (?, ?, ?)"
	conn, err := db.GetConn().Prepare(sql)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(uuid, username, password)
	if err != nil {
		return err
	}

	return nil
}

// GetUser 判断用户是否存在
func GetUser(username, password string) (userModel *model.UserModel, err error) {
	conn, err := db.GetConn().Prepare("select uuid from table_user where username = ? and password = ? limit 1")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(username, password)
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
