package mapper

import (
	"fmt"

	db "github.com/coderc/onlinedisk-util/db"
	"github.com/coderc/onlinedisk-util/model"
)

const (
	errorUserNotExist = "用户不存在"
)

func InsertUser(uuid int64, username, password string) error {
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

// QueryUser 判断用户是否存在
func QueryUser(username, password string) (userModel *model.UserModel, err error) {
	sql := "select id,uuid,create_time,update_time from table_user where username = ? and password = ? limit 1"
	conn, err := db.GetConn().Prepare(sql)
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
		if err = rows.Scan(
			&userModel.Id,
			&userModel.UUID,
			&userModel.CreateTime,
			&userModel.UpdateTime); err != nil {
			return
		}
		userModel.Username = username
		userModel.Password = password
		return
	}
	err = fmt.Errorf(errorUserNotExist)
	return
}
