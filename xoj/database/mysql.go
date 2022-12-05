package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQLDB = (*gorm.DB)(nil)

func NewMySQLDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "MySqlRoot:370802mysql@tcp(47.100.227.175:3306)/judge_backend?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DontSupportRenameIndex:    true,                                                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                                 // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                                // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		logrus.Error("open mysql error: ", err)
	}
	return db
}

func GetMySQLDB() *gorm.DB {
	if MySQLDB == nil {
		MySQLDB = NewMySQLDB()
	}
	return MySQLDB
}
