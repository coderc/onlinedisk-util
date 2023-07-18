package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/coderc/onlinedisk-util/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	dbConn *MysqlConn
)

type MysqlConn struct {
	conn *sql.DB
}

func GetSingleton() *MysqlConn {
	once.Do(func() {
		dbConn = &MysqlConn{}
	})
	return dbConn
}

func GetConn() *sql.DB {
	return GetSingleton().conn
}

func (m *MysqlConn) Init(config *config.DBConfigStruct) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		panic(err)
	}

	m.conn, err = db.DB()
	if err != nil {
		panic(err)
	}

	m.conn.SetMaxIdleConns(config.MaxIdleConns)
	m.conn.SetMaxOpenConns(config.MaxOpenConns)
	m.conn.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)

	err = m.conn.Ping()
	if err != nil {
		panic(err)
	}
}
