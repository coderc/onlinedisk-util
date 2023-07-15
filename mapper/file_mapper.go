package mapper

import (
	"github.com/coderc/onlinedisk-util/db"
	"github.com/coderc/onlinedisk-util/model"
)

func InsertFile(file *model.FileModel) error {
	sql := "insert into table_file (`uuid`, `sha1`, `user_id`, `name`, `size`, `path`) values (?, ?, ?, ?, ?, ?)"
	conn, err := db.GetConn().Prepare(sql)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(file.UUID, file.SHA1, file.UserId, file.Name, file.Size, file.Path)
	return err
}

// QueryFileByUUID 根据uuid查询文件meta信息
func QueryFileByUUID(uuids []int64) (map[int64]*model.FileModel, error) {
	sql := "select id,uuid,sha1,user_id,name,size,path,create_time,update_time from table_file where uuid in (?)"
	conn, err := db.GetConn().Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(uuids)
	if err != nil {
		return nil, err
	}

	fileModels := make(map[int64]*model.FileModel, len(uuids))
	for rows.Next() {
		fileModel := &model.FileModel{}
		if err = rows.Scan(
			&fileModel.Id,
			&fileModel.UUID,
			&fileModel.SHA1,
			&fileModel.UserId,
			&fileModel.Name,
			&fileModel.Size,
			&fileModel.Path,
			&fileModel.CreateTime,
			&fileModel.UpdateTime); err != nil {
			return fileModels, err
		}
		fileModels[fileModel.UUID] = fileModel
	}
	return fileModels, nil
}

func QueryFileBySHA1(sha1 string) (*model.FileModel, error) {
	sql := "select id,uuid,sha1,user_id,size,path,create_time,update_time from table_file where sha1 = ? limit 1"
	conn, err := db.GetConn().Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(sha1)
	if err != nil {
		return nil, err
	}

	fileModel := &model.FileModel{}
	if rows.Next() {
		if err = rows.Scan(
			&fileModel.Id,
			&fileModel.UUID,
			&fileModel.SHA1,
			&fileModel.UserId,
			&fileModel.Size,
			&fileModel.Path,
			&fileModel.CreateTime,
			&fileModel.UpdateTime); err != nil {
			return nil, err
		}
		return fileModel, nil
	}
	return nil, nil
}
