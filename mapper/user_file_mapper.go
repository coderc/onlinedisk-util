package mapper

import (
	"github.com/coderc/onlinedisk-util/db"
	"github.com/coderc/onlinedisk-util/model"
)

func InsertUserFile(userFileModel *model.UserFileModel) error {
	sql := "insert into table_user_file(user_id, file_id, file_name) values(?, ?, ?)"
	conn, err := db.GetConn().Prepare(sql)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(userFileModel.UserId, userFileModel.FileId, userFileModel.FileName)
	return err
}

func QueryUserFileByUserId(userId int64) ([]*model.FileModel, error) {
	sql := "select table_file.uuid as file_uuid, table_file.sha1 as sha1, table_user_file.file_name as file_name, table_file.path as path, table_file.size as size, table_user_file.create_time as create_time, table_user_file.update_time as update_time from table_user_file left join table_file on table_user_file.file_id = table_file.uuid where table_user_file.user_id = ?"
	conn, err := db.GetConn().Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(userId)
	if err != nil {
		return nil, err
	}

	fileModels := make([]*model.FileModel, 0)
	for rows.Next() {
		fileModel := &model.FileModel{}
		if err = rows.Scan(
			&fileModel.UUID,
			&fileModel.SHA1,
			&fileModel.Name,
			&fileModel.Path,
			&fileModel.Size,
			&fileModel.CreateTime,
			&fileModel.UpdateTime); err != nil {
			return nil, err
		}
		fileModel.UserId = userId
		fileModels = append(fileModels, fileModel)
	}
	return fileModels, nil
}
